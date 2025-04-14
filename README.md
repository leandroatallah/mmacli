# MMA CLI Game

A command-line interface mixed martial arts game written in Go where two players engage in turn-based combat using mixed martial arts moves.

## Features

- Turn-based combat system
- Initiative system to determine who plays first
- Defense system with different strategies
- Critical hit and fail system
- Health tracking
- In-game time tracking

## How to Play

1. Start the game
2. Enter names for both fighters
3. Each round:
   - Roll for initiative (if applicable)
   - Choose your attack move
   - Opponent chooses defense strategy
   - Damage is calculated
   - Health status is displayed

### Combat Mechanics

- **Special Rules:**
  - Rolling a 6 triggers a Critical Hit with bonus damage
  - Rolling a 1 results in a Critical Fail (miss)

## Requirements

- Go 1.23.6 or higher

## Installation

```bash
git clone   `git@github.com:leandroatallah/mmacli.git`
cd mmacli
go build
```

