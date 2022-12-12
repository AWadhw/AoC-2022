package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

func roundPoint(x int, y int) int { //where x is us and y is opponent
    var roundPoint int = 0;
    if(x == y){
        roundPoint = 3    
    }else if(x == (y+1)){ //for a win
        roundPoint = 6
    }else if((x == 1) && (y == 3)){
        roundPoint = 6
    }

    return x + roundPoint
    
}

func main(){

    m := make(map[byte]int)
    m['X'] = 1; //Rock
    m['A'] = 1; //Rock
    m['Y'] = 2; //Paper
    m['B'] = 2; //Paper
    m['Z'] = 3; //Scissors
    m['C'] = 3; //Scissors

    f, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    var lineBuffer string
    var perLineScore int
    totalScore := 0
    for scanner.Scan(){
        lineBuffer = scanner.Text()
        if(lineBuffer == ""){
            continue; //don't want to iterate blanks
        }
        fmt.Println(lineBuffer)
        fmt.Printf("%c\n", lineBuffer[0])
        fmt.Printf("%c\n", lineBuffer[2])
        
        perLineScore = roundPoint(m[lineBuffer[2]], m[lineBuffer[0]])
        totalScore += perLineScore
        fmt.Println(perLineScore)
       // fmt.Println(m["X"])
       // fmt.Println(m["Y"])
    }
    fmt.Println(totalScore)
}
