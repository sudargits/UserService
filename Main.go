package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/logger"
)

func main()  {
	fmt.Print("USER MANAGEMENT START")
	api  := iris.New()
	api.Use(logger.New())

	api.Get("/", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK,iris.Map{"status":true})
	})
	user := api.Party("/user")
	user.Get("/",getMyProfile)
	user.Get("/:id",getUser)
	user.Post("/login",login)
	user.Post("/forgot",forgot)
	user.Post("/register",register)
	user.Post("/update/password",updatePassword)
	user.Post("/update/profile",updateProfile)

	admin := api.Party("/admin")
	admin.Post("/create",createUser)
	admin.Get("/list/:page",listUser)
	admin.Get("/:id",getUserAdmin)
	admin.Post("/delete",deleteUser)
	admin.Post("/update/field",updateField)



	api.Listen(":8080")
}

