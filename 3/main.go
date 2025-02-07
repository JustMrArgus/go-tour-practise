package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
	ISBN   int
}

func main() {
	first_book := Book{"Harry Potter", "Edgar Po", 1956, 12345}
	second_book := Book{"Spider Man", "John Lee", 1986, 36942}
	third_book := Book{"Batman", "Super Man", 2001, 39139}
	fourth_book := Book{"Joker", "Hawkeye", 2013, 12308}
	fifth_book := Book{"Cat book", "Marry Poppins", 22021, 14092}

	fmt.Println(first_book.Title)
	fmt.Println(second_book)
	fmt.Println(third_book.Year, third_book.Author)

	first_pointer := &first_book
	first_pointer.Title = "Sherlock Holms"
	fmt.Println(first_book.Title)

	var books [5]Book
	books[0] = first_book
	books[1] = second_book
	books[2] = third_book
	books[3] = fourth_book
	books[4] = fifth_book
	fmt.Println(books)

	sliced_books := books[0:3]
	sliced_books[0] = Book{"SuperGirl", "Some girl", 2024, 12387}
	fmt.Println(sliced_books)
	fmt.Println(books)

	var books_slice []Book
	books_slice = append(books_slice, first_book, second_book, fifth_book)
	fmt.Println(books_slice)
	books_slice = append(books_slice, third_book)
	fmt.Println(books_slice)

	fmt.Println(len(books_slice), cap(books_slice))

	matrix := [][]int{
		{18, 0, 0},
		{0, 1, 21},
		{0, 0, 1},
	}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			fmt.Printf("%d	", matrix[x][y])
		}
		fmt.Println("")
	}

	for i, v := range books {
		book_number := i + 1
		fmt.Printf("Book number %d - %s\n", book_number, v.Title)
	}

	books_map := make(map[string]int)
	books_map["a"] = 1
	books_map["b"] = 2
	books_map["c"] = 3
	books_map[books[0].Title] = books[0].Year

	fmt.Println(books_map)
	fmt.Println(books_map["a"])

	delete(books_map, "b")
	books_map[books[0].Title] = 1984

	fmt.Println(books_map)
	fmt.Println(books_map[books[0].Title])

	elem, ok := books_map["a"]
	fmt.Println(elem, ok)

}
