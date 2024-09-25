package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World!"})
	})

	app.Listen(":8080")
}
