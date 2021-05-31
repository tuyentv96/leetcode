func isValid(s string) bool {
	var arr []int32
	for _, r := range s {
		switch r {
		case '{':
			arr = append(arr, r)
		case '(':
			arr = append(arr, r)
		case '[':
			arr = append(arr, r)
		case '}':
			if len(arr) == 0 || arr[len(arr)-1] != '{' {
				return false
			}

			arr = arr[:len(arr)-1]
		case ')':
			if len(arr) == 0 || arr[len(arr)-1] != '(' {
				return false
			}

			arr = arr[:len(arr)-1]
		case ']':
			if len(arr) == 0 || arr[len(arr)-1] != '[' {
				return false
			}

			arr = arr[:len(arr)-1]
		}
	}

	if len(arr) > 0 {
		return false
	}
	return true
}
