package parser

import (
	"encoding/json"
	"fmt"
	"github.com/ypapax/players/by_name"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/ypapax/players/team"
	"github.com/ypapax/players/team_resp"

	"github.com/sirupsen/logrus"
)

func GetPlayers(urlTemplate string, teams []string, timeout time.Duration) ([]team.Player, error) {
	var requiredTeams = make(map[string]struct{})
	var foundTeams = make(map[string]team.Team)
	var foundPlayers = make(map[string]team.Player)
	for _, t := range teams {
		requiredTeams[t] = struct{}{}
	}
	teamID := 1
	for {
		logrus.Tracef("teamID %+v", teamID)
		if err := func(teamID int) error {
			u := fmt.Sprintf(urlTemplate, teamID)
			b, _, err := req(u, timeout)
			if err != nil {
				logrus.Error(err)
				return err
			}
			var tr team_resp.TeamResponse
			if err := json.Unmarshal(b, &tr); err != nil {
				logrus.Error(err)
				return err
			}
			if _, ok := requiredTeams[tr.Data.Team.Name]; !ok {
				logrus.Infof("team '%+v' is not in our list", tr.Data.Team.Name)
				return nil
			}

			foundTeams[tr.Data.Team.Name] = tr.Data.Team

			for _, p := range tr.Data.Team.Players {
				if _, ok := foundPlayers[p.Id]; !ok {
					p.Teams = make(map[int]team.Team)
					foundPlayers[p.Id] = p
				}
				foundPlayers[p.Id].Teams[tr.Data.Team.Id] = tr.Data.Team
			}
			logrus.Printf("foundTeams: %+v, requiredTeams: %+v", len(foundTeams), len(requiredTeams))
			return nil
		}(teamID); err != nil {
			logrus.Error(err)
			return nil, err
		}
		teamID++
		if len(foundTeams) == len(requiredTeams) {
			break
		}
	}
	var result []team.Player
	for _, p := range foundPlayers {
		result = append(result, p)
	}
	sort.Sort(by_name.ByName(result))
	return result, nil
}

func req(url string, requestTimeout time.Duration) ([]byte, int, error) {
	logrus.Infof("requesting %+v", url)
	client := http.Client{
		Timeout: requestTimeout,
	}
	res, err := client.Get(url)
	if err != nil {
		logrus.Errorf("err: %+v\n", err)
		return nil, 0, err
	}
	if res.StatusCode > 399 || res.StatusCode < 200 {
		err := fmt.Errorf("not good status code %+v requesting %+v", res.StatusCode, url)
		logrus.Errorf("err: %+v\n", err)
		return nil, 0, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("err: %+v\n", err)
		return nil, res.StatusCode, err
	}
	return b, res.StatusCode, nil
}
