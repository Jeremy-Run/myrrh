package schedule

import (
	"fmt"
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

func ExpressionEvaluation(root *Expression) bool {
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
			fmt.Printf(">>> [user.%s() %s %d] execute result is: %t \n",
				expr.Element.Feature, expr.Element.Relation, expr.Element.Value, expr.Result)
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

type ExpressionDescription struct {
	Element *Element
	Logic   string
	Exprs   []*ExpressionDescription
	Desc    string
}

func SimplifyExpression(root *ExpressionDescription) string {
	if root == nil {
		return ""
	}
	stack := NewStack()
	queue := []*ExpressionDescription{root}

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
		cur := stack.Pop().(*ExpressionDescription)
		childsLen := len(cur.Exprs)
		if childsLen == 0 {
			cur.Desc = fmt.Sprintf("(user.%s() %s %d)",
				cur.Element.Feature, cur.Element.Relation, cur.Element.Value)
			continue
		}

		for index, expr := range cur.Exprs {
			if index == 0 {
				cur.Desc = "("
			}

			if expr.Logic == "" {
				if index == childsLen-1 {
					cur.Desc += fmt.Sprintf("(user.%s() %s %d)",
						expr.Element.Feature, expr.Element.Relation, expr.Element.Value)
				} else {
					cur.Desc += fmt.Sprintf("(user.%s() %s %d) %s ",
						expr.Element.Feature, expr.Element.Relation, expr.Element.Value, cur.Logic)
				}

			} else {
				if index == childsLen-1 {
					cur.Desc += expr.Desc
				} else {
					cur.Desc += expr.Desc + " " + cur.Logic + " "
				}

			}

			if index == childsLen-1 {
				cur.Desc += ")"
			}
		}
	}
	return root.Desc
}
