package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// FormatFilename reads question info from question struct
// and returns a unified filename
func FormatFilename(question Question) string {
	basename := fmt.Sprintf("%04d_%s", question.ID, strings.Replace(question.Slug, "-", "_", -1))
	ext := ""

	switch question.Type {
	case database:
		ext = ".sql"
	case shell:
		ext = ".sh"
	default:
		if _, found := question.Code[PrimaryLanguage]; found {
			ext = ".go"
		} else if _, found := question.Code[SecondaryLanguage]; found {
			ext = ".py"
		}
	}

	return basename + ext
}

// ParseSolutionFilename parses the filename of a solution.
// -1 will be returned if filename is not a solution filename.
func ParseSolutionFilename(filename string) int {
	pattern := regexp.MustCompile("^(\\d+)_[0-9a-zA-Z_]*?\\.(go|py|sql|sh)$")
	filename = strings.Trim(filename, "")

	match := pattern.FindStringSubmatch(filename)
	if match == nil {
		return -1
	}

	switch match[2] {
	case "go":
		if strings.HasSuffix(filename, "_test.go") {
			return -1
		}
	}

	problemID, _ := strconv.Atoi(match[1])
	return problemID
}

// GenerateQuestionCode returns a boilerplate for a question
func GenerateQuestionCode(question Question) string {
	problemURL := "Source: https://leetcode.com/problems/" + question.Slug

	switch question.Type {
	case database:
		testCase := formatDatabaseTestCase(question.TestCase)
		return "" +
			"-- " + problemURL + "\n" +
			testCase +
			"\n"
	case shell:
		testCase := formatGeneralTestCase(question.TestCase, "# ")
		return "" +
			"# " + problemURL + "\n" +
			"\n" +
			"# Test Case:\n" +
			testCase
	}

	if _, found := question.Code["golang"]; found {
		testCase := formatGeneralTestCase(question.TestCase, "")
		return "" +
			fmt.Sprintf("package problem%04d\n", question.ID) +
			"\n" +
			"/*\n" +
			problemURL + "\n" +
			"Test Case:\n" +
			testCase +
			"*/\n" +
			"\n" +
			question.Code["golang"]
	} else if _, found := question.Code["python"]; found {
		testCase := formatGeneralTestCase(question.TestCase, "")
		return "" +
			"'''\n" +
			problemURL + "\n" +
			"Test Case:\n" +
			testCase +
			"'''\n" +
			"\n" +
			question.Code["python"] + "\n" +
			"\n" +
			"if __name__ == '__main__':\n" +
			"    solver = Solution()\n"
	}

	panic("Cannot generate code for question")
}

// ParseFolderName parses the filename of a solution.
// -1 will be returned if the given filename is not a directory
func ParseFolderName(filename string) int {
	pattern := regexp.MustCompile("^(\\d+)-[0-9a-zA-Z\\-]*$")
	filename = strings.Trim(filename, "")

	match := pattern.FindStringSubmatch(filename)
	if match == nil {
		return -1
	}

	problemID, _ := strconv.Atoi(match[1])
	return problemID
}

// FormatDirectoryName reads question info from question struct
// and returns a unified directory name
func FormatDirectoryName(question Question) string {
	return fmt.Sprintf("%04d-%s", question.ID, question.Slug)
}

func formatGeneralTestCase(rawTestCase string, linePrefix string) string {
	result := ""

	lines := strings.Split(strings.Replace(rawTestCase, "\r", "", -1), "\\n")
	for _, line := range lines {
		result += linePrefix + line + "\n"
	}

	return result
}

func formatDatabaseTestCase(rawTestCase string) string {
	result := ""
	var testCase map[string]interface{}
	json.Unmarshal([]byte(rawTestCase), &testCase)

	// Construct the table
	tables := make(map[string][][]string)
	maxLength := make(map[string][]int)
	headers := testCase["headers"].(map[string]interface{})
	for tableName, rawRows := range testCase["rows"].(map[string]interface{}) {
		table := [][]string{toStringSlice(headers[tableName].([]interface{}))}

		for _, column := range table[0] {
			maxLength[tableName] = append(maxLength[tableName], len(column)+2)
		}

		for _, rawRow := range rawRows.([]interface{}) {
			row := toStringSlice(rawRow.([]interface{}))
			table = append(table, row)

			for index, column := range row {
				if len(column)+2 > maxLength[tableName][index] {
					maxLength[tableName][index] = len(column) + 2
				}
			}
		}
		tables[tableName] = table
	}

	// Format the table output
	for tableName, rows := range tables {
		result += "\n"
		result += "-- Table: " + tableName + "\n"
		for index, row := range rows {
			result += "-- | "
			for indexCol, col := range row {
				result += leftPadding(col, maxLength[tableName][indexCol]-2) + " | "
			}
			result += "\n"

			if index == 0 {
				result += "-- | "
				for indexCol := range row {
					result += strings.Repeat("-", maxLength[tableName][indexCol]-2) + " | "
				}
				result += "\n"
			}
		}
	}

	if argument, found := testCase["argument"]; found {
		result += "-- Argument:\n"
		switch argument.(type) {
		case []interface{}:
			result += "-- ["
			for index, item := range argument.([]interface{}) {
				result += fmt.Sprintf("%v", item)
				if index != len(argument.([]interface{}))-1 {
					result += ", "
				}
			}
			result += "]\n"
		case string:
			lines := strings.Split(strings.Replace(argument.(string), "\r", "", -1), "\n")
			for _, line := range lines {
				result += "-- " + line + "\n"
			}
		default:
			result += fmt.Sprintf("-- %v\n", argument)
		}
	}

	return result
}

func leftPadding(text string, length int) string {
	return fmt.Sprintf(fmt.Sprintf("%%-%ds", length), text)
}

func toStringSlice(src []interface{}) []string {
	buffer := make([]string, 0, len(src))
	for _, cell := range src {
		buffer = append(buffer, fmt.Sprintf("%v", cell))
	}
	return buffer
}
