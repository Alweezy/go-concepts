package main

//func main() {
//	emptyMap := map[int]string{}
//
//	fmt.Println(emptyMap)
//
//	contentMap := map[int]string{
//		1: "Alex",
//		2: "Bob",
//		3: "Alice",
//	}
//
//	fmt.Println(contentMap)
//}

// creating maps uisng make
//func main() {
//	scoreMap := make(map[string]float64)
//
//	scoreMap["Bob"] = 76.5
//	scoreMap["Alvin"] = 90.6
//	scoreMap["Alice"] = 90.8
//
//	fmt.Println("Scores", scoreMap)
//
//	// get the length of a map using inbuilt function
//	fmt.Printf("The size of the map: %d", len(scoreMap))
//
//}

//func main() {
//	// accessing the contents of a map
//	dummyMap := map[string]int{"Joe": 100, "Bill": 200, "Dickens": 500}
//	fmt.Println(dummyMap["Bill"])
//}

// adding/updating and deleting items from a map
//func main() {
//	sampleMap := map[int]string{10: "Bill", 20: "Cyntia", 30: "Calvin"}
//
//	fmt.Println("Map before", sampleMap)
//	sampleMap[70] = "Donda"
//	sampleMap[80] = "Ye"
//
//	fmt.Printf("SampleMap after adding: %v\n", sampleMap)
//
//	// updating by key
//
//	sampleMap[70] = "Winnie"
//	fmt.Println("Sample map post update: ", sampleMap)
//
//	// delete by key
//	delete(sampleMap, 70)
//
//	fmt.Println("Sample map after deleting item: ", sampleMap)
//}

// iterating over a map
//func main() {
//	// pretty much similar to iterating over an array
//	sampleMap := map[int]string{
//		10: "Abed",
//		20: "Noah",
//		30: "Eve",
//		40: "Noela",
//	}
//
//	for _, name := range sampleMap {
//		fmt.Println(name)
//	}
//
//}

//sorting map by keys

//func main() {
//	unsortedMap := map[int]string{
//		30: "Bill",
//		50: "Kim",
//		20: "Alex",
//		10: "Jack",
//	}
//
//	// a slice to hold keys
//	mapKeys := make([]int, 0)
//
//	for key := range unsortedMap {
//		mapKeys = append(mapKeys, key)
//	}
//
//	fmt.Println("MapKeys populated", mapKeys)
//	// we then sort the keys
//	sort.Ints(mapKeys)
//
//	fmt.Println("MapKeys sorted", mapKeys)
//	for _, key := range mapKeys {
//		fmt.Println(key, unsortedMap[key])
//	}
//}

//sort map by values

//func main() {
//	unsortedMap := map[int]string{
//		30: "Bill",
//		50: "Kim",
//		20: "Alex",
//		10: "Jack",
//	}
//
//	// a slice to hold keys
//	mapValues := make([]string, 0)
//
//	for _, value := range unsortedMap {
//		mapValues = append(mapValues, value)
//	}
//
//	fmt.Println("MapValues populated", mapValues)
//	// we then sort the values
//	sort.Strings(mapValues)
//
//	fmt.Println("MapValues sorted", mapValues)
//	for _, value := range mapValues {
//		fmt.Println(value)
//	}
//}

// truncating a map
//func main() {
//	sampleMap := map[int]string{
//		10: "Kim",
//		20: "Jim",
//		30: "Alex",
//	}
//
//	for key, value := range sampleMap {
//		fmt.Printf("Key: %d Value:%s\n", key, value)
//	}
//
//	for key := range sampleMap {
//		delete(sampleMap, key)
//	}
//
//	fmt.Printf("Sample map after deleting keys .... : %v", sampleMap)
//
//}

// truncating using make
//func main() {
//	sampleMap := map[int]string{
//		10: "Kim",
//		20: "Jim",
//		30: "Alex",
//	}
//
//	for key, value := range sampleMap {
//		fmt.Printf("Key: %d Value:%s\n", key, value)
//	}
//
//	sampleMap = make(map[int]string)
//
//	fmt.Printf("Sample map after deleting keys .... : %v", sampleMap)
//
//}

// merging maps
//
//func main() {
//	firstMap := map[int]string{
//		10: "Kim",
//		20: "Jim",
//		30: "Alex",
//	}
//
//	secondMap := map[int]string{
//		40: "Nana",
//		50: "White",
//		60: "Jean",
//	}
//
//	for key, value := range secondMap {
//		firstMap[key] = value
//	}
//
//	fmt.Printf("First map after update: %v\n", firstMap)
//
//}
