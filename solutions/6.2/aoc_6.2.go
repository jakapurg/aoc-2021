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
	var fishesI[] uint8
	tmpInt := 0
	for _, val := range fishes{
		tmpInt,_ = strconv.Atoi(val)
		fishesI = append(fishesI,uint8(tmpInt))
	}
	fishNum := [9]int{0, 0,0,0,0,0,0,0,0}
	for _,val := range fishesI{
		fishNum[val]++
	}
	fmt.Println(fishNum)
	newFish:=[9]int{0,0,0,0,0,0,0,0,0}

	for day:=0;day<256;day++{
		newFish = [9]int{0,0,0,0,0,0,0,0,0}
		for i,val := range fishNum{
			
			if(i==0){
				
				newFish[6]=val
				newFish[8]=val
			}else
			{
				newFish[i-1]+=val
			}
			fishNum = newFish
			
		}
	}
	count:=0
	for _,val := range fishNum{
		count+=val
	}
	fmt.Println(count)
	
}