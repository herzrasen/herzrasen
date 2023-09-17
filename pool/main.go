package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/herzrasen/common/gender"
	"github.com/herzrasen/common/player"
	"github.com/herzrasen/pool/config"
	"github.com/herzrasen/pool/random"
	"os"
)

var args struct {
	Num            int     `arg:"-n,--num" default:"1"`
	Gender         string  `arg:"-g,--gender" default:"male"`
	MinAttrValue   float32 `arg:"--min-attr-value" default:"0.0"`
	MaxAttrValue   float32 `arg:"--max-attr-value" default:"5.0"`
	MinYearOfBirth uint32  `arg:"--min-year-of-birth" default:"2000"`
	MaxYearOfBirth uint32  `arg:"--max-year-of-birth" default:"2006"`
	MinHeight      uint32  `arg:"--min-height" default:"160"`
	MaxHeight      uint32  `arg:"--max-height" default:"200"`
	MinWeight      uint32  `arg:"--min-height" default:"55000"`
	MaxWeight      uint32  `arg:"--max-height" default:"100000"`
	NameConfig     string  `arg:"--name-config" default:"names.json"`
	OutputFile     string  `arg:"-o,--output-file"`
}

func main() {
	arg.MustParse(&args)
	nameConfig, err := config.FromFile(args.NameConfig)
	if err != nil {
		panic(err)
	}
	r := random.Random{
		Gender:     gender.FromString(args.Gender),
		NameConfig: nameConfig,
		MinYear:    args.MinYearOfBirth,
		MaxYear:    args.MaxYearOfBirth,
		MinHeight:  args.MinHeight,
		MaxHeight:  args.MaxHeight,
		MinWeight:  args.MinWeight,
		MaxWeight:  args.MaxWeight,
		MinAttr:    args.MinAttrValue,
		MaxAttr:    args.MaxAttrValue,
	}
	var players []player.Player
	for i := 0; i < args.Num; i++ {
		p := r.Player()
		players = append(players, p)
	}
	var out bytes.Buffer
	data, err := json.Marshal(&players)
	if err != nil {
		panic(err)
	}
	if err = json.Indent(&out, data, "", "  "); err != nil {
		panic(err)
	}

	if args.OutputFile != "" {
		err := os.WriteFile(args.OutputFile, out.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(out.String())
	}
}
