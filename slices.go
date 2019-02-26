package main

import "fmt"

func main() {
	arr := []int{0,1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("arr[2:5]",arr[2:5])

	fmt.Println("after update slice")
	s1:=arr[2:5]
	s1[0]=100

	fmt.Println("s1:",s1)
	fmt.Println("arr:",arr)
}
