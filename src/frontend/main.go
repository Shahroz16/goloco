package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"time"
)

const (
	port            = "8080"
	cookieMaxAge    = 60 * 60 * 48
	cookiePrefix    = "goloco_"
	cookieSessionID = cookiePrefix + "session-id"
)

const (
	AD_SERVICE_ADDR             = "AD_SERVICE_ADDR"
	LOCATION_SERVICE_ADDR       = "LOCATION_SERVICE_ADDR"
	RECOMMENDATION_SERVICE_ADDR = "RECOMMENDATION_SERVICE_ADDR"
	SEARCH_SERVICE_ADDR         = "SEARCH_SERVICE_ADDR"
)

type frontendServer struct {
	locationSvcAddr string
	locationSvcConn *grpc.ClientConn

	suggestionSvcAddr string
	suggestionSvcConn *grpc.ClientConn

	adSvcAddr string
	adSvcConn *grpc.ClientConn

	searchSvcAddr string
	searchSvcConn *grpc.ClientConn
}

type ctxKeySessionID struct{}

func main() {
	// setting logging
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	}
	log.Out = os.Stdout

	// setting port
	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}
	addr := os.Getenv("LISTEN_ADDR")

	svc := new(frontendServer)

	log.Infof(os.Getenv(LOCATION_SERVICE_ADDR))

	// setting address of each service
	//mapAddressToEnv(&svc.adSvcAddr, AD_SERVICE_ADDR)
	mapAddressToEnv(&svc.locationSvcAddr, LOCATION_SERVICE_ADDR)
	//mapAddressToEnv(&svc.searchSvcAddr, SEARCH_SERVICE_ADDR)
	mapAddressToEnv(&svc.suggestionSvcAddr, RECOMMENDATION_SERVICE_ADDR)

	// making a connection with each service
	mapConnection(ctx, svc.locationSvcAddr, &svc.locationSvcConn)
	//mapConnection(ctx, svc.adSvcAddr, &svc.adSvcConn)
	//mapConnection(ctx, svc.searchSvcAddr, &svc.searchSvcConn)
	mapConnection(ctx, svc.suggestionSvcAddr, &svc.suggestionSvcConn)

	r := mux.NewRouter()
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/add_location", svc.addLocation).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc("/get_suggestion", svc.getSuggestions).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })

	var handler http.Handler = r
	handler = &logHandler{log: log, next: handler} // add logging
	handler = ensureSessionID(handler)             // add session ID

	log.Infof("starting server on " + addr + ":" + srvPort)
	log.Fatal(http.ListenAndServe(addr+":"+srvPort, handler))

}

func mapAddressToEnv(serviceAddress *string, envKey string) {
	if address := os.Getenv(envKey); address != "" {
		*serviceAddress = address
	} else {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}

}

func mapConnection(ctx context.Context, addr string, conn **grpc.ClientConn) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("grpc: failed to connect %s", addr))
	}
}
