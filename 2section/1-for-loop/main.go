package main

import "fmt"

func main() {

	//for --- only way to loop

	// C-style loop
	for i := 1; i <= 10; i++ {
		// fmt.Println(i)
	}


	//while-style

	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}
	fmt.Println(" ----- Infinite loop ------")

	counter := 0

	for{
		fmt.Println("Counter", counter)
		counter ++

		if counter >=5 {
			break
		}
	}


	fmt.Println(" ----- Skipping numbers ------")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println(" ----- array ------")

	items := [3]string {"Go", "Python", "Java"}

	for index, value := range items {
		fmt.Println(index, value)
	}
}