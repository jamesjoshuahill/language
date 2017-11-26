package stats

import "strings"

type Summary struct {
	Count       int      `json:"count"`
	Top5Words   []string `json:"top5words"`
	Top5Letters []string `json:"top5letters"`
}

type stats struct {
	wordsSeen     int
	wordFrequency map[string]int
}

func New() *stats {
	return &stats{
		wordFrequency: make(map[string]int),
	}
}

func (l *stats) Record(language string) {
	words := strings.Fields(language)
	for _, word := range words {
		if _, exists := l.wordFrequency[word]; !exists {
			l.wordsSeen++
			l.wordFrequency[word] = 1
		}
		l.wordFrequency[word] += 1
	}
}

func (l *stats) Summary() Summary {
	return Summary{
		Count: l.wordsSeen,
	}
}
