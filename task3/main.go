/*
Удаляем папку Data со всеми созданными файлами, запускаем программу
rm -rf Data/ && mkdir Data && go run main.go
Считаем сколько файлов в папке Data
ls Data/| wc - l
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const CountFiles = 1000000
const DirectoryName = "Data/"

// Число уже созданных файлов
var cntExistsFiles int = 0

// Описатели открытых файлов
var errorTooManyOpenFiles = errors.New("Ошибка открытия большого числа файлов")

func main() {
	err := MakeBlockFiles(DirectoryName, CountFiles)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	fmt.Printf("Файлов создано: %d\n", cntExistsFiles)
}

func MakeBlockFiles(dirName string, CountFiles int) error {
	cntIteration := 0

	// Открываются блоки файлов до возникновения ошибки или до достижения
	// открытых файлов числа CountFiles
	for cntExistsFiles < CountFiles {
		err := MakeFiles(dirName, &cntExistsFiles, CountFiles-cntExistsFiles)
		if err != nil {
			// Если ошибка не связана с числом открытых файлов, то выходим
			if !errors.Is(err, errorTooManyOpenFiles) {
				return err
			}
		}

		cntIteration++
		fmt.Printf("Количество итераций: %d. Число созданных файлов: %d\n", cntIteration, cntExistsFiles)
	}

	return nil
}

// Создаёт в директории dirName, пустые файлы в количестве cnt, числол успешно созданных
// файлов возвращает в cntExistsFiles
func MakeFiles(dirName string, cntExistsFiles *int, cnt int) (err error) {
	for i := 0; i < cnt; i++ {
		f, err := os.Create(fmt.Sprintf("%s/%d.txt", dirName, *cntExistsFiles+1))
		if err != nil {
			if strings.Contains(err.Error(), "too many open files") {
				return fmt.Errorf("%w, %s", errorTooManyOpenFiles, err)
			}

			return err
		}

		*cntExistsFiles = *cntExistsFiles + 1

		defer func(f *os.File) {
			f.Close()
		}(f)
	}

	return nil
}
