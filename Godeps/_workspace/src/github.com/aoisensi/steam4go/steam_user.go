package steam4go

import (
	"encoding/json"
	"net/url"
	"strings"
)

//PersonaState const
type PersonaState int

const (
	PersonaStateOffline        PersonaState = 1
	PersonaStateOnline                      = 2
	PersonaStateBusy                        = 3
	PersonaStateAway                        = 4
	PersonaStateSnooze                      = 5
	PersonaStateLookingToTrade              = 6
	PersonaStateLookingToPlay               = 7
)

//Player date
type Player struct {
	//public data
	SteamID                  SteamID `json:",string"`
	PersonaName              string
	ProfileURL               string
	Avater                   string
	AvaterMedium             string
	AvaterFull               string
	PersonaState             PersonaState
	CommunityVisibilityState int
	ProfileState             int
	LastLogoff               int64
	CommentPermission        int

	//private data
	RealName       string
	PrimaryClanID  string
	TimeCreated    int
	GameID         AppID `json:",string"`
	GameServerIP   string
	GameExtraInfo  string
	CityID         int // DEPRECATED
	LocCountryCode string
	LocStateCode   string
	LocCityID      int
}

//Relationship date
type Relationship string

//Friend date
type Friend struct {
	SteamID      SteamID `json:",string"`
	Relationship string
	FriendSince  uint64 `json:"friend_since"`
}

//GetPlayerSummary is binded GetPlayerSummaries
func (p *SteamAPI) GetPlayerSummary(steamid SteamID) (*Player, error) {
	players, err := p.GetPlayerSummaries([]SteamID{steamid})
	if err != nil {
		return nil, err
	}
	return &players[0], nil
}

//GetPlayerSummaries is ISteamUser/GetPlayerSummaries/v2
func (p *SteamAPI) GetPlayerSummaries(steamids []SteamID) ([]Player, error) {
	params := url.Values{}
	ids := make([]string, len(steamids))
	for i, id := range steamids {
		ids[i] = id.String()
	}
	params.Add("steamids", strings.Join([]string(ids), ","))
	url := p.genURL("ISteamUser", "GetPlayerSummaries", ver2, params)
	body, err := getJSONFromURL(url)
	if err != nil {
		return nil, err
	}
	var r struct{ Response struct{ Players []Player } }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return r.Response.Players, nil
}

//GetFriendList is  ISteamUser/GetFriendList/v1
func (p *SteamAPI) GetFriendList(steamid SteamID) ([]Friend, error) {
	params := url.Values{}
	params.Add("steamid", steamid.String())
	params.Add("relationship", "all")
	url := p.genURL("ISteamUser", "GetFriendList", ver1, params)
	body, err := getJSONFromURL(url)
	if err != nil {
		return nil, err
	}
	var r struct{ FriendsList struct{ Friends []Friend } }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return r.FriendsList.Friends, nil
}
