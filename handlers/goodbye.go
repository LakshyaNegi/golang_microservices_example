package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Bye"))

	//g.l.Println("Goodbye World")
	// d, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Oops", http.StatusBadRequest)
	// 	return
	// }
	// log.Printf("Data %s\n", d)

	// fmt.Fprintf(rw, "Hello %s", d)
}
