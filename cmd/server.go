package main

import "1008001/splitwiser/internal/store"

type Server struct {
	listenAddr string
	store      store.DB
}

func NewServer(listenAddr string, store store.DB) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}
