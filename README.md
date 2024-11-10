# Program Rekursif dalam Go

--- 

## 1. Perhitungan Deret Fibonacci

Program menghitung deret Fibonacci hingga suku tertentu menggunakan rekursif.

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

- `fibonacci(n int)` adalah fungsi rekursif yang menghitung nilai Fibonacci ke-`n`. Jika `n` lebih kecil atau sama dengan 1, fungsi mengembalikan `n`. Jika tidak, fungsi menjumlahkan dua nilai Fibonacci sebelumnya.

### Contoh Input dan Output

**Input:**
```
10
```

**Output:**
```
n 0 1 2 3 4 5 6 7 8 9 10
Sn 0 1 1 2 3 5 8 13 21 34 55
```

---

## 2. Menampilkan Pola Bintang

Program menampilkan pola bintang berbentuk segitiga.

```go
func printBintang(n int) {
    if n > 0 {
        printBintang(n - 1)
        fmt.Println(bintang(n))
    }
}

func bintang(n int) string {
    if n == 1 {
        return "*"
    }
    return "*" + bintang(n-1)
}
```

- `printBintang(n int)` mencetak pola bintang secara rekursif dari baris atas hingga baris bawah.
- `bintang(n int) string` menghasilkan baris dengan `n` bintang menggunakan rekursif.

### Contoh Input dan Output

**Input:**
```
5
```

**Output:**
```
*
**
***
****
*****
```

---

## 3. Mencari Faktor Bilangan

Program mencari dan mencetak semua faktor dari bilangan `N`.

```go
func printFaktor(n, i int) {
    if i > n {
        return
    }
    if n%i == 0 {
        fmt.Printf("%d ", i)
    }
    printFaktor(n, i+1)
}
```

- `printFaktor(n, i int)` memeriksa apakah `i` adalah faktor dari `n`. Jika ya, mencetak `i`, lalu memanggil dirinya dengan `i+1` hingga mencapai `n`.

### Contoh Input dan Output

**Input:**
```
12
```

**Output:**
```
1 2 3 4 6 12
```

---

## 4. Menampilkan Urutan Angka (Menurun dan Menaik)

Program menampilkan bilangan dari `N` hingga 1, lalu kembali dari 1 ke `N`.

```go
func printDesc(n int) {
    if n < 1 {
        return
    }
    fmt.Printf("%d ", n)
    printDesc(n - 1)
}

func printAsc(n, curr int) {
    if curr > n {
        return
    }
    fmt.Printf("%d ", curr)
    printAsc(n, curr+1)
}
```

- `printDesc(n int)` mencetak bilangan dari `N` hingga 1.
- `printAsc(n, curr int)` mencetak bilangan dari 1 hingga `N` menggunakan `curr` sebagai bilangan awal.

### Contoh Input dan Output

**Input:**
```
5
```

**Output:**
```
5 4 3 2 1 1 2 3 4 5
```

---

## 5. Urutan Bilangan Ganjil

Program menampilkan semua bilangan ganjil dari `1` hingga `N`.

```go
func printGanjil(n, current int) {
    if current > n {
        return
    }
    if current%2 != 0 {
        fmt.Printf("%d ", current)
    }
    printGanjil(n, current+1)
}
```

- `printGanjil(n, current int)` mencetak `current` jika ganjil, kemudian memanggil dirinya dengan `current+1`.

### Contoh Input dan Output

**Input:**
```
10
```

**Output:**
```
1 3 5 7 9
```

---

## 6. Perhitungan Pangkat

Program menghitung hasil dari `x` pangkat `y` menggunakan rekursif.

```go
func pangkat(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * pangkat(x, y-1)
}
```

- `pangkat(x, y int)` menghitung `x` pangkat `y` dengan mengalikan `x` secara rekursif hingga `y` mencapai 0.

### Contoh Input dan Output

**Input:**
```
2 3
```

**Output:**
```
8
```