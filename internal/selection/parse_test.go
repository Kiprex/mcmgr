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