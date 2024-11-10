package playfairCipher

func Verify(result *string) {
	for i, n := range *result {
		if string(n) == PLACEHOLDER {
			if i != len(*result)-1 {
				if i > 0 && i < len(*result)-1 {
					if (*result)[i-1] == (*result)[i+1] {
						*result = (*result)[:i] + (*result)[i+1:]
					}
				}
			} else if i == len(*result)-1 {
				if string(n) == PLACEHOLDER {
					*result = (*result)[:i]
				}
			}
		}
	}
}
