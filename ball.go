package mypong

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const RADIUS = 2

var ballColor = color.RGBA{255, 255, 255, 255}

type Ball struct {
	x, y   float32
	dx, dy float32 // speed per tick
}

func NewBall() *Ball {
	b := &Ball{
		x:  XMAX / 2,
		y:  YMAX / 2,
		dx: 1.5 * rand.Float32(),
		dy: 1.2 + rand.Float32(),
	}
	return b
}

// Update implements ebiten.Game.
func (ball *Ball) Update(p *Pong) error {

	if ball.x < 0 && ball.dx < 0 {
		ball.dx = -ball.dx
	}
	if ball.x > XMAX && ball.dx > 0 {
		ball.dx = -ball.dx
	}
	if ball.y < 0 && ball.dy < 0 {
		ball.dy = -ball.dy
	}
	if ball.y > YMAX && ball.dy > 0 {
		ball.dy = -ball.dy
		// penalize for letting a ball go out of bounds
		p.score = p.score - 20
		// add new ball to the list
		b2 := NewBall()
		b2.dx = 1.3*ball.dx + rand.Float32()
		b2.dy = 1.3*ball.dy + rand.Float32()
		p.balls = append(p.balls, b2)
	}

	ball.x += ball.dx
	ball.y += ball.dy
	return nil
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, ball.x, ball.y, RADIUS, ballColor, true)
}
