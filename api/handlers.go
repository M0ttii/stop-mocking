package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowerUser struct {
	Pk string `json:"pk"`
}

type FollowersResponse struct {
	Users     []FollowerUser `json:"users"`
	NextMaxID string         `json:"next_max_id"`
	Status    string         `json:"status"`
}

type Follower struct {
	UserID string `json:"pk"`
}

func generateRandomUserID() string {
	return fmt.Sprintf("%d", 62000000000+rand.Int63n(999999999))
}

func generateNextMaxID() string {
	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)
	return base64.StdEncoding.EncodeToString(randomBytes)
}

func FollowersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.Query("user_id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
			return
		}

		users := make([]FollowerUser, 22)
		for i := range users {
			users[i] = FollowerUser{
				Pk: generateRandomUserID(),
			}
		}

		response := FollowersResponse{
			Users:     users,
			NextMaxID: generateNextMaxID(),
			Status:    "ok",
		}

		c.JSON(http.StatusOK, response)
	}
}

type Session struct {
	Csrftoken string `json:"csrftoken"`
	DsUserID  string `json:"ds_user_id"`
}

type AddBestiesParam struct {
	Session           Session `json:"session"`
	FollowersResponse string  `json:"followers_response"`
	SerializedProxy   string  `json:"serialized_proxy"`
}

type FriendshipStatus struct {
	Following       bool `json:"following"`
	IncomingRequest bool `json:"incoming_request"`
	IsBestie        bool `json:"is_bestie"`
	IsBlockingReel  bool `json:"is_blocking_reel"`
	IsFeedFavorite  bool `json:"is_feed_favorite"`
	IsMutingReel    bool `json:"is_muting_reel"`
	IsPrivate       bool `json:"is_private"`
	IsRestricted    bool `json:"is_restricted"`
	OutgoingRequest bool `json:"outgoing_request"`
}

type BestiesResponse struct {
	FriendshipStatuses map[string]FriendshipStatus `json:"friendship_statuses"`
	Status             string                      `json:"status"`
}

func BestiesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := c.Request.ParseMultipartForm(10 << 20)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Error parsing form data"})
			return
		}

		addParam := c.Request.FormValue("add")
		if addParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Missing add parameter"})
			return
		}

		var userIDs []string
		err = json.Unmarshal([]byte(addParam), &userIDs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid add parameter format"})
			return
		}

		friendshipStatuses := make(map[string]FriendshipStatus)
		for _, userID := range userIDs {
			friendshipStatuses[userID] = FriendshipStatus{
				Following:       false,
				IncomingRequest: false,
				IsBestie:        true,
				IsBlockingReel:  false,
				IsFeedFavorite:  false,
				IsMutingReel:    false,
				IsPrivate:       false,
				IsRestricted:    false,
				OutgoingRequest: false,
			}
		}

		response := BestiesResponse{
			FriendshipStatuses: friendshipStatuses,
			Status:             "ok",
		}

		c.JSON(http.StatusOK, response)
	}
}
