package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func generateINT(s string) int {
	arr := strings.Split(s, "-")
	start, err := strconv.Atoi(arr[0])
	end, eror := strconv.Atoi(arr[1])

	if eror != nil {
		panic(eror)
	}
	if err != nil {
		panic(err)
	}

	return rand.Intn(end-start) + start
}
func generateVARCHAR(s string) string {
	return "STRKA"
}

func selector(types []string, ranges []string) string {

	var result string = ""

	for i, tupe := range types {
		switch tupe {
		case "INT":
			fmt.Sprintf(result, "%v, ", generateINT(ranges[i]))
		case "VARCHAR":
			fmt.Sprintf(result, "%v, ", generateVARCHAR(ranges[i]))
		default:
			fmt.Println("NOT WORKING!")
		}
	}
	return result
}

func main() {
	result := "INSERT INTO "
	fmt.Println("==============================Генератор INSERT SQL-запроса==============================\n")
	fmt.Println("Справочник: После каждого заполненного поля ставьте пробел, как разделитель")
	fmt.Println("=====Формат=====\n\n'Название таблицы'\n'Введите названия столбцов через /(только для вставки, если вы не хотите вставлять значение в столбец - не указывайте его тут)'\nДалее вводите столбцы в формате INT/VARCHAR/TEXT/TIMESTAMP=1-100, 20, 50, 2007-08-0713:00:00|2008-08-0713:00:00")

	columns := strings.ReplaceAll(os.Args[2], "/", ", ")

	result = fmt.Sprintf(result+"%v (%v) VALUES ", os.Args[1], columns)

	columns_types := strings.Split(os.Args[3], "=")[0]
	columns_ranges := strings.Split(os.Args[3], "=")[1]

	types, ranges := strings.Split(columns_types, "/"), strings.Split(strings.Trim(columns_ranges, " "), ",")
	result = fmt.Sprintf(result+"(%v)", selector(types, ranges))

	fmt.Println(result)
}
