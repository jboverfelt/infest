package main

import (
	"log"
	"strings"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/jboverfelt/infest/models"
	"github.com/markbates/pop/nulls"
)

func fetchClosures(url string) ([]models.Closure, error) {
	doc, err := goquery.NewDocument(url)

	if err != nil {
		return nil, err
	}

	return doScrape(doc)
}

func doScrape(doc *goquery.Document) ([]models.Closure, error) {
	var closures []models.Closure
	var overallErr error

	doc.Find("#dnn_ContentPane tbody tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		var name string
		var address string
		var reason string
		var closureDate time.Time
		var reopenDate nulls.Time
		var reopenComments nulls.String

		s.Find("td").EachWithBreak(func(i int, s *goquery.Selection) bool {
			switch i {
			case 0:
				name = s.Text()
			case 1:
				address = s.Text()
			case 2:
				reason = s.Text()
			case 3:
				tmpClosureDate, err := time.Parse(closureTimeFmt, s.Text())
				if err != nil {
					overallErr = err
					return false
				}

				closureDate = tmpClosureDate
			case 4:
				reopenTxt := strings.TrimSpace(s.Text())

				// don't attempt to set if there's no value
				if reopenTxt == "" {
					break
				}

				reopenComments = nulls.NewString(reopenTxt)

				datePart := parseReopenDate(reopenTxt)

				tmpReopenDate, err := time.Parse(closureTimeFmt, datePart)
				if err != nil {
					log.Printf("Could not parse reopen date from string %s", reopenTxt)
					break
				}

				reopenDate = nulls.NewTime(tmpReopenDate)
			}

			return true
		})

		if overallErr != nil {
			return false
		}

		closures = append(closures, models.Closure{
			Name:           name,
			Address:        address,
			Reason:         reason,
			ClosureDate:    closureDate,
			ReopenDate:     reopenDate,
			ReopenComments: reopenComments,
		})

		return true
	})

	if overallErr != nil {
		return nil, overallErr
	}

	return closures, nil
}

func parseReopenDate(s string) string {
	idx := strings.IndexFunc(s, unicode.IsDigit)

	if idx != -1 {
		return s[idx:]
	}

	return ""
}
