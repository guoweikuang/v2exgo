package v2ex

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidUsername = errors.New("v2ex: invalid user name")
	ErrInvalidNodename = errors.New("v2ex: invalid node name")
)


type Stat struct {
	TopicMax *int `json:"topic_max"`
	MemberMax *int `json:"member_max"`
}

type Info struct {
	Title string `json:"title"`
	Slogan string `json:"slogan"`
	Description string `json:"description"`
	Domain string `json:"domain"`
}

type SiteService interface {
	GetStats() (stats * Stat, err error)
	GetInfo() (stats * Info, err error)
}

func GetAPIData(apiUrl string, v interface{}) (err error) {
	resp, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println("error")
		return
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(v)
	return
}

