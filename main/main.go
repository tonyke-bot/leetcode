package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func create(id int, forceUpdate bool) {
	questions := getQuestions(forceUpdate)
	question, ok := questions[id]
	if !ok {
		fmt.Printf("Problem #%d not found\n", id)
	}

	if question.PaidOnly {
		fmt.Println("This problem is paid user only")
	}

	cwd, _ := os.Getwd()
	newFilename := path.Join(cwd, fmt.Sprintf("../%03d_%s.go", question.ID, strings.Replace(question.Slug, "-", "_", -1)))
	packageName := "leetcode"

	if _, err := os.Stat(newFilename); err != nil && os.IsExist(err) {
		panic(fmt.Sprintf("File %s already existed", newFilename))
	}

	sourceURL := "https://leetcode.com/problems/" + question.Slug
	sourceCode := fmt.Sprintf("package %s\n\n// Source: %s\n\n%s\n/*\nTest Case:\n%s\n*/",
		packageName, sourceURL, question.Code, question.TestCase)
	err := ioutil.WriteFile(newFilename, []byte(sourceCode), 0755)
	if err != nil {
		panic(fmt.Sprintf("Cannot write source code to %v", newFilename))
	}
}

func printDefault() {
	fmt.Println("usage: ./go.sh {new|update} [-f]")
	fmt.Println("command:")
	fmt.Println("  new: create a new solution")
	fmt.Println("  update: update README.md")
	fmt.Println("flag:")
	fmt.Println("  -f: force update problems cache")
}

func main() {
	flag.Parse()

	mode := flag.Arg(0)
	forceUpdate := flag.Arg(1) == "-f"

	switch mode {
	case "new":
		problemID := 0
		fmt.Print("problem #:")
		_, err := fmt.Scan(&problemID)
		if err != nil {
			fmt.Println("Invalid problem id")
			printDefault()
			return
		}
		create(problemID, forceUpdate)
	case "update":
		Format(forceUpdate)

	default:
		fmt.Println("unexpected mode: ", mode)
		printDefault()
		return
	}
}
