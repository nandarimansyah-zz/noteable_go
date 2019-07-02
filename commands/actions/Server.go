package actions

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var tmpl = `
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
  Noteable-app using GO
 ___________________________
  - Port      : %s
  - Env       : %s
  - Debug     : %s
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
`

// ServerListen the app server
func ServerListen(handler http.Handler) {

	fmt.Println(fmt.Sprintf(tmpl,
		os.Getenv("PORT"),
		os.Getenv("env"),
		os.Getenv("debug"),
	))

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	log.Error(err)

}
