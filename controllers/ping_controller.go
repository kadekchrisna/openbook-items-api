package controllers

import "net/http"

const (
	PONG = "PONG"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct {
}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(PONG))
	return
}
