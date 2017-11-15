package main

import (
	"fmt"
	"net/http"
)

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
