package main

import(
        "fmt"
        "bufio"
        "os"
)

func main(){

        f, err := os.Open("day3_input.txt")
        if err != nil{
                panic(err)
        }

        scanner := bufio.NewScanner(f)
        scanner.Split(bufio.ScanRunes)
        var cmds []string

        for scanner.Scan(){
                x := scanner.Text()
                if x == ">" || x == "<" || x == "^" || x == "v"{
                        cmds = append(cmds, x)
                }
        }

        var houses = make(map[string]int)
        robo := false
        currentX := 0
        currentY := 0
        RcurrentX := 0
        RcurrentY := 0

        idx:= fmt.Sprintf("%d x%d", currentX, currentY)
        houses[idx] = houses[idx] + 1

        for _ , cmd := range cmds{
                if robo{
                        switch{
                                case cmd == "<":
                                        RcurrentX = RcurrentX - 1
                                case cmd == ">":
                                        RcurrentX = RcurrentX + 1
                                case cmd == "^":
                                        RcurrentY = RcurrentY + 1
                                case cmd == "v":
                                       RcurrentY = RcurrentY - 1
                        }
                        idx := fmt.Sprintf("%d x %d", RcurrentX, RcurrentY)
                        houses[idx] = houses[idx] + 1
                        robo = false

                } else {
                         switch{
                                 case cmd == "<":
                                         currentX = currentX - 1
                                 case cmd == ">":
                                         currentX = currentX + 1
                                 case cmd == "^":
                                         currentY = currentY + 1
                                 case cmd == "v":
                                         currentY = currentY - 1
                        }
                        idx := fmt.Sprintf("%d x %d", currentX, currentY)
                        houses[idx] = houses[idx] + 1
                        robo = true
                }
       }

        var visited int
        for _ , el := range houses{
                if el > 0{
                        visited = visited + 1
                }
        }
        fmt.Println(visited)
}
