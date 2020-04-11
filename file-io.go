package main

import "fmt"
import "os"
import "log"
import "bufio"
import "strings"

func main() {

    f,err := os.Open("myapp.log")

    if err !=nil {
        log.Fatal(err)
    }

    defer f.Close()

    r:= bufio.NewReader(f)

    for {
        s,err:= r.ReadString('\n')
        if err !=nil {
            break
        }

        if strings.Contains(s, "ERROR"){
            fmt.Println(s)
        }
    }


}



