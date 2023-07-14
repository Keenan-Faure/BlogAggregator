package main

import (
	"blog/internal/database"
	"encoding/json"
	"errors"
	"net/http"
	"objects"
)

// valiates a user
func UserValidation(user objects.RequestBodyUser) error {
	if user.Name == "" {
		return errors.New("Empty Name not allowed")
	}
	return nil
}

// decodes the request body
func DecodeUserRequestBody(r *http.Request) (objects.RequestBodyUser, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyUser{}
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}

// converts a database user to an object
func DatabaseUserToObject(dbUser database.User) objects.ResponseBodyUser {
	return objects.ResponseBodyUser{
		ID:        dbUser.ID.String(),
		CreatedAt: dbUser.CreatedAt.String(),
		UpdatedAt: dbUser.UpdatedAt.String(),
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

// validates a feed
func FeedValidation(feed objects.RequestBodyFeed) error {
	if feed.Name == "" {
		return errors.New("Empty Name not allowed")
	}
	if feed.URL == "" {
		return errors.New("Empty URL not allowed")
	}
	return nil
}

// decodes the request body
func DecodeFeedRequestBody(r *http.Request) (objects.RequestBodyFeed, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyFeed{}
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}

// converts a database user to an object
func DatabaseFeedToObject(feed database.Feed) objects.ResponseBodyFeed {
	return objects.ResponseBodyFeed{
		ID:        feed.ID.String(),
		Name:      feed.Name,
		URL:       feed.Url,
		CreatedAt: feed.CreatedAt.String(),
		UpdatedAt: feed.UpdatedAt.String(),
		UserID:    feed.UserID.String(),
	}
}
