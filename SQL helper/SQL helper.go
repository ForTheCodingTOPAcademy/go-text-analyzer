package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	result := "INSERT INTO "
	fmt.Println("==============================Генератор INSERT SQL-запроса==============================\n")
	fmt.Println("Справочник: После каждого заполненного поля ставьте пробел, как разделитель")
	fmt.Println("=====Формат=====\n\n'Название таблицы'\n'Введите названия столбцов через /(только для вставки, если вы не хотите вставлять значение в столбец - не указывайте его тут)'\nДалее вводите столбцы в формате -INT/VARCHAR/TEXT/TIMESTAMP=(1-100, 20, 50, 2007-08-0713:00:00|2008-08-0713:00:00)")
	fmt.Println(os.Args[1:len(os.Args)])
	columns := strings.ReplaceAll(os.Args[2], "/", ", ")
	result = fmt.Sprintf(result+"%v (%v) VALUES ", os.Args[1], columns)
	fmt.Println(result)
	columns_types := strings.Split(os.Args[3], "=")[0]
	columns_ranges := strings.Split(os.Args[3], "=")[1]
	fmt.Println(columns_types, columns_ranges)
}
