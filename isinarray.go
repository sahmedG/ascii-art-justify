package asciiART

/*
	inputs:
		num			number to be tested if it is in the range

	output:
		bool		(true or false)
*/

func IsInRange(sub_ranges [][]int, num int) bool {

	for i := 0; i < len(sub_ranges); i++ {
		if num < sub_ranges[i][1] && num >= sub_ranges[i][0] {
			return true
		}
	}
	return false
}
