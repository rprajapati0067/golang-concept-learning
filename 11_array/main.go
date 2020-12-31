package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string
	fmt.Println('A')

	books[0] = "Iron man"
	books[1] = "Harry Potter"
	books[2] = "Armageddon"
	books[3] = books[0] + " 2nd Edition"

	fmt.Printf("books: %T\n", books)
	fmt.Println("books: ", books)
	fmt.Printf("books: %q\n", books)
	fmt.Printf("books: %#v\n", books)

	//var (
	//	wBooks [winter]string
	//)
}
