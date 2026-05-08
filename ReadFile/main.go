package main

import (
	"io"
	"os"
)

func main() {
	// 1
	// fmt.Print(os.Args[1])
	// Run   ->   go run main.go myFile.txt
	// O/P   ->   myFile.txt

	// 2
	// Reading the file
	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(400)
	}
	io.Copy(os.Stdout, f)
	// O/P    ->  Hi this is myFile
}

// 3
// Run   ->   go build main.go
// Run   ->   ./main main.go
// O/P   ->	{
// package main

// import (
// 		"io"
// 		"os"
// )

// func main() {
// 		// 1
// 		// fmt.Print(os.Args[1])
// 		// Run   ->   go run main.go myFile.txt
// 		// O/P   ->   myFile.txt

// 		// 2
// 		// Reading the file
// 		f, err := os.Open(os.Args[1])
// 		if err != nil {
// 				os.Exit(400)
// 		}
// 		io.Copy(os.Stdout, f)
// 		// O/P    ->  Hi this is myFile
// }

// // 3
// // Run   ->   go build main.go
// // Run   ->	 ./main main.go
//				}
