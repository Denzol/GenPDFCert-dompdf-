package main

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"os/exec"
)

//Структура для заполнения шаблона
type data struct {
	Title   string
	Student string
	Course  string
	Mentors string
	Date    string
}

// Заполняем шаблон данными из структуры
func ParseTemplate(templateFileName string, data interface{}) {

	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		errors.New("Template parsing error!")
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, data)
	if err != nil {
		errors.New("Template generation error!")
	}

	readBuf, err := ioutil.ReadAll(buf)
	if err != nil {
		errors.New("Reading file error!")
	}

	err = ioutil.WriteFile(templateFileName, readBuf, 0777)
	if err != nil {
		errors.New("Writing file error!")
	}

}

func main() {

	templateData := data{
		Title:   "Certificate Golang School",
		Student: "Khramtsov Denis",
		Course:  "Become a gopher",
		Mentors: "Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling",
		Date:    "08.09.2022",
	}
	//Вызываем функцию для заполнения шаблона данными из структуры
	ParseTemplate("./dompdf/template.html", templateData)

	//Имитация запуска PHP в командной строке с передачей двух значений:библиотеки dompdf и заполненного шаблона данных html
	cmd := exec.Command("php", "./dompdf/dompdf.php", "./dompdf/template.html")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	//Считывание данных из сгенерированного файла PDF для перезаписи в нужную нам папку
	streamPDFbytes, err := ioutil.ReadFile("./dompdf/template.pdf")
	if err != nil {
		log.Fatal(err)
	}
	//Перезапись данных файла PDF в нужную нам папку
	err = ioutil.WriteFile("./Certificates/Certificate.pdf", streamPDFbytes, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
