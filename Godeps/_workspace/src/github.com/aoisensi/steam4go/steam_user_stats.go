package steam4go

import (
	"encoding/json"
	"net/url"
)

//PlayerStats date
type PlayerStats struct {
	SteamID  SteamID `json:",string"`
	GameName string
	Stats    []struct {
		Name  string
		Value int
	}
	Achievements []struct {
		APIName  string
		Name     string
		Achieved int
	}
}

//GetUserStatsForGame is ISteamUserStats/GetUserStatsForGame/v2
func (p *SteamAPI) GetUserStatsForGame(steamid SteamID, appid AppID) (*PlayerStats, error) {
	params := url.Values{}
	params.Add("steamid", steamid.String())
	params.Add("appid", appid.String())
	url := p.genURL("ISteamUserStats", "GetUserStatsForGame", ver2, params)
	body, err := getJSONFromURL(url)
	if err != nil {
		return nil, err
	}
	var r struct{ PlayerStats PlayerStats }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.PlayerStats, nil
}

//GetPlayerAchievements is ISteamUserStats/GetPlayerAchievements/v1
func (p *SteamAPI) GetPlayerAchievements(steamid SteamID, appid AppID) (*PlayerStats, error) {
	params := url.Values{}
	params.Add("steamid", steamid.String())
	params.Add("appid", appid.String())
	url := p.genURL("ISteamUserStats", "GetPlayerAchievements", ver1, params)
	body, err := getJSONFromURL(url)
	if err != nil {
		return nil, err
	}
	var r struct{ PlayerStats PlayerStats }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.PlayerStats, nil
}
