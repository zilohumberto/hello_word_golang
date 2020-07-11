package hash_tables

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) string{
	words := make(map[string]int)
	for _, _magazine := range magazine{
		words[_magazine] ++
	}
	for _, _note := range note{
		count, ok := words[_note]
		if !ok || count==0{
			return "No"
		}
		words[_note] = count-1
	}
	return "Yes"
}