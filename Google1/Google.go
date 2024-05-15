package Google1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://accounts.google.com/o/oauth2/v2/auth",
	TokenURL:  "https://oauth2.googleapis.com/token",
	AuthStyle: oauth2.AuthStyleInParams,
}

func GoogleLogin(c *gin.Context) {
	oauthConf := &oauth2.Config{
		ClientID:     os.Getenv("Client_ID"),
		ClientSecret: os.Getenv("Client_Secret"),
		RedirectURL:  os.Getenv("RedirectURL"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  Endpoint.AuthURL,
			TokenURL: Endpoint.TokenURL,
		},
		Scopes: []string{"https://www.googleapis.com/auth/drive.metadata.readonly"},
	}

	url := oauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleToken(c *gin.Context) {
	oauthConf := &oauth2.Config{
		ClientID:     os.Getenv("Client_ID"),
		ClientSecret: os.Getenv("Client_Secret"),
		RedirectURL:  os.Getenv("RedirectURL"),
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
	code := c.Query("code")
	token, err := oauthConf.Exchange(context.TODO(), code)
	if err != nil {
		// Handle error, e.g., log it and return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}
	fmt.Println("Access Token:", token.AccessToken)
	fmt.Println("Refresh Token:", token.RefreshToken)
	userInfo, err := GoogleUserInfo(token.AccessToken)
	if err != nil {
		// Handle error
		fmt.Println("Error getting user info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken, "Refresh Token": token.RefreshToken, "user_info": userInfo})

}
func GoogleUserInfo(accessToken string) (map[string]interface{}, error) {
	userInfoURL := "https://www.googleapis.com/oauth2/v3/userinfo"

	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	fmt.Println("User Info:", resp.Body)
	return userInfo, nil
}
