package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct{
	XMLName xml.Name `xml:"post"`
	Id string         `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
	Xml string `xml:",innerxml"`
}

type Author struct{
	Id string `xml: "Id, attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil{
		fmt.Println("cannot open the xml file")
		return
	}
	defer xmlFile.Close()
	xmlData,  err := ioutil.ReadAll(xmlFile)
	if err != nil{
		fmt.Println("unable read data from the opened file handle")
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
