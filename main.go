package main

// Written by Neil Skinner April 2016
// Refactored by Neil Skinner November 2017
import ( 
	"encoding/csv"
	"flag"
	"fmt"
	"net/http"
	"os" 
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

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
		jiraRSSactivityFeed, jiraIssueWorklog := whoTimeVars(whotime)
		changeIdqueryresult, changeidmap := closemVars(closem, rollbaseSessionKey)
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
						workrate := PayGroupRateReturn(getUserGroups(jiraCreds, sessionCookie, "https://evolutionpayroll.atlassian.net/rest/api/2/user?username="+perworklog.Author.Name+"&expand=groups", httpclient), jiraCreds)
						subexp := (workrate * perworklog.TimeSpentSeconds) / 3600
						worklogdata = append(worklogdata, billingType{
							serviceBureauStr: searchresp.Issues[i].Fields.Customfield11903.Value,
							keyStr: searchresp.Issues[i].Key,
							billStr: billable.Value,
							whoStr: perworklog.Author.Name,
							rate: workrate,
							timeInt: perworklog.TimeSpent,
							exp: subexp,
						})
						totalTimeSpent += perworklog.TimeSpentSeconds
						totalDueMoneys += subexp
					}
					if len(searchresp.Issues[i].Fields.Subtasks) > 0 {
						for _, subTask := range searchresp.Issues[i].Fields.Subtasks {
							subTaskWorklogURL := jiraIssueWorklog + subTask.Key + "/worklog"
							subTaskWorklog := getIssueWorklog(jiraCreds, sessionCookie, subTaskWorklogURL, httpclient)
							for _, perWorklog := range subTaskWorklog.Worklogs {
								rate := PayGroupRateReturn(getUserGroups(jiraCreds, sessionCookie, "https://evolutionpayroll.atlassian.net/rest/api/2/user?username="+perWorklog.Author.Name+"&expand=groups", httpclient), jiraCreds)
								exp := (rate * perWorklog.TimeSpentSeconds) / 3600
								worklogdata = append(worklogdata, billingType{
									serviceBureauStr: searchresp.Issues[i].Fields.Customfield11903.Value,
									keyStr: searchresp.Issues[i].Key,
									billStr: billable.Value,
									whoStr: perWorklog.Author.Name,
									rate: rate,
									timeInt: perWorklog.TimeSpent,
									exp: exp,
								})
								totalTimeSpent += perWorklog.TimeSpentSeconds
								totalDueMoneys += exp
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
							timedata = append(timedata, []string{searchresp.Issues[i].Key, activityhistory.Entry[u].Content.Text, activityhistory.Entry[u].Author.Name})
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
					serviceBureaus = append(serviceBureaus, searchresp.Issues[i].Fields.Customfield11903.Value)
					worklogdata = append(worklogdata, billingType{
						keyStr: searchresp.Issues[i].Key,
						summaryStr: searchresp.Issues[i].Fields.Summary,
						exp: totalDueMoneys,
						serviceBureauStr: searchresp.Issues[i].Fields.Customfield11903.Value,
						timeInt: fmt.Sprintf("%13s", Round(totalTimeSpentinMinutes, time.Minute)),
					})
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
							caseid := rollbaseSelectSingleQuery(rollbaseSessionKey, "select%20id%20from%20case8%20where%20case_number_5%20=%20\""+ifmulticase[numbofbl]+"\"")
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
