package productfetch

import (
	"encoding/json"
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
	Valid  bool
}

type ConfigShopify struct {
	APIKey      string
	APIPassword string
	Version     string
	Url         string
	Valid       bool
}

const PRODUCT_FETCH_LIMIT = "50" // limit on products to fetch

// Initiates the connection string for woocommerce
func InitConfigWoo(store_name, api_key, api_secret string) ConfigWoo {
	validation := validateConfigWoo(store_name, api_key, api_secret)
	if !validation {
		log.Println("Error setting up connection string for WooCommerce")
	}
	return ConfigWoo{
		Key:    api_key,
		Secret: api_secret,
		Url:    "https://" + store_name + "/wc-api/v3/products?consumer_key=" + api_key + "&consumer_secret=" + api_secret,
		Valid:  validation,
	}
}

// Fetches products from WooCommerce defined on the wooConfig url

func (wooConfig *ConfigWoo) FetchProducts() (objects.WooProducts, error) {
	httpClient := http.Client{
		Timeout: time.Second * 20,
	}
	req, err := http.NewRequest(http.MethodGet, wooConfig.Url+"&filter[limit]="+PRODUCT_FETCH_LIMIT, nil)
	if err != nil {
		log.Println(err)
		return objects.WooProducts{}, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return objects.WooProducts{}, err
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return objects.WooProducts{}, err
	}
	products := objects.WooProducts{}
	err = json.Unmarshal(respBody, &products)
	if err != nil {
		log.Println(err)
		return objects.WooProducts{}, err
	}
	return products, nil
}

// Initiates the connection string for shopify
func InitConfigShopify(store_name, api_key, api_password, version string) ConfigShopify {
	validation := validateConfigShopify(store_name, api_key, api_password)
	if !validation {
		log.Println("Error setting up connection string for Shopify")
	}
	return ConfigShopify{
		APIKey:      api_key,
		APIPassword: api_password,
		Version:     version,
		Url:         "https://" + api_key + ":" + api_password + "@" + store_name + ".myshopify.com/admin/api/" + version + "/products.json",
		Valid:       validation,
	}
}

func (shopifyConfig *ConfigShopify) FetchProducts() (objects.ShopifyProducts, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, shopifyConfig.Url+"?limit="+PRODUCT_FETCH_LIMIT, nil)
	if err != nil {
		log.Println(err)
		return objects.ShopifyProducts{}, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return objects.ShopifyProducts{}, err
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return objects.ShopifyProducts{}, err
	}
	products := objects.ShopifyProducts{}
	err = json.Unmarshal(respBody, &products)
	if err != nil {
		log.Println(err)
		return objects.ShopifyProducts{}, err
	}
	return products, nil
}

// validates the data for WooCommerce
func validateConfigWoo(store_name, api_key, api_password string) bool {
	if store_name == "" {
		log.Println("invalid store name")
		return false
	}
	if api_key == "" || !utils.CheckStringWithWord(api_key, "ck_") {
		log.Println("invalid api key")
		return false
	}
	if api_password == "" || !utils.CheckStringWithWord(api_password, "cs_") {
		log.Println("invalid api password")
		return false
	}
	return true
}

// validates the data for Shopify
func validateConfigShopify(store_name, api_key, api_password string) bool {
	if store_name == "" {
		log.Println("invalid store name")
		return false
	}
	if api_key == "" {
		log.Println("invalid api key")
		return false
	}
	if api_password == "" {
		log.Println("invalid api password")
		return false
	}
	return true
}
