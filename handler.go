package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jboverfelt/infest/models"
	"github.com/markbates/pop"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// Env represents handler dependencies
type Env struct {
	DB *pop.Connection
}

// Handler represents an HTTP handler that can return errors
// and access dependencies in a type-safe way
type Handler struct {
	*Env
	h func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// NewHandler allocates a new handler
func NewHandler(e *Env, handlerFunc func(e *Env, w http.ResponseWriter, r *http.Request) error) Handler {
	return Handler{e, handlerFunc}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.h(h.Env, w, r)
	if err != nil {
		log.Printf("ERROR: %v", err)
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func all(e *Env, w http.ResponseWriter, r *http.Request) error {
	var closures []models.Closure
	params := r.URL.Query()

	baseQuery := e.DB.PaginateFromParams(params)

	addQueryParams(baseQuery, params)

	err := baseQuery.All(&closures)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(closures)

	if err != nil {
		return err
	}

	return nil
}

func addQueryParams(query *pop.Query, params url.Values) {
	supportedParams := map[string]string{
		"sort":      params.Get("sort"),
		"name":      params.Get("name"),
		"reason":    params.Get("reason"),
		"startDate": params.Get("startDate"),
		"endDate":   params.Get("endDate"),
	}

	if supportedParams["sort"] != "" {
		addSort(query, supportedParams["sort"])
	}

	if supportedParams["name"] != "" {
		query.Where("name = ?", supportedParams["name"])
	}

	if supportedParams["reason"] != "" {
		query.Where("reason LIKE '%' || ? || '%'", supportedParams["reason"])
	}

	if supportedParams["startDate"] != "" {
		tmpTime, err := time.Parse(closureTimeFmt, supportedParams["startDate"])

		if err == nil {
			query.Where("closureDate >= ?", tmpTime)
		}
	}

	if supportedParams["endDate"] != "" {
		tmpTime, err := time.Parse(closureTimeFmt, supportedParams["endDate"])

		if err == nil {
			query.Where("closureDate <= ?", tmpTime)
		}
	}
}

func addSort(query *pop.Query, sortVal string) {
	supportedSortFields := map[string]bool{
		"name":           true,
		"address":        true,
		"reason":         true,
		"closureDate":    true,
		"reopenDate":     true,
		"reopenComments": true,
	}
	parts := strings.Split(sortVal, ".")

	if len(parts) == 2 {
		if strings.ToLower(parts[1]) == "asc" || strings.ToLower(parts[1]) == "desc" {
			if _, ok := supportedSortFields[parts[0]]; ok {
				query.Order(strings.Join(parts, " "))
			}
		}
	}
}
