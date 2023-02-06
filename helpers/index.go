package helpers

func IsNil(data string) bool {
	return data == ""
}

func IsEmpty(data interface{}) bool {
	return data == nil || data == ""
}
