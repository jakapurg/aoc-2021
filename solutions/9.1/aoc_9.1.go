package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"gonum.org/v1/gonum/mat"
	"math"
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

func main() {
	file, _ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text,scanner.Text())
	}
	fmt.Println(text)
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
	var riskLevel float64
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
				riskLevel+=tmpCell+1
			}
		}
	}
	fmt.Println(riskLevel)



	
	
	
}