package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func StartConsole(library *services.Library) {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choiceStr := scanner.Text()
		choice, err := strconv.ParseInt(strings.TrimSpace(choiceStr), 10, 64)
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			addBook(library)
		case 2:
			removeBook(library)
		case 3:
			borrowBook(library)
		case 4:
			returnBook(library)
		case 5:
			listAvailableBooks(library)
		case 6:
			listBorrowedBooks(library)
		case 7:
			return
		default:
			fmt.Println("Invalid choice")
			StartConsole(library)
		}
	}

}

func addBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	library.AddBook(book)
	fmt.Println("Book added successfully")

}

func removeBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	library.RemoveBook(id)

}

func borrowBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to borrow: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.
			Println("Book borrowed successfully")
	}

}

func returnBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to return: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully")
	}

}

func listAvailableBooks(library *services.Library) {
	books := library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}

}

func listBorrowedBooks(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	books := library.ListBorrowedBooks(memberID)
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}

}
