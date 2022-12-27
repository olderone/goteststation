package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	Db  *sql.DB
	err error
)

type Sale struct {
	Widget_id int
	Qty       int
	Street    string
	City      string
	State     string
	Zip       int
	Sale_date string
}

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(10.140.254.29:31234)/test")
	if err != nil {
		fmt.Println("Open mysql error:", err)
		os.Exit(1)
	}
}

func main() {
	var (
		table = "sale"
	)
	sale := &Sale{
		Widget_id: 9, 
		Qty: 80, 
		Street: "Huanghe South Road", 
		City: "Anyang Henan", 
		State: "China", 
		Zip: 455000, 
		Sale_date: "2020-03-24"}

	/*ret, _ := sale.GetRecordById()
	if ret != nil {
		fmt.Println(*ret)
	}*/
	createSale := `CREATE TABLE IF NOT EXISTS` + table + `(
		Widget_id INT(10) NOT NULL AUTO_INCREMENT,
		qty INT(10) NOT NULL default 80,
		street VARCHAR(64) NOT NULL DEFAULT "",
		city VARCHAR(64) NOT NULL DEFAULT "",
		state VARCHAR(64) NOT NULL DEFAULT "",
		zip INT(10) NOT NULL default 0,
		created DATE NULL DEFAULT NULL,
		PRIMARY KEY(Widget_id)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	err := sale.createTable(createSale)
	if err != nil {
		log.Panic(err)
	}
	/*sale := &Sale{}

	ret2, _ := sale.GetAllRecord()
	if ret2 != nil {
		for _, record := range ret2 {
			fmt.Println(*record)
		}
	}*/

}

func (s *Sale)createTable(createSale string) error {

	smt, err := Db.Prepare(createSale)
	if err != nil || smt == nil {
		log.Println("err:", err)
		return err
	}
	_, err = smt.Exec()
	if err != nil {
		return err
	}
	return nil
}

//删除
func (s *Sale) DelRecordById() (err error) {
	sqlStr := "delete  from sales where widget_id = ?"
	inStmt, _ := Db.Prepare(sqlStr)
	res, err := inStmt.Exec(2)
	if err != nil {
		fmt.Println("Error del record :", err)
		return errors.New("Error del record ")
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Error del record :", err)
		return errors.New("Error del record ")
	}
	fmt.Println("del success", n)
	return nil
}


//更新
func (s *Sale) UpdateRecordById() (err error) {
	sqlStr := "update sales set qty = ? where widget_id = ?"
	inStmt, _ := Db.Prepare(sqlStr)
	res, err := inStmt.Exec(s.Qty, s.Widget_id)
	if err != nil {
		return errors.New("update error")
	}
	n, _ := res.RowsAffected()
	fmt.Println(n)
	return nil
}


func (s *Sale) GetAllRecord() (ret []*Sale, err_ error) {
	sqlStr := "select * from sales"
	inStmt, _ := Db.Prepare(sqlStr)

	rows, err_ := inStmt.Query()

	if err_ != nil {
		fmt.Println("Error get all: ", err_)
		return nil, err_
	}

	ret = make([]*Sale, 0)

	for rows.Next() {
		record := &Sale{}

		err_ = rows.Scan(&record.Widget_id, &record.Qty, &record.Street, &record.City, &record.State, &record.Zip, &record.Sale_date)

		if err_ != nil {
			fmt.Println("Error get record: ", err_)
			continue
		}

		ret = append(ret, record)
	}

	return ret, nil
}


func (s *Sale) GetRecordById() (ret *Sale, err_ error) {
	sqlStr := "select * from sales where widget_id = ?"
	inStmt, _ := Db.Prepare(sqlStr)

	row := inStmt.QueryRow(s.Widget_id)

	if row == nil {
		fmt.Println("No such record with id = ", s.Widget_id)
		return nil, errors.New("No such record with id = " + fmt.Sprintf("%d", s.Widget_id))
	}

	ret = &Sale{}

	err_ = row.Scan(&ret.Widget_id, &ret.Qty, &ret.Street, &ret.City, &ret.State, &ret.Zip, &ret.Sale_date)

	return ret, err_
}
