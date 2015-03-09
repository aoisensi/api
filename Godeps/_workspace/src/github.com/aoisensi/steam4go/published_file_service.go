package steam4go

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//PublishedFileService date
type PublishedFileService struct {
	Total                int
	PublishedFileDetails []struct {
		Result                int
		PublishedFileID       uint64  `json:",string"`
		Creator               SteamID `json:",string"`
		CreatorAppID          AppID
		ConsmerAppID          AppID
		Filename              string
		FileSize              int64 `json:",string"`
		PreviewFileSize       int64 `json:",string"`
		PreviewURL            string
		URL                   string
		HContentPreview       uint64
		Title                 string
		ShortDescription      string
		TimeCreated           int
		TimeUpdated           int
		Visibility            int
		Flags                 int
		WorkshopFile          bool
		WorkshopAccepted      bool
		NumCommentsDeveloper  int
		NumCommentsPublic     int
		Banned                bool
		BanReason             string
		Banner                SteamID `json:",string"`
		CanBeDeleted          bool
		Incompatible          bool
		AppName               string
		FileType              int
		CanSubscribe          bool
		Favorited             int
		Followers             int
		LifetimeSubscriptions int
		LifetimeFavorited     int
		LifetimeFollowers     int
		Views                 int
		SpoilerTag            bool
		NumChildren           int
		NumReports            int
		Previews              []struct {
			PreviewID string
			SortOrder int
			URL       string
			Size      int64
			Filename  string
		}
		Tags []struct {
			Tag       string
			AdminOnly bool
		}
		KVTags []struct {
			Key, Value string
		}
		VoteData struct {
			Score     float32
			VotesUp   int
			VotesDown int
		}
	}
}

//QueryFilesArgs is QueyFiles func's args
type QueryFilesArgs url.Values

//NewQueryFilesArgs is make new QueryFiles args
func NewQueryFilesArgs() QueryFilesArgs {
	q := QueryFilesArgs(url.Values{})
	q.add("return_vote_data", "1")
	q.add("return_tags", "1")
	q.add("return_kv_tags", "1")
	q.add("return_previews", "1")
	q.add("return_children", "1")
	q.add("return_short_description", "1")
	return q
}

func (p QueryFilesArgs) add(key, value string) {
	url.Values(p).Add(key, value)
}

func (p QueryFilesArgs) set(key, value string) {
	url.Values(p).Set(key, value)
}

func (p QueryFilesArgs) setBool(key string, value bool) {
	if value {
		p.set(key, "1")
	} else {
		p.set(key, "0")
	}
}

/*
	query_type, page                           int
	creator_appid, appid                       AppID
	requiredtags, excludedtags                 string
	match_all_tags                             bool
	required_flags, omitted_flags, search_text string
	filetype                                   int
	child_publishedfileid                      uint64
	days                                       int
	include_recent_votes_only                  bool
    return_vote_data                           bool
	return_tags, return_kv_tags                bool
	return_previews, return_children           bool
	return_short_description                   bool
*/

//SetQueryType is add query_type arg
func (p QueryFilesArgs) SetQueryType(queryType int) {
	p.set("query_type", strconv.Itoa(queryType))
}

//SetPage is add page arg
func (p QueryFilesArgs) SetPage(page int) {
	p.set("page", strconv.Itoa(page))
}

//SetAppID is add appid arg
func (p QueryFilesArgs) SetAppID(appid AppID) {
	p.set("appid", appid.String())
}

//SetRequiredTags is add requiredtags arg
//u need to SetQueryType to 1
func (p QueryFilesArgs) SetRequiredTags(tags []string) {
	key := "requiredtags[]"
	p.set(key, tags[0])
	for _, value := range tags[1:] {
		p.add(key, value)
	}
}

//SetRequiredTag is add only solo tag
//u need to SetQueryType to 1
func (p QueryFilesArgs) SetRequiredTag(tag string) {
	p.set("requiredtags[]", tag)
}

//SetMatchAllTags is add match_all_tags arg
func (p QueryFilesArgs) SetMatchAllTags(value bool) {
	p.setBool("match_all_tags", value)
}

//SetNumPerPage is add numperpage arg
func (p QueryFilesArgs) SetNumPerPage(num int) {
	p.set("numperpage", strconv.Itoa(num))
}

//TODO more

//QueryFiles is IPublishedFileService/QueryFiles/v1
func (p *SteamAPI) QueryFiles(args QueryFilesArgs) (*PublishedFileService, error) {
	if args == nil {
		args = NewQueryFilesArgs()
	}

	url := p.genURL("IPublishedFileService", "QueryFiles",
		ver1, url.Values(args))
	body, err := getJSONFromURL(url)
	if err != nil {
		return nil, err
	}
	var r struct{ Response PublishedFileService }
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.Response, nil
}
