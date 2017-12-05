package stats

import (
	"sync"
)

type Summary struct {
	Count       int      `json:"count"`
	Top5Words   []string `json:"top5words"`
	Top5Letters []string `json:"top5letters"`
}

type stats struct {
	mutex           *sync.RWMutex
	wordFrequency   map[string]int
	letterFrequency map[rune]int
	top5Words       *Top5
	top5Letters     *Top5
}

func NewStats() *stats {
	return &stats{
		mutex:           new(sync.RWMutex),
		wordFrequency:   make(map[string]int),
		letterFrequency: make(map[rune]int),
		top5Words:       NewTop5(),
		top5Letters:     NewTop5(),
	}
}

func (s *stats) Record(word string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.recordWord(word)
	s.recordLetters(word)
}

func (s *stats) Summary() Summary {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return Summary{
		Count:       len(s.wordFrequency),
		Top5Words:   s.top5Words.List(),
		Top5Letters: s.top5Letters.List(),
	}
}

func (s *stats) recordWord(word string) {
	if _, exists := s.wordFrequency[word]; !exists {
		s.wordFrequency[word] = 0
	}
	s.wordFrequency[word] += 1

	s.top5Words.Submit(word, s.wordFrequency[word])
}

func (s *stats) recordLetters(word string) {
	for _, rune := range word {
		if _, exists := s.letterFrequency[rune]; !exists {
			s.letterFrequency[rune] = 0
		}
		s.letterFrequency[rune] += 1

		s.top5Letters.Submit(string(rune), s.letterFrequency[rune])
	}
}
