package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

type ServerOpts struct {
	Port int
	Host string
}

func StartServer(s *ServerOpts) {
	srv := &http.Server{Addr: fmt.Sprintf("%s:%d", s.Host, s.Port), Handler: Router}

	go func() {
		logrus.Infof("Unravelling Cloud Agent started listening on %s:%d", s.Host, s.Port)
		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				logrus.Infof("Server shutting down...")
			} else {
				logrus.Fatalf("Could not listen on %s:%d. %v\n", s.Host, s.Port, err)
			}
			logrus.Error(err)
			return
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	// Create a context with a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the Shutdown method to gracefully shut down the server.
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Error:", err)
	}
}

// Mount attaches a sub-router to the main router
func Mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
