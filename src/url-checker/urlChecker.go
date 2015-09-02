package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
	"url-checker/strategy"
)

var csvPath string
var useAPI bool

//Create a wait group so we can fire off all the goroutines and wait for them to complete
var consumer_wg sync.WaitGroup

func init() {
	flag.StringVar(&csvPath, "csvPath", "urls.csv", "The path of the csv to parse")
	flag.BoolVar(&useAPI, "api", false, "Use the help-cms api")
}

func main() {
	flag.Parse()
	log.Println("Starting...")
	startTime := time.Now()

	var urlStrategy strategy.UrlStrategy
	if useAPI {
		log.Println("Using the API to retrieve urls...")
		urlStrategy = strategy.UrlStrategy(&strategy.ApiUrls{})
	} else {
		log.Println("Using a csv to retrieve the urls...")
		urlStrategy = strategy.UrlStrategy(&strategy.CsvUrls{Path: csvPath})
	}

	urlsToHit, err := urlStrategy.Get()
	if err != nil {
		log.Fatalln(err)
	}

	//Main channel to produce what will essentially be the queue of urls to hit
	var urls = make(chan string, len(urlsToHit))
	//Unroll the []string into the queue channel
	for _, val := range urlsToHit {
		urls <- val
	}
	close(urls)
	urlParsingTimeTaken := time.Since(startTime)

	log.Println("Reaching out and touching the urls")
	var numConsumers int
	if runtime.NumCPU()*2 > len(urlsToHit) {
		numConsumers = len(urlsToHit)
	} else {
		numConsumers = runtime.NumCPU() * 2
	}

	var failedUrl = make(chan string, len(urlsToHit))

	for c := 0; c < numConsumers; c++ {
		consumer_wg.Add(1)
		go consumer(urls, failedUrl)
	}

	//Sit here until all the goroutines are completed, then continue to show the failures
	consumer_wg.Wait()

	timeTaken := time.Since(startTime)
	log.Println("Finished and took:", timeTaken)
	//Don't need this channel any more so close it
	close(failedUrl)
	noFailedUrls := len(failedUrl)
	for response := range failedUrl {
		color.Red("%s\n", response)
	}

	log.Printf("There where %d urls crawled in %s with %d failures. The api took %s to return and parse",
		len(urlsToHit),
		timeTaken,
		noFailedUrls,
		urlParsingTimeTaken)
}

func consumer(urls chan string, failedUrls chan string) {
	for url := range urls {
		hitUrl(url, failedUrls)
		//Check to see if we have drained the channel,
		//if so then signal that this consumer is done
		if len(urls) == 0 {
			consumer_wg.Done()
			return
		}
	}
}

//Hit the chosen URL using a GET request and signal its completion using the waitGroup
//Prints out the URL and the Status code to the command line
func hitUrl(url string, failedUrl chan string) {
	req, err := http.NewRequest("GET", url, nil)
	//Add me some akamai debug headers
	req.Header.Add("Pragma", "akamai-x-cache-on, akamai-x-check-cacheable")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		failedUrl <- err.Error()
		return
	}

	//trim out everthing that isn't the cache hit type
	fromIndex := strings.Index(resp.Header.Get("X-Cache"), " from")
	cacheHit := "No Cache header"
	if fromIndex > 0 {
		cacheHit = resp.Header.Get("X-Cache")[:fromIndex]
	}

	responseString := fmt.Sprintf("Is cacheable? %s  \t Cache hit: %s   \t Url : %s Status: %s \n", resp.Header.Get("X-Check-Cacheable"), cacheHit, url, resp.Status)
	if resp.StatusCode == http.StatusOK {
		color.Green(responseString)
	} else {
		failedUrl <- responseString
	}
}
