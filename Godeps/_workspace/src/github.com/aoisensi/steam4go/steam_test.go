package steam4go

import (
	"bufio"
	"os"
	"testing"
)

const (
	SteamIDTest SteamID = 76561198049739081
	AppIDTF2    AppID   = 440
	AppIDDota2  AppID   = 570
	AppIDKF     AppID   = 1250
	AppIDPD2    AppID   = 218620
	AppIDTest   AppID   = AppIDPD2
	APIKeyPath          = "apikey.txt"
)

func TestGetAppList(t *testing.T) {
	api := NewSteamAPI(LoadAPIKey())
	_, err := api.GetAppList()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetPlayerSummaries(t *testing.T) {
	api := NewSteamAPI(LoadAPIKey())
	_, err := api.GetPlayerSummary(SteamIDTest)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetUserStatsForGame(t *testing.T) {
	api := NewSteamAPI(LoadAPIKey())
	_, err := api.GetUserStatsForGame(SteamIDTest, AppIDTest)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetFriendList(t *testing.T) {
	api := NewSteamAPI(LoadAPIKey())
	_, err := api.GetFriendList(SteamIDTest)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetPlayerAchievements(t *testing.T) {
	api := NewSteamAPI(LoadAPIKey())
	_, err := api.GetPlayerAchievements(SteamIDTest, AppIDTest)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestQueryFiles(t *testing.T) {
	api := NewSteamAPI(LoadAPIKey())
	args := NewQueryFilesArgs()
	_, err := api.QueryFiles(args)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetServerInfo(t *testing.T) {
	api := NewSteamAPI(LoadAPIKey())
	_, err := api.GetServerInfo()
	if err != nil {
		t.Error(err)
		return
	}
}

func LoadAPIKey() string {
	file, err := os.Open(APIKeyPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}
