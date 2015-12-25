package herohunt

type Game struct {
	Players map[string]*Player
}

func NewGame() *Game {
	g := new(Game)
	g.Players = map[string]*Player{}
	return g
}

func (g *Game) AddPlayer(ID string) error {
	p := NewPlayer()
	p.ID = ID
	g.Players[ID] = p
	return nil
}

func (g *Game) Start() error {
	return nil
}

func (g *Game) Stop() error {
	return nil
}

func (g *Game) Wait() {

}
