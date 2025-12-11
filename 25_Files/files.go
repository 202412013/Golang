package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("25) Files/example.txt")
	// error na hoi toh nil return kare che naitar err return karse
	if err != nil {
		// log the error
		// program ne panic mode ma nakhi daie
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("FileName ->", fileInfo.Name())
	fmt.Println("File or Directory ->", fileInfo.IsDir()) // file che folder nai etle false
	fmt.Println("File size ->", fileInfo.Size())
	fmt.Println("File Mode ->", fileInfo.Mode())
	fmt.Println("file modified at ->", fileInfo.ModTime())

	// read file
	defer f.Close()

	buf := make([]byte, fileInfo.Size())

	// IMP line
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	println("data", d, buf)

	// Read data
	for i := 0; i < len(buf); i++ {
		println("Data ->", i, string(buf[i]))
	}

	defer f.Close()

	// Simple trick to read data for small files
	file, error := os.ReadFile("25) Files/example.txt")

	if error != nil {
		panic(error)
	}

	fmt.Print("Data -> ", string(file))

	defer f.Close()

	// Read Folders
	dir, er := os.Open(".")
	if er != nil {
		panic(er)
	}

	defer dir.Close()

	// return slice
	// n > 0 -> atmost etlij files ya folder dekhadse
	// n <= 0 -> returns all the files or folder(directory)
	fileInfos, er := dir.ReadDir(1)

	for _, fi := range fileInfos {
		fmt.Println("All Files and Folder ->", fi.Name(), fi.IsDir())
	}

	// Create a file
	f2, err2 := os.Create("25) Files/example2.txt")
	if err2 != nil {
		panic(err2)
	}

	defer f2.Close()

	f2.WriteString("hi go")
	// Append into file
	f2.WriteString(" -> Hi Golang") // append

	// Replace whole content Homework
	// Clear the file completely
	err3 := f2.Truncate(0)
	if err3 != nil {
		panic(err3)
	}

	// Move cursor to the beginning
	_, err3 = f2.Seek(0, 0)
	if err3 != nil {
		panic(err3)
	}

	f2.WriteString("Hello Golang \nThank You")

	// Bytes array and write into files
	bytes := []byte("\nHello Golang language")
	f2.Write(bytes) // Write function help kare che slice array parameter ma levama

	// Streaming technique user(IMP) and read, write from one file to Another File
	sourceFile, err := os.Open("25) Files/example2.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("25) Files/example3.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()
	// Default Size 4096
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("Written to new File Successfully")

	// delete a File
	errors := os.Remove("25) Files/example3.txt")
	if errors != nil {
		panic(errors)
	}

	fmt.Println("Files Deleted Successfully")

	// // Copy content
	// _, err = io.Copy(destFile, sourceFile)
	// if err != nil {
	// 	panic(err)
	// }

	// // Force write to disk
	// err = destFile.Sync()
	// if err != nil {
	// 	panic(err)
	// }

}
