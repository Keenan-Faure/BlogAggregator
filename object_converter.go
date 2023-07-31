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

// converts a []database.Posts into objects.ResponseBodyPosts
func DatabaseToPosts(feeds []database.Post) []objects.ResponseBodyPosts {
	response := []objects.ResponseBodyPosts{}
	for _, value := range feeds {
		response = append(response, objects.ResponseBodyPosts{
			ID:          value.ID,
			FeedID:      value.FeedID,
			Url:         value.Url,
			Title:       value.Title,
			Description: value.Description.String,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			PublishedAt: value.PublishedAt,
		})
	}
	return response
}

// converts a []database.Bookmark into objects.ResponseBodyBookmark
func DatabaseToBookmark(bookmarks []database.Bookmark) []objects.ResponseBodyBookmark {
	response := []objects.ResponseBodyBookmark{}
	for _, value := range bookmarks {
		response = append(response, objects.ResponseBodyBookmark{
			ID:        value.ID,
			PostID:    value.PostID,
			UserID:    value.UserID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}
	return response
}

// converts a []database.Liked into objects.ResponseBodyLiked
func DatabaseToLiked(liked []database.Liked) []objects.ResponseBodyLiked {
	response := []objects.ResponseBodyLiked{}
	for _, value := range liked {
		response = append(response, objects.ResponseBodyLiked{
			ID:        value.ID,
			PostID:    value.PostID,
			UserID:    value.UserID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}
	return response
}
