package main

import (
	"github.com/kataras/iris"
	"gopkg.in/asaskevich/govalidator.v4"
	"fmt"
	"strings"
	"strconv"
)

func getUser(ctx *iris.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:"failed",Prod:"Id must be integer"}})
	}else{
		user,err := getUserFromDB(User{Id:id},connect().db)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:err.Error(),Prod:err.Error()}})
		}else{
			ctx.JSON(iris.StatusOK,iris.Map{"status":true,"message":MessageDevel{Devel:"success",Prod:"Success"},"content":user})
		}
	}
}
func getMyProfile(ctx *iris.Context) {

}
func register(ctx *iris.Context) {
	var name = ctx.FormValue("username")
	var email = ctx.FormValue("email")
	var password = ctx.FormValue("password")

	dUser :=&MUser{Email:email, Username:name,Password:password}
	result, err := govalidator.ValidateStruct(dUser)

	if err != nil {
		mv := strings.Split(err.Error(),";")
		ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:err.Error(),Prod:"Error Validation"},"content":mv})
	}else{
		fmt.Print(result)
		var db = connect()
		id, err := insertUserDB(User{Username:name,Email:email,Password:password},db.db)
		if err != nil {
			if strings.Contains(err.Error(),"1062") {
				var message = ""
				if strings.Contains(err.Error(),"username"){
					message = "Duplicate username"
				}else{
					message = "Duplicate email"
				}

				ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:err.Error(),Prod:message}})
			}else{
				ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:err.Error(),Prod:err.Error()}})
			}
		}else{
			if id > 0 {
				ctx.JSON(iris.StatusOK,iris.Map{"status":true,"message":MessageDevel{Devel:"success",Prod:"Success"}})
			}else{
				ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:"failed",Prod:"Failed"}})
			}
		}
	}
}
func updateProfile(ctx *iris.Context) {
	var name = ctx.FormValue("name")
	var email = ctx.FormValue("email")
	i, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:"failed",Prod:"Id must be integer"}})
	}else{
		var db = connect()
		var dataUser = User{Id:i,Username:name,Email:email}
		id,err := updateUserDB(dataUser,db.db)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:err.Error(),Prod:err.Error()}})
		}else{
			if id > 0 {
				ctx.JSON(iris.StatusOK,iris.Map{"status":true,"message":MessageDevel{Devel:"success",Prod:"Success"}})
			}else{
				ctx.JSON(iris.StatusBadRequest,iris.Map{"status":false,"message":MessageDevel{Devel:"no data change",Prod:"no data change"}})
			}
		}
	}

}

func login(ctx *iris.Context) {

}

func forgot(ctx *iris.Context) {

}
func updatePassword(ctx *iris.Context) {

}


