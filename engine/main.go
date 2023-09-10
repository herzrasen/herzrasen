package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/herzrasen/engine/match"
	"github.com/herzrasen/engine/team"
)

var args struct {
	HomeTeamFile string `arg:"--home-team-file,required"`
	AwayTeamFile string `arg:"--away-team-file,required"`
}

func main() {
	arg.MustParse(&args)

	homeTeam, err := team.FromFile(args.HomeTeamFile)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(homeTeam)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
	awayTeam, err := team.FromFile(args.AwayTeamFile)
	if err != nil {
		panic(err)
	}
	m := match.NewMatch(homeTeam, awayTeam)
	fmt.Println("%v", m)
}
