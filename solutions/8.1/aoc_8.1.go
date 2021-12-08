package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main(){
	file,_ := os.Open("reports.txt")
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text[] string
	for scanner.Scan(){
		text = append(text,strings.Split(scanner.Text(),"|")[1])
	}
	var codes[] string
	count:=0
	leng:=0
	for _,val := range text {
		codes = strings.Fields(val)
		for _,val2 := range codes {
			leng=len(val2)
			if(leng<5 || leng==7){
				count++
			}
		}
	}
	fmt.Println(count)
}
