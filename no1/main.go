package no1

import "fmt"

// 1. Apa output dari kode berikut?
func main() {
	var a int = 10
	var b float64 = 5.5
	fmt.Println(a + int(b))
}

// a) 15.5
//  b) 15
//  c) Error
//  d) 10

// answer: b) 15

// Penjelasan: Pada kode tersebut, variabel `a` bertipe int dan `b` bertipe float64. Ketika kita menjumlahkan `a` dan `b`, kita perlu mengkonversi `b` menjadi int terlebih dahulu menggunakan `int(b)`. Setelah konversi, hasilnya adalah 10 + 5 = 15. Jadi outputnya adalah 15.
// Jika kita tidak melakukan konversi, maka akan terjadi error karena kita tidak bisa menjumlahkan tipe data yang berbeda secara langsung.

var a int = 10
var b float64 = 5.5
fmt.Println(a + int(b))
