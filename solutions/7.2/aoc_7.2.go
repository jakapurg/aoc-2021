package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func main(){
	file,_ := os.Open("reports.txt")
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text[] string
	for scanner.Scan(){
		text = append(text,scanner.Text())
	}
	s:=strings.Split(text[0],",")
	var pos = make([]int,len(s))
	for i,val := range s{
		pos[i],_ = strconv.Atoi(val)
	}
	fuel:= 0
	smallestF:=math.MaxInt
	mostOptimateI:=0
	diff:=0
	for x:=0;x<len(pos);x++{
		fuel=0
		for y:=0;y<len(pos);y++{
			diff=int(math.Abs(float64(pos[y]-pos[x])))
			for z:=diff-1;z>0;z--{
				diff+=z
			}
			fuel+=diff
		}
		if(fuel<smallestF){
			smallestF=fuel
			mostOptimateI=x
		}
	}
	fmt.Println(smallestF,mostOptimateI)
	

}