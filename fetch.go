package main

import (
	"blog/internal/database"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"objects"
	"productfetch"
	"strings"
	"sync"
	"time"
	"utils"

	"github.com/google/uuid"
)

const fetch_time_xml = 60 * time.Second     // 60 seconds
const fetch_time_woo = 50 * time.Second     // 60 seconds
const fetch_time_shopify = 25 * time.Second // 60 seconds

const n_feeds_to_fetch = 10 // number of feeds to fetch from the database

// initiates the worker to fetch data
func FetchWorker(
	dbconfig dbConfig,
	shopConf productfetch.ConfigShopify,
	wooConf productfetch.ConfigWoo) {
	go LoopXML(dbconfig, fetch_time_xml)
	if shopConf.Url != "" {
		go LoopJSONShopify(dbconfig, shopConf, fetch_time_shopify)
	}
	if wooConf.Url != "" {
		go LoopJSONWoo(dbconfig, wooConf, fetch_time_woo)
	}
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
func LoopJSONShopify(
	dbconfig dbConfig,
	shopifyConfig productfetch.ConfigShopify,
	interval time.Duration) {
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		shopifyProds, err := shopifyConfig.FetchProducts()
		if err != nil {
			log.Println("Shopify > Error fetching next products to process:", err)
			continue
		}
		ProcessShopifyProducts(dbconfig, shopifyProds)
	}
}

// adds shopify products to the database
func ProcessShopifyProducts(dbconfig dbConfig, products objects.ShopifyProducts) {
	for _, value := range products.Products {
		for _, sub_value := range value.Variants {
			_, err := dbconfig.DB.CreateShopifyProduct(context.Background(), database.CreateShopifyProductParams{
				ID:        uuid.New(),
				Title:     value.Title,
				Sku:       sub_value.Sku,
				Price:     sub_value.Price,
				Qty:       int32(sub_value.InventoryQuantity),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			})
			if err != nil {
				if strings.ReplaceAll(err.Error(), "\"", "") == "pq: duplicate key value violates unique constraint shopify_sku_key" {
					_, errUpdate := dbconfig.DB.UpdateShopifyProducts(
						context.Background(),
						database.UpdateShopifyProductsParams{
							Title:     value.Title,
							Price:     sub_value.Price,
							Qty:       int32(sub_value.InventoryQuantity),
							UpdatedAt: time.Now().UTC(),
							Sku:       sub_value.Sku,
						})
					if errUpdate != nil {
						fmt.Println(sub_value.Sku)
						log.Fatal("Error updating shopify product: ", errUpdate)
					}
				} else {
					log.Println("Error creating shopify product:", err)
				}
			}
		}
	}
	log.Printf("From Shopify %d products were collected", len(products.Products))
}

// loop function that uses Goroutine to run
// a function each interval
func LoopJSONWoo(
	dbconfig dbConfig,
	wooConfig productfetch.ConfigWoo,
	interval time.Duration) {
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		wooProds, err := wooConfig.FetchProducts()
		if err != nil {
			log.Println("WooCommerce > Error fetching next products to process:", err)
			continue
		}
		ProcessWooProducts(dbconfig, wooProds)
	}
}

// adds woocommerce products to the database
func ProcessWooProducts(dbconfig dbConfig, products objects.WooProducts) {
	for _, value := range products.Products {
		if len(value.Variants) == 0 {
			fmt.Println(value.Price)
			fmt.Println(value.Sku)
			_, err := dbconfig.DB.CreateWooProduct(context.Background(), database.CreateWooProductParams{
				ID:        uuid.New(),
				Title:     value.Title,
				Sku:       value.Sku,
				Price:     value.Price,
				Qty:       int32(value.StockQuantity),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			})
			if err != nil {
				if strings.ReplaceAll(err.Error(), "\"", "") == "pq: duplicate key value violates unique constraint woocommerce_sku_key" {
					_, errUpdate := dbconfig.DB.UpdateWooProducts(
						context.Background(),
						database.UpdateWooProductsParams{
							Title:     value.Title,
							Price:     value.Price,
							Qty:       int32(value.StockQuantity),
							UpdatedAt: time.Now().UTC(),
							Sku:       value.Sku,
						})
					if errUpdate != nil {
						log.Fatal("Error updating woo product: ", errUpdate)
					}
				} else {
					log.Println("Error creating woo product:", err)
				}
			}
		}
		for _, sub_value := range value.Variants {
			_, err := dbconfig.DB.CreateWooProduct(context.Background(), database.CreateWooProductParams{
				ID:        uuid.New(),
				Title:     value.Title,
				Sku:       sub_value.Sku,
				Price:     sub_value.Price,
				Qty:       int32(sub_value.StockQuantity),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			})
			if err != nil {
				if strings.ReplaceAll(err.Error(), "\"", "") == "pq: duplicate key value violates unique constraint woocommerce_sku_key" {
					_, errUpdate := dbconfig.DB.UpdateWooProducts(
						context.Background(),
						database.UpdateWooProductsParams{
							Title:     value.Title,
							Price:     sub_value.Price,
							Qty:       int32(sub_value.StockQuantity),
							UpdatedAt: time.Now().UTC(),
							Sku:       sub_value.Sku,
						})
					if errUpdate != nil {
						log.Fatal("Error updating product: ", errUpdate)
					}
				} else {
					log.Println("Error creating product:", err)
				}
			}
		}
	}
	log.Printf("From WooCommerce %d products were collected", len(products.Products))
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
