// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package build

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/opt/exec"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/opt"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/xform"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

type execPlan struct {
	root exec.Node

	// outputCols maps columns in the output set of a relational expression to
	// indices in the result columns of the ExecNode.
	//
	// The reason we need to keep track of this (instead of using just the
	// relational properties) is that the relational properties don't force a
	// single "schema": any ordering of the output columns is possible. We choose
	// the schema that is most convenient: for scans, we use the table's column
	// ordering. Consider:
	//   SELECT a, b FROM t WHERE a = b
	// and the following two cases:
	//   1. The table is defined as (k INT PRIMARY KEY, a INT, b INT). The scan will
	//      return (k, a, b).
	//   2. The table is defined as (k INT PRIMARY KEY, b INT, a INT). The scan will
	//      return (k, b, a).
	// In these two cases, the relational properties are effectively the same.
	//
	// An alternative to this would be to always use a "canonical" schema, for
	// example the output columns in increasing index order. However, this would
	// require a lot of otherwise unnecessary projections.
	outputCols opt.ColMap
}

func (b *Builder) buildRelational(ev xform.ExprView) (execPlan, error) {
	switch ev.Operator() {
	case opt.ScanOp:
		return b.buildScan(ev)

	case opt.SelectOp:
		return b.buildSelect(ev)

	default:
		panic(fmt.Sprintf("unsupported relational op %s", ev.Operator()))
	}
}

func (b *Builder) buildScan(ev xform.ExprView) (execPlan, error) {
	tblIndex := ev.Private().(opt.TableIndex)
	md := ev.Metadata()
	tbl := md.Table(tblIndex)
	node, err := b.factory.ConstructScan(tbl)
	if err != nil {
		return execPlan{}, err
	}
	res := execPlan{root: node}
	for i := 0; i < tbl.NumColumns(); i++ {
		res.outputCols.Set(int(md.TableColumn(tblIndex, i)), i)
	}
	return res, nil
}

func (b *Builder) buildSelect(ev xform.ExprView) (execPlan, error) {
	input, err := b.buildRelational(ev.Child(0))
	if err != nil {
		return execPlan{}, err
	}
	// We need to be able to create IndexedVars for the columns of the input plan.
	// We get the maximum value from the map to find out how many we need to
	// support. In most cases this is the same as the Len() of the map; it is TBD
	// if we will have cases where this doesn't hold.
	// TODO(radu): Should we just remember the number of columns in execPlan?
	num := 0
	if max, ok := input.outputCols.MaxValue(); ok {
		num = max + 1
	}
	ctx := buildScalarCtx{
		ivh:     tree.MakeIndexedVarHelper(nil /* no container */, num),
		ivarMap: input.outputCols,
	}
	filter := b.buildScalar(&ctx, ev.Child(1))
	node, err := b.factory.ConstructFilter(input.root, filter)
	if err != nil {
		return execPlan{}, err
	}
	return execPlan{
		root: node,
		// A filtering node does not modify the schema.
		outputCols: input.outputCols,
	}, nil
}
