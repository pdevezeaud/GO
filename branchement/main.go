package main

import "fmt"

func main() {
	// day of month
	day := 18

	if day < 15 {
		// % represente la variable pour le formatage et recupere le day
		fmt.Printf("we're in the first half of the month (day=%d)\n", day)

	} else if day == 18 {
		fmt.Printf("we're in the special day (day= %d)\n ", day)
	} else {
		fmt.Printf("we're in the first half of the month (day=%d)\n", day)
	}

	year, month, day := 2018, 5, 20
	fmt.Printf("Date=%d/%d/%d\n", year, month, day)

	if year == 2009 && month == 11 && day == 10 {
		fmt.Println("This is the first release of GO")
	} else if year == 2009 || month == 11 || day == 10 {

		println("At least one part is right....:/")
	} else {
		fmt.Println("Just another day...")
	}

	if count := 12; count > 10 {
		fmt.Printf("we have enough count. (got=%d)\n", count)
	} else {
		fmt.Printf("Not enough. (got=%d)\n", count)
	}

}
