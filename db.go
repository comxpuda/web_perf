package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"syreclabs.com/go/faker"
	"time"
)

type Doctor struct {
	ID      int64
	Name    string
	Age     int
	Sex     int
	AddTime time.Time
}

func insert() int64 {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.202:3306)/sbtest?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("db connect error", err)
		return -1
	}

	defer db.Close()

	result, err := db.Exec("insert into doctor_tb(name,age,sex,addTime) values(?,?,?,Now())", faker.Name().Name(), faker.Number().NumberInt(2), faker.Number().Positive(3))
	if err != nil {
		fmt.Println("db insert error", err)
		return -1
	}
	newID, _ := result.LastInsertId()
	i, _ := result.RowsAffected()
	fmt.Printf("new ID：%d , affected rows：%d \n", newID, i)
	return newID
}

func query() []Doctor {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.202:3306)/sbtest?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("db connect error", err)
		return []Doctor{}
	}
	defer db.Close()

	rows, err := db.Query("select * from doctor_tb")
	if err != nil {
		fmt.Println("db query error", err)
		return []Doctor{}
	}
	var docList []Doctor
	for rows.Next() {
		var doc Doctor
		rows.Scan(&doc.ID, &doc.Name, &doc.Age, &doc.Sex, &doc.AddTime)
		docList = append(docList, doc)
	}
	return docList
}
