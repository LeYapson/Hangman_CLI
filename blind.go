package main

func Blind(s string) string {
	array := []rune(s)
	for i := 0; i < len(array); i++{
		if array[i] >= 'a' && array[i] <= 'z' {
			array[i] -= 52
	}
	}
	return string(array)
}
