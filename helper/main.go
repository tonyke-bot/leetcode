package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

// QuestionInfos type implements series of function to make sorted by problem # possible
type QuestionInfos []Question

func (qInfos QuestionInfos) Len() int           { return len(qInfos) }
func (qInfos QuestionInfos) Swap(i, j int)      { qInfos[i], qInfos[j] = qInfos[j], qInfos[i] }
func (qInfos QuestionInfos) Less(i, j int) bool { return qInfos[i].ID < qInfos[j].ID }

func format() {
	readmePath := "../README.md"
	readmeTableStartTag := []byte("<!-- OVERVIEW START -->")
	readmeTableEndTag := []byte("<!-- OVERVIEW END -->")

	files, err := ioutil.ReadDir("../")
	if err != nil {
		panic(fmt.Sprintf("Cannot ReadDir of .., error: %v", err))
	}

	// Generate Table
	questions := getQuestions()
	var questionInfos []Question
	for _, f := range files {
		id := ParseSolutionFilename(f.Name())
		if id == -1 {
			continue
		}

		question, found := questions[id]
		if !found {
			panic(fmt.Sprintf("Problem #%d is not found", id))
		}
		questionInfos = append(questionInfos, question)
	}
	sort.Sort(QuestionInfos(questionInfos))

	tableContent := "#|Name|Difficulty|Tags\n" +
		"-:|----|----------|----\n"
	for _, question := range questionInfos {
		filename := FormatFilename(question)
		tags := make([]string, len(question.Tags))

		for _, tag := range question.Tags {
			tags = append(tags, "`"+tag.Name+"`")
		}

		tableContent += strings.Join([]string{
			strconv.Itoa(question.ID),
			fmt.Sprintf("[%s](https://leetcode.com/problems/%s) [[Solution](./%s)]", question.Title, question.Slug, filename),
			question.Level,
			strings.TrimLeft(strings.Join(tags, " "), " "),
		}, "|") + "\n"
	}

	// Subsititude
	readme, _ := ioutil.ReadFile(readmePath)
	startPos := bytes.Index(readme, readmeTableStartTag)
	endPos := bytes.LastIndex(readme, readmeTableEndTag)
	switch {
	case startPos == -1:
		panic("Cannot find table start tag in README.md")
	case endPos == -1:
		panic("Cannot find table end tag in README.md")
	case startPos > endPos:
		panic("Cannot find table placeholder")
	}

	newReadme := string(readme[:startPos+len(readmeTableStartTag)]) + "\n" +
		tableContent +
		string(readme[endPos:])

	// WriteFile
	err = ioutil.WriteFile(readmePath, []byte(newReadme), 0755)
	if err != nil {
		panic(fmt.Sprintf("Fails to write README.MD, error: %v", err))
	}
}

func create(id int) {
	questions := getQuestions()
	question, ok := questions[id]
	if !ok {
		fmt.Printf("Problem #%d not found\n", id)
		return
	}

	if question.PaidOnly {
		fmt.Println("This problem is paid user only")
		return
	}

	filename := FormatFilename(question)
	newFilename := path.Join("..", filename)
	if _, err := os.Stat(newFilename); !os.IsNotExist(err) {
		fmt.Printf("File %s already existed\n", filename)
		return
	}

	sourceCode := GenerateQuestionCode(question)
	err := ioutil.WriteFile(newFilename, []byte(sourceCode), 0755)
	if err != nil {
		fmt.Printf("Cannot write source code to %v\n", newFilename)
		panic(err)
	}
}

func update() {
	questions, err := DownloadQuestions()
	if err != nil {
		panic(err)
	}
	SaveQuestions(CacheFilename, questions)
}

func getQuestions() (questions map[int]Question) {
	questions, err := LoadQuestions(CacheFilename)
	if err != nil {
		fmt.Printf("Fail to load questions from cache, error:%v\n", err)
		fmt.Println("Please run `./go.sh update` manually")
	}
	return questions
}

func printDefault() {
	fmt.Println("usage: ./go.sh {new|update|format} ")
	fmt.Println("command:")
	fmt.Println("  new (problemID): create a new solution")
	fmt.Println("  format: format README.md")
	fmt.Println("  update: update question cache")
}

func main() {
	flag.Parse()
	mode := flag.Arg(0)

	switch mode {
	case "new":
		problemID, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			fmt.Println("invalid problem #: ", flag.Arg(1))
		}

		create(problemID)
	case "format":
		format()

	case "update":
		update()

	default:
		fmt.Println("unexpected mode: ", mode)
		printDefault()
		return
	}
}
