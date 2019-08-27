package team

import (
	"fmt"
	"sort"
	"strings"
)

type Team struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

type Player struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Teams map[int]Team
}

func (p Player) String() string {
	var teamNames []string
	for _, t := range p.Teams {
		teamNames = append(teamNames, t.Name)
	}
	sort.Strings(teamNames)
	return fmt.Sprintf("%+v; %+v; %+v", p.Name, p.Age, strings.Join(teamNames, ", "))
}
