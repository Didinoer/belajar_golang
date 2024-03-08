package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1
	fmt.Println("assalamualaikum")

	//2
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)

	//3
	fmt.Println(len("assalamualaikum"))
	fmt.Println(("assalamualaikum"[7]))
	//konversi int ke string
	var a = "assalamualaikum"
	var b = string(a[0])
	fmt.Println(b)

	// 4
	// dalam golang semua variable harus digunakan
	var (
		fn = "a"
		mn = "b"
		ln = "c"
	)
	var nama string
	nama = "golox"
	//atau
	var nams = "golox"
	fmt.Println(nama)
	nama = "lopus"
	fmt.Println(nama)
	if nama == "golox" {
		fmt.Println("golox nih boss")
	} else if nama == "lopus" {
		fmt.Println("lopus nih boss")
	} else {
		fmt.Println("nggak ada ang hadir")
	}

	// 5 constanta
	// constant variable tidak dapat diubah tapi boleh jika tidak digunakan
	const name = "kuolsui"
	fmt.Println(name)

	// 6 array
	var aray = [...]string{
		"merah",
		"kuning",
		"hijau",
		"biru",
	}
	fmt.Println(aray)
	//di golang tidak dapat menghapus data dalam aray hanya bisa dikosongkan

	// 7 slice
	var hari = [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	//cara mengubah data dalam aray
	slicehari := hari[5:] //data diambil dari value ke lima sampai akhir
	slicehari[0] = "sabtubaru"
	slicehari[1] = "minggubaru"
	fmt.Println(hari)
	fmt.Println(reflect.TypeOf(hari))
	//cara tambah data kedalam array mengunakan slice
	slicehari2 := hari[0:]
	slicehari3 := append(slicehari2, "hari senin lagi")
	fmt.Println(slicehari3)
	fmt.Println(reflect.TypeOf(slicehari3))

	// 8 map

	person := map[string]string{
		"nama":  " didi nurahman",
		"nim":   "17210135",
		"kelas": "17.6c.29",
	}
	person["nama"] = "didi noer"
	delete(person, "nim")
	fmt.Println(person["nama"])
	fmt.Println(person)

}
