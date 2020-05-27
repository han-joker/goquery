package goquery

type Builder struct {

}

func newBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Table() *Builder {
	return b
}

func (b *Builder) Select() *Builder {
	return b
}

func (b *Builder) AddSelect() *Builder {
	return b
}

func (b *Builder) Join() *Builder {
	return b
}

func (b *Builder) LeftJoin() *Builder {
	return b
}

func (b *Builder) RightJoin() *Builder {
	return b
}

func (b *Builder) JoinSub() *Builder {
	return b
}

func (b *Builder) CrossJoin() *Builder {
	return b
}

func (b *Builder) Where() *Builder {
	return b
}

func (b *Builder) OrWhere() *Builder {
	return b
}

func (b *Builder) WhereBetween() *Builder {
	return b
}

func (b *Builder) WhereNotBetween() *Builder {
	return b
}

func (b *Builder) OrWhereBetween() *Builder {
	return b
}

func (b *Builder) OrWhereNotBetween() *Builder {
	return b
}

func (b *Builder) WhereIn() *Builder {
	return b
}

func (b *Builder) WhereNotIn() *Builder {
	return b
}

func (b *Builder) OrWhereIn() *Builder {
	return b
}

func (b *Builder) OrWhereNotIn() *Builder {
	return b
}

func (b *Builder) WhereNull() *Builder {
	return b
}

func (b *Builder) WhereNotNull() *Builder {
	return b
}

func (b *Builder) OrWhereNull() *Builder {
	return b
}

func (b *Builder) OrWhereNotNull() *Builder {
	return b
}

func (b *Builder) WhereDate() *Builder {
	return b
}

func (b *Builder) WhereMonth() *Builder {
	return b
}

func (b *Builder) WhereDay() *Builder {
	return b
}

func (b *Builder) WhereYear() *Builder {
	return b
}

func (b *Builder) WhereTime() *Builder {
	return b
}

func (b *Builder) WhereColumn() *Builder {
	return b
}

func (b *Builder) OrWhereColumn() *Builder {
	return b
}

func (b *Builder) WhereExists() *Builder {
	return b
}

func (b *Builder) WhereJsonContains() *Builder {
	return b
}

func (b *Builder) WhereJsonLength() *Builder {
	return b
}

func (b *Builder) OrderBy() *Builder {
	return b
}

func (b *Builder) Latest() *Builder {
	return b
}

func (b *Builder) Oldest() *Builder {
	return b
}

func (b *Builder) InRandomOrder() *Builder {
	return b
}

func (b *Builder) Reorder() *Builder {
	return b
}

func (b *Builder) GroupBy() *Builder {
	return b
}

func (b *Builder) Having() *Builder {
	return b
}

func (b *Builder) Limit() *Builder {
	return b
}

func (b *Builder) Offset() *Builder {
	return b
}

func (b *Builder) Skip() *Builder {
	return b
}

func (b *Builder) Take() *Builder {
	return b
}

func (b *Builder) When() *Builder {
	return b
}

func (b *Builder) SelectRaw() *Builder {
	return b
}

func (b *Builder) WhereRaw() *Builder {
	return b
}

func (b *Builder) OrWhereRaw() *Builder {
	return b
}

func (b *Builder) HavingRaw() *Builder {
	return b
}

func (b *Builder) OrHavingRaw() *Builder {
	return b
}

func (b *Builder) OrderByRaw() *Builder {
	return b
}

func (b *Builder) GroupByRaw() *Builder {
	return b
}

func (b *Builder) Union() *Builder {
	return b
}

func (b *Builder) SharedLock() *Builder {
	return b
}

func (b *Builder) LockForUpdate() *Builder {
	return b
}
