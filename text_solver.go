package main

import (
  "fmt"
)


func solve(bo [][]int) bool {
  find := findEmpty(bo)
  if find == nil {
    return true
  }
  row := find[0]
  col := find[1]

  for i := 1; i<10; i++ {
    if valid(bo, i, find) {
      bo[row][col] = i

      if solve(bo) {
        return true
      }

      bo[row][col] = 0
    }

  }

  return false
}

func valid(bo [][]int, num int, pos []int) bool{
  // Check row
  for i := 0; i < len(bo[0]); i++ {
    if bo[pos[0]][i] == num && pos[1] != 1 {
      return false
    }
  }

  // Check column
  for i := 0; i < len(bo[0]); i++ {
    if bo[i][pos[1]] == num && pos[0] != i {
      return false
    }
  }

  // Check box
  box_x := pos[1]/3
  box_y := pos[0]/3

  for i := box_y*3; i < box_y*3 + 3; i++ {
    for j := box_x*3; j < box_x*3 + 3; j++ {
      if bo[i][j] == num && i != pos[0] && j != pos[1] {
        return false
      }
    }
  }
  return true
}

func printBoard(bo [][]int){
  for i := range bo[0] {
    if i%3 == 0 && i!=0{
      fmt.Println("- - - - - - - - - - - -")
    }
    for j := range bo[0] {
      if j%3 == 0 && j!= 0 {
        fmt.Printf(" | ")
      }
      if j == 8 {
        fmt.Printf("%d", bo[i][j])
      }else {
        fmt.Printf("%d ", bo[i][j])
      }
    }
    fmt.Println()
  }
}

func findEmpty(bo [][]int) []int{
  for i := range bo[0] {
    for j := range bo[0] {
      if bo[i][j] == 0 {
        return []int{i,j}
      }
    }
  }
  return nil
}

func main() {
  board := [][]int{
    {7,8,0,4,0,0,1,2,0},
    {6,0,0,0,7,5,0,0,9},
    {0,0,0,6,0,1,0,7,8},
    {0,0,7,0,4,0,2,6,0},
    {0,0,1,0,5,0,9,3,0},
    {9,0,4,0,6,0,0,0,5},
    {0,7,0,3,0,0,0,1,2},
    {1,2,0,0,0,7,4,0,0},
    {0,4,9,2,0,6,0,0,7}}
  //printBoard(board)
  //solve(board)
  printBoard(board)
  fmt.Println("\n Solving puzzle. . . \n")
  solve(board)
  printBoard(board)
}
