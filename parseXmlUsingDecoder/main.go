package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}
type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}
type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

/*
func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("cannot open the file for reading")
		return
	}

	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	for {
		thisToken, err := decoder.Token()
		if err == io.EOF{
			fmt.Println("reached the end of the file")
			break
		}
		if err != nil {
			fmt.Println("cannot decode the line")
			fmt.Println(err)
			return
		}

		switch se := thisToken.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment"{
				var comment Comment
				decoder.DecodeElement(&comment, &se)
			}
		}


	}
}
*/

func main() {
	xmlFile, err := os.Open("post2.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}
		//fmt.Printf("%v",t)
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)

			}

		}
	}
}