package http

import (
	"fmt"
	"github.com/guoweikuang/v2exgo"
)

var _ v2ex.SiteService = &RawV2exService{}

type RawV2exService struct {
	ApiUrl string
}

func NewRawV2exService(apiUrl string) *RawV2exService {
	return &RawV2exService{
		ApiUrl: apiUrl,
	}
}

func (s *RawV2exService) GetStats() (*v2ex.Stat, error) {
	url := s.ApiUrl + "/site/stats.json"
	fmt.Println(url)
	stat := &v2ex.Stat{}
	err := v2ex.GetAPIData(url, stat)
	return stat, err
}

func (s *RawV2exService) GetInfo() (*v2ex.Info, error)  {
	url := s.ApiUrl + "/site/info.json"

	info := &v2ex.Info{}
	err := v2ex.GetAPIData(url, info)
	return info, err
}