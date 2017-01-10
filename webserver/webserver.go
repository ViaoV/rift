package webserver

import (
	"fmt"
	"net/http"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("webserver")

// Start starts the client web server
func Start(port int32) {
	var (
		address string
		//	clientFS = http.Dir("/webclient")
		err error
	)
	address = fmt.Sprintf(":%d", port)
	http.Handle("/client", http.StripPrefix("/client", http.FileServer(http.Dir("./client"))))
	log.Infof("Starting webserver on port %d", port)
	err = http.ListenAndServe(address, nil)
	log.Fatal(err)
}
