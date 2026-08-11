package selection

import (
	"testing"
	//"fmt"
	"slices"
)


type testCase struct {
	Name string
	Input string
	Max int
	Want []int
	WantErr bool
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
	if !slices.Equal(result, want){
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
	}


	for _, tt := range tests{

		t.Run(tt.Name, func(t *testing.T){
			// запускаем функцию
    		result, err := Parse(tt.Input, tt.Max)

			// проверка работоспособности самой функции
			if err != nil && !tt.WantErr{
				t.Fatalf("Parse(%q, %d): unexpected error: %v", tt.Input, tt.Max, err)
			}
			if tt.WantErr && err == nil{
				t.Fatalf("Parse(%q, %d): expected error, got result: %v", tt.Input, tt.Max, result)
			}
			// проверка соотвествия результата ожиданиям
			if !tt.WantErr && !slices.Equal(result, tt.Want){
				t.Errorf("Parse(%q, %d) = %v, want %v", tt.Input, tt.Max, result, tt.Want)
			}
		})

		
	}
}