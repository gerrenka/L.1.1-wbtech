package main

import (
	"fmt"
	"time"
)

func main(){
    var N int
    fmt.Scan(&N)
    ch:=make(chan int)
    timelimit:= time.After(time.Duration(N)*time.Second)
    go func(){
        counter:=1
        for{
            select{
            case <-timelimit:
                close(ch)
                return
            default:
                ch<-counter
                counter++
                time.Sleep(500*time.Millisecond)
            }
        }
    }()
    for val:=range ch{
        fmt.Println("Received:",val)
    }
    fmt.Println("Program finished.")




}
