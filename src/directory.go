package battleShip

import(
	"fmt"
	"github.com/google/uuid"
)

type Directory struct{
	GameSlice []*Game
}

func indexOf(slice []*Game, element uuid.UUID) int {
    for i:=0; i < len(slice);i++ {
        if element == slice[i].Player1.ID {
            return i
        }
    }
    return -1 // Return -1 if the element is not found
}

func (d *Directory) NewGame (new *Game){
	d.GameSlice = append(d.GameSlice, new)
}

//Really would only be used for testing without a UX design
func (d *Directory) Display(){
	if (len(d.GameSlice)) ==0{
		fmt.Println("No Games Available")
		return
	}
	for i:= 0;i <  (len(d.GameSlice)); i++{
		fmt.Println(d.GameSlice[i].Player1.Name , "'s Lobby\n" , "Players:" , d.GameSlice[i].PlayerCount)
	}
}

func (d *Directory) RemoveGame(id uuid.UUID){
	index := indexOf(d.GameSlice, id)
	d.GameSlice = append(d.GameSlice[:index], d.GameSlice[index+1:]...)
}