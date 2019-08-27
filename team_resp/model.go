package team_resp

import "github.com/ypapax/players/team"

type TeamResponse struct {
	Data struct{
		Team team.Team `json:"team"`
	}
}
