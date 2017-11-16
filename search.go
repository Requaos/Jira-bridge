package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//SearchResponse is the struct for recieving 'Search' data from JIRA via JSON from the RESTapi
type SearchResponse struct {
	Expand string `json:"expand"`
	Issues []struct {
		Expand string `json:"expand"`
		Fields struct {
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"aggregateprogress"`
			Aggregatetimeestimate         interface{}   `json:"aggregatetimeestimate"`
			Aggregatetimeoriginalestimate int           `json:"aggregatetimeoriginalestimate"`
			Aggregatetimespent            interface{}   `json:"aggregatetimespent"`
			Assignee                      interface{}   `json:"assignee"`
			Components                    []interface{} `json:"components"`
			Created                       string        `json:"created"`
			Creator                       struct {
				Active     bool `json:"active"`
				AvatarUrls struct {
					One6x16   string `json:"16x16"`
					Two4x24   string `json:"24x24"`
					Three2x32 string `json:"32x32"`
					Four8x48  string `json:"48x48"`
				} `json:"avatarUrls"`
				DisplayName  string `json:"displayName"`
				EmailAddress string `json:"emailAddress"`
				Key          string `json:"key"`
				Name         string `json:"name"`
				Self         string `json:"self"`
				TimeZone     string `json:"timeZone"`
			} `json:"creator"`
			Customfield10000 string      `json:"customfield_10000"`
			Customfield10001 string      `json:"customfield_10001"`
			Customfield10002 interface{} `json:"customfield_10002"`
			Customfield10003 interface{} `json:"customfield_10003"`
			Customfield10004 interface{} `json:"customfield_10004"`
			Customfield10005 interface{} `json:"customfield_10005"`
			Customfield10006 string      `json:"customfield_10006"`
			Customfield10007 []string    `json:"customfield_10007"`
			Customfield10008 string      `json:"customfield_10008"`
			Customfield10100 interface{} `json:"customfield_10100"`
			Customfield10200 string      `json:"customfield_10200"`
			Customfield10300 interface{} `json:"customfield_10300"`
			Customfield10400 interface{} `json:"customfield_10400"`
			Customfield10501 interface{} `json:"customfield_10501"`
			Customfield10600 interface{} `json:"customfield_10600"`
			Customfield10601 []struct {
				ID    string `json:"id"`
				Self  string `json:"self"`
				Value string `json:"value"`
			} `json:"customfield_10601"`
			Customfield10700 interface{} `json:"customfield_10700"`
			Customfield10701 interface{} `json:"customfield_10701"`
			Customfield10800 string      `json:"customfield_10800"`
			Customfield10900 []struct {
				ID    string `json:"id"`
				Self  string `json:"self"`
				Value string `json:"value"`
			} `json:"customfield_10900"`
			Customfield11100 interface{}   `json:"customfield_11100"`
			Customfield11300 interface{}   `json:"customfield_11300"`
			Customfield11400 string        `json:"customfield_11400"`
			Customfield11500 interface{}   `json:"customfield_11500"`
			Customfield11600 string        `json:"customfield_11600"`
			Customfield11601 interface{}   `json:"customfield_11601"`
			Customfield11602 []interface{} `json:"customfield_11602"`
			Customfield11603 interface{}   `json:"customfield_11603"`
			Customfield11604 interface{}   `json:"customfield_11604"`
			Customfield11605 interface{}   `json:"customfield_11605"`
			Customfield11606 interface{}   `json:"customfield_11606"`
			Customfield11901 struct {
				ID    string `json:"id"`
				Self  string `json:"self"`
				Value string `json:"value"`
			} `json:"customfield_11901"`
			Customfield11903 struct {
				ID    string `json:"id"`
				Self  string `json:"self"`
				Value string `json:"value"`
			} `json:"customfield_11903"`
			Customfield12100 string      `json:"customfield_12100"`
			Description      string      `json:"description"`
			Duedate          interface{} `json:"duedate"`
			Environment      interface{} `json:"environment"`
			FixVersions      []struct {
				Archived    bool   `json:"archived"`
				Description string `json:"description"`
				ID          string `json:"id"`
				Name        string `json:"name"`
				Released    bool   `json:"released"`
				Self        string `json:"self"`
			} `json:"fixVersions"`
			Issuelinks []interface{} `json:"issuelinks"`
			Issuetype  struct {
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				ID          string `json:"id"`
				Name        string `json:"name"`
				Self        string `json:"self"`
				Subtask     bool   `json:"subtask"`
			} `json:"issuetype"`
			Labels     []interface{} `json:"labels"`
			LastViewed string        `json:"lastViewed"`
			Parent     struct {
				Fields struct {
					Issuetype struct {
						AvatarID    int    `json:"avatarId"`
						Description string `json:"description"`
						IconURL     string `json:"iconUrl"`
						ID          string `json:"id"`
						Name        string `json:"name"`
						Self        string `json:"self"`
						Subtask     bool   `json:"subtask"`
					} `json:"issuetype"`
					Priority struct {
						IconURL string `json:"iconUrl"`
						ID      string `json:"id"`
						Name    string `json:"name"`
						Self    string `json:"self"`
					} `json:"priority"`
					Status struct {
						Description    string `json:"description"`
						IconURL        string `json:"iconUrl"`
						ID             string `json:"id"`
						Name           string `json:"name"`
						Self           string `json:"self"`
						StatusCategory struct {
							ColorName string `json:"colorName"`
							ID        int    `json:"id"`
							Key       string `json:"key"`
							Name      string `json:"name"`
							Self      string `json:"self"`
						} `json:"statusCategory"`
					} `json:"status"`
					Summary string `json:"summary"`
				} `json:"fields"`
				ID   string `json:"id"`
				Key  string `json:"key"`
				Self string `json:"self"`
			} `json:"parent"`
			Priority struct {
				IconURL string `json:"iconUrl"`
				ID      string `json:"id"`
				Name    string `json:"name"`
				Self    string `json:"self"`
			} `json:"priority"`
			Progress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"progress"`
			Project struct {
				AvatarUrls struct {
					One6x16   string `json:"16x16"`
					Two4x24   string `json:"24x24"`
					Three2x32 string `json:"32x32"`
					Four8x48  string `json:"48x48"`
				} `json:"avatarUrls"`
				ID              string `json:"id"`
				Key             string `json:"key"`
				Name            string `json:"name"`
				ProjectCategory struct {
					Description string `json:"description"`
					ID          string `json:"id"`
					Name        string `json:"name"`
					Self        string `json:"self"`
				} `json:"projectCategory"`
				Self string `json:"self"`
			} `json:"project"`
			Reporter struct {
				Active     bool `json:"active"`
				AvatarUrls struct {
					One6x16   string `json:"16x16"`
					Two4x24   string `json:"24x24"`
					Three2x32 string `json:"32x32"`
					Four8x48  string `json:"48x48"`
				} `json:"avatarUrls"`
				DisplayName  string `json:"displayName"`
				EmailAddress string `json:"emailAddress"`
				Key          string `json:"key"`
				Name         string `json:"name"`
				Self         string `json:"self"`
				TimeZone     string `json:"timeZone"`
			} `json:"reporter"`
			Resolution     interface{} `json:"resolution"`
			Resolutiondate interface{} `json:"resolutiondate"`
			Status         struct {
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				ID             string `json:"id"`
				Name           string `json:"name"`
				Self           string `json:"self"`
				StatusCategory struct {
					ColorName string `json:"colorName"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					Name      string `json:"name"`
					Self      string `json:"self"`
				} `json:"statusCategory"`
			} `json:"status"`
			Subtasks []struct {
				Fields struct {
					Issuetype struct {
						AvatarID    int    `json:"avatarId"`
						Description string `json:"description"`
						IconURL     string `json:"iconUrl"`
						ID          string `json:"id"`
						Name        string `json:"name"`
						Self        string `json:"self"`
						Subtask     bool   `json:"subtask"`
					} `json:"issuetype"`
					Priority struct {
						IconURL string `json:"iconUrl"`
						ID      string `json:"id"`
						Name    string `json:"name"`
						Self    string `json:"self"`
					} `json:"priority"`
					Status struct {
						Description    string `json:"description"`
						IconURL        string `json:"iconUrl"`
						ID             string `json:"id"`
						Name           string `json:"name"`
						Self           string `json:"self"`
						StatusCategory struct {
							ColorName string `json:"colorName"`
							ID        int    `json:"id"`
							Key       string `json:"key"`
							Name      string `json:"name"`
							Self      string `json:"self"`
						} `json:"statusCategory"`
					} `json:"status"`
					Summary string `json:"summary"`
				} `json:"fields"`
				ID   string `json:"id"`
				Key  string `json:"key"`
				Self string `json:"self"`
			} `json:"subtasks"`
			Summary              string        `json:"summary"`
			Timeestimate         interface{}   `json:"timeestimate"`
			Timeoriginalestimate interface{}   `json:"timeoriginalestimate"`
			Timespent            interface{}   `json:"timespent"`
			Updated              string        `json:"updated"`
			Versions             []interface{} `json:"versions"`
			Votes                struct {
				HasVoted bool   `json:"hasVoted"`
				Self     string `json:"self"`
				Votes    int    `json:"votes"`
			} `json:"votes"`
			Watches struct {
				IsWatching bool   `json:"isWatching"`
				Self       string `json:"self"`
				WatchCount int    `json:"watchCount"`
			} `json:"watches"`
			Workratio int `json:"workratio"`
		} `json:"fields"`
		ID   string `json:"id"`
		Key  string `json:"key"`
		Self string `json:"self"`
	} `json:"issues"`
	MaxResults int `json:"maxResults"`
	StartAt    int `json:"startAt"`
	Total      int `json:"total"`
}

func getSearchContents(settings ServerSettings, session string, search string, client *http.Client, starting string, length string) SearchResponse {
	resturl := search + length + starting
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
	checkError("Error creating http request during getSearchContents() call", err)
	req.Header.Set("Cookie", session)
	req.Header.Set("X-Ausername", settings.userName)
	req.SetBasicAuth(settings.userName, settings.passWord)
	req.Header.Set("Host", settings.serverName)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//fmt.Println(res)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(body)
	var y SearchResponse
	err = json.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	return y
}
