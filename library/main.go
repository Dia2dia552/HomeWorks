package main

import (
	"fmt"
)

type Book struct {
	Title string
}

type Database struct {
	books map[string]Book
}

type Shelf struct {
	books map[string]Book
}

type LibraryManager interface {
	RetrieveBook(title string) Book
	ReturnBook(book Book)
}

type Manager struct {
	db    Database
	shelf Shelf
}
type Client struct {
	manager LibraryManager
}

func (m *Manager) RetrieveBook(title string) Book {
	book, found := m.db.books[title]
	if !found {
		fmt.Println("Книга не знайдена у бібліотеці")
		return Book{}
	}

	delete(m.db.books, title)
	fmt.Println("Книгу видали з полиці:", title)
	return book
}

func (m *Manager) ReturnBook(book Book) {
	m.shelf.books[book.Title] = book
	fmt.Println("Книгу повернуто на полицю:", book.Title)
}

func (c *Client) BorrowBook(title string) Book {
	return c.manager.RetrieveBook(title)
}

func (c *Client) ReturnBook(book Book) {
	c.manager.ReturnBook(book)
}

func main() {
	db := Database{
		books: make(map[string]Book),
	}

	shelf := Shelf{
		books: make(map[string]Book),
	}

	manager := Manager{
		db:    db,
		shelf: shelf,
	}

	client := Client{
		manager: &manager,
	}

	requestedBookTitle := "Маленький принц"
	retrievedBook := client.BorrowBook(requestedBookTitle)
	if retrievedBook.Title != "" {
		fmt.Println("Клієнт отримав книгу:", retrievedBook.Title)
	}

	wrongBook := Book{Title: "Неправильна книга"}
	client.ReturnBook(wrongBook)

	correctBook := Book{Title: "Маленький принц"}
	client.ReturnBook(correctBook)

	fmt.Println("База даних:", db.books)
	fmt.Println("Книжкова шафа:", shelf.books)
}
