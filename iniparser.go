package main

import "bufio"
import "os"
import "fmt"
// import "bytes"
import "strings"
// import "regexp"

const bufferSize = 4096

func checkForError(e error, errorMsg string) {
    if e != nil {
		fmt.Printf(errorMsg)
        panic(e)
    }
}


func getFileSize(filename string) int {
	// try opening the file, returns the file and error
	f, e := os.Open(filename)
	checkForError(e, "Cannot open file")

	fileData := make([]byte, bufferSize)
	fileSize, e := f.Read(fileData)
	checkForError(e, "Cannot read from file")
	
	f.Close()
	return fileSize
}


func rmvComments(filename string, size int) {
	var iniData string
	// lineCounter := 0

	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)

	for {
		line, _ := reader.ReadString('\n')
		i := strings.Index(line, ";")
		if i == 0 {
			continue
		} else if i > 0 {
			line = line[0:i-1]
		}

		if line == "" {
            break
        }

		// lineCounter+=1
		iniData += line
	}
	os.WriteFile(filename , []byte(iniData) , 0644 )
	f.Close()
}

func parse(filename string, size int) {
	lineCounter := 0
	var section string

	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)

	data := make(map[string]map[string]string)

	for lineCounter < size {
		line, _ := reader.ReadString('\n')
		// strings.Split(line , "\t")
		i := strings.Index(line, "[")
		j := strings.Index(line, "]")
		if i >= 0 && j >= 0 {
			data[line[i+1:j]] = make(map[string]string)
			section = line[i+1:j]
		} 
		
		k := strings.Index(line, "=")
		if k >= 0 {
			data[section][line[0:k-1]] = line[k+1:]
		}
		
		lineCounter+=1
	}

	fmt.Println(data)
	
}

func main() {
	fileSize := getFileSize("input.ini")
	rmvComments("input.ini", fileSize)
	parse("input.ini", fileSize)
}
