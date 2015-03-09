package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aoisensi/steam4go"
)

var (
	statusCache     status
	statusCacheTime time.Time
	steam           *steam4go.SteamAPI
	steamid         steam4go.SteamID
)

type status struct {
	Steam struct {
		Success    bool   `json:"success"`
		Status     int    `json:"status"`
		GameName   string `json:"game_name"`
		LastOnline int64  `json:"last_online"`
		Update     int64  `json:"update"`
	} `json:"steam"`
	Update int64 `json:"update"`
}

func initStatus() {
	steam = steam4go.NewSteamAPI(os.Getenv("STEAM_API"))
	id, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)
	steamid = steam4go.SteamID(id)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	if time.Now().Sub(statusCacheTime) > time.Minute {
		statusSteam()
		statusCache.Update = time.Now().Unix()
	}
	data, _ := json.Marshal(&statusCache)
	w.Write(data)
}

func statusSteam() {
	s, err := steam.GetPlayerSummary(steamid)
	if err != nil {
		return
	}
	v := statusCache.Steam
	v.Status = int(s.PersonaState)
	v.GameName = s.GameExtraInfo
	v.LastOnline = s.LastLogoff
	v.Success = true
	v.Update = time.Now().Unix()
	statusCache.Steam = v
}
