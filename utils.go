package utils

func InSlice(str string, slices []string) bool {
	var res bool = false
	for _, slice := range slices {
		if slice == str {
			res = true
			break
		}
	}
	return res
}

func InSliceInt(dig int, slices []int) bool {
	var res bool = false
	for _, slice := range slices {
		if slice == dig {
			res = true
			break
		}
	}
	return res
}
