package steam4go

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	baseURL = "http://api.steampowered.com"
	ver1    = "v1"
	ver2    = "v2"
)

//SteamID date
type SteamID uint64

//AppID date
type AppID uint32

//SteamAPI main class
type SteamAPI struct {
	key string
}

//NewSteamAPI is gen steam api instance
func NewSteamAPI(key string) *SteamAPI {
	return &SteamAPI{key: key}
}

func (p *SteamAPI) genURL(ifname, mtname, ver string, params url.Values) string {
	url := strings.Join([]string{baseURL, ifname, mtname, ver}, "/")
	params.Add("key", p.key)
	return url + "?" + params.Encode()
}

func (p SteamID) String() string {
	return strconv.FormatUint(uint64(p), 10)
}

func (p AppID) String() string {
	return strconv.FormatUint(uint64(p), 10)
}

func getJSONFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
