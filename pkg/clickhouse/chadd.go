package clickhouse

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"strconv"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

func (c *Ch)AddUser() error  {

	defer func() {
		_ = c.db.Close()
	}()

	//数据预处理写入
	tx, err := c.db.Begin()
	if err != nil {
		fmt.Printf("clickhouse connect err %s", err.Error())
		return err
	}
	stmt, err := tx.Prepare("insert into max.user (id,name,age) values(?,?,?)")
	if err != nil {
		fmt.Printf("clickhouse tx err %s", err.Error())
		return err
	}
	for i := 10; i < 20; i++ {
		if _, err := stmt.Exec(i, "n"+strconv.Itoa(i), i+10); err != nil {
			fmt.Printf("clickhouse stmt err %s", err.Error())
			return err
		}
	}
	_ = tx.Commit()
	return nil
}