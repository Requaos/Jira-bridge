package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

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
