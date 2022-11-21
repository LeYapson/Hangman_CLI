package main

func updateFound(found []string, word string, letter string) bool { // fin du hangman (savoir si toutes les lettres ont été trouvés)
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
