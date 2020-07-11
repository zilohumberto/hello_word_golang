package hashmultimaps

func twoStrings(s1 string, s2 string) string {
	words := make(map[int32]int32)
	for _, word := range s1{
		words[word] ++
	}
	for _, word := range s2{
		_, ok := words[word]
		if ok{
			return "YES"
		}
	}
	return "NO"
}
