package chapi

import "os"

var (
	apiKey  = os.Getenv("CH_API_KEY")
	baseUrl = "https://api.companieshouse.gov.uk"
)

func SetApiKey(k string) {
	apiKey = k
}

func ApiKey() string {
	return apiKey
}

func BaseUrl() string {
	return baseUrl
}

func SetBaseUrl(u string) {
	baseUrl = u
}
