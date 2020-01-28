package main

import (
	"context"
	"fmt"
	pg "github.com/goloco/src/frontend/genproto"
	"github.com/pkg/errors"
	"time"
)

func (fe *frontendServer) getLocations(ctx context.Context) (*pg.AllLocationsResponse, error) {
	return pg.NewLocationServiceClient(fe.locationSvcConn).GetAllLocations(ctx, &pg.EmptyMessageRequest{})
}

func (fe *frontendServer) saveLocation(ctx context.Context, location *pg.Location) (*pg.LocationResponse, error) {
	return pg.NewLocationServiceClient(fe.locationSvcConn).SaveLocation(ctx, &pg.LocationRequest{Location: location})
}

func (fe *frontendServer) getRecommendation(ctx context.Context) (*pg.ListSuggestionsResponse, error) {
	allLocations, _ := fe.getLocations(ctx)
	lastIds := make([] string, 0)
	size := len(allLocations.Location)

	for i := size - 1; i >= size-5; i-- {
		lastIds = append(lastIds, fmt.Sprint((*allLocations.Location[i]).GetId()))
	}
	request := pg.ListSuggestionsRequest{UserId: "1", LocationIds: lastIds}
	return pg.NewSuggestionServiceClient(fe.suggestionSvcConn).ListSuggestions(ctx, &request)
}

func (fe *frontendServer) getAd(ctx context.Context, ctxKeys []string) ([]*pg.Ad, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	resp, err := pg.NewAdServiceClient(fe.adSvcConn).GetAds(ctx, &pg.AdRequest{
		ContextKeys: ctxKeys,
	})

	if err != nil {
		return nil, err
	}

	return resp.GetAds(), errors.Wrap(err, "failed to get ads")
}
