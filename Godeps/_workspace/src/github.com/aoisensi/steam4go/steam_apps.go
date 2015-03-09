package steam4go

import (
	"encoding/json"
)

const (
	getAppListURL          = "http://api.steampowered.com/ISteamApps/GetAppList/v2"
	getServersAtAddressURL = "http://api.steampowered.com/ISteamApps/GetServersAtAddress/v1"
	upToDateCheckURL       = "http://api.steampowered.com/ISteamApps/UpToDateCheck/v1"
)

//App date
type App struct {
	AppID AppID
	Name  string
}

//Server date
type Server struct {
	Addr     string
	GmsIndex int
	AppID    AppID
	GameDir  string
	Region   int
	Secure   bool
	Lan      bool
	GamePort int
	SpecPort int
}

//Servers date
type Servers struct {
	Success bool
	Servers []Server
}

//UpToDate date
type UpToDate struct {
	Success           bool
	UpToDate          bool   `json:"up_to_date"`
	VersionIsListable bool   `json:"version_is_listable"`
	RequiredVersion   uint32 `json:"required_version"`
	Message           string
}

//GetAppList is ISteamApps/GetAppList/v2
func (p *SteamAPI) GetAppList() ([]App, error) {
	body, err := getJSONFromURL(getAppListURL)
	if err != nil {
		return nil, err
	}
	var r struct{ AppList struct{ Apps []App } }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return r.AppList.Apps, nil
}

//GetServersAtAddress is ISteamApps/GetServersAtAddress/v1
func (p *SteamAPI) GetServersAtAddress() (*Servers, error) {
	body, err := getJSONFromURL(getServersAtAddressURL)
	if err != nil {
		return nil, err
	}
	var r struct{ Response Servers }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.Response, nil
}

/*
//UpToDateCheck is ISteamApps/UpToDateCheck/v1
func (p *SteamAPI) UpToDateCheck(appid AppID, version uint32) (*UpToDate, error) {
	body, err := getJSONFromURL(getServersAtAddressURL)
	if err != nil {
		return nil, err
	}
	var r struct{ Response UpToDate }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.Response, nil
}
*/
