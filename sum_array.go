package main


import "fmt"


func sumArray(arr ...int) int {
    sum := 0
    for _, v := range arr{
        sum += v
    }
    return sum
}

func main(){
    
    x := []int{2,3,4,6}
    fmt.Printf("Array sum(%v)=%d\n",x, sumArray(x...))


    y := []int{100, 5, -2}
    fmt.Printf("Array sum(%v)=%d\n", y, sumArray(y...))
}
