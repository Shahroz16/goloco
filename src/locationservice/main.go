package main

import (
	"context"
	"fmt"
	golocopb "github.com/goloco/src/locationservice/genproto"
	"github.com/goloco/src/locationservice/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	listenPort = "3550"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout
}

type LocationService struct {
	repo goloco_repo.Repository
}

func (LocationService) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (LocationService) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}

func (l LocationService) SaveLocation(ctx context.Context, request *golocopb.LocationRequest) (*golocopb.LocationResponse, error) {
	locationModel := goloco_repo.GetLocationModel(request, false)

	model, err := l.repo.CreateLocation(locationModel)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Internet error %v \n", err))
	}

	return model.GetLocationResponse(), nil
}

func (l LocationService) GetLocation(ctx context.Context, req *golocopb.GetLocationLocationRequest) (*golocopb.LocationResponse, error) {
	location, err := l.repo.GetLocation(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Location not found %v \n", err))
	}

	return location.GetLocationResponse(), nil
}

func (l LocationService) UpdateLocation(ctx context.Context, req *golocopb.LocationRequest) (*golocopb.LocationResponse, error) {
	locationModel := goloco_repo.GetLocationModel(req, true)
	updateResponse, err := l.repo.UpdateLocation(locationModel)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Internet error %v \n", err))
	}
	return updateResponse.GetLocationResponse(), nil
}

func (l LocationService) DeleteLocation(ctx context.Context, request *golocopb.DeleteLocationLocationRequest) (*golocopb.DeletedLocationId, error) {
	response, err := l.repo.DeleteLocation(request.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Internet error %v \n", err))
	}

	return &golocopb.DeletedLocationId{Id: response}, nil
}

func (l LocationService) GetAllLocations(context.Context, *golocopb.EmptyMessageRequest) (*golocopb.AllLocationsResponse, error) {
	locations, err := l.repo.GetLocations()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internet error while getting locations %v \n", err))
	}

	responseLocation := make([] *golocopb.Location, 0)
	for _, location := range locations {
		responseLocation = append(responseLocation, location.GetLocationResponse().Location)
	}

	return &golocopb.AllLocationsResponse{Location: responseLocation}, nil
}

func (l LocationService) GetAllLocationsStream(ctx *golocopb.EmptyMessageRequest, stream golocopb.LocationService_GetAllLocationsStreamServer) error {
	locations, err := l.repo.GetLocations()
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internet error while getting locations %v \n", err))
	}

	for _, location := range locations {
		if steamErr := stream.Send(location.GetLocationResponse()); steamErr != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while streaming %v \n", err))
		}
	}

	return nil
}

func main() {

	// setting port
	srvPort := listenPort
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}
	postgresPath := "172.17.0.2"

	//if os.Getenv("POSTGRES_ADDR") != "" {
	//	postgresPath = os.Getenv("POSTGRES_ADDR")
	//}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2)

	store := goloco_repo.GormStore{
		Host:     postgresPath,
		Port:     5432,
		User:     "postgres",
		Password: "",
		Dbname:   "goloco",
	}

	db := store.Connect()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", srvPort))
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	svc := &LocationService{
		repo: goloco_repo.NewRepo(db, *log),
	}

	healthpb.RegisterHealthServer(srv, svc)
	golocopb.RegisterLocationServiceServer(srv, svc)
	reflection.Register(srv)

	go func() {
		if errr := srv.Serve(listener); errr != nil {
			log.Fatalf("failed to serve: %v", errr)
		}

	}()

	<-sigs

}
