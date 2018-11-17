package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestDownloadQuestions(t *testing.T) {
	questions, err := DownloadQuestions()
	if err != nil {
		t.Errorf("DownloadQuestions fails, error: %v", err)
		return
	}

	if len(questions) == 0 {
		t.Error("Get zero-length questions slice")
		return
	}
}

func TestQuestionsCache(t *testing.T) {
	questions := map[int]Question{
		1: Question{ID: 1, Title: "One Sum", Slug: "one-sum", Level: "Easy", TestCase: "1234"},
		2: Question{ID: 2, Title: "two-sum"},
	}

	tempFilename := fmt.Sprintf("%s%vleetcode-test-%v.json", os.TempDir(), os.PathSeparator, time.Now().UnixNano())
	err := SaveQuestions(tempFilename, questions)
	if err != nil {
		t.Fatal("SaveQuestions fails, error:", err)
	}

	actualQuestions, err := LoadQuestions(tempFilename)
	if err != nil {
		t.Fatal("LoadQuestions fails, error:", err)
	}

	if len(questions) != len(actualQuestions) {
		t.Fatalf("Load %d questions instead of %d questions", len(questions), len(actualQuestions))
	}

	expectedQuestion1 := questions[1]
	actualQuestion1 := actualQuestions[1]
	if !reflect.DeepEqual(expectedQuestion1, actualQuestion1) {
		t.Fatalf("Found differences in question #1, \n  Expected: %v\n  Actual: %v", expectedQuestion1, actualQuestion1)
	}

	expectedQuestion2 := questions[2]
	actualQuestion2 := actualQuestions[2]
	if !reflect.DeepEqual(expectedQuestion2, actualQuestion2) {
		t.Fatalf("Found differences in question #1, \n  Expected: %v\n  Actual: %v", expectedQuestion2, actualQuestion2)
	}

}
