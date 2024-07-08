package iteration

const count = 5

func Repeat(character string) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
