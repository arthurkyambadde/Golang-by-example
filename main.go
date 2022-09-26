package main

import (
	"fmt"

	"github.com/arthurkyambadde/goWithExamples/Sums"
)

func main() {
	fmt.Println(Sums.Sum(2, 3))
	fmt.Println(Sums.Sum(4, 6, 8))
	fmt.Println(Sums.Sum(3, 6, 9, 4, 6, 7, 8, 9))

	ages := []int{2, 6, 9, 0, 3}

	fmt.Println(Sums.Sum(ages...))
}
