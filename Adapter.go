package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string
	Email string `json:"email"`
	Created_date string `json:"created_date"`
}
type MSUser struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Created_date string `json:"created_date"`
}

type MUser struct {
	Username string `valid:"alphanum,required"`
	Password string `valid:"required"`
	Email string `valid:"email"`
}
type MessageDevel struct {
	Devel string `json:"devel"`
	Prod string `json:"prod"`
}

type ConnectorDB struct {
	db *sql.DB
	err error
}
func connect() ConnectorDB  {
	db, err := sql.Open("mysql", "root:gitsgits@/service_user_management?charset=utf8")
	checkErr(err)
	fmt.Print("connected")
	return ConnectorDB{ db:db, err:err }
}

func checkErr(err error) {
	if err != nil {
		//panic(err)
		fmt.Print(err.Error())
	}
}

func insertUserDB(data User,db *sql.DB) (int64,error){
	defer db.Close()
	// insert
	stmt, err := db.Prepare("INSERT service_user SET username=?,password=md5(?),email=?")
	res, err := stmt.Exec(data.Username, data.Password, data.Email)
	checkErr(err)
	if err == nil {
		id,err := res.LastInsertId()
		checkErr(err)
		return id,err
	}else{
		return 0,err
	}


}
func updateUserDB(data User,db *sql.DB) (int64,error){
	defer db.Close()
	// update
	stmt, err := db.Prepare("UPDATE service_user SET username=?,email=? where id=?")
	res, err := stmt.Exec(data.Username, data.Email, data.Id)
	checkErr(err)
	if err == nil {
		id, err := res.RowsAffected()
		checkErr(err)
		return id,err
	}else{
		return 0,err
	}
}
func getUserFromDB(data User,db *sql.DB) (MSUser,error){
	defer db.Close()
	var result  =  MSUser{}
	var err = db.QueryRow("select id,username,email,created_date from service_user where id = ?",data.Id).
		Scan(&result.Id,&result.Username,&result.Email,&result.Created_date)
	if err != nil {
		return result,err
	}else{
		return result,nil
	}
}