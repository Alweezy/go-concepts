package main

//type Address struct {
//	state   string
//	city    string
//	zipCode string
//}

// Nested structs

// Author struct
//type Author struct {
//	name   string
//	branch string
//	year   int
//}

// HR struct
//type HR struct {
//	details Author
//}

//func main() {
// instantiating a struct
//address := Address{
//	state:  "Abuja",
//	city:    "Nairobi",
//	zipCode: "00100",
//}
//using new keyword

//address := new(Address)
//address.state = "New York"
//address.city = "New York City"
//address.zipCode = "80100"

//address := &Address{}
//address.state = "New York"
//address.city = "New York City"
//address.zipCode = "80100"

//fmt.Println(*address)

//person := HR{
//	details: Author{
//		name:   "Alex",
//		branch: "Western",
//		year:   2012,
//	},
//}
//
//fmt.Println(person)
//}

/**Adding methods to struct*/

//type Author struct {
//	name   string
//	branch string
//	salary int
//}
//
//func (a Author) show() {
//	fmt.Println("Name: ", author.name)
//	fmt.Println("Branch", author.branch)
//	fmt.Println("Salary", author.salary)
//}
//
//func main() {
//	a := Author{
//		"Alex White",
//		"Western",
//		200,
//	}
//
//	a.show()
//}

// Pointer receiver receivers
//
//type Author struct {
//	name   string
//	branch string
//}
//
//func (a *Author) show(aBranch string) {
//	(*a).branch = aBranch
//}
//
//func main() {
//	author := Author{
//		"Wilkerson",
//		"Western",
//	}
//
//	fmt.Println("Author details before change", author)
//
//	author2 := &author
//	author2.show("Eastern")
//	fmt.Println("Author details post change", author)
//}

// copying structs by reference

//type Student struct {
//	Name      string
//	StudentID int64
//}
//
//func main() {
//	firstStudent := Student{
//		"Alice",
//		22347,
//	}
//
//	secondStudent := &firstStudent
//
//	fmt.Println("First Student", firstStudent)
//	fmt.Println("Second Student", *secondStudent)
//
//	secondStudent.StudentID = 209506
//
//	fmt.Println("First Student", firstStudent)
//	fmt.Println("Second Student", *secondStudent)
//}

// Anonymous structs, useful for creating a one time use struct

//func main() {
//	student := struct {
//		Name      string
//		StudentID int64
//		Age       int
//	}{
//		"Jim",
//		120098,
//		30,
//	}
//
//	fmt.Println("Student: ", student)
//
//}
