package productfetch

import (
	"errors"
	"log"
	"utils"
)

type ConfigWoo struct {
	StoreName string
	Key       string
	Secret    string
}

type ConfigShopify struct {
	StoreName   string
	APIKey      string
	APIPassword string
	Version     string
}

// Initiates the connection string for woocommerce
func (configWoo *ConfigWoo) InitConfig() (ConfigWoo, error) {
	store_name := utils.LoadEnv("store_name")
	api_key := utils.LoadEnv("api_key")
	api_secret := utils.LoadEnv("api_password")

	validation := validateConfig(store_name, api_key, api_secret)
	if validation != nil {
		log.Fatal("Error setting up connection string for WooCommerce")
		return ConfigWoo{}, validation
	}
	return ConfigWoo{
		StoreName: "https://" + store_name + "/wc-api/v3",
		Key:       api_key,
		Secret:    api_secret,
	}, nil

}

// Initiates the connection string for shopify
func (configShopify *ConfigShopify) InitShopifyURL() (ConfigShopify, error) {
	store_name := utils.LoadEnv("store_name")
	api_key := utils.LoadEnv("api_key")
	api_password := utils.LoadEnv("api_password")

	validation := validateConfig(store_name, api_key, api_password)
	if validation != nil {
		log.Fatal("Error setting up connection string for Shopify")
		return ConfigShopify{}, validation
	}
	return ConfigShopify{
		StoreName:   store_name,
		APIKey:      api_key,
		APIPassword: api_password,
		Version:     "2023-07",
	}, nil
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
