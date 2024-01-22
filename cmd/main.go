package main

import "projectTemplate/app"

func main() {
	a := app.NewApp(":8080")
	a.Register()
	a.Start()
}
