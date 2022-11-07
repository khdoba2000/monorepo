package utils

// InEnums ...
func InEnums(str string, enums []string) bool {
	for _, enum := range enums {
		if enum == str {
			return true
		}
	}
	return false
}