package goquery

import (
	"fmt"
	"strings"
)

// -- 构造查询 --
func (q *Query) clearBuilder() {
	q.table = ""
}


func (q *Query) Table(tables ...string) *Query {
	ts := make([]string, len(tables))
	for i, t := range tables {
		ts[i] = quoteTable(t)
	}
	q.table = strings.Join(ts, ", ")
	return q
}

func (q *Query) Select() *Query {
	return q
}

func (q *Query) AddSelect() *Query {
	return q
}

func (q *Query) Join() *Query {
	return q
}

func (q *Query) LeftJoin() *Query {
	return q
}

func (q *Query) RightJoin() *Query {
	return q
}

func (q *Query) JoinSub() *Query {
	return q
}

func (q *Query) CrossJoin() *Query {
	return q
}

func (q *Query) Where(field string, operator string, operand ...interface{}) *Query {

	cond := Cond{
		left: Field(field),
		operator: CompareOp(operator),
		right: operand,
	}

	// 合并条件
	w := q.whereCond.mergeCond(cond, "and").stringify()
	fmt.Println(w)
	return q
}

func (q *Query) OrWhere() *Query {
	return q
}

func (q *Query) WhereBetween() *Query {
	return q
}

func (q *Query) WhereNotBetween() *Query {
	return q
}

func (q *Query) OrWhereBetween() *Query {
	return q
}

func (q *Query) OrWhereNotBetween() *Query {
	return q
}

func (q *Query) WhereIn() *Query {
	return q
}

func (q *Query) WhereNotIn() *Query {
	return q
}

func (q *Query) OrWhereIn() *Query {
	return q
}

func (q *Query) OrWhereNotIn() *Query {
	return q
}

func (q *Query) WhereNull() *Query {
	return q
}

func (q *Query) WhereNotNull() *Query {
	return q
}

func (q *Query) OrWhereNull() *Query {
	return q
}

func (q *Query) OrWhereNotNull() *Query {
	return q
}

func (q *Query) WhereDate() *Query {
	return q
}

func (q *Query) WhereMonth() *Query {
	return q
}

func (q *Query) WhereDay() *Query {
	return q
}

func (q *Query) WhereYear() *Query {
	return q
}

func (q *Query) WhereTime() *Query {
	return q
}

func (q *Query) WhereColumn() *Query {
	return q
}

func (q *Query) OrWhereColumn() *Query {
	return q
}

func (q *Query) WhereExists() *Query {
	return q
}

func (q *Query) WhereJsonContains() *Query {
	return q
}

func (q *Query) WhereJsonLength() *Query {
	return q
}

func (q *Query) OrderBy() *Query {
	return q
}

func (q *Query) Latest() *Query {
	return q
}

func (q *Query) Oldest() *Query {
	return q
}

func (q *Query) InRandomOrder() *Query {
	return q
}

func (q *Query) Reorder() *Query {
	return q
}

func (q *Query) GroupBy() *Query {
	return q
}

func (q *Query) Having() *Query {
	return q
}

func (q *Query) Limit() *Query {
	return q
}

func (q *Query) Offset() *Query {
	return q
}

func (q *Query) Skip() *Query {
	return q
}

func (q *Query) Take() *Query {
	return q
}

func (q *Query) When() *Query {
	return q
}

func (q *Query) SelectRaw() *Query {
	return q
}

func (q *Query) WhereRaw() *Query {
	return q
}

func (q *Query) OrWhereRaw() *Query {
	return q
}

func (q *Query) HavingRaw() *Query {
	return q
}

func (q *Query) OrHavingRaw() *Query {
	return q
}

func (q *Query) OrderByRaw() *Query {
	return q
}

func (q *Query) GroupByRaw() *Query {
	return q
}

func (q *Query) Union() *Query {
	return q
}

func (q *Query) SharedLock() *Query {
	return q
}

func (q *Query) LockForUpdate() *Query {
	return q
}

