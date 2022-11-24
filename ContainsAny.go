package main

func containsAny(s string, chars []string) bool { //verifie si la lettre donnée est bien dans le mot
	for _, ch := range s {
		for _, ch2 := range chars {
			if string(ch) == ch2 {
				return true
			}
		}
	}
	return false
}
