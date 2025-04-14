package main

import (
	"mmacli/models"
	"mmacli/myfmt"
	"mmacli/utils"
)

func PlayersInitiative() (playerIndex int, isDraw bool) {
	myfmt.PrintDelay("\n# Players roll initiative\n\n")
	for {
		type p struct {
			name string
			roll int
		}
		p1 := p{models.GetPlayerByIndex(1).Name, utils.RollDice()}
		p2 := p{models.GetPlayerByIndex(2).Name, utils.RollDice()}

		for _, p := range []p{p1, p2} {
			myfmt.PrintDelay("- %s rolls: ", p.name)
			myfmt.PrintDelay("%d\n", p.roll)
		}

		if p1.roll > p2.roll {
			myfmt.PrintDelay("- %s (player 1) is next to play\n\n", p1.name)
			return 1, false
		} else if p2.roll > p1.roll {
			myfmt.PrintDelay("- %s (player 2) is next to play\n\n", p2.name)
			return 2, false
		}
		myfmt.PrintDelay("- Draw...\n")
		return 0, true
	}
}
