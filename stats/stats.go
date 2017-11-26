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
	top5Words     *Top5
}

func NewStats() *stats {
	return &stats{
		wordFrequency: make(map[string]int),
		top5Words:     NewTop5(),
	}
}

func (s *stats) Record(language string) {
	words := strings.Fields(language)
	for _, word := range words {
		s.recordWord(word)
		s.top5Words.Submit(word, s.wordFrequency[word])
	}
}

func (s *stats) Summary() Summary {
	return Summary{
		Count:     s.wordsSeen,
		Top5Words: s.top5Words.List(),
	}
}

func (s *stats) recordWord(word string) {
	if _, exists := s.wordFrequency[word]; !exists {
		s.wordsSeen++
		s.wordFrequency[word] = 0
	}
	s.wordFrequency[word] += 1
}
