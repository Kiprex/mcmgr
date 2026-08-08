package selection

import (
	"testing"
	//"fmt"
	"slices"
)



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

func TestParseSingleNumReturnsOneResult(t *testing.T) {
	// вводим исходные и ожидаемые  выходные данные
    input := "1"
	max := 3
	want := []int{0}


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

func TestParseNumBiggerThenMaxError(t *testing.T){
	// вводим исходные и ожидаемые  выходные данные
	input := "4"
	max := 3
	
	// запускаем функцию
    _, err := Parse(input, max)

	if err == nil{
		t.Fatalf("Parse(%q, %d): expected error, got nil", input, max)
	}

}