package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)
func main(){
	file,_ := os.Open("reports.txt")
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text[] string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	fishes := strings.Split(text[0],",")
	var fishesI[] int
	tmpInt := 0
	for _, val := range fishes{
		tmpInt,_ = strconv.Atoi(val)
		fishesI = append(fishesI,tmpInt)
	}
	var newFish[] int
	for day:=0;day<80;day++{
		newFish=nil
		for _,val := range fishesI{
			if(val==0){
				newFish=append(newFish,6)
				newFish=append(newFish,8)
			}else
			{
				newFish=append(newFish,val-1)
			}
		}
		fishesI = newFish
	}
	fmt.Println(len(fishesI))
}