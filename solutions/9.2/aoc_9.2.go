package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"gonum.org/v1/gonum/mat"
	"math"
	"sort"
)
func printMatrix(A *mat.Dense){
	fmt.Printf("%v\n\n", mat.Formatted(A, mat.Prefix(" "), mat.Excerpt(2)))
}

func MinOf(vars ...float64) float64 {
    min := vars[0]

    for _, i := range vars {
        if min > i {
            min = i
        }
    }

    return min
}

func contains(s [][]int, str []int) bool {
	for _, v := range s {
		if v[0] == str[0] && v[1]==str[1] {
			return true
		}
	}

	return false
}
func unique(coordinates [][]int) [][]int {
	var coordinatesU [][]int

	for _,val := range coordinates{
		if(!contains(coordinatesU,val)){
			coordinatesU = append(coordinatesU,val)
		}
	}
	return coordinatesU
}

func basinLeng(x,y,m,n int, matrix *mat.Dense)[][]int{
	tmpCell:=matrix.At(x,y)
	left:=math.MaxFloat64
	up:=math.MaxFloat64
	right:=math.MaxFloat64
	down:=math.MaxFloat64
	var coordinates [][]int
	tmpC:= []int{x,y}

	coordinates = append(coordinates,tmpC)
	if(y!=0){
		left =  matrix.At(x,y-1)
	}
	if(x!=0){
		up = matrix.At(x-1,y)
	}
	if(x+1<m){
		down = matrix.At(x+1,y)
	}
	if(y+1<n){
		right = matrix.At(x,y+1)
	}
	if(left>tmpCell && left<math.MaxFloat64 && left!=9){
		coordinates = append(coordinates,basinLeng(x,y-1,m,n,matrix)...)
	}
	if(up>tmpCell && up<math.MaxFloat64 && up!=9){

		coordinates = append(coordinates,basinLeng(x-1,y,m,n,matrix)...)
	}
	if(down>tmpCell && down<math.MaxFloat64 && down!=9){

		coordinates = append(coordinates,basinLeng(x+1,y,m,n,matrix)... )
	}
	if(right>tmpCell && right<math.MaxFloat64 && right!=9){

		coordinates =append(coordinates, basinLeng(x,y+1,m,n,matrix)...)
	}
	return coordinates


}
func main() {
	file, _ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text,scanner.Text())
	}
	n:=len(text[0])
	m:=len(text)

	matrix := mat.NewDense(m,n,nil)
	for j, each_ln := range text {
		numberStrings := strings.Split(each_ln,"")
		var numbers [] float64
		for _, x := range numberStrings {
			number,_ := strconv.Atoi(x)
			numbers = append(numbers,float64(number))
		}
		matrix.SetRow(j,numbers)
	}
	var tmpCell float64
	var left float64
	var up float64
	var down float64
	var right float64
	var min float64
	var basins []int
	for x:=0;x<m;x++{
		for y:=0;y<n;y++{
			left=math.MaxFloat64
			up=math.MaxFloat64
			right=math.MaxFloat64
			down=math.MaxFloat64
			
			tmpCell = matrix.At(x,y)
			if(x!=0){
				left =  matrix.At(x-1,y)
			}
			if(y!=0){
				up = matrix.At(x,y-1)
			}
			if(x+1<m){
				down = matrix.At(x+1,y)
			}
			if(y+1<n){
				right = matrix.At(x,y+1)
			}
			min = MinOf(left,up,down,right)
			if(min>tmpCell){
				xt:=unique(basinLeng(x,y,m,n,matrix))
				
				basins = append(basins,len(xt))
			}
		}
	}
	sort.Ints(basins)
	score:=1
	for i:=1;i<4;i++{
		score=score*basins[len(basins)-i]
	}
	fmt.Println(score)



	
	
	
}