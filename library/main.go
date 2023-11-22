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

// LibraryManager визначає методи для управління книгами у бібліотеці (Принцип розділення інтерфейсів - ISP).
type LibraryManager interface {
	RetrieveBook(title string) Book
	ReturnBook(book Book)
}

// Manager реалізує інтерфейс LibraryManager та управляє книгами у базі даних та на полиці (Принцип єдиної відповідальності - SRP).
type Manager struct {
	db    Database
	shelf Shelf
}
type Client struct {
	manager LibraryManager
}

// RetrieveBook отримує книгу від менеджера (або з бази даних, або з полиці) за назвою (Принцип єдиної відповідальності - SRP).
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

// ReturnBook повертає книгу на полицю, керовану менеджером (Принцип єдиної відповідальності - SRP).
func (m *Manager) ReturnBook(book Book) {
	m.shelf.books[book.Title] = book
	fmt.Println("Книгу повернуто на полицю:", book.Title)
}

// BorrowBook бере книгу від менеджера через інтерфейс LibraryManager (Принцип інверсії залежностей - DIP).
func (c *Client) BorrowBook(title string) Book {
	return c.manager.RetrieveBook(title)
}

// ReturnBook повертає книгу менеджеру через інтерфейс LibraryManager (Принцип інверсії залежностей - DIP).
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
