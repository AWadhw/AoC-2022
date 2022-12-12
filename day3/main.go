package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

func findCommon(myString string) byte{
    strLen := len(myString)
    m := make(map[byte]int)
    var toReturn byte = '*'
    for i:=0; i < strLen/2 ; i++{
       m[myString[i]]++
    }
    for i:= strLen/2; i < strLen; i++{
        if(m[myString[i]] > 0){
            toReturn = myString[i]
            return toReturn
        }
    }
    return toReturn
}

func getCharVal(myChar byte) int{
    return 0
}

func main(){
    f, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    var lineBuffer string
    var commonScore int
    for scanner.Scan(){
        lineBuffer = scanner.Text()
        fmt.Println(lineBuffer)
        myCommon := findCommon(lineBuffer)
        //fmt.Printf("my a is %i and z is %i and A is %i and Z is %i", 'a', 'z', 'A', 'Z')
        fmt.Printf("my commmon is %c %i \n", myCommon, myCommon)
        if(myCommon >= 97){
            commonScore += int(myCommon) - 96
        }else{
            commonScore += int(myCommon) - 38
        }
        fmt.Println("Score for this line is:", commonScore)
    }
}
