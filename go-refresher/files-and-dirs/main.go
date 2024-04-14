package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//func main() {
//	file, err := os.Create("sampleFile.txt")
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		log.Println(file)
//	}
//}
//

// make a directory in go
//func main() {
//	currentDir, err := os.Getwd()
//	currentDir = currentDir + "/files-and-dirs"
//	if err != nil {
//		log.Fatal("an error occurred: ", err)
//	}
//
//	fmt.Println("current directory", currentDir)
//	mkDirErr := os.Mkdir(currentDir+"/sampleDir", 0755)
//
//	if mkDirErr != nil {
//		log.Fatal(mkDirErr)
//	}
//}

//rename a directory in go, behaves like move dir

//func main() {
//	file, err := os.Create("text.txt")
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		log.Println(file)
//	}
//
//	file.Close()
//
//	err = os.Rename("text.txt", "temp.txt")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}

//rename folder

//func main() {
//	//err := os.Mkdir("targetFolder", 0755)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	err := os.Rename("targetFolder", "newFolder")
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// check if file exists

//func main() {
//	//fileInfo, err := os.Stat("sample.txt")
//	folderInfo, err := os.Stat("newFolder")
//
//	if os.IsNotExist(err) {
//		log.Fatal("\nfolder does not exist!")
//	}
//	log.Println(folderInfo)
//}

//func main() {
//	file, err := os.Create("simpleText.txt")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	file.WriteString(
//		"Some information into the file\n" +
//			"More information into file\n",
//	)
//	stat, err := os.Stat("simpleText.txt")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("File permissions: %s\n", stat.Mode())
//	fmt.Printf("File size: %d\n", stat.Size())
//	fmt.Printf("Date modified: %v\n", stat.ModTime())
//
//	err = file.Close()
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// reading a file in go

//func main() {
////	file, err := os.Open("simpleText.txt")
////
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	defer file.Close()
////
////	scanner := bufio.NewScanner(file)
////
////	for scanner.Scan() {
////		fmt.Println(scanner.Text())
////	}
////
////	if err := scanner.Err(); err != nil {
////		log.Fatal(err)
////	}
////
////}

//func main() {
//	fmt.Println("Enter file name: ")
//	var fileName string
//	_, err := fmt.Scanln(&fileName)
//	if err != nil {
//		log.Fatalf("error getting file name: %s", err)
//	}
//
//	fmt.Println("Enter text: ")
//	inputReader := bufio.NewReader(os.Stdin)
//
//	input, _ := inputReader.ReadString('\n')
//	createFile(fileName, input)
//	readFile(fileName)
//
//}
//
//func createFile(fileName, text string) {
//	file, err := os.Create(fileName)
//
//	if err != nil {
//		log.Fatalf("File creation failed %s", err)
//	}
//
//	fileLength, err := file.WriteString(text)
//
//	if err != nil {
//		log.Fatalf("Error writing to file %s", err)
//	}
//
//	fmt.Printf("\nFile name: %s", file.Name())
//	fmt.Printf("\nFile length %d bytes", fileLength)
//}
//
//func readFile(fileName string) {
//	fmt.Println("\nReading file ...")
//	data, err := os.ReadFile(fileName)
//
//	if err != nil {
//		log.Panicf("reading file error: %s", err)
//	}
//
//	fmt.Printf("\nFile name: %s", fileName)
//	fmt.Printf("\nFile size: %d bytes", len(data))
//	fmt.Printf("\nFile content: %s", string(data))
//}

// deleting files and directories

//func main() {
//	fileError := os.Remove("temp.txt")
//
//	if fileError != nil {
//		log.Println(fileError)
//	}
//
//	folderError := os.Remove("sampleDir")
//
//	if folderError != nil {
//		log.Println(folderError)
//	}
//}

// Remove all files and dirs

//func main() {
//	currentDir, err := os.Getwd()
//	if err != nil {
//		log.Fatal("Error getting current directory")
//	}
//
//	targetFolderPath := currentDir + "/targetFolder"
//
//	err = os.RemoveAll(targetFolderPath)
//
//	if err != nil {
//		log.Fatalf("\nError deleting folder: %s", err)
//	}
//}

// iterating over files at a certain path

func main() {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	targetPath := currentDir + "/sampleFolder"
	iterate(targetPath)

}

func iterate(filePath string) {
	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if info.IsDir() {
			fmt.Printf("Folder Name: %s\n", info.Name())
		} else {
			fmt.Printf("File Name: %s\n", info.Name())
		}

		return nil
	})
}
