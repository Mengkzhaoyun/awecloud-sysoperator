package common

type array2Map struct{}

// Array2Map , Support Contains Append Method to Array
var Array2Map = array2Map{}

// Contains , 检查是否包含某String
func (a *array2Map) Contains(slice []string, arg string) bool {
	if len(slice) == 0 {
		return false
	}

	xMap := TurnArray2Map(slice)
	_, ok := xMap[arg]
	if ok {
		return true
	}
	return false
}

// Append , 检查是否包含某String，失败追加string，成功则不追加string
func (a *array2Map) Append(slice []string, arg string) []string {
	if len(slice) == 0 {
		return append(slice, arg)
	}

	xMap := TurnArray2Map(slice)
	_, ok := xMap[arg]
	if ok {
		return append(slice, arg)
	}
	return slice
}

// Remove , 在元素中检查并删除string
func (a *array2Map) Remove(slice []string, arg string) []string {
	if len(slice) == 0 {
		return slice
	}

	result := []string{}
	for _, str := range slice {
		if str == arg {
			continue
		}
		result = append(result, str)
	}
	return result
}

// TurnArray2Map ,
func TurnArray2Map(slice []string) map[string]struct{} {
	result := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		result[s] = struct{}{}
	}
	return result
}
