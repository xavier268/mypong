package mypong

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const XMAX, YMAX = 320, 240 // logical game space dimensions

func init() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

type Pong struct {
	paddles []*Paddle
	balls   []*Ball
	paused  bool
	score   int
	ticks   int
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

	UpdateAudioBackground() // sounds even if paused ...

	if p.paused {
		return nil
	}

	p.ticks += 1

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

	// every n ticks, increase ball speed
	if p.ticks%(60*9) == 0 { // 10 sec

		// increase ball speed, adjust paddle height, until a certain speed
		for _, ball := range p.balls {
			if math.Abs(float64(ball.dy)) < 10 {
				ball.dx = ball.dx * 1.05
				ball.dy = ball.dy * 1.1
				paddleHeight = max(3*ball.dy, paddleHeight)
			}
		}
	}

	// every n ticks, decrease paddle size
	if p.ticks%(60*11) == 0 { // 10 sec
		// change paddle size
		paddleWidth = paddleWidth * 0.8
		if paddleWidth <= 12 {
			paddleWidth = 80
			paddleColor = color.RGBA{255, 255, 255, 255} // white color
		} else {
			paddleColor.B = max(0, paddleColor.B>>1)
			paddleColor.G = max(0, paddleColor.G>>1)
		}
	}

	// every n ticks, forget a few balls
	if p.ticks%(60*43) == 0 {

		if len(p.balls) > 1 {
			n := max(1, (len(p.balls) / 2))
			n = min(n, len(p.balls)-1)
			p.balls = p.balls[n:]
		}
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
	if p.score >= 50000 {
		p.paused = true // game over, no need to draw anything else on screen.
		ebitenutil.DebugPrint(screen, "\n\nYOU WON !")
	}

}
