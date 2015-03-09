package steam4go

import (
	"encoding/json"
	"net/url"
	"time"
)

//SteamWebAPIUtil date
type SteamWebAPIUtil struct {
	APIList struct {
		Interface struct {
			Name   string
			Method []struct {
				Name       string
				Version    int
				HTTPMethod string
				Parameters []struct {
					Name        string
					Type        string
					Optional    bool
					Description string
				}
			}
		}
	}
}

//GetSupportedAPIList is ISteamWebAPIUtil/GetSupportedAPIList/v1
func (p *SteamAPI) GetSupportedAPIList() (*SteamWebAPIUtil, error) {
	params := url.Values{}
	url := p.genURL("ISteamWebAPIUtil", "GetSupportedAPIList", ver1, params)
	body, err := getJSONFromURL(url)
	if err != nil {
		return nil, err
	}
	var r SteamWebAPIUtil
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

//GetServerInfo is ISteamWebAPIUtil/GetServerInfo/v1
func (p *SteamAPI) GetServerInfo() (time.Time, error) {
	params := url.Values{}
	url := p.genURL("ISteamWebAPIUtil", "GetServerInfo", ver1, params)
	body, err := getJSONFromURL(url)
	if err != nil {
		return time.Time{}, err
	}
	var r struct {
		ServerTime       int64
		ServerTimeString string
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(r.ServerTime, 0), nil
}
