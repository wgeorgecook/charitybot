package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// startServer spins up an http listener for this service on the
// port and path specified
func startServer() {
	log.Info("New server started")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", parseResponse)
	log.Fatal(http.ListenAndServe(cfg.Port, router))
}

// parseResponse captures the get params off the incoming request
// We use this to get data following a successful donation
// Example:
// localhost:8080?invoker=@wgeorgecook&honorary=@charityyeti&inReplyToTweetID=1197178917825630210&donationValue=5
func parseResponse(w http.ResponseWriter, r *http.Request) {

	replyToTweetID, _ := strconv.ParseInt(r.URL.Query()["inReplyToTweetID"][0], 10, 64)

	tweet := successfulDonationData{
		r.URL.Query()["invoker"][0],
		r.URL.Query()["honorary"][0],
		replyToTweetID,
		r.URL.Query()["donationValue"][0],
	}

	log.Infow("Endpoint hit")

	if tweet.invoker == "" || tweet.honorary == "" || tweet.donationValue == "" || tweet.inReplyToTweetID == 0 {
		fmt.Fprintf(w, "All requests must include 'invoker', 'honorary', and 'inReplyToTweetID', and 'donationValue' params")
	} else {
		fmt.Fprintf(w, fmt.Sprintf("{Data: { invoker: %v, honorary: %v, replyToURL: %v, donationValue: %v}}", tweet.invoker, tweet.honorary, tweet.inReplyToTweetID, tweet.donationValue))
		err := respondToDonation(tweet)

		if err != nil {
			log.Error(err)
		}
	}
}
