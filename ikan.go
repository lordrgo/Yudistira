package main
import "fmt"
import "os"
import "os/exec"
type ArrData struct {
    namaBarang string
    namaSuplier string
    peringatan string
    kodeBarang int
    jumlah int
}

var barang [20]ArrData

func main () {
    menu()
}

func inputBarang() int {
    var inputNama, inputSuplier, sudah string
    var inputKode, inputJumlah int
    ulang := false
    i:=0
    
    for (!ulang) {
        fmt.Println("GUNAKAN '_' UNTUK SPASI")
        fmt.Print("NAMA BARANG :")
        fmt.Scanln(&inputNama)
        fmt.Print("NAMA SUPPLIER :")
        fmt.Scanln(&inputSuplier)
        fmt.Print("KODE BARANG (4 DIGIT) : ")
        fmt.Scanln(&inputKode)
        fmt.Print("JUMLAH : ")
        fmt.Scanln(&inputJumlah)
        
        if (inputJumlah >=5) {
            barang[i].peringatan = "CUKUP"
        } else {
            barang[i].peringatan = "PESAN BARANG LAGI!"
        }
        barang[i].namaBarang = inputNama
        barang[i].namaSuplier = inputSuplier
        barang[i].kodeBarang = inputKode
        barang[i].jumlah = inputJumlah
        
        fmt.Print("INPUT BARANG LAGI? (ya/tidak) ")
        fmt.Scanln(&sudah)
        if (sudah == "tidak") {
            ulang =  true
        }
        i++
    }
    return i
}

func cekBarang(banyakBarang *int) {
    i:=0
    for (i<*banyakBarang) {
        fmt.Println("NAMA BARANG : ", barang[i].namaBarang)
        fmt.Println("SUPPLIER : ", barang[i].namaSuplier)
        fmt.Println("KODE : ", barang[i].kodeBarang)
        fmt.Println("Kondisi logistik = ", barang[i].peringatan)
        fmt.Println("-------------------------------")
        i++
    }
}

func cariBarang(nama *string, banyakBarang *int) int {
    i:=0
    
    for(i<*banyakBarang) {
        if (barang[i].namaBarang == *nama) {
            return i
        }
        i++
    }
    return -1
}

func tambahBarang (banyakBarang *int) {
    var namaBrg string
    var jml int
    
    fmt.Print("Nama barang yang ingin ditambah : ")
    fmt.Scanln(&namaBrg)
    
    var index int = cariBarang(&namaBrg, banyakBarang)
    
    if (index==-1) {
        fmt.Println("BARANG tidak ditemukan")
    } else {
        fmt.Print("Masukkan jumlah yang ingin ditambah : ")
        fmt.Scanln(&jml)
        barang[index].jumlah = barang[index].jumlah + jml
        if (barang[index].jumlah>=5) {
            barang[index].peringatan = "CUKUP"
        } else if (barang[index].jumlah<5) {
            barang[index].peringatan = "PESAN BARANG LAGI"
        }
        
        fmt.Println("Barang sudah ditambahkan")
    }
}

func keluarBarang (banyakBarang *int) {
    var namaBrg string
    var jml int
    
    fmt.Print("Nama barang yang ingin dikeluarkan : ")
    fmt.Scanln(&namaBrg)
    
    var index int = cariBarang(&namaBrg, banyakBarang)
    
    if index==-1 {
        fmt.Println("BARANG tidak ditemukan")
    } else {
        fmt.Print("Masukkan jumlah yang ingin dikeluarkan : ")
        fmt.Scanln(&jml)
        if ((barang[index].jumlah - jml) < 0) {
            barang[index].jumlah = 0
        } else {
            barang[index].jumlah = barang[index].jumlah - jml
        }
        
        if (barang[index].jumlah>=5) {
            barang[index].peringatan = "CUKUP"
        } else if (barang[index].jumlah<5) {
            barang[index].peringatan = "PESAN BARANG LAGI"
        }
        
        fmt.Println("Barang sudah dikeluarkan.")
    }
}

func menu() {
    var pilih int
    var ulang bool=false
    var nBarang = inputBarang()
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
    
    for (!ulang) {
        fmt.Println("1. Tambah jumlah Barang")
        fmt.Println("2. Keluarkan Barang")
        fmt.Println("3. Lihat Barang")
        fmt.Println("4. TUTUP")
        
        fmt.Print("PILIHAN = ")
        fmt.Scanln(&pilih)
        
        if (pilih==1) {
            tambahBarang(&nBarang)
        } else if (pilih==2) {
            keluarBarang(&nBarang)
        } else if (pilih==3) {
            cekBarang(&nBarang)
        } else {
            ulang = true
        }
    }
    fmt.Println()
    fmt.Println("YUDIS KNTL")
}