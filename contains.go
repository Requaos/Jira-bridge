package main

//PayGroupContains checks if a string is in an array of strings
func PayGroupContains(s UserWithGroups, e string) bool {
	for _, a := range s.Groups.Items {
		if a.Name == e {
			return true
		}
	}
	return false
}

// PayGroupRateReturn returns the appropriate
func PayGroupRateReturn(s UserWithGroups, v ServerSettings) int {
	var i int
	if PayGroupContains(s, "ManagerConsult") {
		i = v.mRate
	} else if PayGroupContains(s, "CustomProgramming") {
		i = v.pRate
	} else if PayGroupContains(s, "TechProgSupport") {
		i = v.tRate
	} else {
		i = v.dRate
	}
	return i
}