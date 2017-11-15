package main

//Written by Neil Skinner April 2016
import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

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

//FilterResponse is the struct for recieving 'Filter' data from JIRA via JSON from the RESTapi
type FilterResponse struct {
	Description string `json:"description"`
	Favourite   bool   `json:"favourite"`
	ID          string `json:"id"`
	Jql         string `json:"jql"`
	Name        string `json:"name"`
	Owner       struct {
		Active     bool `json:"active"`
		AvatarUrls struct {
			One6x16   string `json:"16x16"`
			Two4x24   string `json:"24x24"`
			Three2x32 string `json:"32x32"`
			Four8x48  string `json:"48x48"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Key         string `json:"key"`
		Name        string `json:"name"`
		Self        string `json:"self"`
	} `json:"owner"`
	SearchURL        string `json:"searchUrl"`
	Self             string `json:"self"`
	SharePermissions []struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"sharePermissions"`
	SharedUsers struct {
		Endindex   int           `json:"end-index"`
		Items      []interface{} `json:"items"`
		Maxresults int           `json:"max-results"`
		Size       int           `json:"size"`
		Startindex int           `json:"start-index"`
	} `json:"sharedUsers"`
	Subscriptions struct {
		Endindex   int           `json:"end-index"`
		Items      []interface{} `json:"items"`
		Maxresults int           `json:"max-results"`
		Size       int           `json:"size"`
		Startindex int           `json:"start-index"`
	} `json:"subscriptions"`
	ViewURL string `json:"viewUrl"`
}

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

//UserWithGroups is the response from JIRA of the user API with 'groups' expanded
type UserWithGroups struct {
	Active           bool `json:"active"`
	ApplicationRoles struct {
		Items []interface{} `json:"items"`
		Size  int           `json:"size"`
	} `json:"applicationRoles"`
	AvatarUrls struct {
		One6x16   string `json:"16x16"`
		Two4x24   string `json:"24x24"`
		Three2x32 string `json:"32x32"`
		Four8x48  string `json:"48x48"`
	} `json:"avatarUrls"`
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
	Expand       string `json:"expand"`
	Groups       struct {
		Items []struct {
			Name string `json:"name"`
			Self string `json:"self"`
		} `json:"items"`
		Size int `json:"size"`
	} `json:"groups"`
	Key      string `json:"key"`
	Locale   string `json:"locale"`
	Name     string `json:"name"`
	Self     string `json:"self"`
	TimeZone string `json:"timeZone"`
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

//LoginResponse is the struct for recieving 'Auth Cookie' data from JIRA via JSON from the RESTapi
type LoginResponse struct {
	LoginInfo struct {
		FailedLoginCount    int    `json:"failedLoginCount"`
		LastFailedLoginTime string `json:"lastFailedLoginTime"`
		LoginCount          int    `json:"loginCount"`
		PreviousLoginTime   string `json:"previousLoginTime"`
	} `json:"loginInfo"`
	Session struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"session"`
}

//SessionID is a struct for BL auth creds
type SessionID struct {
	XMLName   xml.Name `xml:"resp"`
	Status    string   `xml:"status,attr"`
	SessionID string   `xml:"sessionId"`
}

//Logoutres is the basic XML response from the BL logout api command
type Logoutres struct {
	XMLName xml.Name `xml:"resp"`
	Status  string   `xml:"status,attr"`
}

//Queryresponse is a struct for BL selectQuery
type Queryresponse struct {
	Status string     `xml:"status,attr"`
	Col    [][]string `xml:"row>col"`
}

//QuerySingleresponse is a struct for BL selectQuery
type QuerySingleresponse struct {
	XMLName xml.Name `xml:"resp"`
	Rows    struct {
		XMLName xml.Name `xml:"row"`
		Cols    struct {
			XMLName xml.Name `xml:"col"`
			Col     string   `xml:",chardata"`
		}
	}
}

//DataField is a struct for getting names of BL objects via ID
type DataField struct {
	XMLName xml.Name `xml:"resp"`
	Status  string   `xml:"status,attr"`
	Field   struct {
		XMLName xml.Name `xml:"field"`
		Name    string   `xml:"name,attr"`
		Field   string   `xml:",chardata"`
	}
}

//Updateresp is a response struct for status of a BL object update
type Updateresp struct {
	XMLName xml.Name `xml:"resp"`
	Status  string   `xml:"status,attr"`
	Msg     struct {
		XMLName xml.Name `xml:"Msg"`
		Msg     string   `xml:",chardata"`
	}
}

//TriggerResponse is a response struct for status of the runTrigger() BL object command
type TriggerResponse struct {
	XMLName xml.Name `xml:"resp"`
	Status  string   `xml:"status,attr"`
	Error   struct {
		XMLName xml.Name `xml:"err"`
		Error   string   `xml:",chardata"`
	}
}

//Feed is the RSS activity feed struct
type Feed struct {
	Xmlns     string `xml:"xmlns,attr"`
	Atlassian string `xml:"atlassian,attr"`
	ID        struct {
		XMLName xml.Name `xml:"id"`
		ID      string   `xml:",chardata"`
	}
	Title          Title   `xml:"title"`
	TimezoneOffset string  `xml:"timezone-offset"`
	Updated        string  `xml:"updated"`
	Link           Link    `xml:"link"`
	Entry          []Entry `xml:"entry"`
}

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

//Content rss feed struct
type Content struct {
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
}

//Div html thingy
type Div struct {
	Class string `xml:"class,attr"`
	P     P      `xml:"p"`
}

//P html thingy
type P struct {
	P string `xml:",chardata"`
}

//Ul html thingy
type Ul struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}

//TitleObjectEntry rss feed stuff
type TitleObjectEntry struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

//Summary rss feed stuff
type Summary struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

//LinkObjectEntry rss feed stuff
type LinkObjectEntry struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

//Generator rss feed stuff
type Generator struct {
	URI string `xml:"uri,attr"`
}

//Title rss feed stuff
type Title struct {
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
	A    []struct {
		XMLName xml.Name `xml:"a"`
		Href    string   `xml:"href,attr"`
		Class   string   `xml:"class,attr"`
		Text    string   `xml:",chardata"`
	}
}

//Link rss feed stuff
type Link struct {
	Title string `xml:"title,attr"`
	Href  string `xml:"href,attr"`
	Rel   string `xml:"rel,attr"`
}

//LinkAuthorEntry rss feed stuff
type LinkAuthorEntry struct {
	Rel    string `xml:"rel,attr"`
	Href   string `xml:"href,attr"`
	Height string `xml:"height,attr"`
	Width  string `xml:"width,attr"`
	Media  string `xml:"media,attr"`
}

//Category rss feed stuff
type Category struct {
	Term string `xml:"term,attr"`
}
type billingType struct {
	timeInt, keyStr, whoStr, billStr, totalStr, serviceBureauStr, summaryStr string
	exp, rate                                                                int
}

//ServerSettings is a container for all of the configurable email servrer settings
type ServerSettings struct {
	serverName string
	userName   string
	passWord   string
	blTenantID string
	blUser     string
	blPass     string
	pRate      int
	tRate      int
	dRate      int
	mRate      int
}

func main() {
	//creating a commandline option to change the filterid, but declaring the default to be AJ's 'SNAP NEXT' filter
	//filterid in jira can be seen at the end of the url of the filter
	filterid := flag.Int("f", 18800, "JIRA Filter ID")
	jiraCreds := getSettings()
	//maxSearchItemsReturned is the maximum number of items to return in a JIRA search
	maxSearchItemsReturned := flag.Int("max", 201, "Maximum search items to return in JIRA query")
	file := flag.String("file", "results", "filename for the output")
	closem := flag.Bool("close", false, "wether or not to close BL cases")
	helper := flag.Bool("?", false, "Prints the syntax helper")
	whotime := flag.Bool("time", false, "gets 'Time Estimate' history")
	branches := flag.Bool("branch", false, "gets branch info")
	flag.Parse()
	if *helper != true {
		var Maxres = "&maxResults=" + strconv.Itoa(*maxSearchItemsReturned)
		httpclient := &http.Client{}
		filename := *file + ".csv"
		timefilename := "activity.csv"
		var sessionCookie string
		jiraFilterURL := "https://evolutionpayroll.atlassian.net/rest/api/2/filter/" + strconv.Itoa(*filterid)
		loginclient := jiralogin(jiraCreds, httpclient).Session
		sessionCookie = loginclient.Name + "=" + loginclient.Value
		var rollbaseSessionKey string
		if *closem != true {
			// section for grabing change_ids from releasenote in BL, limited by &maxrows param in rollbaseSelectQuery(), have not yet implemented a loop to grab all
			rollbaseSessionKey = rollbaselogin(jiraCreds.blUser, jiraCreds.blPass, jiraCreds.blTenantID)
		}
		// fmt.Println(sessionCookie)
		// changeIdqueryresult := rollbaseSelectQuery(rollbaseSessionKey, "select%20change_id%20from%20release_note%20order%20by%20release_date%20desc")
		// var f map[string]string
		// f = make(map[string]string)
		// //fmt.Print("\n[Grabbing Defects]Working")
		// for d := range changeIdqueryresult.Rows {
		// 	for e := range changeIdqueryresult.Rows[d].Cols {
		// 		fieldname := rollbaseGetDataField(rollbaseSessionKey, "name", changeIdqueryresult.Rows[d].Cols[e].Col)
		// 		f[fieldname.Field.Field] = changeIdqueryresult.Rows[d].Cols[e].Col
		// 		//fmt.Println("ChangeId:", fieldname.Field.Field, " ChangeId#Id:", changeIdqueryresult.Rows[d].Cols[e].Col)
		// 		fmt.Print(".")
		// 	}
		// }
		var ChangeIDExists bool
		var changeidmap map[string]string
		var val string
		var jiraRSSactivityFeed string
		var jiraIssueWorklog string
		if *whotime != false {
			jiraRSSactivityFeed = "https://evolutionpayroll.atlassian.net/activity?maxResults=20&streams=issue-key+IS+"
			jiraIssueWorklog = "https://evolutionpayroll.atlassian.net/rest/api/2/issue/"
		}
		if *closem != false {
			changeIdqueryresult := rollbaseSelectQuery(rollbaseSessionKey, "select%20id,%20name%20from%20change_id%20order%20by%20updatedAt%20desc")
			changeidmap = getchangeIDsfromRollbase(rollbaseSessionKey, changeIdqueryresult)
		}
		filtersearchstring := getFilterContents(jiraCreds, sessionCookie, httpclient, jiraFilterURL)
		var data [][]string
		if *whotime != false {
			data = [][]string{{"Jira Issue", "Priority", "BL#", "Reso#", "Service Bureau", "Release Note", "Hours", "Billable"}}
		} else if *branches != false {
			data = [][]string{{"Jira Issue", "Priority", "BL#", "Reso#", "Epic", "Fix Verion", "Release Note", "Banch Name"}}
		} else {
			data = [][]string{{"Jira Issue", "Priority", "BL#", "Reso#", "Epic", "Fix Verion", "Release Note"}}
		}
		var timedata = [][]string{{"Jira Issue", "Change Content", "Whom"}}
		var serviceBureaus []string
		var worklogdata []billingType
		var precalc float64
		var totalissues int
		precalc = float64(getSearchContents(jiraCreds, sessionCookie, filtersearchstring.SearchURL, httpclient, "&startAt=0", Maxres).Total) / float64(*maxSearchItemsReturned)
		if precalc == float64(int(precalc)) {
			totalissues = int(precalc)
		} else {
			totalissues = 1 + int(precalc)
		}
		for l := 0; l < totalissues; l++ {
			startingat := "&startAt=" + strconv.Itoa(l**maxSearchItemsReturned)
			searchresp := getSearchContents(jiraCreds, sessionCookie, filtersearchstring.SearchURL, httpclient, startingat, Maxres)
			for i := range searchresp.Issues {
				var totalTimeSpent int
				var totalDueMoneys int
				var branchName IssueBranchResponse
				var epicname EpicIssueResponse
				if *closem != false {
					val, ChangeIDExists = changeidmap[searchresp.Issues[i].Key]
					fmt.Println(searchresp.Issues[i].Key + " ID: " + val)
				}
				if searchresp.Issues[i].Fields.Customfield10008 != "" {
					epicsearch := "https://evolutionpayroll.atlassian.net/rest/api/2/issue/" + searchresp.Issues[i].Fields.Customfield10008
					epicname = getEpicIssueContents(jiraCreds, sessionCookie, epicsearch, httpclient)
				} else {
					epicname.Fields.Customfield10009 = "No Epic Link"
				}
				if *branches != false {
					branchesEndpoint := "https://evolutionpayroll.atlassian.net/rest/dev-status/latest/issue/detail?issueId=" + searchresp.Issues[i].ID + "&applicationType=bitbucket&dataType=branch"
					branchName = getBranchIssueContents(jiraCreds, sessionCookie, branchesEndpoint, httpclient)
				}
				var totalTimeSpentinMinutes time.Duration
				if *whotime != false {
					project := strings.Split(searchresp.Issues[i].Key, "-")
					projectname := project[0]
					billable := searchresp.Issues[i].Fields.Customfield11901
					issueWorklog := jiraIssueWorklog + searchresp.Issues[i].Key + "/worklog"
					workLog := getIssueWorklog(jiraCreds, sessionCookie, issueWorklog, httpclient)
					for _, perworklog := range workLog.Worklogs {
						var worklogdatablock billingType
						worklogdatablock.serviceBureauStr = searchresp.Issues[i].Fields.Customfield11903.Value
						worklogdatablock.keyStr = searchresp.Issues[i].Key
						worklogdatablock.billStr = billable.Value
						worklogdatablock.whoStr = perworklog.Author.Name
						payGroups := getUserGroups(jiraCreds, sessionCookie, "https://evolutionpayroll.atlassian.net/rest/api/2/user?username="+perworklog.Author.Name+"&expand=groups", httpclient)
						if PayGroupContains(payGroups, "ManagerConsult") {
							worklogdatablock.rate = jiraCreds.mRate
						} else if PayGroupContains(payGroups, "CustomProgramming") {
							worklogdatablock.rate = jiraCreds.pRate
						} else if PayGroupContains(payGroups, "TechProgSupport") {
							worklogdatablock.rate = jiraCreds.tRate
						} else {
							worklogdatablock.rate = jiraCreds.dRate
						}
						worklogdatablock.timeInt = perworklog.TimeSpent
						worklogdatablock.exp = (worklogdatablock.rate * perworklog.TimeSpentSeconds) / 3600
						worklogdata = append(worklogdata, worklogdatablock)
						totalTimeSpent += perworklog.TimeSpentSeconds
						totalDueMoneys += worklogdatablock.exp
					}
					if len(searchresp.Issues[i].Fields.Subtasks) > 0 {
						for _, subTask := range searchresp.Issues[i].Fields.Subtasks {
							subTaskWorklogURL := jiraIssueWorklog + subTask.Key + "/worklog"
							subTaskWorklog := getIssueWorklog(jiraCreds, sessionCookie, subTaskWorklogURL, httpclient)
							for _, perWorklog := range subTaskWorklog.Worklogs {
								var worklogdatablock billingType
								worklogdatablock.serviceBureauStr = searchresp.Issues[i].Fields.Customfield11903.Value
								worklogdatablock.keyStr = searchresp.Issues[i].Key
								worklogdatablock.billStr = billable.Value
								worklogdatablock.whoStr = perWorklog.Author.Name
								payGroups := getUserGroups(jiraCreds, sessionCookie, "https://evolutionpayroll.atlassian.net/rest/api/2/user?username="+perWorklog.Author.Name+"&expand=groups", httpclient)
								if PayGroupContains(payGroups, "ManagerConsult") {
									worklogdatablock.rate = jiraCreds.mRate
								} else if PayGroupContains(payGroups, "CustomProgramming") {
									worklogdatablock.rate = jiraCreds.pRate
								} else if PayGroupContains(payGroups, "TechProgSupport") {
									worklogdatablock.rate = jiraCreds.tRate
								} else {
									worklogdatablock.rate = jiraCreds.dRate
								}
								worklogdatablock.timeInt = perWorklog.TimeSpent
								worklogdatablock.exp = (worklogdatablock.rate * perWorklog.TimeSpentSeconds) / 3600
								worklogdata = append(worklogdata, worklogdatablock)
								totalTimeSpent += perWorklog.TimeSpentSeconds
								totalDueMoneys += worklogdatablock.exp
							}
						}
					}
					issueFeed := jiraRSSactivityFeed + searchresp.Issues[i].Key + "&streams=key+IS+" + projectname + "&os_authType=basic&title=undefined"
					activityhistory := getIssueHistory(jiraCreds, issueFeed, httpclient)
					//fmt.Println(activityhistory)
					for u := range activityhistory.Entry {
						if activityhistory.Entry[u].Content.Text != "" {
							fmt.Println("//")
							fmt.Println(activityhistory.Entry[u].Content.Text)
							fmt.Println("//")
							activitycontents := activityhistory.Entry[u].Content.Text
							activityauthorname := activityhistory.Entry[u].Author.Name
							var timedatablock = []string{searchresp.Issues[i].Key, activitycontents, activityauthorname}
							timedata = append(timedata, timedatablock)
						}
					}

					fmt.Print("Original Estimate:")
					fmt.Println(searchresp.Issues[i].Fields.Aggregatetimeoriginalestimate)
					fmt.Print("Actual Time:")
					fmt.Println(totalTimeSpent)
					if totalTimeSpent > searchresp.Issues[i].Fields.Aggregatetimeoriginalestimate {
						totalTimeSpentinMinutes = time.Duration(time.Duration(searchresp.Issues[i].Fields.Aggregatetimeoriginalestimate) * time.Second)
						fmt.Println("Overage Detected!")
					} else {
						totalTimeSpentinMinutes = time.Duration(time.Duration(totalTimeSpent) * time.Second)
					}
					var worklogdatablockTotal billingType
					worklogdatablockTotal.keyStr = searchresp.Issues[i].Key
					worklogdatablockTotal.summaryStr = searchresp.Issues[i].Fields.Summary
					worklogdatablockTotal.exp = totalDueMoneys
					worklogdatablockTotal.serviceBureauStr = searchresp.Issues[i].Fields.Customfield11903.Value
					worklogdatablockTotal.timeInt = fmt.Sprintf("%13s", Round(totalTimeSpentinMinutes, time.Minute))
					serviceBureaus = append(serviceBureaus, searchresp.Issues[i].Fields.Customfield11903.Value)
					worklogdata = append(worklogdata, worklogdatablockTotal)
				}
				fmt.Println("//")
				fmt.Println(searchresp.Issues[i].Key)
				fmt.Println(searchresp.Issues[i].Fields.Priority.Name)
				fmt.Println(searchresp.Issues[i].Fields.Customfield10800)
				fmt.Println(searchresp.Issues[i].Fields.Customfield11600)
				fmt.Println(epicname.Fields.Customfield10009)
				if len(searchresp.Issues[i].Fields.FixVersions) > 0 {
					fmt.Println(searchresp.Issues[i].Fields.FixVersions[0].Name)
				} else {
					fmt.Println("No FixVersion")
				}
				fmt.Println(searchresp.Issues[i].Fields.Customfield11400)
				fmt.Println("//")
				var datablock []string
				if *whotime != false {
					datablock = []string{searchresp.Issues[i].Key, searchresp.Issues[i].Fields.Priority.Name, searchresp.Issues[i].Fields.Customfield10800, searchresp.Issues[i].Fields.Customfield11600, searchresp.Issues[i].Fields.Customfield11903.Value, searchresp.Issues[i].Fields.Customfield11400, fmt.Sprintf("%13s", Round(totalTimeSpentinMinutes, time.Minute)), searchresp.Issues[i].Fields.Customfield11901.Value}
				} else if *branches != false {
					if len(branchName.Detail) > 0 {
						if len(branchName.Detail[0].Branches) > 0 {
							datablock = []string{searchresp.Issues[i].Key, searchresp.Issues[i].Fields.Priority.Name, searchresp.Issues[i].Fields.Customfield10800, searchresp.Issues[i].Fields.Customfield11600, epicname.Fields.Customfield10009, "No FixVersion", searchresp.Issues[i].Fields.Customfield11400, branchName.Detail[0].Branches[0].Name}
						} else {
							datablock = []string{searchresp.Issues[i].Key, searchresp.Issues[i].Fields.Priority.Name, searchresp.Issues[i].Fields.Customfield10800, searchresp.Issues[i].Fields.Customfield11600, epicname.Fields.Customfield10009, "No FixVersion", searchresp.Issues[i].Fields.Customfield11400, "No Branch"}
						}
					}
				} else {
					if len(searchresp.Issues[i].Fields.FixVersions) > 0 {
						if len(searchresp.Issues[i].Fields.FixVersions) > 1 {
							datablock = []string{searchresp.Issues[i].Key, searchresp.Issues[i].Fields.Priority.Name, searchresp.Issues[i].Fields.Customfield10800, searchresp.Issues[i].Fields.Customfield11600, epicname.Fields.Customfield10009, searchresp.Issues[i].Fields.FixVersions[0].Name + "/" + searchresp.Issues[i].Fields.FixVersions[1].Name, searchresp.Issues[i].Fields.Customfield11400}
						} else {
							datablock = []string{searchresp.Issues[i].Key, searchresp.Issues[i].Fields.Priority.Name, searchresp.Issues[i].Fields.Customfield10800, searchresp.Issues[i].Fields.Customfield11600, epicname.Fields.Customfield10009, searchresp.Issues[i].Fields.FixVersions[0].Name, searchresp.Issues[i].Fields.Customfield11400}
						}
					} else {
						datablock = []string{searchresp.Issues[i].Key, searchresp.Issues[i].Fields.Priority.Name, searchresp.Issues[i].Fields.Customfield10800, searchresp.Issues[i].Fields.Customfield11600, epicname.Fields.Customfield10009, "No FixVersion", searchresp.Issues[i].Fields.Customfield11400}
					}
				}
				data = append(data, datablock)
				if *closem != false {
					fmt.Println("Close Flagged!")
					if ChangeIDExists != true {
						fmt.Println("A change_id may not exist for " + searchresp.Issues[i].Key)
					}
					if searchresp.Issues[i].Fields.Customfield10800 != "" {
						fmt.Println("BL# Field is not empty for " + searchresp.Issues[i].Key)
						ifmulticase := strings.Split(searchresp.Issues[i].Fields.Customfield10800, ";")
						fmt.Println("Whole Field: \"" + searchresp.Issues[i].Fields.Customfield10800 + "\" and just in case there are more than one delimited by \";\" ")
						for numbofbl := range ifmulticase {
							fmt.Println("This is each single BL# " + ifmulticase[numbofbl])
							selectcomposedid := "select%20id%20from%20case8%20where%20case_number_5%20=%20\""
							selectcomposedid += ifmulticase[numbofbl]
							selectcomposedid += "\""
							caseid := rollbaseSelectSingleQuery(rollbaseSessionKey, selectcomposedid)
							fmt.Println("This is the caseId before the workflow runs:" + caseid.Rows.Cols.Col)
							rollbaseCloseCaseWorkFlow(rollbaseSessionKey, httpclient, caseid.Rows.Cols.Col, val)
						}
					}
				}
				//section for checking if a Reso# exists in BL that is listed in JIRA for an issue
				//fmt.Println("\n" + searchresp.Issues[i].Key + ":")
				//fmt.Println(searchresp.Issues[i].Self)
				// if searchresp.Issues[i].Fields.Customfield11600 != "" {
				// 	defectnumber := searchresp.Issues[i].Fields.Customfield11600
				// 	val, ok := f[defectnumber]
				// 	if ok != false {
				// 		fmt.Println("got a hit: " + defectnumber)
				// 		fmt.Println(val)
				// 	}
				// }
			}
			fmt.Print("\n\nNumber of issues in this filter: ")
			fmt.Println(searchresp.Total)
			//fmt.Println(searchresp.StartAt)
		}
		if *closem != false {
			rollbaselogout(rollbaseSessionKey)
		}
		if *whotime != false {
			//This block outputs the pdfs
			RemoveDuplicates(&serviceBureaus)
			pdf := gofpdf.New("P", "mm", "A4", "")
			header := []string{"Jira Issue", "Summary", "Time Spent", "Whom"}
			fancyTable := func() {
				w := []float64{25, 70, 40, 40}
				wSum := 0.0
				for _, v := range w {
					wSum += v
				}
				// Colors, line width and bold font
				pdf.SetFillColor(255, 0, 0)
				pdf.SetTextColor(255, 255, 255)
				pdf.SetDrawColor(128, 0, 0)
				pdf.SetLineWidth(.3)
				pdf.SetFont("", "B", 0)
				// 	Header
				for j, str := range header { //This is giving each member of the header it's own cell and cell specifications
					pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
				}
				pdf.Ln(-1)
				// Color and font restoration
				pdf.SetFillColor(224, 235, 255)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetFont("", "", 0)
				// 	Data
				fill := false
				for _, c := range worklogdata {

					pdf.CellFormat(w[0], 6, c.keyStr, "LR", 0, "C", fill, 0, "")
					if len(c.summaryStr) > 33 {
						pdf.CellFormat(w[1], 6, c.summaryStr[:33], "LR", 0, "L", fill, 0, "")
					} else {
						pdf.CellFormat(w[1], 6, c.summaryStr, "LR", 0, "L", fill, 0, "")
					}
					pdf.CellFormat(w[2], 6, c.timeInt, "LR", 0, "C", fill, 0, "")
					pdf.CellFormat(w[3], 6, c.whoStr, "LR", 0, "C", fill, 0, "")
					pdf.Ln(-1)
					fill = !fill

				}
				pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
			}
			pdf.SetFont("Arial", "", 9)
			//This block is for actually outputting the pdf file, it's wonky right now so give it a better format before turning it back on
			pdf.AddPage()
			fancyTable()
			err := pdf.OutputFileAndClose("BureauReport.pdf")
			checkError("Cannot write to file", err)

			timefile, err := os.Create(timefilename)
			checkError("Cannot create file", err)
			defer timefile.Close()

			timewriter := csv.NewWriter(timefile)

			for _, value := range timedata {
				err := timewriter.Write(value)
				checkError("Cannot write to file", err)
			}
			defer timewriter.Flush()
		}
		file, err := os.Create(filename)
		checkError("Cannot create file", err)
		defer file.Close()

		writer := csv.NewWriter(file)

		for _, value := range data {
			err := writer.Write(value)
			checkError("Cannot write to file", err)
		}
		defer writer.Flush()
	} else {
		fmt.Print("\njira-bridge [-f filterID] [-max maxrows] [-close] [-?]\n\n\n")
		fmt.Println("-f      allows you to change the filter this tool pulls from")
		fmt.Print("         Default: 18800 \"SNAP NEXT\"\n\n")
		fmt.Println("-max    allows you to change how many rows are returned per")
		fmt.Println("        search iteration, speed versus # of api calls")
		fmt.Print("         Default: 201\n\n")
		fmt.Print("-time   creates a PDF with the time spent and if it's billable\n\n")
		fmt.Println("-file   allows you to change output filename")
		fmt.Print("         Default: results\n\n")
		fmt.Print("-?     Prints this message block\n\n")
		fmt.Println("-close  allows you to link an existing 'change_id' any BL cases attached")
		fmt.Print("         Default: off\n\n")
		fmt.Println("(1) Example: jira-bridge -f 19015 -close")
		fmt.Print("(2) Example: jira-bridge -f 18515 -max 200 -file snapnextReleased\n\n")
	}
}

func getSettings() ServerSettings {
	var settings ServerSettings
	jiraLoginCreds := "jira-bridge.ini"
	//add a commented example here:
	//line 0:JiraServer=evolutionpayroll.atlassian.net
	//line 1:JiraUsername=
	//line 2:JiraPassword=
	//line 3:BureauLinkTenant=13924598
	//line 4:BureauLinkUser=
	//line 5:BureauLinkPassword=
	//Line 6:ProgrammingRate=
	//Line 7:TechConsultRate=
	//Line 8:DefaultRate=
	//Line 9:ManagerRate=
	if _, err := os.Stat(jiraLoginCreds); err == nil {
		f, _ := os.Open(jiraLoginCreds)
		scanner := bufio.NewScanner(f)
		// Set the Split method to ScanWords.
		scanner.Split(bufio.ScanWords)
		var lines []string
		// Scan all words from the file.
		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
		}
		if lines != nil {
			settings.serverName = lines[0][11:]
			settings.userName = lines[1][13:]
			settings.passWord = lines[2][13:]
			settings.blTenantID = lines[3][17:]
			settings.blUser = lines[4][15:]
			settings.blPass = lines[5][19:]
			settings.pRate, _ = strconv.Atoi(lines[6][16:])
			settings.tRate, _ = strconv.Atoi(lines[7][16:])
			settings.dRate, _ = strconv.Atoi(lines[8][12:])
			settings.mRate, _ = strconv.Atoi(lines[9][12:])
		}
	}
	return settings
}

//PayGroupContains checks if a string is in an array of strings
func PayGroupContains(s UserWithGroups, e string) bool {
	for _, a := range s.Groups.Items {
		if a.Name == e {
			return true
		}
	}
	return false
}

//RemoveDuplicates is a neat little function for removing duplicate strings from the array of strings
func RemoveDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

//Round rounds a time.Duration into a human readable duration ie.- 1hr33m
func Round(d, r time.Duration) time.Duration {
	if r <= 0 {
		return d
	}
	neg := d < 0
	if neg {
		d = -d
	}
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	if neg {
		return -d
	}
	return d
}

func rollbaseCloseCaseWorkFlow(session string, httpclient *http.Client, caseid string, val string) {
	if val != "" {
		msg := rollbaseGetUpdateField(session, caseid, "change_id", val)
		fmt.Println(msg.Status)
	} else {
		fmt.Println("is there a change_id for this JIRA ticket yet?")
		//rollbaseRunTrigger(session, httpclient, "Case", caseid, "autoclosestatus")
		//rollbaseRunTrigger(session, httpclient, "Case", caseid, "autorelease_email")
		//rollbaseRunTrigger(session, httpclient, "Case", caseid, "casecloseddate")
	}
}

func rollbaselogin(login string, password string, custid string) string {
	resturl := "https://apps.bureaulink.com/rest/api/"
	resturl += "login?loginName="
	resturl += login
	resturl += "&password="
	resturl += password
	resturl += "&custID="
	resturl += custid
	res, err := http.Get(resturl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(body)
	var y SessionID
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	//fmt.Println(y)
	return y.SessionID
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func rollbaselogout(session string) {
	resturl := "https://apps.bureaulink.com/rest/api/logout"
	res, err := http.Get(resturl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var y Logoutres
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
}

func rollbaseRunTrigger(session string, client *http.Client, objname string, id string, triggerid string) TriggerResponse {
	resturl := "https://apps.bureaulink.com/rest/api/runtrigger?sessionId="
	resturl += session
	resturl += "&id=" + id
	resturl += "&triggerId="
	resturl += triggerid
	resturl += "&checkValidation=true"
	req, err := http.NewRequest("PUT", resturl, nil)
	req.Header.Set("Content-Type", "application/xml")
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
	fmt.Println(body)
	var y TriggerResponse
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	fmt.Println(y.Status + " " + y.Error.Error)
	//logout(session)
	return y
}

func rollbaseGetDataField(session string, fieldname string, id string) DataField {
	resturl := "https://apps.bureaulink.com/rest/api/getDataField?sessionId="
	resturl += session
	resturl += "&output=xml"
	resturl += "&fieldName="
	resturl += fieldname
	resturl += "&id=" + id
	//fmt.Println(resturl)
	res, err := http.Get(resturl)
	if err != nil {
		panic(err)
	}
	//fmt.Println(res)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var y DataField
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	//logout(session)
	return y
}

func rollbaseSelectQuery(session string, query string) Queryresponse {
	resturl := "https://apps.bureaulink.com/rest/api/selectQuery?sessionId="
	resturl += session
	resturl += "&output=xml&maxRows=201&startRow=0"
	resturl += "&query="
	resturl += query
	res, err := http.Get(resturl)
	if err != nil {
		panic(err)
	}
	//fmt.Println(res)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var y Queryresponse
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	//logout(session)
	return y
}

func rollbaseSelectSingleQuery(session string, query string) QuerySingleresponse {
	resturl := "https://apps.bureaulink.com/rest/api/selectQuery?sessionId="
	resturl += session
	resturl += "&output=xml&maxRows=1&startRow=0"
	resturl += "&query="
	resturl += query
	fmt.Println(resturl)
	res, err := http.Get(resturl)
	if err != nil {
		panic(err)
	}
	//fmt.Println(res)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var y QuerySingleresponse
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	//logout(session)
	return y
}

func getchangeIDsfromRollbase(rollbaseSessionKey string, changeIdqueryresult Queryresponse) map[string]string {
	var f map[string]string
	f = make(map[string]string)
	var name string
	var id string
	fmt.Print("[Grabbing a list of published change_ids from rollbase] Status:" + changeIdqueryresult.Status + "\nWorking.")
	for d := range changeIdqueryresult.Col {
		//columns := len(changeIdqueryresult.Col[d])
		//fmt.Println("# of Columns: " + strconv.Itoa(columns))
		//fmt.Print("BL \"change_id\" ")
		//fmt.Println(d)
		for e := range changeIdqueryresult.Col[d] {
			//fmt.Println(e)
			if d%2 == 0 {
				id = changeIdqueryresult.Col[d][e]
				//fmt.Print("ID# " + id)
			} else {
				name = changeIdqueryresult.Col[d][e]
				f[name] = id
				//fmt.Print("Name: " + name)
			}
			// fieldname := rollbaseGetDataField(rollbaseSessionKey, "name", changeIdqueryresult.Rows[d].Cols[e].Col)
			// f[fieldname.Field.Field] = changeIdqueryresult.Rows[d].Cols[e].Col
			//fmt.Println("ChangeId:", fieldname.Field.Field, " ChangeId#Id:", changeIdqueryresult.Rows[d].Cols[e].Col)
			fmt.Print(".")
		}
	}
	fmt.Println("")
	return f
}

func rollbaseGetUpdateField(session string, id string, fieldname string, fieldvalue string) Updateresp {
	resturl := "https://apps.bureaulink.com/rest/api/update2?sessionId="
	resturl += session
	resturl += "&id=" + id
	resturl += "&useIds=true"
	resturl += "&"
	resturl += fieldname
	resturl += "="
	resturl += fieldvalue

	//fmt.Println(resturl)
	res, err := http.Get(resturl)
	if err != nil {
		panic(err)
	}
	//fmt.Println(res)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var y Updateresp
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	//logout(session)
	return y
}

func getIssueHistory(settings ServerSettings, search string, client *http.Client) Feed {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
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
	var y Feed
	err = xml.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	return y
}

func getUserGroups(settings ServerSettings, session string, search string, client *http.Client) UserWithGroups {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
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
	var y UserWithGroups
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

func getIssueContents(settings ServerSettings, session string, search string, client *http.Client) IssueResponse {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
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

func getSearchContents(settings ServerSettings, session string, search string, client *http.Client, starting string, length string) SearchResponse {
	resturl := search + length + starting
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
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

func getFilterContents(settings ServerSettings, session string, client *http.Client, jiraFilter string) FilterResponse {
	resturl := jiraFilter
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
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
	var y FilterResponse
	err = json.Unmarshal(body, &y)
	if err != nil {
		panic(err)
	}
	return y
}

func jiralogin(settings ServerSettings, client *http.Client) LoginResponse { //This was for cookie based auth, but I couldn't get it right
	url := "https://evolutionpayroll.atlassian.net/rest/auth/1/session"
	//fmt.Println("URL:>", url)
	var jsonStr = []byte(`{"username":` + settings.userName + `, "password":` + settings.passWord + `}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var sessioncookiegrabber LoginResponse
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &sessioncookiegrabber)
	if err != nil {
		panic(err)
	}
	//fmt.Println("response Body:", string(body))
	return sessioncookiegrabber
}
