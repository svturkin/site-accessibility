package services

import (
	"fmt"
	"log"
	"site-accessibility/helpers"
	"site-accessibility/modules/site-accessibility-check/dto"
)

type SiteAccessibilityService struct{}

type SiteAccessibilityInterface interface {
	GetInfoByUrl(url string) (int64, error)
	GetFastestUrl() (string, error)
	GetSlowestUrl() (string, error)
}

func SiteAccessibilityServiceHandler() SiteAccessibilityInterface {
	svc := &SiteAccessibilityService{}

	return svc
}

func (service *SiteAccessibilityService) GetInfoByUrl(url string) (int64, error) {
	siteList, err := helpers.ReadJsonFile("siteList.json")
	if err != nil {
		return 0, err
	}

	for _, site := range siteList {
		if site.Name == url {
			return site.TimeinMs, nil
		}
	}

	return 0, fmt.Errorf("site %s not found", url)
}

func (service *SiteAccessibilityService) GetFastestUrl() (string, error) {
	siteList, err := helpers.ReadJsonFile("siteList.json")
	if err != nil {
		return "", err
	}

	var fastestUrl dto.Site
	for _, site := range siteList {
		if site.TimeinMs > 0 && (fastestUrl.TimeinMs == 0 || site.TimeinMs < fastestUrl.TimeinMs) {
			fastestUrl = site
		}
	}

	return fastestUrl.Name, nil
}

func (service *SiteAccessibilityService) GetSlowestUrl() (string, error) {
	siteList, err := helpers.ReadJsonFile("siteList.json")
	if err != nil {
		return "", err
	}

	var slowestUrl dto.Site
	for _, site := range siteList {
		if site.TimeinMs > slowestUrl.TimeinMs {
			slowestUrl = site
		}
	}

	log.Println(slowestUrl)
	return slowestUrl.Name, nil
}
