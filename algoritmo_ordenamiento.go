 package main
 import "fmt"
 func main(){
	 arreglo_desordenado := []int{412,12,123,1,2324,3,155,23412,2}
	
	 fmt.Println(arreglo_desordenado)
	
	 arreglo_ordenado := insertion_sort(arreglo_desordenado)
	
	 fmt.Println(arreglo_ordenado)
	}
	//i=1
	//j es igual a i 
	// 412 > 12
 func insertion_sort(arreglo []int) []int {
	 for i := 1; i<len(arreglo); i++{
		j :=i
		for j > 0 && arreglo[j-1]>arreglo[j]{
			fmt.Printf("%d se esta intercambiando con %d \n",arreglo[j-1],arreglo[j] )
 			swap(j-1,j,&arreglo);
			j--;
		}
	 }
	 return arreglo

 }
 func swap(previo int,actual int,punter_arreglo *[]int){
	 //actual 1
	 //previo 0
	 //arreglo en previo igual a 412
	 //arreglo de actual igual a 412
	 //[412,12,123,1,2324,3,155,23412,2]
	 arreglo:= *punter_arreglo
	 copia:= arreglo[actual]
	arreglo[actual] = arreglo[previo]
	 arreglo[previo]= copia	
}