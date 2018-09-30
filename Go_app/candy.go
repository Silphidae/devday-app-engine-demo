package main

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	candyKind    = "Candy"
	listPageSize = 100
)

type Candy struct {
	ID      int64 `datastore:"-"`
	Text    string
	Created string
}

type CandyList struct {
	Candies []Candy `json:"items"`
}

func routeCandiesList(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	q := datastore.NewQuery(candyKind).Order("Created")

	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr != "" {
		cursor, err := datastore.DecodeCursor(cursorStr)
		if err != nil {
			log.Errorf(c, "Could not decode cursor '%v': '%v'", cursorStr, err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		q = q.Start(cursor)
	}

	var cList CandyList
	t := q.Run(c)

	for {
		var ci Candy
		key, err := t.Next(&ci)
		if err == datastore.Done {
			break
		}
		if err != nil {
			log.Errorf(c, "Could not fetch next item: %v", err)
			break
		}
		ci.ID = key.IntID()
		cList.Candies = append(cList.Candies, ci)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(mustJSON(cList))
}

func routeCandyCreate(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	type CandyReq struct {
		Text string `json:"text"`
	}

	dec := json.NewDecoder(r.Body)

	var candyReq CandyReq

	if err := dec.Decode(&candyReq); err != nil {
		log.Errorf(c, "Decoding JSON failed while creating project: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	key := datastore.NewIncompleteKey(c, candyKind, nil)
	var candy Candy
	candy.Text = candyReq.Text
	JavascriptISOString := "2006-01-02T15:04:05.999Z07:00"
	candy.Created = time.Now().UTC().Format(JavascriptISOString)

	key, err := datastore.Put(c, key, &candy)

	if err != nil && key != nil {
		log.Errorf(c, "Adding candy failed: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	candy.ID = key.IntID()

	w.Write(mustJSON(candy))
}
