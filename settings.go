package main

import (
	"bufio"
	"os"
	"strconv"
)

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