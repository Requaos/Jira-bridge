To launch with standard parameters:
	(filter 18800 "Release Note Grabber Tool", 201 file max, filename "results")
	- Double click on jira-bridge.exe
	- Wait for file results.csv to appear

To launch with custom prarameters:
	- Open the folder containing script
	- Shift+Right click
	- Choose "Open command window here"
	- enter the command parameters (see below)


jira-bridge [-f filterID] [-max maxrows] [-?]


-f      allows you to change the filter this tool pulls from
         Default: 18800 "SNAP NEXT"

-max    allows you to change how many rows are returned per
        search iteration, speed versus # of api calls
         Default: 201

-time   creates a PDF with the time spent and if it's billable (bool)

-file   allows you to change output filename
         Default: results

-?     Prints this message block

-close  allows you to link an existing 'change_id' any BL cases attached
         Default: off

(1) Example: jira-bridge -f 19015 -close
(2) Example: jira-bridge -f 18515 -max 200 -file snapnextReleased

Config File Name: jira-bridge.ini

Config File Contents:
JiraServer=evolutionpayroll.atlassian.net
JiraUsername=
JiraPassword=
BureauLinkTenant=13924598
BureauLinkUser=
BureauLinkPassword=
ProgrammingRate=
TechConsultRate=
DefaultRate=
ManagerRate=
