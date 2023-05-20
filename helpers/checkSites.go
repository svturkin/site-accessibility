package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"site-accessibility/modules/site-accessibility-check/dto"
)

var sites = make([]dto.Site, 0)

func sendRequest(site *dto.Site) {
	startTime := time.Now()

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://" + site.Name)
	if err != nil {
		log.Println(err)
		site.IsOn = false
		site.TimeinMs = 0
		return
	}
	defer resp.Body.Close()

	elapsedTime := time.Since(startTime).Milliseconds()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		site.IsOn = true
		site.TimeinMs = elapsedTime
	} else {
		site.IsOn = false
		site.TimeinMs = 0
	}
}

func ReadJsonFile(fileName string) ([]dto.Site, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &sites)
	if err != nil {
		return nil, err
	}

	return sites, nil
}

func writeToJsonFile(fileName string) {
	data, err := json.Marshal(sites)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckNewData() {
	sites, _ := ReadJsonFile("siteList.json")

	var wg sync.WaitGroup
	wg.Add(len(sites))

	for i := range sites {
		go func(site *dto.Site) {
			defer wg.Done()
			sendRequest(site)
		}(&sites[i])
	}

	wg.Wait()

	writeToJsonFile("siteList.json")
}
