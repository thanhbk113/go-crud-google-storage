package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"my-app/models"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func EnvGet(name string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	return os.Getenv(name)
}

func GetKeys() models.Keys {

	url := "https://storage.googleapis.com/mmt-app/keys.json?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=mmt-app%40project-18-11-369002.iam.gserviceaccount.com%2F20221122%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20221122T034801Z&X-Goog-Expires=900&X-Goog-Signature=43cfd0ca8a99403a384127d4dd53d575eed804da61a97db5c7d128ac2ffc83b1bd97e08d8a5f07bbe851eb4b0303fe387c687097d9834e038c27c3f73cc5d3cdf0166574aff74881cd4147141f75a1f58cca2635735ad5fd61fdcf5003205f6dfa89980805d7d437c4860be47f7b2c7f4217315b4c1480872b96aba0d086773bd49e25a2681f830bc88c0b9fa9b27861e82086be01cb10435801603ffca65766abae7978740b4675169a699d7e7d191cd2f92943bdc2d259970d84da51deb1a4f06c5a1c5862de13a9a86b08dd11aeaee105eef4ca501bd83abf7edeff9b4d6d4b981064dfcdcc0fece5d998ff9f4d3c4cf23d02e53e83421e13b6100c7b6735&X-Goog-SignedHeaders=host"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data models.Keys
	json.Unmarshal(body, &data)
	// KeysGet := models.Keys{
	// 	Type:                    config.EnvGet("GOOGLE_CLOUD_TYPE"),
	// 	ProjectID:               config.EnvGet("GOOGLE_CLOUD_PROJECT_ID"),
	// 	PrivateKeyID:            config.EnvGet("GOOGLE_CLOUD_PRIVATE_KEY_ID"),
	// 	PrivateKey:              config.EnvGet("GOOGLE_CLOUD_PRIVATE_KEY"),
	// 	ClientEmail:             config.EnvGet("GOOGLE_CLOUD_CLIENT_EMAIL"),
	// 	ClientID:                config.EnvGet("GOOGLE_CLOUD_CLIENT_ID"),
	// 	AuthURI:                 config.EnvGet("GOOGLE_CLOUD_AUTH_URI"),
	// 	TokenURI:                config.EnvGet("GOOGLE_CLOUD_TOKEN_URI"),
	// 	AuthProviderX509CertURL: config.EnvGet("GOOGLE_CLOUD_AUTH_PROVIDER_X509_CERT_URL"),
	// 	ClientX509CertURL:       config.EnvGet("GOOGLE_CLOUD_CLIENT_X509_CERT_URL"),
	// }

	return data
}
