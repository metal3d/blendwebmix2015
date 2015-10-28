package data

import "time"

// START STRUCT OMIT
type Book struct {
	Title     string
	Author    string
	Published time.Time
}

// END STRUCT OMIT

// START BOOKSAVE OMIT

func (book *Book) Save() {
	d := getMongo()
	defer d.Close()

	set := Book{
		Title:  book.Title,
		Author: book.Author,
	}

	d.DB("books").C("book").Upsert(set, book)
}

// END BOOKSAVE OMIT

// START BOOKGET OMIT

func (book *Book) Get(title string) {
	d := getMongo()
	defer d.Close()
	d.DB("books").C("book").Find(map[string]string{
		"title": title,
	}).One(book)
}

//END BOOKGET OMIT
