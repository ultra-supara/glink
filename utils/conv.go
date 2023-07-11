package utils

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/ultra-supara/glink/static"
)

func PrintTxtFromTxt(fp string) {
	text := readFile(filepath.Join("txt", fp))
	fmt.Println(text)
}

func TxtFromText(fp string) string {
	text := readFile(filepath.Join("txt", fp))
	return text
}

func readFile(fp string) string {
	file, err := static.Txt.Open(fp)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//todo: fileを読み込んで出力する
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.String()
}
