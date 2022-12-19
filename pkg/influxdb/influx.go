package influxdb

import (
	"fmt"
	"log"
	"time"
)

const (
	MyDB     = "max"
	username = ""
	password = ""
)

func Influx() (res []client.Result, err1 error) {
	fmt.Println("connect to influx .....")

	// 连接数据库
	conn, err := connInflux()
	if err != nil {
		log.Fatal("连接失败，err:", err)
		return
	}

	// 写数据
	err = WritesPoints(conn)
	if err != nil {
		log.Fatal("写入数据失败，err:", err)
		return
	}

	cmd := "select * from cpu_usage"
	res, err = QueryDB(conn, cmd)
	if err != nil {
		log.Fatal("查询数据失败，err:", err)
		return
	}
	return res,nil
}

// 连接数据库
func connInflux() (cli client.Client, err error) {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.34.11.33:8086",
		Username: username,
		Password: password,
	})

	return conn, err
}

//查询数据
func QueryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

//写数据
func WritesPoints(cli client.Client) error{
	//选择数据库
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal("client.NewBatchPoints,err:", err)
		return err
	}

	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   20.1,
		"system": 43.3,
		"user":   86.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal("client.NewPoint,err:", err)
	}
	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		log.Fatal("conn.Write,err:", err)
		return err
	}

	return nil
}