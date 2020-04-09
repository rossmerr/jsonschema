package jsonschema

func Contains(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func Filter(a []string, delegate func(string) bool) []string {
	list := []string{}
	for _, v := range a {
		if delegate(v) {
			list = append(list, v)
		}
	}
	return list
}
