package goquery

const (
	PROD = 1 << iota
	DEV
	DEBUG
)

const defaultDriverName = "mysql"
const defaultMode = DEV

var compareOps = []string{
	"=", "<", ">", "<=", ">=", "<>", "!=", "<=>",
	"like", "not like",
	"in", "not in",
	"between", "not between",
	"is null", "is not null",
}

var logicOs = []string{
	"and", "or", "not", "xor",
}