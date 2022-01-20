package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

// Парсим URL сайта, чтобы использлвать как имя для нового файла
func fileNameParse(site string) string {
	urls := strings.Split(site, "/")
	for _, val := range urls {
		fmt.Println(val)
	}
	return urls[2] + ".html"
}

func download(site string) {
	// Отправляем get запрос
	resp, err := http.Get(site)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fileName := fileNameParse(site)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Записываем в файл тело ответа
	_, err = io.Copy(file, resp.Body)
}

func main() {
	site := flag.String("s", "https://www.youtube.com/", "site")

	flag.Parse()

	// Проверяем, что аргументов передан сайт
	if ok, err := regexp.MatchString("^(http|https)://", *site); ok == true && err == nil {
		download(*site)
	} else {
		log.Fatal("invalid url")
	}

}
