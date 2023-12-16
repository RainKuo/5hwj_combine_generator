package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type MysqlDriver struct {
	db *sql.DB
}

func ConnectDB() *MysqlDriver {
	db, err := sql.Open("mysql", "root:xiaopang12121@tcp(124.220.163.61:3306)/rpg_data")
	if err != nil {
		panic(err.Error())
	}
	md := &MysqlDriver{db: db}
	md.db.SetMaxOpenConns(1000)
	md.db.SetMaxIdleConns(5)
	return &MysqlDriver{db: db}
}

func (md *MysqlDriver) Disconnect() {
	md.db.Close()
}

func FormatSqlValues(values []string) string {
	newValues := []string{}
	for _, str := range values {
		newValues = append(newValues, "'"+str+"'")
	}
	return strings.Join(newValues, ",")
}

func (md *MysqlDriver) Exec(sql string) error {
	_, err := md.db.Exec(sql)
	if err != nil {
		fmt.Printf("Exec SQL failed,ERR: %s", err)
	}
	return err
}

func (md *MysqlDriver) Insert(tableName string, fields []string, values []string) error {
	if len(fields) != len(values) {
		return fmt.Errorf("fields len not equal to values")
	}
	sql := fmt.Sprintf("insert into %s (%s) values(%s)", tableName, strings.Join(fields, ","), FormatSqlValues(values))
	// fmt.Println(sql)
	return md.Exec(sql)
}

// Query md.Query("select * from table where id = ?", 10086) 问号占位符
func (md *MysqlDriver) Query(sqlstr string, args ...interface{}) []map[string]string {
	rows, err := md.db.Query(sqlstr, args...)
	if err != nil {
		fmt.Printf("Execute query error: %s", err)
		return nil
	}
	defer rows.Close()

	results := make([]map[string]string, 0)

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		row := make(map[string]string)
		values := make([]interface{}, len(columns))
		for i := range columns {
			values[i] = new(string)
		}
		rows.Scan(values...)
		for i, col := range columns {
			val := *(values[i].(*string))
			row[col] = val
		}
		results = append(results, row)

	}
	return results
}
