package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main(){

    f, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

//    for scanner.Scan(){
//        lineText := scanner.Text()
//        if (lineText == "") {
//            fmt.Println("EMPTYLINE")
//        }else {
//            fmt.Println(lineText)
//        }
//    }

    maxCal := 0
    calBuffer := 0

    for scanner.Scan(){
        lineText := scanner.Text()
        if(lineText != ""){ 
            perLineCal, _ := strconv.Atoi(lineText)
            calBuffer += perLineCal
            if(calBuffer > maxCal){
                maxCal = calBuffer
            }
        }else if(lineText == ""){
            calBuffer = 0;
        }
    }

    fmt.Println(maxCal)


    if err := scanner.Err(); err!= nil {
        log.Fatal(err)
    }
}
