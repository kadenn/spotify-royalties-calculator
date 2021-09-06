package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestHello(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message":"hello"}`, w.Body.String())
}

func TestGetFeaturedPlaylists(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/featured_playlists", nil)
	router.ServeHTTP(w, req)

	res := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, res["title"])
	assert.NotEmpty(t, res["playlists"])

}

func TestGetAlbumTracksByID(t *testing.T) {
	// ALBUM URL: https://open.spotify.com/album/7M0Zg2A3mrTOOqfVyRUjb8?si=qcy1NozGQnWpuGa2Mj41Xw&dl_branch=1

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album_tracks/7M0Zg2A3mrTOOqfVyRUjb8", nil)
	router.ServeHTTP(w, req)

	res := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res["albumID"], "7M0Zg2A3mrTOOqfVyRUjb8")
	assert.Len(t, res["tracks"], 11)
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "0cFyMyGiXySrooKVF8qWnH", "name": "Amniotic"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "3lchYylgQrEHvabP3rLAjX", "name": "Sirens"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "22DI1dE9Ii05gLRkeDnF1H", "name": "Frozen"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "5OKoDYrRmSbPJAGxdN8UO3", "name": "Swallow"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "4sONjuAg3sv8wrJEU5xX1X", "name": "Riverman"})
}

func TestGetAlbumTracksByWrongID(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album_tracks/ASDASDASDASDASD", nil)
	router.ServeHTTP(w, req)

	res := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"message":"couldn't get album tracks"}`, w.Body.String())
}

func TestGetPlaylistTracksByID(t *testing.T) {
	// PLAYLIST URL: https://open.spotify.com/playlist/2YOCwzVIgGHzqjaYWlBUS3?si=4d7cbddf84ae4c13

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/playlist_tracks/2YOCwzVIgGHzqjaYWlBUS3", nil)
	router.ServeHTTP(w, req)

	res := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res["playlistID"], "2YOCwzVIgGHzqjaYWlBUS3")
	assert.Len(t, res["tracks"], 100)
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "4fFoT8ncHrvYD0CcsxI4XX", "name": "Galuchat"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "0cX5GAhqkM8kcM4vuRTeji", "name": "Leonardo"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "6KqtrO5LMg2fQ3Aes3Yz8M", "name": "Panama"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "4R111miDcCg5dU1q1Ewfya", "name": "Kissing Your Shadow - MBNN Remix"})
	assert.Contains(t, res["tracks"], map[string]interface{}{"trackID": "3C8jZgiUP0pP1l58837NzH", "name": "Turn Around"})

}

func TestGetPlaylistTracksByWrongID(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/playlist_tracks/ASDASDASDASDASD", nil)
	router.ServeHTTP(w, req)

	res := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"message":"couldn't get playlist tracks"}`, w.Body.String())
}

func TestGetTrackByID(t *testing.T) {
	// TRACK URL: https://open.spotify.com/track/4TXYAETrGC53xgXZ7ykNEl?si=f5caaeefdf264968

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/track/4TXYAETrGC53xgXZ7ykNEl", nil)
	router.ServeHTTP(w, req)

	res := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res["trackID"], "4TXYAETrGC53xgXZ7ykNEl")
	assert.Equal(t, res["name"], "Viva la vida")
	assert.GreaterOrEqual(t, res["totalRoyalties"], float64(0))
	assert.GreaterOrEqual(t, res["playCount"], float64(0))
	assert.Len(t, res["artists"], 3)
	assert.Len(t, res["images"], 3)
}

func TestGetTrackByWrongID(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/track/ASDASDASDASDASD", nil)
	router.ServeHTTP(w, req)

	res := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"message":"couldn't get track"}`, w.Body.String())
}

func TestEstimatePlayCount(t *testing.T) {
	assert.GreaterOrEqual(t, estimatePlayCount(0), 0)
	assert.GreaterOrEqual(t, estimatePlayCount(20), 0)
	assert.GreaterOrEqual(t, estimatePlayCount(40), 0)
	assert.GreaterOrEqual(t, estimatePlayCount(60), 0)
	assert.GreaterOrEqual(t, estimatePlayCount(80), 0)
	assert.GreaterOrEqual(t, estimatePlayCount(100), 0)
}

func TestEstimateTotalRoyalties(t *testing.T) {
	assert.Equal(t, estimateTotalRoyalties(123), float64(0.369))
	assert.Equal(t, estimateTotalRoyalties(123123), float64(369.369))
	assert.Equal(t, estimateTotalRoyalties(123123123), float64(369369.369))
	assert.Equal(t, estimateTotalRoyalties(1000000), float64(3000))
	assert.Equal(t, estimateTotalRoyalties(10000000), float64(30000))
	assert.Equal(t, estimateTotalRoyalties(100000000), float64(300000))
}
