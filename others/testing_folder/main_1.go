package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Bible struct {
	XMLName xml.Name `xml:"XMLBIBLE"`
	Version string `xml:"biblename,attr"`
	Book []Book `xml:"BIBLEBOOK"`
}

type Book struct {
	Name string `xml:"bname,attr"`
	ListBooks []string `xml:"-"`
	Chapter []Chapter `xml:"CHAPTER"`
}

type Chapter struct {
	Name string `xml:"cnumber,attr"`
	Verse []Verse `xml:"VERS"`
}

type Verse struct {
	Number string `xml:"vnumber,attr"`
	Text string `xml:",chardata"`
}


func main() {
	version := "nkjv"
	bible := open(version)

	// Get the Translation
	fmt.Println("Bible Translation:", bible.Version)

	// fmt.Printf("%v - %v\n", bible.Book[0].Chapter[1].Verse[1].Number, bible.Book[0].Chapter[1].Verse[1].Text)
	fmt.Println(bible.Book)

	// Get the Passage
	// fmt.Printf("%v - %v\n", bible.Content.Books[65].Chapters[0].Verses[0].Number, bible.Content.Books[65].Chapters[0].Verses[0].Content)

	// for i := 0; i < len(bible.Content.Books); i++ {
	// 	fmt.Println(bible.Content.Books[i].Name)
	// }

}

func open(version string) Bible {
	// Open the file
	xmlFile, err := os.Open("./data/bibles/" + version + ".xml")

	if err != nil {
		fmt.Println("Error Opening XML file:", err)
	}

	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)

	if err != nil {
		fmt.Println("Error reading XML data:", err)
	}

	var bible Bible

	xml.Unmarshal(byteValue, &bible)

	return bible
}

// func (b *Bible) getBooks() {
// 	for i := 0; i < len(b.Content.Books); i++ {
// 		b.Content.ListBooks = append(b.Content.ListBooks, b.Content.Books[i].Name)
// 	}
// }

	// for i := 0; i < len(b.Content.Books); i++ {
	// 	b.Content.ListBooks = append(b.Content.ListBooks, b.Content.Books[i].Name)
	// }

func (b *Bible) getBooks() {
	for i := 0; i < len(b.Book.Chapter); i++ {
		b.Book.ListBooks = append(b.Book.ListBooks, b.Content.Books[i].Name)
	}
}


func standardizeBook(s string) string {
	// Converts any type of book to the the standard format
	// word := "2Pet"
	matched, err := regexp.MatchString(`^\s*[1-3]\s*[A-Za-z]`, s)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if matched {
		fmt.Printf("%s matched\n", s)
		regx := regexp.MustCompile(`\d+`)
		prefix := regx.FindString(s)

		b := strings.SplitAfterN(s, prefix, 2)
		for i, v := range b {
			b[i] = strings.TrimSpace(v)
		}
		return strings.Join(b, " ")
	}
	return s
}

// User Input

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the Book Chapter and verse: ")
	// passage, _ := reader.ReadString('\n')
	// passage = strings.TrimSpace(passage)
	// fmt.Println(passage)


	{Rev 66 22 [20 29 22 11 14 17 17 13 21 11 19 17 18 20 8 21 18 24 21 15 27 21]}




	// Book Name

	// for _, book := range bible.Content.Books {
	// 	books = append(books, BookInfo{
	// 		Name:     bookName[i],
	// 		Position: i + 1,
	// 		Chapters: len(book.Chapters),
	// 	})
	// 	for _, chapter := range book.Chapters {
	// 		books[i].VersesPerChapter = append(books[i].VersesPerChapter, len(chapter.Verses))
	// 	}
	// 	i++
	// }
	// if books == nil {
	// 	return books, nil
	// }
	// return books, nil

	SL := map[string]map[string]string{}

	dict := map[string]map[string]int
	{
		"Genesis": {
			"Position": 1,
			"Chapters": 50,
			"VersesPerChapter": {
				"1": 46,
				"2": 37,
				"3":29,
				"4": 49, 
				"5":33, 
				"6": 25, 
				"7" :26, 
				"8":20, 
				"9": 29, 
				"10": 22, 
				"11" :32, 
				"12" :32, 
				"13": 18,
				"14": 29, 
				"15":23, 
				"16": 22, 
				"17": 20, 
				"18": 22, 
				"19": 21, 
				"20": 20, 
				"21": 23, 
				"22" :30, 
				"23": 25, 
				"24" :22, 
				"25": 19, 
				"26": 19, 
				"27": 26, 
				"28" :68, 
				"29" :29, 
				"30": 20, 
				"31": 30, 
				"32": 52, 
				"33": 29, 
				"34": 12,
			}
		},
	}