/*Create a two dimensional array of integers. 5 rows and 2 columns. First
column of 5 rows are 10, 12, 15, 19, 24. Popopulate the 2nd column with
double of first column values*/

/*step1=create 2D array
step2=5rows 2columns
step3=1st colum of 5 rows are 10, 12, 15, 19, 24.
step4=Popopulate the 2nd column with double of first column values*/

package main

import (
	"fmt"
)

func main() {
	var numbers [5][2]int
	numbers[0][0] = 10
	numbers[1][0] = 12
	numbers[2][0] = 15
	numbers[3][0] = 19
	numbers[4][0] = 24

	numbers[0][1] = numbers[0][0] * 2
	numbers[1][1] = numbers[1][0] * 2
	numbers[2][1] = numbers[2][0] * 2
	numbers[3][1] = numbers[3][0] * 2
	numbers[4][1] = numbers[4][0] * 2
	fmt.Println(numbers)

}
