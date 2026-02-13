package main 

import "fmt"

func main() {

	tmp := 25

	if  tmp > 30 {
		fmt.Println("Greater than 30")
	} else{
		fmt.Println("greater is less than 30")
	}

	// Else if condition statement

	score := 85
	if score >= 90{
		fmt.Println("Grade is A")
	} else if score >=80{
		fmt.Println("Grade B")
	}else{
		fmt.Println("Grade F")
	}


	// 	Map is a like a dictionary or array depending on the language 
	userAccess := map[string] bool{
		"jane": true,
		"john" : true,
	}

	if hasAccess, ok:= userAccess["jane"]; ok && hasAccess{
		fmt.Println("Jane can access the system")
	}
}


