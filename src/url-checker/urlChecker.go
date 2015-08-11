package main

import (
	"encoding/csv"
	"flag"
	"github.com/fatih/color"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var csvPath string

func init() {
	flag.StringVar(&csvPath, "csvPath", "urls.csv", "The path of the csv to parse")
}

func main() {
	flag.Parse()

	urlsToHit, err := readUrls(csvPath)
	if err != nil {
		log.Fatalln(err)
	}

	//Create a wait group so we can fire off all the goroutines and wait for them to complete
	var wg sync.WaitGroup

	for key, val := range urlsToHit {
		//Fire off as many go routines as we have URLs
		//pewpewpewpewpewpew
		wg.Add(1)
		//index[0] will always contain the one and only element of a valid URL csv
		go hitUrl(val[0], &wg)

		//Some *super* naive throttling
		//every 10 urls pause for a second before adding another 10 urls
		if key > 0 && key%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	//Sit here until all the goroutines are completed, then exit nicely
	wg.Wait()
}

//Creates an array of the urls that are present in the CSV
func readUrls(path string) ([][]string, error) {
	csvfile, err := os.Open(path)
	if err != nil {
		return [][]string{}, err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = 1

	CSVdata, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return CSVdata, nil
}

//Hit the chosen URL using a GET request and signal its completion using the waitGroup
//Prints out the URL and the Status code to the command line
func hitUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		color.Red("%s\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		color.Red("Url : %s Status: %s\n", url, resp.Status)
	} else {
		color.Green("Url : %s Status: %s\n", url, resp.Status)
	}
}
