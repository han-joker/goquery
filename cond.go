package goquery

import (
	"fmt"
	"strings"
)

type Cond struct {
	left interface{}
	operator interface{}
	right interface{}
}

type Field string

type LogicOp string
type CompareOp string

func (c Cond) mergeCond(target Cond, logic LogicOp) Cond {
	if c.right == nil && c.operator == nil {
		return target
	}

	oldC := c
	newC := Cond{
		left:     oldC,
		operator: logic,
		right:    target,
	}

	return newC
}

func (c Cond) stringify() string {
	fmt.Println(c)
	// 条件为比较运算
	if op, ok := c.operator.(CompareOp); ok {
		switch strings.Trim(strings.ToLower(string(op)), " ") {
		case "=":
			return "(" + expStringify(c.left) + " = ?)"
		case "!=":
			return "(" + expStringify(c.left) + " != ?)"
		case ">":
			return "(" + expStringify(c.left) + " > ?)"
		case "<":
			return "(" + expStringify(c.left) + " < ?)"
		case ">=":
			return "(" + expStringify(c.left) + " >= ?)"
		case "<=":
			return "(" + expStringify(c.left) + " <= ?)"
		case "<=>":
			return "(" + expStringify(c.left) + " <=> ?)"
		case "like":
			return "(" + expStringify(c.left) + " LIKE ?)"
		case "not like":
			return "(" + expStringify(c.left) + " NOT LIKE ?)"
		case "between":
			return "(" + expStringify(c.left) + " BETWEEN ? AND ?)"
		case "not between":
			return "(" + expStringify(c.left) + " NOT BETWEEN ? AND ?)"
		case "in":
			return "(" + expStringify(c.left) + " in (?)"
		case "not in":
			return "(" + expStringify(c.left) + " NOT IN (?)"
		}
	} else if op, ok := c.operator.(LogicOp); ok {
		switch op {
		case "and":
		case "or":
		}
	}
	return ""
}

func expStringify(exp interface{}) string {
	switch e := exp.(type) {
	case Field:
		return quoteField(string(e))
	}
	return ""
}
