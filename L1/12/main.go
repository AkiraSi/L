package main

func main() {
	inp := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool) // типа сет

	for _, str := range inp { // добавление в сет элемента
		set[str] = true
	}
}
