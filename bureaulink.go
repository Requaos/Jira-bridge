package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func rollbaseRunTrigger(session string, client *http.Client, objname string, id string, triggerid string) TriggerResponse {
	resturl := fmt.Sprintf("https://apps.bureaulink.com/rest/api/runtrigger?sessionId=%s&id=%s&triggerId=%s&checkValidation=true", session, id, triggerid)
	req, err := http.NewRequest("PUT", resturl, nil)
	checkError("Error creating http request during runTrigger() call", err)
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
