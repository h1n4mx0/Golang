package main
import "fmt"

/**bai 1: hien thi day so tu 1-100, 
moi so cach nhau boi dau phay, 
bo dau phay cuoi cung
bo qua so 6, 48, 75, 89**/

func printOneToHundred() {
	var xhtml = ""
	for i:= 1; i <= 100; i++ {
		if i == 6 || i == 48 || i == 75 || i == 89 {
			continue // bo qua so 6, 48, 75, 89
		}
		if i == 100 {
			xhtml += fmt.Sprintf("%d", i) // khong in dau phay sau so 100
			break // dung vong lap khi den so 100
		}
		xhtml += fmt.Sprintf("%d,", i)
	}
	fmt.Println(xhtml)
}
