// package main
// import "golang.org/x/tour/pic"
// import "strings"

// // func Pic(dx, dy int) [][]uint8{
// // 	pic:=make([][]uint8, dy)

// // 	for y:=0; y<dy; y++{
// // 		pic[y]=make([]uint8, dx)
// // 		for x:=0; x<dx; x++{
// // 			pic[y][x] = uint8(x^y)
// // 		}
// // 	}
// // 	return pic;
// // }

// func main(){
// 	//  fmt.Println(Pic(4,4))
// 	//  pic.Show(Pic)
// }

package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	word :=""
	
	for _ , ch := range s{
		if ch==' '{
			if word!="" {
				counts[word]++
				word=""
			}
		} else{
			word+=string(ch)
		}
			
	}
	
	if word!=""{
		counts[word]++;
	}
	
	return counts;

}

func main() {
	wc.Test(WordCount)
}
