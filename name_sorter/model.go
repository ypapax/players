package name_sorter

import "github.com/ypapax/players/player"

type NameSorter []player.Player

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name + " " + a[i]. < a[j].Axis }



