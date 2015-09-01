package model

type UrlList struct {
	Urls []struct {
		Visible     bool   `json:"show_in_listings"`
		Url         string `json:"slug"`
		ProductType []struct {
			Type string `json:"term"`
		} `json:"product_type_tags"`
	} `json:"articles"`
}
