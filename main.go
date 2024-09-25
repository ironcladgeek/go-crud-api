package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World!"})
	})

	app.Post("/users", func(ctx iris.Context) {
		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		// Print the user data to the console
		fmt.Printf("Received user: %+v\n", user)

		// Return a success message
		ctx.JSON(iris.Map{"message": "User created successfully", "user": user})
	})

	app.Listen(":8080")
}
