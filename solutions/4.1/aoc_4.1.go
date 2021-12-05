package main
import (
		"fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		"gonum.org/v1/gonum/mat"
)
func every(s []float64, str float64) bool {
	for _, v := range s {
		if v != str {
			return false
		}
	}

	return true
}

func printMatrix(A *mat.Dense){
	fmt.Printf("%v\n\n", mat.Formatted(A, mat.Prefix(" "), mat.Excerpt(2)))
}
func checkWinner(boards []*mat.Dense, dim int) (bool,int) {
	for x,tmpBoard := range boards{
		for d:=0;d<dim;d++{
			tmpRow:=mat.Row(nil,d,tmpBoard)
			if(every(tmpRow,-1)){
				return true,x
			}
		}
		for d:=0;d<dim;d++{
			tmpCol:=mat.Col(nil,d,tmpBoard)
			if(every(tmpCol,-1)){
				return true,x
			}
		}
	}
	return false,-1
}

func main(){
	file, _ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text[] string
	for scanner.Scan() {
		text = append(text,scanner.Text())
	}
	drawnNumbers:=strings.Split(text[0],",")
	numbers := make([]int,len(drawnNumbers))
	for i,val := range drawnNumbers{
		numbers[i],_=strconv.Atoi(val)
	}
	dims:=len(strings.Fields(text[2]))
	
	var boards[] *mat.Dense
	//matrix:= mat.NewDense(dims,dims,nil)
	var tmpNums [] float64
	for _, ln := range text[2:]{
		rowStrings :=strings.Fields(ln)
		
		if(len(rowStrings)>0){
			for _,x := range rowStrings{
				tmpnum,_:=strconv.Atoi(x)
				tmpNums = append(tmpNums,float64(tmpnum))
			}
		}else{
			boards = append(boards,mat.NewDense(dims,dims,tmpNums))
			tmpNums=nil
		}
	}
	fmt.Println(len(boards))
	for _,num := range numbers{
		for i,_ := range boards{
			for d:=0;d<dims;d++{
				rowD:=mat.Row(nil,d,boards[i])
				for k,_:= range rowD{
					if(rowD[k]==float64(num)){
						rowD[k]=-1
					}
				}
				boards[i].SetRow(d,rowD)
			}
		}
		res,in := checkWinner(boards,dims)
		score:=0
		if(res){
			fmt.Println(in)
			for l:=0;l<dims;l++{
				for m:=0;m<dims;m++{
					cell:=boards[in].At(l,m)
					if(cell!=-1){
						score+=int(cell)
					}
				}
			}
			fmt.Println(score*num)
			return
		}
		//CHECK FOR WINNER
	}
}