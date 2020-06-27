package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kadekchrisna/openbook-items-api/client/elasticsearch"
	"github.com/kadekchrisna/openbook-utils-go/logger"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()
	logger.Info(fmt.Sprintf("Server start at %s", "8090"))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
