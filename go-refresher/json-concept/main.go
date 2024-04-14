package old_main

//type Person struct {
//	Name     string `json:"FirstName"`
//	Age      int64
//	Location string
//
//	// We cannot marshal lowercase fields
//	status string
//}
//
//func main() {
//	person := Person{
//		"Ali",
//		25,
//		"Texas",
//		"single",
//	}
//
//	personArray, err := json.Marshal(person)
//
//	if err != nil {
//		log.Fatal("Unable to encode data")
//	}
//
//	fmt.Println(string(personArray))
//
//}

// OmitEmpty

//type Person struct {
//	Name     string `json:"FirstName"`
//	Age      int64
//	Location string `json:"location,omitempty"`
//
//	// We cannot marshal lowercase fields
//	status string
//}
//
//func main() {
//	person := Person{
//		Name:   "Ali",
//		Age:    25,
//		status: "single",
//	}
//
//	personArray, err := json.Marshal(person)
//
//	if err != nil {
//		log.Fatal("Unable to encode data")
//	}
//
//	fmt.Println(string(personArray))
//
//}

// Json to go types, Unmarshal

//type Person struct {
//	Name     string
//	Age      int64
//	Location string
//}
//
//func main() {
//	data := []byte(`{"name": "Joel", "age": 30, "location": "TX"}`)
//	var p Person
//	err := json.Unmarshal(data, &p)
//
//	if err != nil {
//		log.Fatalf("Unable to unmarshal data: %s\n", err)
//	}
//
//	fmt.Printf("Name: %s\n", p.Name)
//	fmt.Printf("Age: %d\n", p.Age)
//	fmt.Printf("Location: %s\n", p.Location)
//
//}

// Streaming encoders and decoders

//func main() {
//	jsonStream := `{"Name":"Jimmie", "Age": 20, "Location": "Texas"}{"Name":"Tye", "Age": 45, "Location": "New Jersey"}`
//	reader := strings.NewReader(jsonStream)
//
//	writer := os.Stdout
//	decoder := json.NewDecoder(reader)
//
//	encoder := json.NewEncoder(writer)
//
//	for {
//		// All keys in json must be string
//		var v map[string]interface{}
//		if err := decoder.Decode(&v); err != nil {
//			log.Println(err)
//			return
//		}
//
//		for k := range v {
//			if k == "Age" {
//				delete(v, k)
//			}
//		}
//
//		if err := encoder.Encode(&v); err != nil {
//			log.Println(err)
//		}
//	}
//}
