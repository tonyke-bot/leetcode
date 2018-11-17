package main

import (
	"fmt"
	"testing"
)

func TestFormatDatabaseTestCase(t *testing.T) {
	result := formatDatabaseTestCase("{\"argument\": 1.24235, \"headers\": {\"Person\": [\"PersonId\", \"LastName\", \"FirstName\"], \"Address\": [\"AddressId\", \"PersonId\", \"City\", \"State\"]}, \"rows\": {\"Person\": [[1, \"Wang\", \"Allen\"]], \"Address\": [[1, 2, \"New York City\", \"New York\"]]}}")
	fmt.Println(result)
}

func TestParseSolutionFilename(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"9_one_sum.go"}, 9},
		{"", args{"10_one_sum.sql"}, 10},
		{"", args{"12_one_sum.sh"}, 12},
		{"", args{"123_one_sum.py"}, 123},
		{"", args{"1_one_sum_test.go"}, -1},
		{"", args{"link_node.py"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseSolutionFilename(tt.args.filename); got != tt.want {
				t.Errorf("ParseSolutionFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
