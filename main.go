package main

import "hara-depo-proj/app"

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":80")
}
