package selection


import(
	"strings"
	"errors"
)



func Parse(input string, max int) ([]int, error){
	if strings.TrimSpace(input) == "" {
		if max <= 0 {
			return []int{}, nil
		}

		selected := make([]int, 0, max)
		for i := 0; i < max; i++ {
			selected = append(selected, i)
		}

		return selected, nil
	}

	return nil, errors.New("selection: not implemented")
}