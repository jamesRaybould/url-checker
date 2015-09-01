package strategy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"url-checker/model"
)

type ApiUrls struct {
	BaseUrl string
}

func (s *ApiUrls) Get() ([]string, error) {

	if s.BaseUrl == "" {
		s.BaseUrl = model.Settings.BaseUrl
	}
	resp, err := http.Get(model.Settings.ApiUrl)
	if err != nil {
		log.Panicln(err)
		return []string{}, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return []string{}, err
	}

	var urlList model.UrlList
	err = json.Unmarshal(contents, &urlList)

	var urls []string
	for _, value := range urlList.Urls {
		if value.Visible && value.ProductType[0].Type != model.Settings.ProductType {
			urls = append(urls, s.BaseUrl+value.Url)
		}
	}

	return urls, nil
}
