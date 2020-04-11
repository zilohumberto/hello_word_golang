package main

import "fmt"

func fibo(n uint64) uint64{
    if n==0{
        return 0
    }
    if n==1 || n==2 {
        return 1
    }
    return fibo(n-1) + fibo(n-2)
}

func main(){
    fmt.Printf("Fibo of %v is %v\n", 5, fibo(5))
    fmt.Printf("Fibo of %v is %v\n", 10, fibo(10))
    fmt.Printf("Fibo of %v is %v\n", 30, fibo(30))
    fmt.Printf("Fibo of %v is %v\n", 50, fibo(50))
}
