package main

import (
	"github.com/b01/diablo-assistant/pkg/server"
	"github.com/kohirens/stdlib/log"
	"os"
)

const authFile = "auth.json"

func main() {
	var mainErr error

	defer func() { // Exit; logging any errors beforehand.
		if mainErr != nil {
			log.Fatf("%v", mainErr.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}()

	mainErr = server.Web("/", ":443", server.HtmlPageLoader)

	return

}
