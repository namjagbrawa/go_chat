package main

import "sync"

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool
}

type muxEntry struct {
	explicit bool
	h        Handler
	pattern  string
}

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}