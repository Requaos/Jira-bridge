package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

//Entry issue struct
type Entry struct {
	Activity string `xml:"activity,attr"`
	ID       struct {
		XMLName xml.Name `xml:"id"`
		ID      string   `xml:",chardata"`
	}
	Title   Title   `xml:"title"`
	Content Content `xml:"content"`
	Author  struct {
		XMLName               xml.Name `xml:"author"`
		Usr                   string   `xml:"usr,attr"`
		LinkAuthorEntry       []Link   `xml:"link"`
		ObjectTypeAuthorEntry string   `xml:"object-type"`
		Name                  string   `xml:"name"`
		Username              string   `xml:"username"`
		Email                 string   `xml:"email"`
		URI                   string   `xml:"uri"`
	}
	Published      string    `xml:"published"`
	UpdatedEntry   string    `xml:"updated"`
	Category       Category  `xml:"category"`
	Link           []Link    `xml:"link"`
	Generator      Generator `xml:"generator"`
	Application    string    `xml:"application"`
	Verb           string    `xml:"verb"`
	TimezoneOffset string    `xml:"timezone-offset"`
	Object         struct {
		XMLName xml.Name `xml:"object"`
		ID      struct {
			XMLName xml.Name `xml:"id"`
			ID      string   `xml:",chardata"`
		}
		Title           Title   `xml:"title"`
		Summary         Summary `xml:"summary"`
		LinkObjectEntry Link    `xml:"link"`
		ObjectType      string  `xml:"object-type"`
	}
}

//IssueResponse is the json datacontainer for JIRA's Issue rest api response
type IssueResponse struct {
	Expand string `json:"expand"`
	Fields struct {
		Aggregateprogress struct {
			Percent  int `json:"percent"`
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Aggregatetimeestimate         int `json:"aggregatetimeestimate"`
		Aggregatetimeoriginalestimate int `json:"aggregatetimeoriginalestimate"`
		Aggregatetimespent            int `json:"aggregatetimespent"`
		Assignee                      struct {
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
		} `json:"assignee"`
		Attachment []struct {
			Author struct {
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
			} `json:"author"`
			Content  string `json:"content"`
			Created  string `json:"created"`
			Filename string `json:"filename"`
			ID       string `json:"id"`
			MimeType string `json:"mimeType"`
			Self     string `json:"self"`
			Size     int    `json:"size"`
		} `json:"attachment"`
		Comment struct {
			Comments []struct {
				Author struct {
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
				} `json:"author"`
				Body         string `json:"body"`
				Created      string `json:"created"`
				ID           string `json:"id"`
				Self         string `json:"self"`
				UpdateAuthor struct {
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
				} `json:"updateAuthor"`
				Updated string `json:"updated"`
			} `json:"comments"`
			MaxResults int `json:"maxResults"`
			StartAt    int `json:"startAt"`
			Total      int `json:"total"`
		} `json:"comment"`
		Components []interface{} `json:"components"`
		Created    string        `json:"created"`
		Creator    struct {
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
		Customfield10001 interface{} `json:"customfield_10001"`
		Customfield10002 interface{} `json:"customfield_10002"`
		Customfield10003 interface{} `json:"customfield_10003"`
		Customfield10004 interface{} `json:"customfield_10004"`
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
		Customfield10700 interface{}   `json:"customfield_10700"`
		Customfield10701 interface{}   `json:"customfield_10701"`
		Customfield10800 string        `json:"customfield_10800"`
		Customfield10900 interface{}   `json:"customfield_10900"`
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
		Customfield11904 interface{} `json:"customfield_11904"`
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
		Issuelinks []struct {
			ID           string `json:"id"`
			OutwardIssue struct {
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
			} `json:"outwardIssue"`
			Self string `json:"self"`
			Type struct {
				ID      string `json:"id"`
				Inward  string `json:"inward"`
				Name    string `json:"name"`
				Outward string `json:"outward"`
				Self    string `json:"self"`
			} `json:"type"`
		} `json:"issuelinks"`
		Issuetype struct {
			AvatarID    int    `json:"avatarId"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Self        string `json:"self"`
			Subtask     bool   `json:"subtask"`
		} `json:"issuetype"`
		Labels     []string `json:"labels"`
		LastViewed string   `json:"lastViewed"`
		Priority   struct {
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
		Resolution struct {
			Description string `json:"description"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Self        string `json:"self"`
		} `json:"resolution"`
		Resolutiondate string `json:"resolutiondate"`
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
		Summary              string      `json:"summary"`
		Timeestimate         int         `json:"timeestimate"`
		Timeoriginalestimate interface{} `json:"timeoriginalestimate"`
		Timespent            interface{} `json:"timespent"`
		Timetracking         struct {
			RemainingEstimate        string `json:"remainingEstimate"`
			RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
		} `json:"timetracking"`
		Updated  string        `json:"updated"`
		Versions []interface{} `json:"versions"`
		Votes    struct {
			HasVoted bool   `json:"hasVoted"`
			Self     string `json:"self"`
			Votes    int    `json:"votes"`
		} `json:"votes"`
		Watches struct {
			IsWatching bool   `json:"isWatching"`
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
		} `json:"watches"`
		Worklog struct {
			MaxResults int           `json:"maxResults"`
			StartAt    int           `json:"startAt"`
			Total      int           `json:"total"`
			Worklogs   []interface{} `json:"worklogs"`
		} `json:"worklog"`
		Workratio int `json:"workratio"`
	} `json:"fields"`
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

//IssueBranchResponse is the branch response
type IssueBranchResponse struct {
	Errors []interface{} `json:"errors"`
	Detail []struct {
		Branches []struct {
			Name                 string `json:"name"`
			URL                  string `json:"url"`
			CreatePullRequestURL string `json:"createPullRequestUrl"`
			Repository           struct {
				Name   string `json:"name"`
				Avatar string `json:"avatar"`
				URL    string `json:"url"`
			} `json:"repository"`
		} `json:"branches"`
		PullRequests []interface{} `json:"pullRequests"`
		Repositories []interface{} `json:"repositories"`
		Instance     struct {
			SingleInstance bool   `json:"singleInstance"`
			BaseURL        string `json:"baseUrl"`
			Name           string `json:"name"`
			TypeName       string `json:"typeName"`
			ID             string `json:"id"`
			Type           string `json:"type"`
		} `json:"_instance"`
	} `json:"detail"`
}

//Worklog is a struct for recieving the worklog chain in an issue
type Worklog struct {
	MaxResults int `json:"maxResults"`
	StartAt    int `json:"startAt"`
	Total      int `json:"total"`
	Worklogs   []struct {
		Author struct {
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
		} `json:"author"`
		Comment          string `json:"comment"`
		Created          string `json:"created"`
		ID               string `json:"id"`
		IssueID          string `json:"issueId"`
		Self             string `json:"self"`
		Started          string `json:"started"`
		TimeSpent        string `json:"timeSpent"`
		TimeSpentSeconds int    `json:"timeSpentSeconds"`
		UpdateAuthor     struct {
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
		} `json:"updateAuthor"`
		Updated string `json:"updated"`
	} `json:"worklogs"`
}

//EpicIssueResponse is how we grab the epic name
type EpicIssueResponse struct {
	Expand string `json:"expand"`
	Fields struct {
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Aggregatetimeestimate         interface{}   `json:"aggregatetimeestimate"`
		Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
		Aggregatetimespent            interface{}   `json:"aggregatetimespent"`
		Assignee                      interface{}   `json:"assignee"`
		Attachment                    []interface{} `json:"attachment"`
		Comment                       struct {
			Comments   []interface{} `json:"comments"`
			MaxResults int           `json:"maxResults"`
			StartAt    int           `json:"startAt"`
			Total      int           `json:"total"`
		} `json:"comment"`
		Components []interface{} `json:"components"`
		Created    string        `json:"created"`
		Creator    struct {
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
		Customfield10000 interface{} `json:"customfield_10000"`
		Customfield10001 interface{} `json:"customfield_10001"`
		Customfield10002 interface{} `json:"customfield_10002"`
		Customfield10003 interface{} `json:"customfield_10003"`
		Customfield10004 interface{} `json:"customfield_10004"`
		Customfield10005 interface{} `json:"customfield_10005"`
		Customfield10006 string      `json:"customfield_10006"`
		Customfield10007 interface{} `json:"customfield_10007"`
		Customfield10008 interface{} `json:"customfield_10008"`
		Customfield10009 string      `json:"customfield_10009"`
		Customfield10010 struct {
			ID    string `json:"id"`
			Self  string `json:"self"`
			Value string `json:"value"`
		} `json:"customfield_10010"`
		Customfield10011 string        `json:"customfield_10011"`
		Customfield10100 interface{}   `json:"customfield_10100"`
		Customfield10200 string        `json:"customfield_10200"`
		Customfield10300 interface{}   `json:"customfield_10300"`
		Customfield10400 interface{}   `json:"customfield_10400"`
		Customfield10501 interface{}   `json:"customfield_10501"`
		Customfield10600 interface{}   `json:"customfield_10600"`
		Customfield10601 interface{}   `json:"customfield_10601"`
		Customfield10700 interface{}   `json:"customfield_10700"`
		Customfield10701 interface{}   `json:"customfield_10701"`
		Customfield10800 interface{}   `json:"customfield_10800"`
		Customfield10900 interface{}   `json:"customfield_10900"`
		Customfield11100 interface{}   `json:"customfield_11100"`
		Customfield11300 interface{}   `json:"customfield_11300"`
		Customfield11400 interface{}   `json:"customfield_11400"`
		Customfield11500 interface{}   `json:"customfield_11500"`
		Customfield11600 interface{}   `json:"customfield_11600"`
		Customfield11601 interface{}   `json:"customfield_11601"`
		Customfield11602 []interface{} `json:"customfield_11602"`
		Customfield11603 interface{}   `json:"customfield_11603"`
		Customfield11604 interface{}   `json:"customfield_11604"`
		Customfield11605 interface{}   `json:"customfield_11605"`
		Customfield11606 interface{}   `json:"customfield_11606"`
		Customfield11800 interface{}   `json:"customfield_11800"`
		Customfield12100 string        `json:"customfield_12100"`
		Description      interface{}   `json:"description"`
		Duedate          interface{}   `json:"duedate"`
		Environment      interface{}   `json:"environment"`
		FixVersions      []interface{} `json:"fixVersions"`
		Issuelinks       []interface{} `json:"issuelinks"`
		Issuetype        struct {
			AvatarID    int    `json:"avatarId"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Self        string `json:"self"`
			Subtask     bool   `json:"subtask"`
		} `json:"issuetype"`
		Labels     []interface{} `json:"labels"`
		LastViewed string        `json:"lastViewed"`
		Priority   struct {
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
		Timetracking         struct{}      `json:"timetracking"`
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
		Worklog struct {
			MaxResults int           `json:"maxResults"`
			StartAt    int           `json:"startAt"`
			Total      int           `json:"total"`
			Worklogs   []interface{} `json:"worklogs"`
		} `json:"worklog"`
		Workratio int `json:"workratio"`
	} `json:"fields"`
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

func getIssueContents(settings ServerSettings, session string, search string, client *http.Client) IssueResponse {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
	checkError("Error creating new http request during getIssueContents() call", err)
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
	var y IssueResponse
	err = json.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	return y
}

func getBranchIssueContents(settings ServerSettings, session string, search string, client *http.Client) IssueBranchResponse {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
	checkError("Error creating new http request during getBranchIssueContents() call", err)
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
	var y IssueBranchResponse
	err = json.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	return y
}

func getEpicIssueContents(settings ServerSettings, session string, search string, client *http.Client) EpicIssueResponse {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
	checkError("Error creating new http request during getEpicIssueContents() call", err)
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
	var y EpicIssueResponse
	err = json.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	return y
}

func getIssueWorklog(settings ServerSettings, session string, search string, client *http.Client) Worklog {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
	checkError("Error creating new http request during getIssueWorklog() call", err)
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
	var y Worklog
	err = json.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	return y
}
