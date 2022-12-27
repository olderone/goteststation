package main

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

func main()  {
	//db := clickhouse.GetDb()
	//defer db.Close()
	source := "tcp://127.0.0.1:9000?debug=true&username=admin&password=topsec!518&database=test"
	connect, err := sqlx.Connect("clickhouse", source)
	if err != nil {
		fmt.Printf("clickhouse open err %s", err.Error())
		return
	}
	defer func() {
		_ = connect.Close()
	}()
	//数据预处理写入
	/*tx, err := connect.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	stmt, err := tx.Prepare("insert into test.hi_test_user (id,name,age) values(?,?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 10; i < 20; i++ {
		if _, err := stmt.Exec(i, "n"+strconv.Itoa(i), i+10); err != nil {
			fmt.Println(err)
		}
	}
	_ = tx.Commit()*/
	var items []User
	if err := connect.Select(&items, "select * from test.hi_test_user limit 10"); err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Printf("id: %d, name: %s, age: %d\n", item.Id, item.Name, item.Age)
	}
}
