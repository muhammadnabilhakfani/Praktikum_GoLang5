package main
import "fmt"

func main() {
	type Book struct {
		title  string
		genre  string
		year   int
		rating float64
	}
	var book1, book2 Book
	fmt.Scan(&book1.title, &book2.title)
	fmt.Scan(&book1.genre, &book2.genre)
	fmt.Scan(&book1.year, &book2.year)
	fmt.Scan(&book1.rating, &book2.rating)

	fmt.Println("--- Book 1 ---")
	fmt.Println(book1.title)
	fmt.Println(book1.genre)
	fmt.Println(book1.year)
	fmt.Println(book1.rating)

	fmt.Println("--- Book 2 ---")
	fmt.Println(book2.title)
	fmt.Println(book2.genre)
	fmt.Println(book2.year)
	fmt.Println(book2.rating)
}

program InformasiBuku
kamus
type Buku <
  judul, genre : string
  tahun : integer
  rating : real
>
buku1, buku2 : Buku
algoritma
input(buku1.judul, buku2.judul, buku1.genre, buku2.genre, buku1.tahun, buku2.tahun, buku1.rating, buku2.rating)
output("--- Book 1 ---")
output(buku1.judul, buku1.genre, buku1.tahun, buku1.rating)
output("--- Book 2 ---")
output(buku2.judul, buku2.genre, buku2.tahun, buku2.rating)
endprogram
