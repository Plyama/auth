package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	VerifiedEmail bool   `json:"verified_email"`
}

func NewGoogleConfig(redirectURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}

func GetGoogleUserInfo(accessToken string) (GoogleUserData, error) {
	var userInfo GoogleUserData

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(accessToken))
	if err != nil {
		return userInfo, err
	}
	defer resp.Body.Close()

	userBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return userInfo, err
	}

	err = json.Unmarshal(userBytes, &userInfo)

	return userInfo, err
}

func GoogleSignUpCallbackURL() string {
	return fmt.Sprintf("http://%s:%s/%s", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("GOOGLE_SIGN_UP_CALLBACK"))
}
