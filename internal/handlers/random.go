package handlers

import (
	"net/http"

	"github.com/TovarischSuhov/go-template/internal/metrics"
)

func GetRandomTopicHandler(w http.ResponseWriter, r *http.Request) {
	id := metrics.GetRandomTopicID()
	topicPath := "https://getpocket.com/read/" + id
	http.Redirect(w, r, topicPath, http.StatusTemporaryRedirect)
}
