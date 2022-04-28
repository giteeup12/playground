package main

import "fmt"

func main() {
	fmt.Println("hello")
	type xxx struct {
		first_name  string
		last_name   string
		middle_init string
		age         int
	}
	var yyyy []xxx
	yyyy = append(yyyy, xxx{"first name", "last name", "middle", 3})
	yyyy = append(yyyy, xxx{"first name1", "last name1", "middle1", 4})
	yyyy = append(yyyy, xxx{"first name2", "last name2", "middle2", 5})
	yyyy = append(yyyy, xxx{"first name3", "last name3", "middle3", 4})
	for sub, tab := range yyyy {
		fmt.Println(sub, tab.first_name, tab.last_name, tab.middle_init, tab.age)
	}
	type qqqq struct {
		astronaut1          string
		cosmonaut           string
		private_citizen     string
		astronaut1_age      int
		cosmonaut_age       int
		private_citizen_age int
	}

	var ssss []qqqq
	ssss = append(ssss, qqqq{"a1", "c1", "p1", 10, 11, 12})
	for sub1, tab_value := range ssss {
		fmt.Println(sub1, tab_value.astronaut1, tab_value.cosmonaut_age+1)
	}

}
