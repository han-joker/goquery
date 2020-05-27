package goquery

const (
	PROD = 1 << iota
	DEV
	DEBUG
)

const defaultDriverName = "mysql"
const defaultMode = DEV
