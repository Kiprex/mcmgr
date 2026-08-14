package selection

import (
	"errors"
	"strconv"
	"strings"
)

// parseSingle читает строку и пытается привести ее к числу, входящему в диапазон от нуля до max.
//
// Параметры:
//   - input: строка для обработки
//   - max: максимальное значение диапазона
//
// Возвращает:
//   - получившиеся число
//   - ошибку при неудавшемся парсинге
func parseSingle(input string, max int) (int, error) {
	singleResult, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	if singleResult > max {
		return -1, errors.New("selection: given number is bigger than items list length")
	} else if singleResult <= 0 {
		return -1, errors.New("selection: given number is smaller than one and cannot be sequence number")
	}
	return singleResult - 1, nil
}

// parseRange читает строку и пытается привести ее диапазону чисел, входящему в диапазон от нуля до max.
//
// Параметры:
//   - input: строка для обработки
//   - max: максимальное значение диапазона
//
// Возвращает:
//   - получившийся срез чисел
//   - ошибку при неудавшемся парсинге
func parseRange(input string, max int) ([]int, error) {
	if strings.Count(input, "-") > 1 {
		return nil, errors.New("Invalid input")
	}
	caseList := strings.Split(input, "-")
	left, err := parseSingle(caseList[0], max)

	if err != nil {
		return nil, err
	}

	right, err := parseSingle(caseList[len(caseList)-1], max)
	if err != nil {
		return nil, err
	}
	if left > right {
		return nil, errors.New("Invalid range edges")
	}
	result := make([]int, right-left+1)
	for i := range result {
		result[i] = left + i
	}
	return result, nil
}

func removeDuplicates(inputRange []int) []int {
	result := make([]int, 0, len(inputRange))

	// map результатов: ключ - элемент
	unsortedResult := make(map[int]bool)

	for _, item := range inputRange {
		if !unsortedResult[item] {
			unsortedResult[item] = true
			result = append(result, item)
		}
	}
	return result
}

func Parse(input string, max int) ([]int, error) {
	formattedInput := strings.TrimSpace(input)

	switch {
	case strings.Contains(formattedInput, ","): //  если сложное выражение
		caseList := strings.Split(formattedInput, ",")
		preResult := []int{}
		for _, item := range caseList {
			preRange, err := Parse(item, max)
			if err != nil {
				return nil, err
			}
			for _, rangeItem := range preRange {
				preResult = append(preResult, rangeItem)
			}
		}
		result := removeDuplicates(preResult)
		return result, nil

	case strings.Contains(formattedInput, "-"): // если диапазон
		return parseRange(formattedInput, max)

	case formattedInput == "":
		if max <= 0 {
			return []int{}, nil
		}

		selected := make([]int, 0, max)
		for i := 0; i < max; i++ {
			selected = append(selected, i)
		}
		return selected, nil
	default:
		result, err := parseSingle(formattedInput, max)
		if err != nil {
			return nil, err
		}
		return []int{result}, nil
	}
}
