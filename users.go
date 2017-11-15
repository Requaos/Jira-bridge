package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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