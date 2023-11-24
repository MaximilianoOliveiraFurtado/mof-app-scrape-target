package mapper

import (
	"github.com/MaximilianoOliveiraFurtado/mof-app-scrape-target/internal/dynamodbmodel"
	"github.com/MaximilianoOliveiraFurtado/mof-app-scrape-target/internal/model"
)

func ToDynamoDBScrapeTarget(scrapeTarget model.ScrapeTarget) dynamodbmodel.ScrapeTarget {
	return dynamodbmodel.ScrapeTarget{
		TargetID:    scrapeTarget.TargetID,
		Method:      scrapeTarget.Method,
		Description: scrapeTarget.Description,
		URL:         scrapeTarget.URL,
		Domain:      scrapeTarget.Domain,
		Version:     scrapeTarget.Version,
		BasePath:    scrapeTarget.BasePath,
		Paths:       scrapeTarget.Paths,
		QueryParams: scrapeTarget.QueryParams,
		Headers:     scrapeTarget.Headers,
		PayloadBody: scrapeTarget.PayloadBody,
	}
}

func FromDynamoDBScrapeTarget(scrapeTarget dynamodbmodel.ScrapeTarget) model.ScrapeTarget {
	return model.ScrapeTarget{
		TargetID:    scrapeTarget.TargetID,
		Method:      scrapeTarget.Method,
		Description: scrapeTarget.Description,
		URL:         scrapeTarget.URL,
		Domain:      scrapeTarget.Domain,
		Version:     scrapeTarget.Version,
		BasePath:    scrapeTarget.BasePath,
		Paths:       scrapeTarget.Paths,
		QueryParams: scrapeTarget.QueryParams,
		Headers:     scrapeTarget.Headers,
		PayloadBody: scrapeTarget.PayloadBody,
	}
}
