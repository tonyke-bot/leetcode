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
	readmeOverviewStartTag := []byte("<!-- OVERVIEW START -->")
	readmeOverviewEndTag := []byte("<!-- OVERVIEW END -->")
	readmeSolutionStartTag := []byte("<!-- SOLUTION START -->")
	readmeSolutionEndTag := []byte("<!-- SOLUTION END -->")

	files, err := ioutil.ReadDir("../")
	if err != nil {
		panic(fmt.Sprintf("Cannot ReadDir of .., error: %v", err))
	}

	levelCount := map[string]int{
		"Easy":   0,
		"Medium": 0,
		"Hard":   0,
	}
	levelTotalCount := map[string]int{
		"Easy":   0,
		"Medium": 0,
		"Hard":   0,
	}
	// Generate Table
	questions := getQuestions()
	var questionInfos []Question
	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		id := ParseFolderName(f.Name())
		if id == -1 {
			continue
		}

		question, found := questions[id]
		if !found {
			panic(fmt.Sprintf("Problem #%d is not found", id))
		}
		levelCount[question.Level]++
		questionInfos = append(questionInfos, question)
	}
	sort.Sort(QuestionInfos(questionInfos))

	for _, q := range questions {
		levelTotalCount[q.Level]++
	}

	overviewTable := "" +
		fmt.Sprintf("![Easy: %d/%d](https://img.shields.io/badge/Easy-%d/%d-green.svg?style=for-the-badge)\n", levelCount["Easy"], levelTotalCount["Easy"], levelCount["Easy"], levelTotalCount["Easy"]) +
		fmt.Sprintf("![Medium: %d/%d](https://img.shields.io/badge/Medium-%d/%d-orange.svg?style=for-the-badge)\n", levelCount["Medium"], levelTotalCount["Medium"], levelCount["Medium"], levelTotalCount["Medium"]) +
		fmt.Sprintf("![Hard: %d/%d](https://img.shields.io/badge/Hard-%d/%d-red.svg?style=for-the-badge)\n", levelCount["Hard"], levelTotalCount["Hard"], levelCount["Hard"], levelTotalCount["Hard"]) +
		"\n"

	solutionTable := "#|Name|Difficulty|Tags\n" +
		"-:|----|----------|----\n"
	for _, question := range questionInfos {
		// Don't use path.Join because this filename is used in HTML
		filename := FormatDirectoryName(question) + "/" + FormatFilename(question)
		tags := make([]string, len(question.Tags))

		for _, tag := range question.Tags {
			tags = append(tags, "`"+tag.Name+"`")
		}

		solutionTable += strings.Join([]string{
			strconv.Itoa(question.ID),
			fmt.Sprintf("[%s](https://leetcode.com/problems/%s) [[Solution](./%s)]", question.Title, question.Slug, filename),
			question.Level,
			strings.TrimLeft(strings.Join(tags, " "), " "),
		}, "|") + "\n"
	}

	// Subsititude
	readme, _ := ioutil.ReadFile(readmePath)
	overviewStartPos := bytes.Index(readme, readmeOverviewStartTag)
	overviewEndPos := bytes.LastIndex(readme, readmeOverviewEndTag)
	switch {
	case overviewStartPos == -1:
		panic("Cannot find table start tag in README.md")
	case overviewEndPos == -1:
		panic("Cannot find table end tag in README.md")
	case overviewStartPos > overviewEndPos:
		panic("Cannot find overview placeholder")
	}
	readme = []byte(string(readme[:overviewStartPos+len(readmeOverviewStartTag)]) +
		"\n" +
		overviewTable +
		string(readme[overviewEndPos:]))

	solutionStartPos := bytes.Index(readme, readmeSolutionStartTag)
	solutionEndPos := bytes.LastIndex(readme, readmeSolutionEndTag)
	switch {
	case solutionStartPos == -1:
		panic("Cannot find table start tag in README.md")
	case solutionEndPos == -1:
		panic("Cannot find table end tag in README.md")
	case solutionStartPos > solutionEndPos:
		panic("Cannot find table placeholder")
	}
	readme = []byte(string(readme[:solutionStartPos+len(readmeSolutionStartTag)]) +
		"\n" +
		solutionTable +
		string(readme[solutionEndPos:]))

	// WriteFile
	err = ioutil.WriteFile(readmePath, readme, 0755)
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

	dirname := path.Join("..", FormatDirectoryName(question))
	if _, err := os.Stat(dirname); !os.IsNotExist(err) {
		fmt.Printf("Solution #%d has already been created\n", question.ID)
		return
	}

	os.Mkdir(dirname, 0755)
	filename := FormatFilename(question)
	newFilename := path.Join(dirname, filename)
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
