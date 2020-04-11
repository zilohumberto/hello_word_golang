package main

import "fmt"

var memo []uint64 

func fibo(x uint64) uint64{
    if x==0 {
        return 0
    }
    if x==1 || x==2 {
        return 1
    }
    if memo[x] != 0{
        return memo[x]
    }
    memo[x] = fibo(x-1) + fibo(x-2)
    return memo[x]
}

func main(){
    memo = make([]uint64, 51)
    fmt.Printf("Fibonacci of %v is %v\n", 30, fibo(30))
    memo = make([]uint64, 51)
    fmt.Printf("Fibonnaci of %v is %v\n", 50, fibo(50))
}
