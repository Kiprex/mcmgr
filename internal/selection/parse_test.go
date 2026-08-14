package selection

import (
	"testing"
	//"fmt"
	"slices"
)

type testCase struct {
	Name    string
	Input   string
	Max     int
	Want    []int
	WantErr bool
}

func runTestCase(tt testCase, t *testing.T) {
	// запускаем функцию
	result, err := Parse(tt.Input, tt.Max)

	// проверка работоспособности самой функции
	if err != nil && !tt.WantErr {
		t.Fatalf("Parse(%q, %d): unexpected error: %v", tt.Input, tt.Max, err)
	}
	if tt.WantErr && err == nil {
		t.Fatalf("Parse(%q, %d): expected error, got result: %v", tt.Input, tt.Max, result)
	}
	// проверка соотвествия результата ожиданиям
	if !tt.WantErr && !slices.Equal(result, tt.Want) {
		t.Errorf("Parse(%q, %d) = %v, want %v", tt.Input, tt.Max, result, tt.Want)
	}
}

func TestParseEmptyInputSelectsAll(t *testing.T) {
	// вводим исходные и ожидаемые  выходные данные
	input := ""
	max := 3
	want := []int{0, 1, 2}

	// запускаем функцию
	result, err := Parse(input, max)

	// проверка работоспособности самой функции
	if err != nil {
		t.Fatalf("Parse(%q, %d): unexpected error: %v", input, max, err)
	}

	// проверка соотвествия результата ожиданиям
	if !slices.Equal(result, want) {
		t.Errorf("Parse(%q, %d) = %v, want %v", input, max, result, want)
	}
}

func TestParseSingleNumsReturnsOneResult(t *testing.T) {

	tests := []testCase{
		{"1 returns 0", "1", 3, []int{0}, false},
		{"2 returns 1", "2", 3, []int{1}, false},
		{"0 returns error", "0", 3, nil, true},
		{"4 (max=3) returns error", "4", 3, nil, true},
		{"3 returns 2", "3", 3, []int{2}, false},
		{"10 returns 9", "10", 10, []int{9}, false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			runTestCase(tt, t)
		})
	}
}
func TestParseMultipleSingleNums(t *testing.T) {
	tests := []testCase{
		{"1,2 returns [0,1]", "1, 2", 3, []int{0, 1}, false},
		{"1,3 returns [0,2]", "1, 3", 3, []int{0, 2}, false},
		{"1,2,3 returns [0,1,2]", "1, 2, 3", 3, []int{0, 1, 2}, false},
		{"1,1 returns [0]", "1, 1", 3, []int{0}, false},
		{"1,2,1 returns [0,1]", "1, 2, 1", 3, []int{0, 1}, false},
		{"1,0,2 returns error", "1, 0, 2", 3, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			runTestCase(tt, t)
		})
	}
}

func TestParseRanges(t *testing.T) {
	tests := []testCase{
		{"1-3 returns [0,1,2]", "1-3", 3, []int{0, 1, 2}, false},
		{"3-1 returns error", "3-1", 3, nil, true},
		{"1-3-5 returns error", "1-3-5", 3, nil, true},
		{"1--3 returns error", "1--3", 3, nil, true},
		{"1- returns error", "1-", 3, nil, true},
		{"-3 returns error", "-3", 3, nil, true},
		{"3-3 returns [2]", "3-3", 3, []int{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			runTestCase(tt, t)
		})
	}
}


func TestParseComplexExpression(t *testing.T){
	tests := []testCase{

		{"1,3-5,7 returns [0,2,3,4,6]", "1,3-5, 7", 7, []int{0,2,3,4,6}, false},
		{"1-3,3-5 returns [0,1,2,3,4]", "1-3, 3-5", 5, []int{0,1,2,3,4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			runTestCase(tt, t)
		})
	}
}