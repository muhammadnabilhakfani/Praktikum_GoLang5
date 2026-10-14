package main
import "fmt"

func main() {
	const (
		diskon1 float64 = 0.10
		diskon2 float64 = 0.20
	)
	var hargaAwal, hargaDiskon1, hargaAkhir float64
	fmt.Scan(&hargaAwal)
	hargaDiskon1 = hargaAwal - (hargaAwal * diskon1)
	hargaAkhir = hargaDiskon1 - (hargaDiskon1 * diskon2)
	fmt.Printf("%.2f\n", hargaAkhir)
}
