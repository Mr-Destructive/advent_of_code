package main

import (
	"log"
    "strings"

	"aoc2022/helpers"
)

func four_set(set string) (bool) {
    is_present := false
    for i:=1; i<len(set); i++ {
        bef, aft, _ := strings.Cut(set, string(set[i]))
        removed_set := bef + aft
        if !strings.ContainsRune(removed_set, rune(set[i])){
            is_present = true
        } else {
            is_present = false 
            break
        }
    }
    return is_present
}

func main() {
	file, err := helpers.ReadInput("inputs/input6.test")
    helpers.HandleError(err)
    for i:=3; i< len(file[0]); i+=1 {
        is_repeated := four_set(string(file[0][i-3:i+1]))
        if is_repeated {
            log.Println(i+1)
            // PART 1
            // 1480
            break
        } 
    }
    for i:=13; i< len(file[0]); i+=1 {
        is_repeated := four_set(string(file[0][i-13:i+1]))
        if is_repeated {
            log.Println(i+1)
            // PART 2
            // 2746
            break
        } 
    }
}
