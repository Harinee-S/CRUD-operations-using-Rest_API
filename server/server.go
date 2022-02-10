package server

import (
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"Project2/log"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func recoverFromPanic() negroni.HandlerFunc {
	return negroni.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
					log.Log.HTTPErrorf(r, "Recovered from panic: %v", err)
					return
				}
			}()
			next(rw, r)
		})
}

func listenAndServe(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != nil {
		log.Log.Fatalf("failed to start web router: %v", err)
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	log.Log.Info("web server shutting down")
	// Finish all API calls being served and shutdown gracefully
	apiServer.Shutdown(context.Background())
	log.Log.Info("web server shutting down")
}

func httpStatLogger() negroni.HandlerFunc {
	return negroni.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			startTime := time.Now()
			next(rw, r)

			response := rw.(negroni.ResponseWriter)
			responseTime := time.Now()
			latency := responseTime.Sub(startTime)

			log.Log.WithFields(
				logrus.Fields{
					"RequestTime":    startTime.Format(time.RFC3339),
					"ResponseTime":   responseTime.Format(time.RFC3339),
					"Latency":        latency,
					"RequestURL":     r.URL.Path,
					"RequestMethod":  r.Method,
					"ResponseStatus": response.Status(),
					"RequestProxy":   r.RemoteAddr,
					"RequestSource":  r.Header.Get("X_FORWARDED-FOR"),
				}).Info("Http Logs")
		})
}

// StartServer - Starts the web server using a fully qualified mux router
func StartServer(router *mux.Router) {
	log.Log.Infof("starting %s ... on port %s", "go-sample", "8080")

	//if config.GetSwaggerEnabled() && strings.ToUpper(config.GetAppEnvironment()) != "PRODUCTION" {
	//	router.PathPrefix(config.GetSwaggerDocsDirectory()).Handler(httpSwagger.WrapHandler)
	//}

	handlerFunc := router.ServeHTTP

	n := negroni.New()
	n.Use(httpStatLogger())
	n.Use(recoverFromPanic())
	n.UseHandlerFunc(handlerFunc)

	portInfo := ":" + strconv.Itoa(8080)
	server := &http.Server{Addr: portInfo, Handler: n}

	go listenAndServe(server)
	waitForShutdown(server)
}
