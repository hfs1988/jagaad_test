package common

func ConvertToMap(vals []string) map[string]bool {
	var valsMap = make(map[string]bool)
	for _, v := range vals {
		valsMap[v] = true
	}

	return valsMap
}

func IsExist(valsMap map[string]bool, val string) bool {
	if _, exist := valsMap[val]; !exist {
		return false
	}

	return true
}
