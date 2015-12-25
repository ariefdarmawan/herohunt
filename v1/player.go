package herohunt

type Player struct {
	Castle
	ID         string
	LoginID    string
	ScreenName string
	MapArea    int
}

func NewPlayer() *Player {
	p := new(Player)
	return p
}
