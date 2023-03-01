package main

import (
	_ "github.com/ClickHouse/clickhouse-go"
	"log"
	"test-station/pkg/clickhouse"
)

func main()  {
	ch, err := clickhouse.NewChDB()
	if err != nil {
		log.Println(err)
	}
	err = ch.GetUsers()
	if err != nil {
		log.Println(err)
	}
}
