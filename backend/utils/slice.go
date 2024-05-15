package utils

func AppendUnique(slice []string, element string) []string {
	for _, e := range slice {
		if e == element {
			return slice // Element already exists, no need to append
		}
	}
	return append(slice, element)
}

func StringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func RemoveElement(slice []string, element string) []string {
	for i, e := range slice {
		if e == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}