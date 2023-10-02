package main

func check(x int) bool {
	for p := 2; p*p <= x; p++ {
		if x%p == 0 {
			return false
		}
	}
	return true
}
