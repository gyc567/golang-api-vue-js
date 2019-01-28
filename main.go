package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
)

func main() {
	// To show API docs.
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

