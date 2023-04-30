package server

import (
	"fmt"
	"github.com/b01/diablo-assistant/internal/cli"
	"github.com/b01/diablo-assistant/pkg/auth"
	"github.com/b01/diablo-assistant/pkg/d3"
	"github.com/kohirens/stdlib"
	"github.com/kohirens/stdlib/log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// HtmlPageLoader An http.HandleFunc function for responding to HTTP request with HTML.
func HtmlPageLoader(w http.ResponseWriter, r *http.Request) {
	var written int
	aFile := "website" + r.URL.Path

	if !stdlib.PathExist(aFile) {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, "<h1>Page Not Found</h1><p>The page %q could not be found</p>", r.URL.Path[1:])
		//return
	}

	battleNetClientId, ok := os.LookupEnv("BATTLE_NET_CLIENT_ID")
	if !ok {
		log.Fatf("")
		//return
	}
	state := fmt.Sprintf("%s_%d", time.Now().Format("220060102150105"), rand.Uint64())
	redirectUri := "https://localhost"
	log.Logf("here")

	loginUrl := fmt.Sprintf(auth.CodeLoginUrlFmt, battleNetClientId, d3.AuthScopes, state, redirectUri)

	written, err1 := fmt.Fprintf(w, "<a href=%q>Login</a>", loginUrl)
	if err1 != nil {
		log.Fatf(cli.Stderr.ResponseNotWritten, err1.Error())
		return
	}

	log.Dbugf(cli.Stderr.PageBytesWritten, written)
}

// Web This is a program blocking call. It runs the Go web server.
// Takes an endpoints is a pattern of endpoints ths server will cover, for
// example "/" or "api/". addr is an TCP IP:Port combination the server will
// listen on. While handler will respond to request sent to this server.
func Web(endpoints, addr string, handler http.HandlerFunc) error {
	// register an HTTP request handler
	http.HandleFunc(endpoints, handler)

	// run the web server
	return http.ListenAndServeTLS(addr, "certs/server.crt", "certs/server.key", nil)
}
