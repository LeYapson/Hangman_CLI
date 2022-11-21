package main

func containsAny(s string, chars []string) bool {      //verifie si la letre donner et bien dans le mot
	for _, ch := range s {
		for _, ch2 := range chars {
			if string(ch) == ch2 {
				return true
			}
		}
	}
	return false
}
