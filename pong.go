package mypong

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const XMAX, YMAX = 320, 240 // logical game space dimensions

type Pong struct {
	paddles []*Paddle
	balls   []*Ball
	paused  bool
	score   int
	start   time.Time
}

func NewPong() *Pong {
	p := &Pong{}
	p.paddles = append(p.paddles, NewPaddle())
	p.balls = append(p.balls, NewBall())
	p.score = 0
	return p
}

var _ ebiten.Game = &Pong{}

// Layout implements ebiten.Game.
// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (*Pong) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return XMAX, YMAX
}

// Update implements ebiten.Game.
func (p *Pong) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.paused = !p.paused
	}

	if p.paused {
		return nil
	}

	for _, paddle := range p.paddles {
		paddle.Update(p)
	}

	for _, ball := range p.balls {

		if err := ball.Update(p); err != nil {
			return err
		}

		// Check ball/paddle collisions
		for _, paddle := range p.paddles {
			paddle.Collide(p, ball)
		}
	}

	if len(p.balls) >= 20 {
		fmt.Println("too many balls - you lost !")
		p.paused = true
	} else {
		p.score = p.score + 1
	}

	// every 10 sec, increase game speed
	if time.Since(p.start) > 10*time.Second {
		p.start = time.Now()
		for _, ball := range p.balls {
			ball.dx, ball.dy = ball.dx*1.1, ball.dy*1.1
		}
		paddleSpeed = paddleSpeed * 1.1
	}

	return nil
}

// Draw implements ebiten.Game.
// Draw draws the game screen.
func (p *Pong) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff}) // clear screen

	// Draw paddles and balls
	for _, ball := range p.balls {
		ball.Draw(screen)
	}
	for _, paddle := range p.paddles {
		paddle.Draw(screen)
	}

	// Print score on screen
	if p.paused {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d\nPAUSED", p.score))
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", p.score))
	}

	if len(p.balls) >= 20 {
		ebitenutil.DebugPrint(screen, "\n\nYOU LOST !")
	}

}
