package main

func blind(word string) []string {
	found := []string{}
	for i := 0; i < len(word); i++ {
		found = append(found, "_")
	}
	return found
}