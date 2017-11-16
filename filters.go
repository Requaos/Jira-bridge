package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func getFilterContents(settings ServerSettings, session string, client *http.Client, jiraFilter string) FilterResponse {
	resturl := jiraFilter
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
	checkError("Error creating new http request during getFilterContents() call", err)
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
