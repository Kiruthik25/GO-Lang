package students

import "fmt"



func total (){

	
students := map[string]map[string]int{
		"John": {
			"Math":    95,
			"Science": 90,
		},
		"David": {
			"Math":    80,
			"Science": 91,
		},
		"Alice": {
			"Math":    98,
			"Science": 97,
		},
	}

	topper :=""

	highestTotal := 0


	for student, subject := range students{

		total := 0

		for _,marks := range subject{

			total+=marks
		}

		if total > highestTotal {
			highestTotal = total
			topper = student
		}
	}

	fmt.Println(topper, highestTotal)
}