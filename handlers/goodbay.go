package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBay struct {
	l *log.Logger
}

func NewGoodBay(l *log.Logger) *GoodBay {
	return &GoodBay{l}
}

func (g *GoodBay) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("GoodBay World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "GoodBay %s", d)
}
