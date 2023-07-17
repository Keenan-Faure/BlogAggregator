package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"objects"
)

// User: validation
func UserValidation(user objects.RequestBodyUser) error {
	if user.Name == "" {
		return errors.New("empty name not allowed")
	}
	return nil
}

// User: decodes the request body
func DecodeUserRequestBody(r *http.Request) (objects.RequestBodyUser, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyUser{}
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}

// Feed: validation
func FeedValidation(feed objects.RequestBodyFeed) error {
	if feed.Name == "" {
		return errors.New("empty name not allowed")
	}
	if feed.URL == "" {
		return errors.New("empty URL not allowed")
	}
	return nil
}

// Feed: decodes the request body
func DecodeFeedRequestBody(r *http.Request) (objects.RequestBodyFeed, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyFeed{}
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}

// FeedFollow: decodes the request body
func DecodeFeedFollowRequestBody(r *http.Request) (objects.RequestBodyFeedFollow, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyFeedFollow{}
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}

// FeedFollow: validation
func FeedFollowValidation(feed objects.RequestBodyFeedFollow) error {
	if feed.FeedID == "" {
		return errors.New("empty feed_id not allowed")
	}
	return nil
}
