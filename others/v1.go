package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bible struct {
	XMLName xml.Name `xml:"osis"`
	Content Content  `xml:"osisText"`
}

type Content struct {
	Metadata Metadata   `xml:"header"`
	BookInfo []BookInfo `xml:"-"`
	Books    []Book     `xml:"div"`
}

type Metadata struct {
	Work Work `xml:"work"`
}

// Contains info about the specific work
type Work struct {
	Translation  string `xml:"title"`
	Abbreviation string `xml:"identifier"`
	Language     string `xml:"language"`
	Description  string `xml:"description"`
}

type Book struct {
	Name     string    `xml:"osisID,attr"`
	Chapters []Chapter `xml:"chapter"`
}

type Chapter struct {
	Name   string   `xml:"osisID,attr"`
	Verses []string `xml:"verse"`
}

type BookInfo struct {
	Name             string
	Position         int
	Chapters         int
	VersesPerChapter []int
}

type Reference struct {
	Book    string
	Chapter int
	Start   int
	End     int
	Extra   []int
}

// func main() {
// 	version := "kjv"
// 	bible := open(version)

// 	s := "1 John 3:10,12"

// 	ref, err := standardize(s)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = bible.ValidateRef(ref)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	bible.reference(ref)

// }
func Guy(s string) {
	fmt.Println(s)
	// version := "kjv"
	// bible := open(version)

	// s = "1 John 3:10,12"

	// ref, err := standardize(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = bible.ValidateRef(ref)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// bible.reference(ref)

}

func (bible *Bible) reference(ref Reference) {
	// var result string
	result := map[int]string{}

	bookInfo, err := bible.GetBookInfo(ref)
	if err != nil {
		fmt.Println(err)
		// return err
	}
	bookIndex := bookInfo.Position - 1
	chapterIndex := ref.Chapter - 1
	startIndex := ref.Start - 1
	endIndex := ref.End - 1

	// Single
	if ref.End == 0 {
		result[ref.Start] = bible.Content.Books[bookIndex].Chapters[chapterIndex].Verses[startIndex]
	}

	// Multiple
	if len(ref.Extra) != 0 {
		result[ref.Start] = bible.Content.Books[bookIndex].Chapters[chapterIndex].Verses[startIndex]

		for i := 0; i < len(ref.Extra); i++ {
			result[ref.Extra[i]] = bible.Content.Books[bookIndex].Chapters[chapterIndex].Verses[ref.Extra[i]-1]
		}
	}
	// Range
	if ref.End != 0 {
		for i := startIndex; i <= endIndex; i++ {
			result[i+1] = bible.Content.Books[bookIndex].Chapters[chapterIndex].Verses[i]
		}
	}

	fmt.Println(result)
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

	bible.SetBookInfo()
	return bible
}

// Sets the all the Books info
func (bible *Bible) SetBookInfo() {
	i := 0
	bName := []string{"Genesis", "Exodus", "Leviticus", "Numbers", "Deuteronomy", "Joshua", "Judges", "Ruth", "1 Samuel", "2 Samuel", "1 Kings", "2 Kings", "1 Chronicles", "2 Chronicles", "Ezra", "Nehemiah", "Esther", "Job", "Psalm", "Proverbs", "Ecclesiastes", "Song of Solomon", "Isaiah", "Jeremiah", "Lamentations", "Ezekiel", "Daniel", "Hosea", "Joel", "Amos", "Obadiah", "Jonah", "Micah", "Nahum", "Habakkuk", "Zephaniah", "Haggai", "Zechariah", "Malachi", "Matthew", "Mark", "Luke", "John", "Acts", "Romans", "1 Corinthians", "2 Corinthians", "Galatians", "Ephesians", "Philippians", "Colossians", "1 Thessalonians", "2 Thessalonians", "1 Timothy", "2 Timothy", "Titus", "Philemon", "Hebrews", "James", "1 Peter", "2 Peter", "1 John", "2 John", "3 John", "Jude", "Revelation"}

	for _, book := range bible.Content.Books {
		bible.Content.BookInfo = append(bible.Content.BookInfo, BookInfo{
			Name:     bName[i],
			Position: i + 1,
			Chapters: len(book.Chapters),
		})
		for _, chapter := range book.Chapters {
			bible.Content.BookInfo[i].VersesPerChapter = append(bible.Content.BookInfo[i].VersesPerChapter, len(chapter.Verses))
		}
		i++
	}
}

// Retrieve Book info of a particular book when you pass in the name
func (bible Bible) GetBookInfo(ref Reference) (BookInfo, error) {
	refBookLength := len(ref.Book)

	for _, book := range bible.Content.BookInfo {
		if refBookLength > len(book.Name) {
			continue
		}

		if strings.EqualFold(book.Name[:refBookLength], ref.Book) {
			return book, nil
		}
	}
	return BookInfo{}, errors.New("Error: Book " + ref.Book + " not found")
}

func (bible Bible) ValidateRef(ref Reference) error {
	bookInfo, err := bible.GetBookInfo(ref)
	if err != nil {
		return err
	}

	// Check Chapters
	if ref.Chapter <= 0 || ref.Chapter > bookInfo.Chapters {
		return fmt.Errorf("chapter %v is not in range of '%v'", ref.Chapter, ref.Book)
	}
	// Check Verse
	if ref.Start <= 0 || ref.Start > bookInfo.VersesPerChapter[ref.Chapter-1] {
		return fmt.Errorf("verse %v is not in range of the chapter", ref.Start)
	}

	// Check Verse
	// if ref.End <= ref.Start || ref.End > bookInfo.VersesPerChapter[ref.Chapter-1] {
	if (ref.End > 0 && ref.End <= ref.Start) || ref.End > bookInfo.VersesPerChapter[ref.Chapter-1] {
		return fmt.Errorf("verse %v is not in range of the chapter", ref.End)
	}

	// for _, extraVerse := range ref.Extra {
	// 	if extraVerse <= 0 || extraVerse > bookInfo.VersesPerChapter[ref.Chapter-1] {
	// 		return fmt.Errorf("verse %v is not in range of the chapter", extraVerse)
	// 	}
	// }
	// bookInfo
	return nil
}

func standardize(s string) (ref Reference, err error) {
	// " 			john 3 10   11" //!Fix adding space as a delimiter
	// john 3:10-12,21 //! this as well
	pattern := `(?P<num>^\s*[1-3]?\s*)(?P<book>[A-Za-z]+)\s*(?P<chapter>\d+)[;: ]+(?P<verse>\d+)(?P<delimiter>[-,]+)?(?P<extraVerse>\d+)?`
	re := regexp.MustCompile(pattern)

	if re.MatchString(s) {
		var num, book, delimiter string
		var chapter, verse, extraVerse int
		match := re.FindStringSubmatch(s)
		for i := 1; i < len(match); i++ {
			if len(match[i]) < 1 { // to skip trimming the empty values and " " delimiter for extra verses //! Need to fix for multiple space in delimiter
				continue
			}
			match[i] = strings.TrimSpace(match[i]) // remove white spaces
		}

		num, book, delimiter = match[1], match[2], match[5]
		chapter, _ = strconv.Atoi(match[3])
		verse, _ = strconv.Atoi(match[4])
		extraVerse, _ = strconv.Atoi(match[6])

		if num != "" {
			book = fmt.Sprintf("%v %v", num, book)
		}
		ref.Book = book
		ref.Chapter = chapter
		ref.Start = verse

		if extraVerse != 0 && extraVerse > verse {
			if delimiter == "-" {
				ref.End = extraVerse
			} else if delimiter == "," {
				ref.Extra = append(ref.Extra, extraVerse)
			}
		}
	} else {
		return ref, errors.New("Error: " + s + " is an invalid passage to reference")
	}
	return ref, nil
}

/*
	Dedicated to my friends, who believe i am cooler and smarter than i actually am. thank you ;)
*/
