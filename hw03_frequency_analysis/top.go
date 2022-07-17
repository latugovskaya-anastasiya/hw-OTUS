package hw03frequencyanalysis

import (
	"bufio"
	"sort"
	"strings"
)

type entity struct {
	count int
	text  string
}

type mapForCounting map[string]*entity

/*Top10 Необходимо написать Go функцию, принимающую на вход строку с текстом и
возвращающую слайс с 10-ю наиболее часто встречаемыми в тексте словами.

Если слова имеют одинаковую частоту, то должны быть отсортированы **лексикографически**.

* Словом считается набор символов, разделенных пробельными символами.

* Если есть более 10 самых частотых слов (например 15 разных слов встречаются ровно 133 раза,
остальные < 100), то следует вернуть 10 лексикографически первых слов.

* Словоформы не учитываем: "нога", "ногу", "ноги" - это разные слова.

* Слово с большой и маленькой буквы считать за разные слова. "Нога" и "нога" - это разные слова.

* Знаки препинания считать "буквами" слова или отдельными словами.
"-" (тире) - это отдельное слово. "нога," и "нога" - это разные слова.
*/
func Top10(text string) []string {
	storage := mapForCounting{}
	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		ok := having(storage, word)

		switch {
		case !ok:
			e := new(entity)
			e.count++
			e.text = word
			storage[word] = e

			continue

		default:
			storage[word].count++
		}
	}

	rtn := makeSlice(storage)

	return rtn
}

func makeSlice(w mapForCounting) []string {
	totalEntitiesCount := len(w)
	entities := make([]*entity, 0, totalEntitiesCount)

	for _, ent := range w {
		entities = append(entities, ent)
	}

	sort.SliceStable(
		entities,
		func(i, j int) bool {
			return entities[i].text < entities[j].text
		},
	)

	sort.SliceStable(
		entities,
		func(i, j int) bool {
			return entities[i].count > entities[j].count
		},
	)

	rtn := make([]string, 0, 10)
	for _, e := range entities {
		rtn = append(rtn, e.text)

		if len(rtn) == 10 {
			break
		}
	}

	return rtn
}

// having func check if the word already stored.
func having(storage map[string]*entity, word string) bool {
	_, ok := storage[word]
	return ok
}
