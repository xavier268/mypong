package mypong

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var paddleColor = color.RGBA{255, 255, 255, 255} // white color
var paddleWidth float32 = 80                     // paddle width
var paddleHeight float32 = 2                     // paddle heigh

type Paddle struct {
	x, y float32 // center
}

// Creates a new default paddle
func NewPaddle() *Paddle {
	p := &Paddle{
		x: XMAX / 2,
		y: YMAX - 10,
	}
	return p
}

func (paddle *Paddle) Update(p *Pong) error {

	// update position from mouse x position
	X, _ := ebiten.CursorPosition()
	paddle.x = float32(X)
	if paddle.x < paddleWidth/2 {
		paddle.x = paddleWidth / 2
	}
	if paddle.x > XMAX-paddleWidth/2 {
		paddle.x = XMAX - paddleWidth/2
	}

	// update y position with time ...
	if p.ticks%(30) == 0 {
		paddle.y = max(YMAX/2, paddle.y-0.3)
		if paddle.y <= YMAX/2+1 { // reset when mid screen is reached
			paddle.y = YMAX - 10
		}
	}

	return nil

}

// Manage ball/paddle interactions
func (paddle *Paddle) Collide(p *Pong, b *Ball) {

	// Check if ball is within paddle
	if b.x > paddle.x-paddleWidth/2 && b.x < paddle.x+paddleWidth/2 {

		// Check if ball is within paddle height
		if b.y > paddle.y-paddleHeight/2 && b.y < paddle.y+paddleHeight/2 {

			// Reverse ball vertical direction
			b.dy = -b.dy
			b.dx = b.dx + (b.x-paddle.x)*paddleWidth/1000. // horizontal speed change with impact position relative to paddle width
			// increase score
			p.score = p.score + 50
		}
	}

}

func (paddle *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, paddle.x-paddleWidth/2, paddle.y-paddleHeight/2, paddleWidth, paddleHeight, paddleColor, true)
}
