package main

func isIsomorphic(s string, t string) bool {
	stMap := make(map[string]string)
	tsMap := make(map[string]string)
	for pos, val := range s {
		if stMap[string(val)] == "" {
			stMap[string(val)] = string(t[pos])
		} else if stMap[string(val)] != string(t[pos]) {
			return false
		}
		if tsMap[string(t[pos])] == "" {
			tsMap[string(t[pos])] = string(val)
		} else if tsMap[string(t[pos])] != string(val) {
			return false
		}
	}
	return true
}

func main() {
	print(isIsomorphic("paper", "title"))
}
