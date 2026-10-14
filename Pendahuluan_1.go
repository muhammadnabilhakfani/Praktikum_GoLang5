package main
import "fmt"

const (
	kuis      float64 = 0.15
	tugas     float64 = 0.15
	praktikum float64 = 0.2
	uts       float64 = 0.25
	uas       float64 = 0.25
)

type Mahasiswa struct {
	Nama           string
	NIM            string
	NilaiTugas     float64
	NilaiKuis      float64
	NilaiPraktikum float64
	NilaiUTS       float64
	NilaiUAS       float64
	NilaiAkhir     float64
}

func main() {
	var mhs Mahasiswa
	fmt.Scan(&mhs.Nama)
	fmt.Scan(&mhs.NIM)
	fmt.Scan(&mhs.NilaiTugas)
	fmt.Scan(&mhs.NilaiKuis)
	fmt.Scan(&mhs.NilaiPraktikum)
	fmt.Scan(&mhs.NilaiUTS)
	fmt.Scan(&mhs.NilaiUAS)

	mhs.NilaiAkhir = kuis*mhs.NilaiKuis + tugas*mhs.NilaiTugas + praktikum*mhs.NilaiPraktikum + uts*mhs.NilaiUTS + uas*mhs.NilaiUAS

	fmt.Println("Data Mahasiswa")
	fmt.Println("Nama :", mhs.Nama)
	fmt.Println("NIM :", mhs.NIM)
	fmt.Println("Nilai Tugas :", mhs.NilaiTugas)
	fmt.Println("Nilai Kuis :", mhs.NilaiKuis)
	fmt.Println("Nilai Praktikum :", mhs.NilaiPraktikum)
	fmt.Println("Nilai UTS :", mhs.NilaiUTS)
	fmt.Println("Nilai UAS :", mhs.NilaiUAS)
	fmt.Println("Nilai Akhir :", mhs.NilaiAkhir)
}
