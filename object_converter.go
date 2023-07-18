package main

import (
	"blog/internal/database"
	"objects"
)

// converts a []database.Feed into []objects.ResponseBodyFeed
func DatabaseToFeeds(feeds []database.Feed) []objects.ResponseBodyFeed {
	response := []objects.ResponseBodyFeed{}
	for _, value := range feeds {
		response = append(response, objects.ResponseBodyFeed{
			ID:            value.ID,
			Name:          value.Name,
			Url:           value.Url,
			CreatedAt:     value.CreatedAt,
			UpdatedAt:     value.UpdatedAt,
			UserID:        value.UserID,
			LastFetchedAt: value.LastFetchedAt.Time,
		})
	}
	return response
}

// converts a database.Feed into objects.ResponseBodyFeed
func DatabaseToFeed(feed database.Feed) objects.ResponseBodyFeed {
	return objects.ResponseBodyFeed{
		ID:            feed.ID,
		Name:          feed.Name,
		Url:           feed.Url,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		UserID:        feed.UserID,
		LastFetchedAt: feed.LastFetchedAt.Time,
	}
}
