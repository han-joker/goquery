package goquery

import (
	"database/sql"
	"log"
)

type Query struct {
	// query-builder
	*Builder
	// pool of idle connections
	db *sql.DB
	// The last query
	lastQuery string
	// The last error
	lastErr error
	// running mode
	mode int
}

// constructor
func New(DSN string) (*Query, error) {
	// 建立连接池
	db, err := sql.Open(defaultDriverName, DSN)
	if err != nil {
		return nil, err
	}
	// DAO 对象
	return &Query{
		db:   db,
		Builder: newBuilder(),
		mode: defaultMode,
	}, nil
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (dao *Query) Ping() error {
	return dao.db.Ping()
}

// Exec executes a query without returning any rows.
func (dao *Query) Exec(query string, args ...interface{}) (sql.Result, error) {
	if dao.mode | DEV > 0 {
		log.Println(query, args)
	}
	return dao.db.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
func (dao *Query) Query(query string, args ...interface{}) ([]map[string]string, error) {
	// # 执行 SQL
	rows, err := dao.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// # 获取通用的数据结构
	// ## 获取结果集中字段的信息
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnLen := len(columns) // 字段数量
	// 依据 colnums 的数量，确定 Scan()中需要的变量数量
	fields := []interface{}{}
	for i := 0; i < columnLen; i++ {
		fields = append(fields, &sql.NullString{})
	}
	// ## 全部记录的变量
	records := []map[string]string{}
	// ## 遍历结果而来
	for rows.Next() {
		err = rows.Scan(fields...)
		if err != nil {
			continue
		}
		// 构造每条记录的map
		record := map[string]string{}
		// 处理该记录的每个字段
		for i := 0; i < columnLen; i++ {
			// 从Scan的结果中，获取每个字段，断言为*sql.NullString
			stringField := fields[i].(*sql.NullString)
			// 处理是否为 NULL
			if stringField.Valid {
				record[columns[i]] = stringField.String
			} else {
				record[columns[i]] = "NULL(NULL)"
			}
		}
		records = append(records, record)
	}

	// # 返回
	return records, nil
}

func (dao *Query) Insert(value map[string]interface{}) {
}

func (dao *Query) InsertOrIgnore() {
}

func (dao *Query) InsertGetId() {
}

func (dao *Query) Update() {
}

func (dao *Query) UpdateOrInsert() {
}

func (dao *Query) Increment() {
}

func (dao *Query) Decrement() {
}

func (dao *Query) Delete() {
}

func (dao *Query) Truncate() {
}

func (dao *Query) Replace() {
}

func (dao *Query) Get() {

}

func (dao *Query) First() {

}

func (dao *Query) Value() {

}

func (dao *Query) Find() {

}

func (dao *Query) Pluck() {

}

func (dao *Query) Chunk() {

}
func (dao *Query) Count() {
}

func (dao *Query) Sum() {
}

func (dao *Query) Avg() {
}

func (dao *Query) Min() {
}

func (dao *Query) Max() {
}

func (dao *Query) GroupConcat() {
}

func (dao *Query) Exists() {
}

func (dao *Query) DoesntExists() {
}

func (dao *Query) LastQuery() string {
	return dao.lastQuery
}

func (dao *Query) Dump() {
}
