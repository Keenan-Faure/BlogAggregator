package productfetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"objects"
	"time"
	"utils"
)

type ConfigWoo struct {
	Key    string
	Secret string
	Url    string
}

type ConfigShopify struct {
	APIKey      string
	APIPassword string
	Version     string
	Url         string
}

const PRODUCT_FETCH_LIMIT = "10" // limit on products to fetch

// Initiates the connection string for woocommerce
func InitConfigWoo() (ConfigWoo, error) {
	store_name := utils.LoadEnv("woo_store_name")
	api_key := utils.LoadEnv("woo_consumer_key")
	api_secret := utils.LoadEnv("woo_consumer_secret")

	validation := validateConfig(store_name, api_key, api_secret)
	if validation != nil {
		log.Fatal("Error setting up connection string for WooCommerce")
		return ConfigWoo{}, validation
	}
	return ConfigWoo{
		Key:    api_key,
		Secret: api_secret,
		Url:    "https://" + store_name + "/wc-api/v3/products?consumer_key=" + api_key + "&consumer_secret=" + api_secret,
	}, nil
}

// Fetches products from WooCommerce defined on the wooConfig url

func (wooConfig *ConfigWoo) FetchProducts() (objects.WooProducts, error) {
	httpClient := http.Client{
		Timeout: time.Second * 20,
	}
	req, err := http.NewRequest(http.MethodGet, wooConfig.Url+"&filter[limit]="+PRODUCT_FETCH_LIMIT, nil)
	if err != nil {
		log.Fatal(err)
		return objects.WooProducts{}, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return objects.WooProducts{}, err
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return objects.WooProducts{}, err
	}
	products := objects.WooProducts{}
	err = json.Unmarshal(respBody, &products)
	if err != nil {
		log.Fatal(err)
		return objects.WooProducts{}, err
	}
	return products, nil
}

// Initiates the connection string for shopify
func InitConfigShopify() (ConfigShopify, error) {
	store_name := utils.LoadEnv("store_name")
	api_key := utils.LoadEnv("api_key")
	api_password := utils.LoadEnv("api_password")
	version := utils.LoadEnv("version")

	validation := validateConfig(store_name, api_key, api_password)
	if validation != nil {
		log.Fatal("Error setting up connection string for Shopify")
		return ConfigShopify{}, validation
	}
	return ConfigShopify{
		APIKey:      api_key,
		APIPassword: api_password,
		Version:     version,
		Url:         "https://" + api_key + ":" + api_password + "@" + store_name + ".myshopify.com/admin/api/" + version + "/products.json",
	}, nil
}

func (shopifyConfig *ConfigShopify) FetchProducts() (objects.ShopifyProducts, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, shopifyConfig.Url+"?limit="+PRODUCT_FETCH_LIMIT, nil)
	fmt.Println(shopifyConfig.Url + "?limit=" + PRODUCT_FETCH_LIMIT)
	if err != nil {
		log.Fatal(err)
		return objects.ShopifyProducts{}, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return objects.ShopifyProducts{}, err
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return objects.ShopifyProducts{}, err
	}
	products := objects.ShopifyProducts{}
	fmt.Println(len(products.Products))
	err = json.Unmarshal(respBody, &products)
	if err != nil {
		log.Fatal(err)
		return objects.ShopifyProducts{}, err
	}
	return products, nil
}

// validates the data for the API config connectors
func validateConfig(store_name, api_key, api_password string) error {
	if store_name == "" {
		return errors.New("invalid store name")
	}
	if api_key == "" {
		return errors.New("invalid api key")
	}
	if api_password == "" {
		return errors.New("invalid api password")
	}
	return nil
}
