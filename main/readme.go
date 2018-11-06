package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var filenamePattern = regexp.MustCompile("(\\d*)_[0-9a-zA-Z_]*?\\.go")
var readmeTableStartTag = []byte("<!-- OVERVIEW START -->")
var readmeTableEndTag = []byte("<!-- OVERVIEW END -->")

// QuestionInfos type implements series of function to make sorted by problem # possible
type QuestionInfos []Question

func (qInfos QuestionInfos) Len() int           { return len(qInfos) }
func (qInfos QuestionInfos) Swap(i, j int)      { qInfos[i], qInfos[j] = qInfos[j], qInfos[i] }
func (qInfos QuestionInfos) Less(i, j int) bool { return qInfos[i].ID < qInfos[j].ID }

// Format README.md
func Format(forceUpdate bool) {
	readmePath := "../README.md"
	questions := getQuestions(forceUpdate)
	solutions := GetSolutionList()
	if len(solutions) == 0 {
		panic("Nothing to be formatted")
	}

	// Generate Table
	var questionInfos []Question
	for _, filename := range solutions {
		id, _ := strconv.Atoi(filenamePattern.FindStringSubmatch(filename)[1])
		question, found := questions[id]
		if !found {
			panic(fmt.Sprintf("Problem #%d is not found", id))
		}
		questionInfos = append(questionInfos, question)
	}
	sort.Sort(QuestionInfos(questionInfos))

	tableContent := "#|Name|Difficulty|Tags\n-|----|----------|--------|----\n"
	for _, question := range questionInfos {
		filename := fmt.Sprintf("%d_%s.go", question.ID, strings.Replace(question.Slug, "-", "_", -1))
		tags := make([]string, len(question.Tags))

		for _, tag := range question.Tags {
			tags = append(tags, "`"+tag+"`")
		}

		tableContent += strings.Join([]string{
			strconv.Itoa(question.ID),
			fmt.Sprintf("[%s](https://leetcode.com/problems/%s/description/) [[Solution](%s)]", question.Title, question.Slug, filename),
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
	err := ioutil.WriteFile(readmePath, []byte(newReadme), 0755)
	if err != nil {
		panic(fmt.Sprintf("Fails to write README.MD, error: %v", err))
	}
}

// GetSolutionList returns a list of solution filename
func GetSolutionList() (solutions []string) {
	files, err := ioutil.ReadDir("../")
	if err != nil {
		panic(fmt.Sprintf("Cannot ReadDir of .., error: %v", err))
	}

	for _, f := range files {
		if filenamePattern.MatchString(f.Name()) {
			solutions = append(solutions, f.Name())
		}
	}

	return
}
