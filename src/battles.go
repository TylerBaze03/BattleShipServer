package battleShip

import (
	"fmt"
	"strings"
	"slices"
)

type Game struct {
	//2 = 2 size boat, 3 = 3 size boat, 4 = 4 size boat, 5= 5 size boat, 0 = empty
	Player1Board [][]string
	//2 = 2 size boat, 3 = 3 size boat, 4 = 4 size boat, 5= 5 size boat, 0 = empty
	Player2Board [][]string

	PlayerCount int

	Player1 *Player
	Player2 *Player
}

type Player struct {
	ID   int
	Name string
	//0 = empty 1 = hit 2 = miss
	MissList [][]int
}

const (
	size = 10
)
var boats = []string{"2","3","4","5"}


func NewGame(p1 *Player) *Game {
	return &Game{Player1Board: make([][]string, size), Player1: p1, PlayerCount: 1}
}

func (g *Game) AddPlayer( player *Player){
	if g.PlayerCount ==2 {
		fmt.Println("Game Full")
		return
	}
	g.PlayerCount++
	g.Player2Board = make([][]string, size)
	g.Player2 = player
}

//0= exit , 1 = create new game
func (g *Game) EndGame() int {
	fmt.Println("Game Has Ended\n Create New Game? y/n")
	
	userIn := ""
	fmt.Scan(&userIn)
	userIn = strings.ToLower(userIn)
	for{
		
		if userIn != "n" && userIn != "y"{
			fmt.Println("Invalid input please try again\nGame Has Ended\n Create New Game? y/n")
		}else{
			break
		}
	}
	if userIn == "y"{
		return 1
	}else{
		return 0
	}
}

// 1 = player 1, 2= player 2  
// 0 = empty/err 1 = hit 2 = miss
func (g *Game) Fire(player int, cords [2]int)int{
	//player 1 firing at player 2
	if player == 1{
		if g.Player2Board[cords[0]][cords[1]] == "0"{
			return 0
		}else if slices.Contains(boats, g.Player2Board[cords[0]][cords[1]]){
			return 1
		}else{
			return 2
		}
	}else{//player 2 firing at player 1
		if g.Player1Board[cords[0]][cords[1]] == "0"{
			return 0
		}else if slices.Contains(boats, g.Player1Board[cords[0]][cords[1]]){
			return 1
		}else{
			return 2
		}
	}
}


