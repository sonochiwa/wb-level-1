package exercise_15

import (
	"fmt"
	"log"
	"os"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string
//
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
//
// func main() {
// 	someFunc()
// }

// Создание большой строки может привести к утечкам памяти
var justString string

func someFunc() {
	// Сохранение большой строки в переменную может привести
	// к утечке памяти, поэтому лучше записать данные в файл
	f, err := os.Create("huge_string.txt") // создаем пустой файл
	if err != nil {
		panic(err)
	}
	defer f.Close() // не забываем закрыть его

	// создаем большую строку, помещаем ее в файл, и после
	// завершения выполнения функции GC почистит память
	err = createHugeString(1<<10, *f)
	if err != nil {
		fmt.Println(err)
	}

	// Читаем первые 100 символов из ранее созданного файла
	v, err := os.ReadFile("huge_string.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Если бы все-таки было необходимо писать данные в переменную,
	// а не в файл, то для более оптимального чтения из большого объема
	// мы бы использовали буферицаю
	// justString = bytes.NewBuffer(v[:100]).String()
	// но, мы записали данные в файл, поэтому спокойно можем записать
	// первые 100 символов из файла в переменную justString
	justString = string(v[:100])
}

func createHugeString(size int, f os.File) error {
	// Создаем массив байт размером 1 << 10
	result := make([]byte, size)

	// Заполняем массив байт значениями
	for i := 0; i < size; i++ {
		result[i] = byte('a')
	}

	// Пишем строку в файл
	_, err := f.WriteString(string(result))
	if err != nil {
		return err
	}

	return nil
}

func Run() {
	someFunc()
	fmt.Println(justString)
}
