package main
import (
		"fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		"gonum.org/v1/gonum/mat"
)
func printMatrix(A *mat.Dense){
	fmt.Printf("%v\n\n", mat.Formatted(A, mat.Prefix(" "), mat.Excerpt(9)))
}
func main(){
	file,_ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text[] string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	lines := mat.NewDense(1000,1000,nil)
	
	beginX:=0
	beginY:=0
	endX:=0
	endY:=0
	for _, val := range text {
		i:= strings.Split(val," -> ")
		begin := strings.Split(i[0],",")
		end := strings.Split(i[1],",")
		beginX,_ = strconv.Atoi(begin[0])
		beginY,_ = strconv.Atoi(begin[1])
		endX,_ = strconv.Atoi(end[0])
		endY,_ = strconv.Atoi(end[1])

		if(beginX==endX || beginY==endY){
			row1:=mat.Row(nil,beginY,lines)

			for(beginX!=endX){
				row1[beginX]++
				if(beginX>endX){
					beginX--
				}else
				{
					beginX++
				}
				if(beginX==endX){
					row1[beginX]++

				}
			}
			lines.SetRow(beginY,row1)
			col1:=mat.Col(nil,beginX,lines)
			for(beginY!=endY){
				
				col1[beginY]++
				if(beginY>endY){
					beginY--
				}else
				{
					beginY++
				}
				if(beginY==endY){
					col1[beginY]++
				}
				}
				lines.SetCol(beginX,col1)
		}
	}
	
	var tmpRow []float64
	count:=0
	for j:=0;j<1000;j++{
		tmpRow = mat.Row(nil,j,lines)
		for k:=0;k<1000;k++{
			if(tmpRow[k]>1){
				count++
			}
		}
	}
	fmt.Println(count)
}