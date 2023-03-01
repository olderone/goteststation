package clickhouse

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	SCHEME="tcp://"
	IP="10.34.11.32"
	PORT="9000"
	USER = "default"
	PASSWORD=""
	DATABASE="max"
)

type Ch struct {
	db 			*sqlx.DB
	Scheme		string
	Ip    		string
	Port   		string
	User 		string
	Pass 		string
	Database 	string
}

func NewChDB() (ch Ch, err error) {
	ch.Scheme = SCHEME
	ch.Ip = IP
	ch.Port=PORT
	ch.User = USER
	ch.Pass = PASSWORD
	ch.Database = DATABASE
	// if does not has secret
	source := fmt.Sprintf("%s%s:%s?debug=true&username=%s&datbase=%s",
		ch.Scheme, ch.Ip, ch.Port, ch.User, ch.Database )
	// if has secret
	//source := "tcp://10.34.11.32:9000?debug=true&username=default&password=''&database=max"
	ch.db, err = sqlx.Connect("clickhouse", source)
	if err != nil {
		fmt.Printf("clickhouse open err %s", err.Error())
		return ch, err
	}
	return ch, nil
}