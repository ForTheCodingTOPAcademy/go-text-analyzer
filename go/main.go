/*Git Tips*/
//git init - используется когда репозиторий GitHub пуст и нужно инициализировать новый
//git clone "src/href" - используется когда нужно выгрузить уже существующий репозиторий с GitHub на нынешний ПК
//git pull - используется когда нужно обновить локальный репозиторий с GitHub репозитория (добавить все отсутствующие изменения)
//git push /ветка/ - отправляет когд на локальный репозиторий GitHub
//git commit - создает коммит
//git remote add origin <ссылка-на-ваш-удаленный-репозиторий> - служит для привязки репозитория на GitHub

/*
Working

go run go/main.go
go run go/main.go go/data.txt
go run go/main.go -top=3 go/data.txt
go run go/main.go -l=1 go/data.txt
go run go/main.go -top=3 -l=1 go/data.txt
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

type TextStats struct {
	WordCount   int
	SymbolCount int
	SpaceCount  int
	LineCount   int
	WordFreq    map[string]int
}

type WordStats struct {
	Word  string
	Count int
}

func analyzeFile(filename string) (TextStats, error) {

	file, err := os.Open(filename)

	if err != nil {
		return TextStats{}, err
	}

	defer file.Close() //Закрытие файла, которое выполниться в любом случае (при хорошем сценарии в конце)

	var textStats TextStats
	textStats.WordFreq = make(map[string]int)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		textStats.WordCount++

		cleanWord := strings.TrimFunc(strings.ToLower(word), func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		if len(cleanWord) > 0 {
			textStats.WordFreq[cleanWord]++
		}

	}

	file.Seek(0, 0)

	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		textStats.LineCount++
		textStats.SymbolCount += len(lineScanner.Text())
		for _, symbol := range lineScanner.Text() {
			if unicode.IsSpace(symbol) {
				textStats.SpaceCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return TextStats{}, err
	}
	if err := lineScanner.Err(); err != nil {
		return TextStats{}, err
	}
	return textStats, nil
}

func getTopWords(wordFreq map[string]int, topN int) []WordStats {
	var wordStats []WordStats
	for word, word_count := range wordFreq {
		wordStats = append(wordStats, WordStats{Word: word, Count: word_count})
	}
	sort.Slice(wordStats, func(i, j int) bool {
		return wordStats[i].Count > wordStats[j].Count
	})
	if len(wordStats) >= topN {
		return wordStats[0:topN]
	}
	return wordStats
}

func printStats(textStats TextStats, wordStats []WordStats, onlyLines bool) {
	if onlyLines {
		fmt.Println("Количество строк в файле: ", textStats.WordCount)
	} else {
		fmt.Printf("Количество слов: %d. Количество строк: %d. Количество символов: %d. Количество пробелов: %d\n", textStats.WordCount, textStats.LineCount, textStats.SymbolCount, textStats.SpaceCount)
	}
	for i := 0; i < len(wordStats); i++ {
		fmt.Printf("%v) Слово: %v. Встречается %v раз/а\n", i+1, wordStats[i].Word, wordStats[i].Count)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Укажите путь к файлу, который хотите проанализировать.")
		os.Exit(1)
	}

	topN := flag.Int("top", 5, "Amount of places at topWords list")
	onlyLines := flag.Bool("l", false, "Give stats as only lineCount")

	flag.Parse()

	textStats, err := analyzeFile(os.Args[len(os.Args)-1])

	if err != nil {
		log.Fatal(err)
	}

	wordStats := getTopWords(textStats.WordFreq, *topN)

	printStats(textStats, wordStats, *onlyLines)

}
