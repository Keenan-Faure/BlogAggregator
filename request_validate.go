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

// Liked: decodes the request body
func DecodeLikedRequestBody(r *http.Request) (objects.RequestBodyLiked, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyLiked{}
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}

// Bookmark: decodes the request body
func DecodeBookmarkRequestBody(r *http.Request) (objects.RequestBodyBookmark, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyBookmark{}
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

// Liked: validation
func LikedValidation(post objects.RequestBodyLiked) error {
	if post.PostID == "" {
		return errors.New("empty feed_id not allowed")
	}
	return nil
}

// Bookmark: validation
func BookmarkValidation(post objects.RequestBodyBookmark) error {
	if post.PostID == "" {
		return errors.New("empty feed_id not allowed")
	}
	return nil
}

// Posts: decode the request body
func DecodePostRequestBody(r *http.Request) (objects.CreatePostParams, error) {
	decoder := json.NewDecoder(r.Body)
	params := objects.CreatePostParams{}
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}
