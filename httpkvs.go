package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type value struct {
	contentType string
	content     []byte
}

type data struct {
	sync.Mutex
	kvs map[string]value
}

func get(d *data, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	d.Lock()
	val, ok := d.kvs[ps.ByName("key")]
	d.Unlock()
	if ok {
		if len(val.contentType) > 0 {
			w.Header().Set("Content-Type", val.contentType)
		}
		w.Write(val.content)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Key '%s' not found\n", ps.ByName("key"))
	}
}

func put(d *data, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var buf []byte = make([]byte, r.ContentLength)
	io.ReadFull(r.Body, buf)
	contentType := r.Header.Get("Content-Type")
	d.Lock()
	d.kvs[ps.ByName("key")] = value{contentType: contentType, content: buf}
	d.Unlock()
	fmt.Fprintf(w, "Key '%s' saved\n", ps.ByName("key"))
}

func main() {
	router := httprouter.New()
	d := new(data)
	d.kvs = make(map[string]value)

	router.GET("/:key", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		get(d, w, r, ps)
	})
	router.PUT("/:key", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		put(d, w, r, ps)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
