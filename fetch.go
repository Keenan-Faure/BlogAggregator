package main

import (
	"blog/internal/database"
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"objects"
	"strings"
	"sync"
	"time"
	"utils"

	"github.com/google/uuid"
)

const fetch_time_xml = 60 * time.Second  // 60 seconds
const fetch_time_json = 45 * time.Second // 60 seconds
const n_feeds_to_fetch = 10              // number of feeds to fetch from the database

// initiates the worker to fetch data
func FetchWorker(dbconfig dbConfig) {
	go LoopXML(dbconfig, fetch_time_xml)
}

// fetches feed(s) from a url
func FetchFeed(url string) (objects.RSS, error) {
	// if !utils.CheckStringWithWord(url, ".xml") {
	// 	return objects.RSS{}, errors.New("unable to parse non-xml feed")
	// }
	resp, err := http.Get(url)
	if err != nil {
		return objects.RSS{}, err
	}
	defer resp.Body.Close()
	body, err_ := io.ReadAll(resp.Body)
	if err_ != nil {
		return objects.RSS{}, err
	}
	result := objects.RSS{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return objects.RSS{}, err
	}
	return result, nil
}

// processes the feeds (loop)
func ProcessFeeds(dbconfig dbConfig, feeds []database.Feed) {
	wg := &sync.WaitGroup{}
	for _, value := range feeds {
		wg.Add(1)
		go process_feed(dbconfig, value, wg)
	}
	wg.Wait()
}

// processes feed (single)
func process_feed(dbconfig dbConfig, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := dbconfig.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: utils.ConvertTimeSQL(time.Now().UTC()),
		ID:            feed.ID,
	})
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}
	rssfeed, err := FetchFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}
	for _, value := range rssfeed.Channel.Item {
		_, err := dbconfig.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       value.Title,
			Url:         value.Link,
			Description: utils.ConvertStringToSQL(value.Description),
			PublishedAt: utils.ConvertStringToTime(value.PubDate),
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.ReplaceAll(err.Error(), "\"", "") != "pq: duplicate key value violates unique constraint posts_url_key" {
				log.Println("Error creating post:", err)
			}
		}
	}
	log.Printf("Feed %s collected, %d posts found", feed.Name, len(rssfeed.Channel.Item))
}

// loop function that uses Goroutine to run
// a function each interval
func LoopJSON(dbconfig dbConfig, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		unprocessedFeeds, err := dbconfig.DB.GetNextFeedsToFetch(context.Background(), n_feeds_to_fetch)
		if err != nil {
			log.Println("Error fetching next feeds to process:", err)
			continue
		}
		ProcessFeeds(dbconfig, unprocessedFeeds)
	}
}

// loop function that uses Goroutine to run
// a function each interval
func LoopXML(dbconfig dbConfig, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		unprocessedFeeds, err := dbconfig.DB.GetNextFeedsToFetch(context.Background(), n_feeds_to_fetch)
		if err != nil {
			log.Println("Error fetching next feeds to process:", err)
			continue
		}
		ProcessFeeds(dbconfig, unprocessedFeeds)
	}
}
