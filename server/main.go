package main

import (
	"context"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

var spotifyClient = createSpotifyClient()

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default()) // Allows all origins.

	router.GET("/", hello)
	router.GET("/track/:id", getTrackByID)
	router.GET("/featured_playlists", getFeaturedPlaylists)
	router.GET("/playlist_tracks/:id", getPlaylistTracksByID)
	router.GET("/album_tracks/:id", getAlbumTracksByID)

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

// getAlbumTracksByID locates the album whose ID value matches the id parameter sent by the client, then returns tracks from this album.
func getAlbumTracksByID(c *gin.Context) {
	id := c.Param("id")
	page, err := spotifyClient.GetAlbumTracks(c, spotify.ID(id))

	if err != nil {
		// If the token has expired it will be refreshed.
		if strings.Contains(err.Error(), "token expired and refresh token is not set") {
			spotifyClient = createSpotifyClient()
			getPlaylistTracksByID(c)
		} else {
			log.Error("couldn't get album tracks:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get album tracks"})
		}
		return
	}

	var tracks []gin.H
	for _, track := range page.Tracks {
		tracks = append(tracks, gin.H{"trackID": track.ID, "name": track.Name})
	}

	c.JSON(http.StatusOK, gin.H{"albumID": id, "tracks": tracks})
}

// getPlaylistTracksByID locates the playlist whose ID value matches the id parameter sent by the client, then returns tracks from this playlist.
func getPlaylistTracksByID(c *gin.Context) {
	id := c.Param("id")
	page, err := spotifyClient.GetPlaylistTracks(c, spotify.ID(id))

	if err != nil {
		// If the token has expired it will be refreshed.
		if strings.Contains(err.Error(), "token expired and refresh token is not set") {
			spotifyClient = createSpotifyClient()
			getPlaylistTracksByID(c)
		} else {
			log.Error("couldn't get playlist tracks:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get playlist tracks"})
		}
		return
	}

	var tracks []gin.H
	for _, playlistTrack := range page.Tracks {
		track := playlistTrack.Track
		tracks = append(tracks, gin.H{"trackID": track.ID, "name": track.Name})
	}

	c.JSON(http.StatusOK, gin.H{"playlistID": id, "tracks": tracks})
}

// getFeaturedPlaylists responds with the list of featured playlists as JSON.
func getFeaturedPlaylists(c *gin.Context) {
	msg, page, err := spotifyClient.FeaturedPlaylists(c, spotify.Limit(12))

	if err != nil {
		// If the token has expired it will be refreshed.
		if strings.Contains(err.Error(), "token expired and refresh token is not set") {
			spotifyClient = createSpotifyClient()
			getFeaturedPlaylists(c)
		} else {
			log.Error("couldn't get featured playlists:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get featured playlists"})
		}
		return
	}

	var playlists []gin.H
	for _, playlist := range page.Playlists {
		playlists = append(playlists, gin.H{"playlistID": playlist.ID, "name": playlist.Name, "owner": playlist.Owner, "images": playlist.Images, "tracks": playlist.Tracks})
	}

	c.JSON(http.StatusOK, gin.H{"title": msg, "playlists": playlists})
}

// getTrack locates the track whose ID value matches the id parameter sent by the client, then returns that track as a response.
func getTrackByID(c *gin.Context) {
	id := c.Param("id")
	track, err := spotifyClient.GetTrack(c, spotify.ID(id))

	if err != nil {
		// If the token has expired it will be refreshed.
		if strings.Contains(err.Error(), "token expired and refresh token is not set") {
			spotifyClient = createSpotifyClient()
			getTrackByID(c)
		} else {
			log.Error("couldn't get track:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get track"})
		}
		return
	}

	// Spotify API doesn't share how many times a song has been played so I wrote a function to estimate it.
	playCount := estimatePlayCount(track.Popularity)

	totalRoyalties := estimateTotalRoyalties(playCount)

	c.JSON(http.StatusOK, gin.H{"trackID": track.ID, "name": track.Name, "artists": track.Artists, "playCount": playCount, "totalRoyalties": totalRoyalties, "images": track.Album.Images})
}

// estimatePlayCount estimates the number of times a track was played based on the popularity number.
func estimatePlayCount(popularity int) int {
	return int(math.Pow(float64(popularity), 3)) + rand.Intn(1000000)
}

// estimateTotalRoyalties estimates the total royalties payable to rights holders for this track.
func estimateTotalRoyalties(playCount int) float64 {
	// Spotify pays artists around £0.003 per stream
	// SOURCE: https://freeyourmusic.com/blog/how-much-does-spotify-pay-per-stream
	return float64(playCount) * 0.003
}

// Make sure you set the SPOTIFY_ID and SPOTIFY_SECRET environment variables prior to running this example.
func createSpotifyClient() *spotify.Client {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Error("couldn't get token:", err)
		return nil
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	log.Info("new token successfully received")

	return client
}
