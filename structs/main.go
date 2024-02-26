package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) updateFirstName(newFirstName string) {
	(*&p.firstName) = newFirstName
}

func main() {
	alex := person{
		firstName: "Doan",
		lastName:  "Tran Ba Dat",
		contact: contactInfo{
			email:   "badatdoan@gmail.com",
			zipCode: 59,
		},
	}
	fmt.Println(alex.firstName)
	fmt.Println(alex.lastName)
	fmt.Println(alex.contact.email)

	var john person
	john.firstName = "John"
	john.lastName = "Waven"
	john.contact.email = "johnwaven@gmail.com"
	john.contact.zipCode = 23

	//Go là ngôn ngữ tham trị
	johnPointer := &john
	johnPointer.updateFirstName("Jimmy")
	fmt.Println("Jimmy new: ", john)

	mySlice := []string{"Hi", "Hello", "How", "Are", "You"}
	updateStrArr("Bye", mySlice)
	fmt.Println(mySlice)

	name := "Bill"
	fmt.Println(*&name) //Bill
	//&name: trỏ đến vùng nhớ trong bộ nhớ chứa giá trị của chuỗi "Bill"
	//*: Lấy giá trị nơi mà con trỏ được trỏ đến

	name2 := "Alex"

	namePointer := &name2

	fmt.Println(&namePointer)
	printPointer(namePointer)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors["red"])

	colors2 := make(map[string]string)
	colors2["white"] = "#fffff"
	colors2["red"] = "#ff0000"
	printMap(colors2)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)

}

func updateStrArr(newStr string, s []string) {
	s[0] = newStr
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key + " is: " + value)
	}
}
