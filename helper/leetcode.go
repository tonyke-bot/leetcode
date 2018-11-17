package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// QuestionType is the definition of question type
type QuestionType int

const (
	normal   QuestionType = 0
	database QuestionType = 1
	shell    QuestionType = 2
)

// CacheFilename stands for a unified cache filename
var CacheFilename = "../problems.cache"

// QuestionTag  the tag of question
type QuestionTag struct {
	Name string
	Slug string
}

// Question is a struct to store leetcode problem
type Question struct {
	ID       int
	ActualID int
	Type     QuestionType
	Title    string
	Slug     string
	Level    string
	Tags     []QuestionTag
	Code     map[string]string
	TestCase string
	PaidOnly bool
}

// DownloadQuestions downloads question from Leetcode and return
// as a Question struct
func DownloadQuestions() (map[int]Question, error) {
	questions := map[int]Question{}
	resp, err := query("query {\n" +
		"    questions: allQuestions {\n" +
		"        actualId: questionId\n" +
		"        id: questionFrontendId\n" +
		"        title: questionTitle\n" +
		"        slug: questionTitleSlug\n" +
		"        level: difficulty\n" +
		"        paidOnly: isPaidOnly\n" +
		"        tags: topicTags { name slug }\n" +
		"        code: codeSnippets { slug: langSlug code }\n" +
		"        testCase: sampleTestCase\n" +
		"    }\n" +
		"}")

	if err != nil {
		return questions, err
	}

	rawQuestions := resp["data"].(map[string]interface{})["questions"].([]interface{})
	for _, rawQuestion := range rawQuestions {
		question := parseQuestion(rawQuestion.(map[string]interface{}))
		questions[question.ID] = question
	}

	return questions, nil
}

// SaveQuestions serializes questions struct and write to file
func SaveQuestions(filename string, questions map[int]Question) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		fmt.Printf("Cannot %s for writing, error: %v\n", filename, err)
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(questions)

	return nil
}

// LoadQuestions loads serialized questions from file and deserialize them into
// Question struct
func LoadQuestions(filename string) (map[int]Question, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	defer file.Close()
	if err != nil {
		fmt.Printf("Cannot %s for reading, error: %v\n", filename, err)
		return nil, err
	}

	var result map[int]Question
	decoder := json.NewDecoder(file)
	decoder.Decode(&result)

	return result, nil
}

func query(content string) (map[string]interface{}, error) {
	csrf := "qBxU4UYyArpJTErnIzkLPyDtBEZibG5rotgOvuvfSnSiiXaM7WYQTejn5upGUAdD"

	reqBody, _ := json.Marshal(map[string]string{"query": content})
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referer", "https://leetcode.com/problems/")
	req.Header.Add("User-Agent", "LeetCode/1.0.0")
	req.Header.Add("x-csrftoken", csrf)
	req.Header.Add("Cookie", "csrftoken="+csrf)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	decoder := json.NewDecoder(gzipReader)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func parseQuestion(rawQuestion map[string]interface{}) Question {
	id, _ := strconv.Atoi(rawQuestion["id"].(string))
	actualID, _ := strconv.Atoi(rawQuestion["actualId"].(string))

	result := Question{
		ID:       id,
		ActualID: actualID,
		Title:    rawQuestion["title"].(string),
		Slug:     rawQuestion["slug"].(string),
		Level:    rawQuestion["level"].(string),
		TestCase: rawQuestion["testCase"].(string),
		PaidOnly: rawQuestion["paidOnly"].(bool),
		Code:     make(map[string]string),
	}

	for _, rawTag := range rawQuestion["tags"].([]interface{}) {
		result.Tags = append(result.Tags, QuestionTag{
			Name: rawTag.(map[string]interface{})["name"].(string),
			Slug: rawTag.(map[string]interface{})["slug"].(string),
		})
	}

	if rawQuestion["code"] != nil {
		for _, codeSnippet := range rawQuestion["code"].([]interface{}) {
			slug := codeSnippet.(map[string]interface{})["slug"].(string)
			code := codeSnippet.(map[string]interface{})["code"].(string)
			result.Code[slug] = code
		}
	}

	if _, found := result.Code["mysql"]; found {
		result.Type = database
	} else if _, found := result.Code["bash"]; found {
		result.Type = shell
	} else {
		result.Type = normal
	}

	return result
}
