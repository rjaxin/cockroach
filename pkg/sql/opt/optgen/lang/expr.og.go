// Code generated by langgen; DO NOT EDIT.

package lang

import (
	"bytes"
	"fmt"
)

type RootExpr struct {
	Defines DefineSetExpr
	Rules   RuleSetExpr
	Src     *SourceLoc
}

func (e *RootExpr) Op() Operator {
	return RootOp
}

func (e *RootExpr) ChildCount() int {
	return 2
}

func (e *RootExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Defines
	case 1:
		return &e.Rules
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *RootExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Defines"
	case 1:
		return "Rules"
	}
	return ""
}

func (e *RootExpr) Value() interface{} {
	return nil
}

func (e *RootExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&RootExpr{Defines: *children[0].(*DefineSetExpr), Rules: *children[1].(*RuleSetExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *RootExpr) Source() *SourceLoc {
	return e.Src
}

func (e *RootExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *RootExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type DefineSetExpr []*DefineExpr

func (e *DefineSetExpr) Op() Operator {
	return DefineSetOp
}

func (e *DefineSetExpr) ChildCount() int {
	return len(*e)
}

func (e *DefineSetExpr) Child(nth int) Expr {
	return (*e)[nth]
}

func (e *DefineSetExpr) ChildName(nth int) string {
	return ""
}

func (e *DefineSetExpr) Value() interface{} {
	return nil
}

func (e *DefineSetExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		typedChildren := make(DefineSetExpr, len(children))
		for i := 0; i < len(children); i++ {
			typedChildren[i] = children[i].(*DefineExpr)
		}
		return &typedChildren
	}
	return accept(e)
}

func (e *DefineSetExpr) Source() *SourceLoc {
	return nil
}

func (e *DefineSetExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *DefineSetExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type RuleSetExpr []*RuleExpr

func (e *RuleSetExpr) Op() Operator {
	return RuleSetOp
}

func (e *RuleSetExpr) ChildCount() int {
	return len(*e)
}

func (e *RuleSetExpr) Child(nth int) Expr {
	return (*e)[nth]
}

func (e *RuleSetExpr) ChildName(nth int) string {
	return ""
}

func (e *RuleSetExpr) Value() interface{} {
	return nil
}

func (e *RuleSetExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		typedChildren := make(RuleSetExpr, len(children))
		for i := 0; i < len(children); i++ {
			typedChildren[i] = children[i].(*RuleExpr)
		}
		return &typedChildren
	}
	return accept(e)
}

func (e *RuleSetExpr) Source() *SourceLoc {
	return nil
}

func (e *RuleSetExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *RuleSetExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type DefineExpr struct {
	Tags   TagsExpr
	Name   StringExpr
	Fields DefineFieldsExpr
	Src    *SourceLoc
}

func (e *DefineExpr) Op() Operator {
	return DefineOp
}

func (e *DefineExpr) ChildCount() int {
	return 3
}

func (e *DefineExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Tags
	case 1:
		return &e.Name
	case 2:
		return &e.Fields
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *DefineExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Tags"
	case 1:
		return "Name"
	case 2:
		return "Fields"
	}
	return ""
}

func (e *DefineExpr) Value() interface{} {
	return nil
}

func (e *DefineExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&DefineExpr{Tags: *children[0].(*TagsExpr), Name: *children[1].(*StringExpr), Fields: *children[2].(*DefineFieldsExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *DefineExpr) Source() *SourceLoc {
	return e.Src
}

func (e *DefineExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *DefineExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type TagsExpr []TagExpr

func (e *TagsExpr) Op() Operator {
	return TagsOp
}

func (e *TagsExpr) ChildCount() int {
	return len(*e)
}

func (e *TagsExpr) Child(nth int) Expr {
	return &(*e)[nth]
}

func (e *TagsExpr) ChildName(nth int) string {
	return ""
}

func (e *TagsExpr) Value() interface{} {
	return nil
}

func (e *TagsExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		typedChildren := make(TagsExpr, len(children))
		for i := 0; i < len(children); i++ {
			typedChildren[i] = *children[i].(*TagExpr)
		}
		return &typedChildren
	}
	return accept(e)
}

func (e *TagsExpr) Source() *SourceLoc {
	return nil
}

func (e *TagsExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *TagsExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type TagExpr string

func (e *TagExpr) Op() Operator {
	return TagOp
}

func (e *TagExpr) ChildCount() int {
	return 0
}

func (e *TagExpr) Child(nth int) Expr {
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *TagExpr) ChildName(nth int) string {
	return ""
}

func (e *TagExpr) Value() interface{} {
	return string(*e)
}

func (e *TagExpr) Visit(accept AcceptFunc) Expr {
	return accept(e)
}

func (e *TagExpr) Source() *SourceLoc {
	return nil
}

func (e *TagExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *TagExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type DefineFieldsExpr []*DefineFieldExpr

func (e *DefineFieldsExpr) Op() Operator {
	return DefineFieldsOp
}

func (e *DefineFieldsExpr) ChildCount() int {
	return len(*e)
}

func (e *DefineFieldsExpr) Child(nth int) Expr {
	return (*e)[nth]
}

func (e *DefineFieldsExpr) ChildName(nth int) string {
	return ""
}

func (e *DefineFieldsExpr) Value() interface{} {
	return nil
}

func (e *DefineFieldsExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		typedChildren := make(DefineFieldsExpr, len(children))
		for i := 0; i < len(children); i++ {
			typedChildren[i] = children[i].(*DefineFieldExpr)
		}
		return &typedChildren
	}
	return accept(e)
}

func (e *DefineFieldsExpr) Source() *SourceLoc {
	return nil
}

func (e *DefineFieldsExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *DefineFieldsExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type DefineFieldExpr struct {
	Name StringExpr
	Type StringExpr
	Src  *SourceLoc
}

func (e *DefineFieldExpr) Op() Operator {
	return DefineFieldOp
}

func (e *DefineFieldExpr) ChildCount() int {
	return 2
}

func (e *DefineFieldExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Name
	case 1:
		return &e.Type
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *DefineFieldExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Name"
	case 1:
		return "Type"
	}
	return ""
}

func (e *DefineFieldExpr) Value() interface{} {
	return nil
}

func (e *DefineFieldExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&DefineFieldExpr{Name: *children[0].(*StringExpr), Type: *children[1].(*StringExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *DefineFieldExpr) Source() *SourceLoc {
	return e.Src
}

func (e *DefineFieldExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *DefineFieldExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type RuleExpr struct {
	Name    StringExpr
	Tags    TagsExpr
	Match   *MatchExpr
	Replace Expr
	Src     *SourceLoc
}

func (e *RuleExpr) Op() Operator {
	return RuleOp
}

func (e *RuleExpr) ChildCount() int {
	return 4
}

func (e *RuleExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Name
	case 1:
		return &e.Tags
	case 2:
		return e.Match
	case 3:
		return e.Replace
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *RuleExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Name"
	case 1:
		return "Tags"
	case 2:
		return "Match"
	case 3:
		return "Replace"
	}
	return ""
}

func (e *RuleExpr) Value() interface{} {
	return nil
}

func (e *RuleExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&RuleExpr{Name: *children[0].(*StringExpr), Tags: *children[1].(*TagsExpr), Match: children[2].(*MatchExpr), Replace: children[3], Src: e.Source()})
	}
	return accept(e)
}

func (e *RuleExpr) Source() *SourceLoc {
	return e.Src
}

func (e *RuleExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *RuleExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type BindExpr struct {
	Label  StringExpr
	Target Expr
	Src    *SourceLoc
}

func (e *BindExpr) Op() Operator {
	return BindOp
}

func (e *BindExpr) ChildCount() int {
	return 2
}

func (e *BindExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Label
	case 1:
		return e.Target
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *BindExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Label"
	case 1:
		return "Target"
	}
	return ""
}

func (e *BindExpr) Value() interface{} {
	return nil
}

func (e *BindExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&BindExpr{Label: *children[0].(*StringExpr), Target: children[1], Src: e.Source()})
	}
	return accept(e)
}

func (e *BindExpr) Source() *SourceLoc {
	return e.Src
}

func (e *BindExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *BindExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type RefExpr struct {
	Label StringExpr
	Src   *SourceLoc
}

func (e *RefExpr) Op() Operator {
	return RefOp
}

func (e *RefExpr) ChildCount() int {
	return 1
}

func (e *RefExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Label
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *RefExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Label"
	}
	return ""
}

func (e *RefExpr) Value() interface{} {
	return nil
}

func (e *RefExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&RefExpr{Label: *children[0].(*StringExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *RefExpr) Source() *SourceLoc {
	return e.Src
}

func (e *RefExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *RefExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type MatchExpr struct {
	Names OpNamesExpr
	Args  ListExpr
	Src   *SourceLoc
}

func (e *MatchExpr) Op() Operator {
	return MatchOp
}

func (e *MatchExpr) ChildCount() int {
	return 2
}

func (e *MatchExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Names
	case 1:
		return &e.Args
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *MatchExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Names"
	case 1:
		return "Args"
	}
	return ""
}

func (e *MatchExpr) Value() interface{} {
	return nil
}

func (e *MatchExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&MatchExpr{Names: *children[0].(*OpNamesExpr), Args: *children[1].(*ListExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *MatchExpr) Source() *SourceLoc {
	return e.Src
}

func (e *MatchExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *MatchExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type OpNamesExpr []OpNameExpr

func (e *OpNamesExpr) Op() Operator {
	return OpNamesOp
}

func (e *OpNamesExpr) ChildCount() int {
	return len(*e)
}

func (e *OpNamesExpr) Child(nth int) Expr {
	return &(*e)[nth]
}

func (e *OpNamesExpr) ChildName(nth int) string {
	return ""
}

func (e *OpNamesExpr) Value() interface{} {
	return nil
}

func (e *OpNamesExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		typedChildren := make(OpNamesExpr, len(children))
		for i := 0; i < len(children); i++ {
			typedChildren[i] = *children[i].(*OpNameExpr)
		}
		return &typedChildren
	}
	return accept(e)
}

func (e *OpNamesExpr) Source() *SourceLoc {
	return nil
}

func (e *OpNamesExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *OpNamesExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type OpNameExpr string

func (e *OpNameExpr) Op() Operator {
	return OpNameOp
}

func (e *OpNameExpr) ChildCount() int {
	return 0
}

func (e *OpNameExpr) Child(nth int) Expr {
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *OpNameExpr) ChildName(nth int) string {
	return ""
}

func (e *OpNameExpr) Value() interface{} {
	return string(*e)
}

func (e *OpNameExpr) Visit(accept AcceptFunc) Expr {
	return accept(e)
}

func (e *OpNameExpr) Source() *SourceLoc {
	return nil
}

func (e *OpNameExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *OpNameExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type MatchAndExpr struct {
	Left  Expr
	Right Expr
	Src   *SourceLoc
}

func (e *MatchAndExpr) Op() Operator {
	return MatchAndOp
}

func (e *MatchAndExpr) ChildCount() int {
	return 2
}

func (e *MatchAndExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return e.Left
	case 1:
		return e.Right
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *MatchAndExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Left"
	case 1:
		return "Right"
	}
	return ""
}

func (e *MatchAndExpr) Value() interface{} {
	return nil
}

func (e *MatchAndExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&MatchAndExpr{Left: children[0], Right: children[1], Src: e.Source()})
	}
	return accept(e)
}

func (e *MatchAndExpr) Source() *SourceLoc {
	return e.Src
}

func (e *MatchAndExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *MatchAndExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type MatchInvokeExpr struct {
	FuncName StringExpr
	Args     ListExpr
	Src      *SourceLoc
}

func (e *MatchInvokeExpr) Op() Operator {
	return MatchInvokeOp
}

func (e *MatchInvokeExpr) ChildCount() int {
	return 2
}

func (e *MatchInvokeExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.FuncName
	case 1:
		return &e.Args
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *MatchInvokeExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "FuncName"
	case 1:
		return "Args"
	}
	return ""
}

func (e *MatchInvokeExpr) Value() interface{} {
	return nil
}

func (e *MatchInvokeExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&MatchInvokeExpr{FuncName: *children[0].(*StringExpr), Args: *children[1].(*ListExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *MatchInvokeExpr) Source() *SourceLoc {
	return e.Src
}

func (e *MatchInvokeExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *MatchInvokeExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type MatchNotExpr struct {
	Input Expr
	Src   *SourceLoc
}

func (e *MatchNotExpr) Op() Operator {
	return MatchNotOp
}

func (e *MatchNotExpr) ChildCount() int {
	return 1
}

func (e *MatchNotExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return e.Input
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *MatchNotExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Input"
	}
	return ""
}

func (e *MatchNotExpr) Value() interface{} {
	return nil
}

func (e *MatchNotExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&MatchNotExpr{Input: children[0], Src: e.Source()})
	}
	return accept(e)
}

func (e *MatchNotExpr) Source() *SourceLoc {
	return e.Src
}

func (e *MatchNotExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *MatchNotExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type MatchAnyExpr struct {
}

func (e *MatchAnyExpr) Op() Operator {
	return MatchAnyOp
}

func (e *MatchAnyExpr) ChildCount() int {
	return 0
}

func (e *MatchAnyExpr) Child(nth int) Expr {
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *MatchAnyExpr) ChildName(nth int) string {
	return ""
}

func (e *MatchAnyExpr) Value() interface{} {
	return nil
}

func (e *MatchAnyExpr) Visit(accept AcceptFunc) Expr {
	return accept(e)
}

func (e *MatchAnyExpr) Source() *SourceLoc {
	return nil
}

func (e *MatchAnyExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *MatchAnyExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type MatchListExpr struct {
	MatchItem Expr
	Src       *SourceLoc
}

func (e *MatchListExpr) Op() Operator {
	return MatchListOp
}

func (e *MatchListExpr) ChildCount() int {
	return 1
}

func (e *MatchListExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return e.MatchItem
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *MatchListExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "MatchItem"
	}
	return ""
}

func (e *MatchListExpr) Value() interface{} {
	return nil
}

func (e *MatchListExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&MatchListExpr{MatchItem: children[0], Src: e.Source()})
	}
	return accept(e)
}

func (e *MatchListExpr) Source() *SourceLoc {
	return e.Src
}

func (e *MatchListExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *MatchListExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type ConstructExpr struct {
	Name Expr
	Args ListExpr
	Src  *SourceLoc
}

func (e *ConstructExpr) Op() Operator {
	return ConstructOp
}

func (e *ConstructExpr) ChildCount() int {
	return 2
}

func (e *ConstructExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return e.Name
	case 1:
		return &e.Args
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *ConstructExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Name"
	case 1:
		return "Args"
	}
	return ""
}

func (e *ConstructExpr) Value() interface{} {
	return nil
}

func (e *ConstructExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&ConstructExpr{Name: children[0], Args: *children[1].(*ListExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *ConstructExpr) Source() *SourceLoc {
	return e.Src
}

func (e *ConstructExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *ConstructExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type ConstructListExpr struct {
	Items ListExpr
	Src   *SourceLoc
}

func (e *ConstructListExpr) Op() Operator {
	return ConstructListOp
}

func (e *ConstructListExpr) ChildCount() int {
	return 1
}

func (e *ConstructListExpr) Child(nth int) Expr {
	switch nth {
	case 0:
		return &e.Items
	}
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *ConstructListExpr) ChildName(nth int) string {
	switch nth {
	case 0:
		return "Items"
	}
	return ""
}

func (e *ConstructListExpr) Value() interface{} {
	return nil
}

func (e *ConstructListExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		return accept(&ConstructListExpr{Items: *children[0].(*ListExpr), Src: e.Source()})
	}
	return accept(e)
}

func (e *ConstructListExpr) Source() *SourceLoc {
	return e.Src
}

func (e *ConstructListExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *ConstructListExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type ListExpr []Expr

func (e *ListExpr) Op() Operator {
	return ListOp
}

func (e *ListExpr) ChildCount() int {
	return len(*e)
}

func (e *ListExpr) Child(nth int) Expr {
	return (*e)[nth]
}

func (e *ListExpr) ChildName(nth int) string {
	return ""
}

func (e *ListExpr) Value() interface{} {
	return nil
}

func (e *ListExpr) Visit(accept AcceptFunc) Expr {
	children := visitExprChildren(e, accept)
	if children != nil {
		typedChildren := ListExpr(children)
		return &typedChildren
	}
	return accept(e)
}

func (e *ListExpr) Source() *SourceLoc {
	return nil
}

func (e *ListExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *ListExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}

type StringExpr string

func (e *StringExpr) Op() Operator {
	return StringOp
}

func (e *StringExpr) ChildCount() int {
	return 0
}

func (e *StringExpr) Child(nth int) Expr {
	panic(fmt.Sprintf("child index %d is out of range", nth))
}

func (e *StringExpr) ChildName(nth int) string {
	return ""
}

func (e *StringExpr) Value() interface{} {
	return string(*e)
}

func (e *StringExpr) Visit(accept AcceptFunc) Expr {
	return accept(e)
}

func (e *StringExpr) Source() *SourceLoc {
	return nil
}

func (e *StringExpr) String() string {
	var buf bytes.Buffer
	e.Format(&buf, 0)
	return buf.String()
}

func (e *StringExpr) Format(buf *bytes.Buffer, level int) {
	formatExpr(e, buf, level)
}
