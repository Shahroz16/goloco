package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	golocopb "github.com/goloco/src/frontend/genproto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"math/rand"
	"net/http"
)

var (
	templates, _ = template.New("").ParseGlob("templates/*.html")
)

func (fe *frontendServer) dummyHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	_, err := w.Write([]byte("Helllo Word"))
	if err != nil {
		log.Error(err)
	}
}

func (fe *frontendServer) addLocation(w http.ResponseWriter, r *http.Request) {
	_ = r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	decoder := json.NewDecoder(r.Body)
	location := &golocopb.Location{}
	err := jsonpb.UnmarshalNext(decoder, location)
	if err != nil {
		fmt.Sprintf("Json could not be unmarshaled. Error: %s", err.Error())
		return
	}
	response, saveErr := fe.saveLocation(r.Context(), location)
	if saveErr != nil {
		fmt.Sprintf("Json could not be unmarshaled. Error: %s", saveErr.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(&response.Location)
}

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	locations, err := fe.getLocations(r.Context())
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve locations"), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
		"locations":  locations.Location,
		//"ad":         fe.chooseAd(r.Context(), []string{}, log),
	}); err != nil {
		log.Error(err)
	}

}

func (fe *frontendServer) getSuggestions(writer http.ResponseWriter, request *http.Request) {
	log := request.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	suggestedLocation, err := fe.getRecommendation(request.Context())
	if err != nil {
		renderHTTPError(log, request, writer, errors.Wrap(err, "could not retrieve suggestions"),
			http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(writer).Encode(suggestedLocation.LocationIds)
}

func (fe *frontendServer) chooseAd(ctx context.Context, ctxKeys []string, log logrus.FieldLogger) *golocopb.Ad {
	ads, err := fe.getAd(ctx, ctxKeys)
	if err != nil {
		log.WithField("error", err).Warn("failed to retrieve ads")
		return nil
	}
	return ads[rand.Intn(len(ads))]
}

func renderHTTPError(log logrus.FieldLogger, r *http.Request, w http.ResponseWriter, err error, code int) {
	log.WithField("error", err).Error("request error")
	errMsg := fmt.Sprintf("%+v", err)

	w.WriteHeader(code)

	templates.ExecuteTemplate(w, "error", map[string]interface{}{
		"session_id":  sessionID(r),
		"request_id":  r.Context().Value(ctxKeyRequestID{}),
		"error":       errMsg,
		"status_code": code,
		"status":      http.StatusText(code)})
}

func sessionID(r *http.Request) string {
	v := r.Context().Value(ctxKeySessionID{})
	if v != nil {
		return v.(string)
	}
	return ""
}
