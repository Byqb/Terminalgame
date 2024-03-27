package tetris

import (
	"math"
	"math/rand"
)

type piece struct {
	shape []vector
	canRotate bool
	color int
}

func (p *piece) rotateBack(){
	ang := math.Pi / 2 * 3
	p.rotateWithAngle(ang)
}

func (p *piece) rotate(){
	ang := math.Pi / 2
	p.rotateWithAngle(ang)
}
	func (p *piece) rotateWithAngle(ang float64){
	if !p.canRotate{
		return
	}

	cos := int(math.Round(math.Cos(ang)))
	sin :=int(math.Round(math.Sin(ang)))

	for i,e := range p.shape {
		ny := e.y*cos -e.x*sin
		nx := e.y*sin -e.x*cos

		p.shape[i] = vector{ny, nx}
	}
}

var pieces = []piece{
	{
		shape: []vector{{0, 0}},
		color: 1,
		canRotate: false,
	},
	{
		// L shap
		shape: []vector{{0, -1}, {0, 0}, {0, 1}, {1, 1}},
		color: 1,
		canRotate: true,
	},
	{
		// oposite L shap
		shape: []vector{{0, -1}, {0, 0}, {0, 1}, {-1, 1}},
		color: 1,		canRotate: true,

	},
	{
		// I shap
		shape: []vector{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
		color: 3,
		canRotate: true,

	},
	{
		// O shap
		shape: []vector{{0, -1}, {1, 0}, {0, 1}, {1, 0}},
		color: 4,
		canRotate: true,
	},
	{
		// + shap
		shape: []vector{{0, -1}, {0, 0}, {0, 1}, {1, 1}},
		color: 5,
		canRotate: true,

	},
	{
		// Z shap
		shape: []vector{{1, -1}, {1, 0}, {0, 0}, {0, 1}},
		color: 6,
		canRotate: true,

	},
	{
		// S shap
		shape: []vector{{0, -1}, {0, 0}, {1, 1}, {1, 1}},
		color: 7,
		canRotate: true,

	},
}


func randomPiece() piece {
	idx := rand.Intn(len(pieces)-1) + 1
	pc := pieces[idx]
	return piece{
		shape:   append([]vector(nil), pc.shape...),
		canRotate:  pc.canRotate,
		color:  pc.color,
	}
}
