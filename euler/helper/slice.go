package helper

func RemoveAndReturnIndex[T any](slice []T, index int) ([]T, T) {
	var item T

	if index < 0 || index >= len(slice) {
		return slice, item
	}

	item = slice[index]
	return append(slice[:index], slice[index+1:]...), item
}
