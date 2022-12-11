package entity

type Team struct {
	ID      string
	Name    string
	Players []*Player
}

func NewTeam(id string, name string) *Team {
	return &Team{
		ID:   id,
		Name: name,
	}
}

func (team *Team) AddPlayer(player *Player) {
	team.Players = append(team.Players, player)
}

func (team *Team) RemovePlayer(player *Player) {
	for idx, playerInTeam := range team.Players {
		if playerInTeam.ID == player.ID {
			team.Players = append(team.Players[:idx], team.Players[idx+1:]...)
			return
		}
	}
}
