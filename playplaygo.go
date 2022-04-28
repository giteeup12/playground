package main

import "fmt"

func main() {

	fmt.Println("hello")
	type xyxy struct {
		first_name  string
		last_name   string
		middle_init string
		age         int
	}
	var yxyx []xyxy
	yxyx = append(yxyx, xyxy{"first name", "last name", "m", 4})
	yxyx = append(yxyx, xyxy{"first name1", "last name2", "m2", 5})
	yxyx = append(yxyx, xyxy{"first name2", "last name3", "m3", 6})
	for sub, sub_value := range yxyx {
		fmt.Println("subscript: ", sub, sub_value.first_name)
	}

}
