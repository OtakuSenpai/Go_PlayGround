package main
 
import ( 
  "fmt"
)

func main() {
  var ( 
    board[3][3] rune
    player rune 
    cpu rune
    winner rune 
    i int
    j int
    row int
    col int
    takenPlayerInput bool = false
    gameWon bool = false 
    found bool = false 
  )
  
  for i=0;i<3;i++ {
    for j=0;j<3;j++ {
      board[i][j] = ' '
    }
  } 

  fmt.Println("This is a tic tac toe game made in Go.")
  fmt.Println("Enter your choice of character(x/o) : ");
  fmt.Scanf("%c", &player)
  
  if player == 'x' || player == 'X' {
    cpu = 'o'
    player = 'x'
  } else {
    cpu = 'x'
    player = 'o'
  } 
    
  
  for gameWon != true {
     
    //Update board
    for i=0;i<3;i++ {
      for j=0;j<3;j++ {
        fmt.Printf("%c", board[i][j])
      }
      fmt.Print("\n")
    }       
   
    //Take player input
    for takenPlayerInput != true {
      fmt.Print("Enter row: ") 
      fmt.Scanf("%d", &row) 
      fmt.Print("\n") 
      fmt.Print("Enter column: ") 
      fmt.Scanf("%d", &col) 
      fmt.Print("\n")
     
      if board[row][col] == cpu || board[row][col] == player {
        fmt.Print("Space already taken,give a new input!\n")  
        continue
      } else if board[row][col] == ' ' {
        board[row][col] = player 
        takenPlayerInput = true     
        break
      }
    }
    
    j = 0
    //Time for Computers turn
    for i=0;i<3;i++ {
      //for Column
      //1-2,1-3,2-3,
      if board[i][j] == player {
        if board[i][j+1] == player {
          board[i][j+2] = cpu
        } else if board[i][j+2] == player {
            board[i][j+1] = cpu
          }
      } else if board[i][j+1] == player {
        if board[i][j+2] == player {
          board[i][j] = cpu
          }
        }          
      }  
    
    i = 0 
    for j=0;j<3;j++ {
      //for Row
      //1-2,1-3,2-3,
      if board[i][j] == player {
        if board[i+1][j] == player {
          board[i+2][j] = cpu
        } else if board[i+2][j] == player {
            board[i+1][j] = cpu
          }
        } else if board[i+1][j] == player {
          if board[i+2][j] == player {
            board[i][j] = cpu
          }
        }
    }  
    
    i = 0
    j = 0
    //For Left Diagonal
    //3-3,2-2,1-1
    if i==j && board[i][j] == player {
      if board[i+1][j+1] == player {
        board[i+2][j+2] = cpu
      } else if board[i+2][j+2] == player {
          board[i+1][j+1] = cpu
        }
      } else if i==j && board[i+1][j+1] == player {
      if board[i+2][j+2] == player {
        board[i][j] = cpu
      }
    }
    
    i = 0
    j = 0
    //For Right Diagonal
    //2-0,1-1,0-2
    if board[i][j+2] == player {
      if board[i+1][j+1] == player {
        board[i+2][j] = cpu
      } else if board[i+2][j] == player {
        board[i+1][j+1] = cpu
      }
    } else if board[i+1][j+1] == player {
      if board[i+2][j+0] == player {
        board[i][j] = cpu 
      }
    }
    //Done 
    
    //Now to check for winning party
    for i=0;i<3;i++ {
      if board[i][j] == board[i][j+1] {
        if board[i][j+1] == board[i][j+2] {
          winner = board[i][j]
          found = true
        } 
      }
    } 
    for j=0;j<3;j++ {
      if board[i][j] == board[i+1][j] { 
        if board[i+1][j] == board[i+2][j] {
          winner = board[i][j]
          found = true
        } 
      }
    }
    i = 0
    j = 0
    if board[i][j] == board[i+1][j+1] {
      if board[i+1][j+1] == board[i+2][j+2] {                                  
        winner = board[i+1][j+1]
        found = true
      }
    } else if board[i+2][j] == board[i+1][j+1] {
      if board[i+1][j+1] == board[i][j+2] {
        winner = board[i+1][j+1] 
        found = true    
      }
    }
    if found == true {
      if winner == player {
        fmt.Println("You won!!!")
        break
      } else if winner == cpu {
        fmt.Println("The computer won !")
        break
      }
    }
    takenPlayerInput = false
  }
}                
