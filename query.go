package goquery

import (
	"database/sql"
	"log"
)

type Query struct {
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
		mode: defaultMode,
	}, nil
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (q *Query) Ping() error {
	return q.db.Ping()
}

// Exec executes a query without returning any rows.
func (q *Query) Exec(query string, args ...interface{}) (sql.Result, error) {
	if q.mode | DEV > 0 {
		log.Println(query, args)
	}
	return q.db.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
func (q *Query) Query(query string, args ...interface{}) ([]map[string]string, error) {
	// # 执行 SQL
	rows, err := q.db.Query(query, args...)
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

func (q *Query) Insert(value map[string]interface{}) {
}

func (q *Query) InsertOrIgnore() {
}

func (q *Query) InsertGetId() {
}

func (q *Query) Update() {
}

func (q *Query) UpdateOrInsert() {
}

func (q *Query) Increment() {
}

func (q *Query) Decrement() {
}

func (q *Query) Delete() {
}

func (q *Query) Truncate() {
}

func (q *Query) Replace() {
}

func (q *Query) Get() {

}

func (q *Query) First() {

}

func (q *Query) Value() {

}

func (q *Query) Find() {

}

func (q *Query) Pluck() {

}

func (q *Query) Chunk() {

}
func (q *Query) Count() {
}

func (q *Query) Sum() {
}

func (q *Query) Avg() {
}

func (q *Query) Min() {
}

func (q *Query) Max() {
}

func (q *Query) GroupConcat() {
}

func (q *Query) Exists() {
}

func (q *Query) DoesntExists() {
}

func (q *Query) LastQuery() string {
	return q.lastQuery
}

func (q *Query) Dump() {
}
