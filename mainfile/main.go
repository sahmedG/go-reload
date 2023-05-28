package main

// importing all needed modules
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"proj" // custom module to call functions
)

func main() {

	args := os.Args[1:] // catching the argument from terminal

	// checking if arguments contains sample & result file
	if len(args) <= 1 {
		fmt.Println("Error: Missing sample file or output file")
	} else {

		e := "ERROR: " //custom error string

		// opening the file using the path from the arguments catching error if the file/path doesn't exist
		readFile, err := os.Open(args[0])
		if err != nil {
			fmt.Println(e + " " + err.Error())
			//cos.Exit(1)
			return
		}

		//creating a scanner for each new line in the file case the file has multiple lines
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}
		readFile.Close() // closing the file after all the lines where scanned

		var arraytest []string
		f, err := os.Create(args[1]) // creating a result file using the argument and throwing error if the argument was empty
		if err != nil {
			fmt.Println(e + " " + err.Error())
			f.Close()
			return
		}
		// main for loop for the program looping throught each lines were scanned from the sample file
		for _, line := range fileLines {
			link := &proj.List{}
			teststr := ""
			found3 := ""
			ar := strings.Fields(line)
			for _, strline := range ar {
				proj.ListPush(link, strline)
			}
			for i := 0; i < proj.ListSize(link); i++ {
				// converting bin to decimal
				if proj.ListAt(link.Head, i).Data == "(bin)" {
					bintodec, err := strconv.ParseInt((proj.ListAt(link.Head, i-1).Data), 10, 32)
					if err != nil {
						fmt.Println(e + " " + err.Error())
					}
					bintodec = proj.BinaryToDecimal(bintodec)
					proj.ListAt(link.Head, i-1).Data = fmt.Sprint(bintodec)
					proj.DeleteNode(link.Head, i)

				} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(bin,") { // to catch any degits after the tag
					test := strings.Split(proj.ListAt(link.Head, i+1).Data, ")")
					iterator, err := strconv.Atoi(test[0])

					if err != nil {
						fmt.Println(e + " " + err.Error())
						os.Exit(1)
					}
					if i-iterator <= 0 {
						iterator = i
					}

					for j := i - iterator; j < i; j++ {
						bintodec,_ := strconv.ParseInt((proj.ListAt(link.Head, j).Data), 10, 32)
						bintodec = proj.BinaryToDecimal(bintodec)
						proj.ListAt(link.Head, j).Data = fmt.Sprint(bintodec)
					}
					proj.DeleteNode(link.Head, i)
					proj.DeleteNode(link.Head, i)

					// hex to decimal case
				} else if proj.ListAt(link.Head, i).Data == "(hex)" {
					hextodec, err := strconv.ParseInt((proj.ListAt(link.Head, i-1).Data), 16, 64)
					if err != nil {
						fmt.Println(e + " " + err.Error())
						os.Exit(0)
					}
					proj.ListAt(link.Head, i-1).Data = fmt.Sprint(hextodec)
					proj.DeleteNode(link.Head, i)
				} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(hex,") { // to catch any degits after the tag
					test := strings.Split(proj.ListAt(link.Head, i+1).Data, ")")
					iterator, err := strconv.Atoi(test[0])
					if err != nil {
						fmt.Println(e + " " + err.Error())
						os.Exit(1)
					}
					if i-iterator <= 0 {
						iterator = i
					}

					for j := i - iterator; j < i; j++ {
						hextodec, _ := strconv.ParseInt((proj.ListAt(link.Head, j).Data), 16, 64)
						proj.ListAt(link.Head, j).Data = fmt.Sprint(hextodec)
					}

					proj.DeleteNode(link.Head, i)
					proj.DeleteNode(link.Head, i)

					// to Upper case
				} else if proj.ListAt(link.Head, i).Data == "(up)" {
					proj.ListAt(link.Head, i-1).Data = strings.ToUpper(proj.ListAt(link.Head, i-1).Data)
					proj.DeleteNode(link.Head, i)
					//goto loop
				} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(up,") {
					//iterator := StringToInt(proj.ListAt(link.Head, i).Data)
					test := strings.Split(proj.ListAt(link.Head, i+1).Data, ")")
					iterator, _err := strconv.Atoi(test[0])
					if err != nil {
						fmt.Println(e + " " + _err.Error())
					}
					if i-iterator <= 0 {
						iterator = i
					}

					for j := i - iterator; j < i; j++ {
						proj.ListAt(link.Head, j).Data = strings.ToUpper(proj.ListAt(link.Head, j).Data)
					}

					proj.DeleteNode(link.Head, i)
					proj.DeleteNode(link.Head, i)
					//goto loop
					// handling cap case
				} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(cap)") {
					proj.ListAt(link.Head, i-1).Data = strings.Title(proj.ListAt(link.Head, i-1).Data)
					proj.DeleteNode(link.Head, i)
					//goto loop
				} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(cap,") {
					//iterator := StringToInt(proj.ListAt(link.Head, i).Data)
					test := strings.Split(proj.ListAt(link.Head, i+1).Data, ")")
					iterator, _err := strconv.Atoi(test[0])
					if err != nil {
						fmt.Println(e + " " + _err.Error())
					}
					if i-iterator <= 0 {
						iterator = i
					}

					for j := i - iterator; j < i; j++ {
						proj.ListAt(link.Head, j).Data = strings.Title(proj.ListAt(link.Head, j).Data)
					}

					proj.DeleteNode(link.Head, i)
					proj.DeleteNode(link.Head, i)
					//goto loop
					// handling tolower case
				} else if proj.ListAt(link.Head, i).Data == "(low)" {
					proj.ListAt(link.Head, i-1).Data = strings.ToLower(proj.ListAt(link.Head, i-1).Data)
					proj.DeleteNode(link.Head, i)
					//goto loop
				} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(low,") {
					//iterator := StringToInt(proj.ListAt(link.Head, i).Data)
					test := strings.Split(proj.ListAt(link.Head, i+1).Data, ")")
					iterator, _err := strconv.Atoi(test[0])
					if err != nil {
						fmt.Println(e + " " + _err.Error())
					}
					if i-iterator <= 0 {
						iterator = i
					}
					for j := i - iterator; j < i; j++ {
						proj.ListAt(link.Head, j).Data = strings.ToLower(proj.ListAt(link.Head, j).Data)
					}
					proj.DeleteNode(link.Head, i)
					proj.DeleteNode(link.Head, i)
					//goto loop
					// Vowels handle
				} else if proj.ListAt(link.Head, i).Data == "a" {
					vowels := "aAeEiIoOuU"
					for _, ltr := range vowels {
						match, _ := regexp.Compile(string(ltr))
						for _, word := range string(proj.ListAt(link.Head, i+1).Data) {
							found := match.MatchString(string(word))
							if found {
								proj.ListAt(link.Head, i).Data = strings.Replace(proj.ListAt(link.Head, i).Data, "a", "an", -1)
							}
							break
						}
					}
				}

				//punctuation check
				checkpunc := []rune(proj.ListAt(link.Head, i).Data)
				for j := 0; j < len(checkpunc); j++ {

					if strings.Contains(string(checkpunc[j]), "'") {

					} else if unicode.IsPunct(checkpunc[0]) && unicode.IsPunct(checkpunc[j]) {
						proj.ListAt(link.Head, i-1).Data = proj.ListAt(link.Head, i-1).Data + string(checkpunc[j])
						proj.ListAt(link.Head, i).Data = string(checkpunc[j+1:])

					}
				}
			}

			counter := 0
			it := link.Head
			found := ""
			found2 := ""

			for it != nil {
				teststr = teststr + it.Data
				if counter < proj.ListSize(link)-1 && it.Next.Data != "," {
					teststr = teststr + " "
				}
				it = it.Next
				counter++
			}
			teststr = strings.ReplaceAll(teststr, "  ", " ")
			match, _ := regexp.Compile("[‘|'|`](.*?)\\s[‘|'|`]")
			found = (match.FindString(teststr))
			if len(found) > 0 {
				firstpunc := (string(found[0]))
				lastpunc := (string(found[len(found)-1]))
				found2 = (strings.ReplaceAll(found, " ‘", "'"))
				found2 = (strings.ReplaceAll(found2, "‘ ", "'"))
				if firstpunc == "â" || lastpunc == "â" {
					firstpunc = "'"
					lastpunc = "'"
				}
				found3 = fmt.Sprintf("%s%s%s", firstpunc, strings.TrimSpace(found2[1:len(found2)-1]), lastpunc)
				found3 = strings.ReplaceAll(teststr, found, found3)
			}
			if found3 == "" {
				found3 = teststr
			}
			arraytest = append(arraytest, found3)
		}
		for _, v := range arraytest {
			fmt.Fprintln(f, v)
			if err != nil {
				fmt.Println(e + " " + err.Error())
				return
			}
		}

		err = f.Close()
		if err != nil {
			fmt.Println(e + " " + err.Error())
			return
		}
		if err == nil {
			fmt.Println("file written successfully")
		}
	}
}
