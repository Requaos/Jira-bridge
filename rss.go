package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

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

func getIssueHistory(settings ServerSettings, search string, client *http.Client) Feed {
	resturl := search
	//fmt.Println(resturl)
	req, err := http.NewRequest("GET", resturl, nil)
	checkError("Error creating http request during getIssueHistory() call", err)
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
