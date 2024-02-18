package web

import (
	"DADS/internal/hlogger"
	"fmt"
	"net/http"
)

func Run() {
	http.HandleFunc("/", sroot)
	http.Handle("/testhandle", newHandler())
	http.HandleFunc("/testquery", queryTestHandler)
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServe(":8080", newHandler())

	// server := &http.Server{
	// 	Addr:			":8080",
	// 	Handler: 		newHandler(),
	// 	ReadTimeout: 	5 * time.Second,
	// 	WriteTimeout: 	5 * time.Second,
	// }
	// server.ListenAndServe()
}

func queryTestHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	message := fmt.Sprintf("Query map: %v §n", q)
	// key1=2&key2=3
	v1, v2 := q.Get("key1"), q.Get("key2")
	if v1 == v2 {
		message = message + fmt.Sprintf("v1 and v2 are equal %s §n", v1)
	} else {
		message = message + fmt.Sprintf("v1 is equal to %s, v2 is equal to %s", v1, v2)
	}
	
	fmt.Fprint(w, message)
}



func sroot(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter is a I/O reader object
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to the Hydra software system")
	logger.Println("Received an http get request on root url")
}

