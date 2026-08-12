package selection


import(
	"strings"
	"errors"
	"strconv"
)


func parseSingle(input string, max int) (int, error){
	singleResult, err := strconv.Atoi(input)
	if err != nil{
		return -1, err
	}
	if singleResult > max{
		return -1, errors.New("selection: given number is bigger than items list length")
	} else if singleResult <= 0{
		return -1, errors.New("selection: given number is smaller than one and cannot be sequence number")
	}
	return singleResult - 1, nil
}




func Parse(input string, max int) ([]int, error){

	formattedInput := strings.TrimSpace(input)
	if formattedInput == "" {
		if max <= 0 {
			return []int{}, nil
		}

		selected := make([]int, 0, max)
		for i := 0; i < max; i++ {
			selected = append(selected, i)
		}
		return selected, nil
	} 
	if len(formattedInput) == 1 {
		result, err := parseSingle(formattedInput, max)
		if err != nil{
			return nil, err
		}
		return []int{result}, nil
	}
	if strings.Contains(formattedInput, ","){
		caseList := strings.Split(formattedInput, ",")

		result := make([]int, 0, len(caseList))

		// map результатов: ключ - элемент
		unsortedResult := make(map[int]bool)

		for _, item := range caseList{
			formattedItem, err := parseSingle(strings.TrimSpace(item), max)
			if err != nil{
				return nil, err
			}
			if !unsortedResult[formattedItem]{
				unsortedResult[formattedItem] = true
				result = append(result, formattedItem)
			}
		}

		return result, nil
	}

	return nil, errors.New("selection: not implemented")
}