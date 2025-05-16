package main
import "fmt"

func subtractProductAndSum(n int) int {
    x := n
    kali := 1
    tambah := 0
    for n > 0 {
        hasil := n % 10
        kali = kali * hasil
        n = n / 10
    }
    for x > 0 {
        hasil := x % 10
        tambah = tambah + hasil
        x = x / 10
    }

    return kali - tambah
}
func main() {
	var n int
	fmt.Scan(&n)
    fmt.Println(subtractProductAndSum(n))
}