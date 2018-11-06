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

// CacheFilename stands for a unified cache filename
var CacheFilename = "../problems.cache"

// Question is a struct to store leetcode problem
type Question struct {
	ID       int
	Title    string
	Slug     string
	Level    string
	Tags     []string
	Code     string
	TestCase string
	PaidOnly bool
}

// DownloadQuestions downloads question from Leetcode and return
// as a Question struct
func DownloadQuestions() (map[int]Question, error) {
	questions := map[int]Question{}
	resp, err := query("query {\n" +
		"    questions: allQuestions {\n" +
		"        id: questionId\n" +
		"        title: questionTitle\n" +
		"        slug: questionTitleSlug\n" +
		"        level: difficulty\n" +
		"        paidOnly: isPaidOnly\n" +
		"        tags: topicTags { name }\n" +
		"        code: codeDefinition\n" +
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

func getQuestions(forceUpdate bool) (questions map[int]Question) {
	var err error
	shouldDownload := forceUpdate

	if !forceUpdate {
		questions, err = LoadQuestions(CacheFilename)
		if err != nil {
			fmt.Printf("Fail to load questions from cache, error:%v\n", err)
			shouldDownload = true
		}
	}

	if shouldDownload {
		questions, err = DownloadQuestions()
		if err != nil {
			panic(err)
		}
		SaveQuestions(CacheFilename, questions)
	}

	return questions
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
	targetLanguage := "golang"
	id, _ := strconv.Atoi(rawQuestion["id"].(string))

	result := Question{
		ID:       id,
		Title:    rawQuestion["title"].(string),
		Slug:     rawQuestion["slug"].(string),
		Level:    rawQuestion["level"].(string),
		TestCase: rawQuestion["testCase"].(string),
		PaidOnly: rawQuestion["paidOnly"].(bool),
	}

	for _, rawTag := range rawQuestion["tags"].([]interface{}) {
		result.Tags = append(result.Tags, rawTag.(map[string]interface{})["name"].(string))
	}

	if rawQuestion["code"] != nil {
		var rawCodes []interface{}
		json.Unmarshal([]byte(rawQuestion["code"].(string)), &rawCodes)

		for _, rawCode := range rawCodes {
			if rawCode.(map[string]interface{})["value"].(string) == targetLanguage {
				result.Code = rawCode.(map[string]interface{})["defaultCode"].(string)
				break
			}
		}
	}
	return result
}
