package selection


import(
	"strings"
	"errors"
	"strconv"
)




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
		singleResult, err := strconv.Atoi(formattedInput)
		if err != nil{
			return nil, err
		}
		if singleResult > max{
			return nil, errors.New("selection: given number is bigger than items list length")
		} else if singleResult <= 0{
			return nil, errors.New("selection: given number is smaller than one and cannot be sequence number")
		}
		return []int{singleResult - 1}, nil
	}

	return nil, errors.New("selection: not implemented")
}