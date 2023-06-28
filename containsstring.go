package asciiART

/*
	This function will test if the letter to be printed is to be colored or not
	inputs:
		sub_str:		letters to be colored
		input_str:		the letter to be tested

	output:
		indicies		positions of each found word [beg, end]
		sub_len			length of the sub-string

*/

func ContainsString(sub_str string, input_str string) ([][]int, int) {
	//index := strings.Index(input_str, sub_str)
	var indicies [][]int
	sub_len := len(sub_str)

	for idx := 0; idx <= len(input_str) && idx+sub_len <= len(input_str); idx++ {
		if sub_str == input_str[idx:idx+sub_len] {
			indicies = append(indicies, []int{idx, idx + sub_len})
		}
	}
	return indicies, sub_len
}
