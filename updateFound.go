package main

func updateFound(found []string, word string, letter string) bool {
	complete := true
	for i, r := range word {
		if letter == string(r) {
			found[i] = letter
		}
		if found[i] == "_" {
			complete = false
		}
	}
	return complete
}
