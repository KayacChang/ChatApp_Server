package utils

// Find TODO
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Remove TODO
func Remove(arr []string, item string) []string {
	i := Find(arr, item)
	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = ""
	return arr[:len(arr)-1]
}
