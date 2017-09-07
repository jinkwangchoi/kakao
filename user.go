package kakao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID                    int64          `json:"id"`
	KAccountEmail         string         `json:"kaccount_email"`
	KAccountEmailVerified bool           `json:"kaccount_email_verified"`
	Properties            UserProperties `json:"properties"`
}

type UserProperties struct {
	Nickname       string `json:"nickname"`
	ThumbnailImage string `json:"thumbnail_image"`
	ProfileImage   string `json:"profile_image"`
	CustomField1   string `json:"custom_field1"`
	CustomField2   string `json:"custom_field2"`
}

func GetUser(accessToken string) (*User, error) {
	req, err := http.NewRequest("GET", `https://kapi.kakao.com/v1/user/me`, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %s", err.Error())
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("client.Do error: %s", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s", err.Error())
	}

	return user, nil
}
