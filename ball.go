package mypong

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const RADIUS = 2

type Ball struct {
	x, y   float32
	dx, dy float32 // speed per tick
	color  color.RGBA
}

func NewBall() *Ball {
	b := &Ball{
		x:     XMAX / 2,
		y:     YMAX / 2,
		dx:    rand.Float32(),
		dy:    (1.1 + rand.Float32()) / 2.,
		color: color.RGBA{255, 255, 255, 255},
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
		PlayAudioBoing()
		// penalize for letting a ball go out of bounds
		p.score = p.score - 100
		// add new ball to the list, based on 50% probability ...
		if rand.Float32() < 0.5 {
			b2 := NewBall()
			b2.dx = 1.1*ball.dx + rand.Float32()
			b2.dy = 1.1*ball.dy + rand.Float32()
			b2.color.G = max(0, ball.color.G/2)
			b2.color.B = max(0, ball.color.B/2)
			p.balls = append(p.balls, b2)
		}

	}

	ball.x += ball.dx
	ball.y += ball.dy
	return nil
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, ball.x, ball.y, RADIUS, ball.color, true)
}
