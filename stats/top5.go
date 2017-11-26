package stats

type Top5 struct {
	entries map[string]int
}

func NewTop5() *Top5 {
	return &Top5{entries: make(map[string]int)}
}

func (t *Top5) Submit(entry string, score int) *Top5 {
	if score == 0 {
		return t
	}
	t.entries[entry] = score
	if len(t.entries) > 5 {
		lowest, _ := t.Lowest()
		delete(t.entries, lowest)
	}
	return t
}

func (t *Top5) List() []string {
	var top5 []string
	for entry := range t.entries {
		top5 = append(top5, entry)
	}
	return top5
}

func (t *Top5) Lowest() (lowestEntry string, lowestScore int) {
	for entry, score := range t.entries {
		if lowestScore == 0 || score < lowestScore {
			lowestEntry = entry
			lowestScore = score
		}
	}
	return
}
