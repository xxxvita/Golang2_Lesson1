package main

import (
	"fmt"
	"os"
)

const CountFiles = 1000000

// Описатели открытых файлов
var fRefs []*os.File

func main() {
	fRefs = make([]*os.File, CountFiles)

	MakeBlockFiles("Data/", CountFiles)
	fmt.Println("Один миллион файлов созданы")
}

func MakeBlockFiles(dirName string, CountFiles int) {
	// Число уже созданных файлов
	var cntExistsFiles int = 0
	for cntExistsFiles != CountFiles {
		MakeFiles(dirName, &cntExistsFiles, CountFiles-cntExistsFiles)
	}
}

// Создаёт в директории dirName, пустые файлы в количестве cnt, числол успешно созданных
// файлов возвращает в cntExistsFiles
func MakeFiles(dirName string, cntExistsFiles *int, cnt int) {
	defer func() {
		if v := recover(); v != nil {
			for i := 0; i < *cntExistsFiles; i++ {
				if fRefs[i] != nil {
					fRefs[i].Close()
					fRefs[i] = nil
				}
			}
		}
	}()

	for i := 0; i < cnt; i++ {
		f, err := os.Create(fmt.Sprintf("%s/%d.txt", dirName, *cntExistsFiles+i))
		if err != nil {
			panic("Ошибка создания файла")
		}

		*cntExistsFiles = *cntExistsFiles + 1
		fRefs[*cntExistsFiles] = f
	}
}
