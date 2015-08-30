package model

type UrlList struct {
	Urls []struct {
		Visible bool   `json:"show_in_listings"`
		Url     string `json:"slug"`
	} `json:"articles"`
}
