package schedule

import (
	"myrrh/business"
	"myrrh/calculation"
)

type Element struct {
	Feature  string
	Value    int
	Relation string
}

type Expression struct {
	Element *Element
	Logic   string
	Exprs   []*Expression
	Result  bool
}

func Evaluation(root *Expression) bool {
	if root == nil {
		return false
	}
	if root.Logic == "" {
		fc := calculation.ExplainRelation(root.Element.Relation)
		if fc == nil {
			return false
		}
		return fc(
			business.Eval(root.Element.Feature),
			root.Element.Value,
		)
	}

	stack := NewStack()
	queue := []*Expression{root}

	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			queue = queue[1:]
			stack.Push(node)

			for _, expr := range node.Exprs {
				if expr.Logic == "" {
					continue
				}
				queue = append(queue, expr)
			}
		}
	}

	for !stack.Empty() {
		cur := stack.Pop().(*Expression)

		var rs []bool
		for _, expr := range cur.Exprs {
			if expr.Logic != "" {
				rs = append(rs, expr.Result)
				continue
			}
			fc := calculation.ExplainRelation(expr.Element.Relation)
			if fc == nil {
				return false
			}
			expr.Result = fc(
				business.Eval(expr.Element.Feature),
				expr.Element.Value,
			)
			rs = append(rs, expr.Result)
		}

		fc := calculation.ExplainLogic(cur.Logic)
		if fc == nil {
			return false
		}
		cur.Result = fc(rs...)
	}
	return root.Result
}
