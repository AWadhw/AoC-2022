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

    first := 0
    second := 0
    third := 0

  //  for scanner.Scan(){
  //      lineText := scanner.Text()
  //      if(lineText != ""){ 
  //          perLineCal, _ := strconv.Atoi(lineText)
  //          calBuffer += perLineCal
  //          if(calBuffer > maxCal){
  //              maxCal = calBuffer
  //          }
  //      }else if(lineText == ""){
  //          calBuffer = 0;
  //      }
  //  }
for scanner.Scan(){
    lineText := scanner.Text()
    if(lineText != ""){
        perLineCal, _ := strconv.Atoi(lineText)
        calBuffer += perLineCal
        
    }else if(lineText == ""){
        fmt.Println(calBuffer)
        fmt.Println("THIS IS...", first, second, third)
        if(calBuffer > first){
            if(first > second){
                if(second > third){
                    third = second
                }
                second = first
            }
            first = calBuffer
        }else if(calBuffer > second){
            if(second > third){
                third = second
            }
            second = calBuffer
        }else if(calBuffer > third){
            third = calBuffer
        }
        calBuffer = 0;
    }
}

fmt.Println("THIS IS...", first, second, third)
total := 0
total = first + second + third
fmt.Println("THE TOTAL IS:")
fmt.Println(total)
fmt.Println("----------------------------------")

    fmt.Println(maxCal)


    if err := scanner.Err(); err!= nil {
        log.Fatal(err)
    }
}
