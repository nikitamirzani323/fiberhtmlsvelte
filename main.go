package main

import (
	"log"

	"github.com/nikitamirzani323/fiberhtmlsvelte/routers"
)

func main() {

	app := routers.Init()
	log.Fatal(app.Listen(":7075"))
}
