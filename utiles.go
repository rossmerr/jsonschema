package jsonschema

// Contains checks a slice to see if it contains the matching string
func Contains(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

// Filter removes from a slice any true conditions from the delegate
func Filter(a []string, delegate func(string) bool) []string {
	list := []string{}
	for _, v := range a {
		if delegate(v) {
			list = append(list, v)
		}
	}
	return list
}
