package clickhouse

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"net"
	"time"
)

func GetConn() (conn driver.Conn, err error) {
	conn, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{"10.34.11.32:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "admin",
			Password: "topsec!518",
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			//dialCount++
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: true,
		Debugf: func(format string, v ...interface{}) {
			fmt.Printf(format, v)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:      time.Duration(10) * time.Second,
		MaxOpenConns:     5,
		MaxIdleConns:     5,
		ConnMaxLifetime:  time.Duration(10) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		BlockBufferSize: 10,
		//MaxCompressionBuffer: 10240,
	})
	if err != nil {
		return conn, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return conn, err
	}
	return conn, nil

}

func GetDb() *sql.DB {
	db := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"10.34.11.32:9000"},
		Auth: clickhouse.Auth{
			Database: "test",
			Username: "admin",
			Password: "topsec!518",
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: true,
		BlockBufferSize: 10,
		//MaxCompressionBuffer: 10240,
	})
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return db
}