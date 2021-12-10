package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func main(){
	file, _ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text [] string
	for scanner.Scan(){
		text = append(text,scanner.Text())
	}
}