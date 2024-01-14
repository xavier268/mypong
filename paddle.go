package mypong

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var paddleColor = color.RGBA{255, 255, 255, 255} // white color
var paddleSpeed = float32(3.5)                   // paddle speed

type Paddle struct {
	x, y float32 // center
	w    float32 // width
	h    float32 // height
}

// Creates a new default paddle
func NewPaddle() *Paddle {
	p := &Paddle{
		x: XMAX / 2, y: YMAX - 10, w: 40, h: 4,
	}
	return p
}

func (paddle *Paddle) Update(p *Pong) error {

	// update position from keys
	if ebiten.IsKeyPressed(ebiten.KeyNumpad4) {
		paddle.x -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyNumpad6) {
		paddle.x += paddleSpeed
	}
	if paddle.x < paddle.w/2 {
		paddle.x = paddle.w / 2
	}
	if paddle.x > XMAX-paddle.w/2 {
		paddle.x = XMAX - paddle.w/2
	}
	return nil
}

// Manage ball/paddle interactions
func (paddle *Paddle) Collide(p *Pong, b *Ball) {

	// Check if ball is within paddle
	if b.x > paddle.x-paddle.w/2 && b.x < paddle.x+paddle.w/2 {

		// Check if ball is within paddle height
		if b.y > paddle.y-paddle.h/2 && b.y < paddle.y+paddle.h/2 {

			// Reverse ball vertical direction
			b.dy = -b.dy
			// increase score
			p.score = p.score + 5
		}
	}

}

func (paddle *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, paddle.x-paddle.w/2, paddle.y-paddle.h/2, paddle.w, paddle.h, paddleColor, true)
}
