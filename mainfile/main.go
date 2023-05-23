package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"proj"
)

func main() {
	args := os.Args[1:]

	// from cat command
	if len(args) < 1 {
		input, err := io.ReadAll(os.Stdin)
		fmt.Println(input)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	for i := 0; i < len(args); i++ {
		data, err := ioutil.ReadFile(args[i])
		e := "ERROR: "
		if err != nil {
			fmt.Println(e + " " + err.Error())
			os.Exit(1)
		}
		// end of cat command
		link := &proj.List{}

		// convert byte array to strings
		dataInStr := string(data)
		ar := strings.Fields(dataInStr)

		for _, v := range ar {
			proj.ListPush(link, v)
		}

		// modify based on func
		// main loop
		for i := 0; i < proj.ListSize(link); i++ {

			// binary to decimal case
			if proj.ListAt(link.Head, i).Data == "(bin)" {
				bintodec, _err := strconv.ParseInt((proj.ListAt(link.Head, i-1).Data), 10, 32)
				if err != nil {
					fmt.Println(e + " " + _err.Error())
				}
				bintodec = proj.BinaryToDecimal(bintodec)
				proj.ListAt(link.Head, i-1).Data = fmt.Sprint(bintodec)
				proj.DeleteNode(link.Head, i)
			}

			// hex to decimal case
			if proj.ListAt(link.Head, i).Data == "(hex)" {
				hextodec, _ := strconv.ParseInt((proj.ListAt(link.Head, i-1).Data), 16, 32)
				proj.ListAt(link.Head, i-1).Data = fmt.Sprint(hextodec)
				proj.DeleteNode(link.Head, i)
			}

			// to Upper case
			if proj.ListAt(link.Head, i).Data == "(up)" {
				proj.ListAt(link.Head, i-1).Data = strings.ToUpper(proj.ListAt(link.Head, i-1).Data)
				proj.DeleteNode(link.Head, i)
			} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(up,") {
				//iterator := proj.StringToInt(proj.ListAt(link.Head, i).Data)
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
			}

			// handling cap case
			if strings.Contains(proj.ListAt(link.Head, i).Data, "(cap)") {
				proj.ListAt(link.Head, i-1).Data = strings.Title(proj.ListAt(link.Head, i-1).Data)
				proj.DeleteNode(link.Head, i)
			} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(cap,") {
				//iterator := proj.StringToInt(proj.ListAt(link.Head, i).Data)
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
			}

			// handling tolower case
			if proj.ListAt(link.Head, i).Data == "(low)" {
				proj.ListAt(link.Head, i-1).Data = strings.ToLower(proj.ListAt(link.Head, i-1).Data)
				proj.DeleteNode(link.Head, i)
			} else if strings.Contains(proj.ListAt(link.Head, i).Data, "(low,") {
				//iterator := proj.StringToInt(proj.ListAt(link.Head, i).Data)
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
			}

			// Vowels handle
			if proj.ListAt(link.Head, i).Data == "a" {
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
			for j := 0; j < len(checkpunc)-1; j++ {
				if strings.Contains(string(checkpunc[j]), "'") {
					if checkpunc[j+1] == 32 {
						proj.ListAt(link.Head, i).Data = strings.Replace(proj.ListAt(link.Head, i).Data, "'", "''", -1)
					}
				} else if strings.Contains(string(checkpunc[j]), ",") {
					proj.ListAt(link.Head, i).Data = strings.Replace(proj.ListAt(link.Head, i).Data, ", ", ",", -1)
					proj.ListAt(link.Head, i).Data = strings.Replace(proj.ListAt(link.Head, i).Data, ",'", ", '", -1)
				} else if unicode.IsPunct(checkpunc[0]) && unicode.IsPunct(checkpunc[j]) {
					proj.ListAt(link.Head, i-1).Data = proj.ListAt(link.Head, i-1).Data + string(checkpunc[j])
					proj.ListAt(link.Head, i).Data = string(checkpunc[j+1:])
				}
			}
		}

		// output for testing
		counter := 0
		it := link.Head
		for it != nil {
			fmt.Print(it.Data)
			if counter < proj.ListSize(link)-1 && it.Next.Data != "," {
				if it.Data == "'" {

				}
				fmt.Print(" ")
			}
			it = it.Next
			counter++
		}
	}
	fmt.Println()
}
