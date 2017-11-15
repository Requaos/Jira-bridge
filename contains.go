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