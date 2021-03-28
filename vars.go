package main

import (
	_ "embed"

	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

const (
	width              = 800 * 2
	height             = 450 * 2
	gravity            = 980
	terrainWidth       = 5.0
	fps                = 30
	particlePlaceSpeed = 180
)

// Quality settings
const (
	size       = 10  // Blur Amount
	quality    = 4   // Blur Quality
	directions = 32  // Blur Directions
	threshold  = 0.5 // Blur Threshold
	iterations = 10  // Physics Quality
)

// Autocalculated
var shader r.Shader
var space *cp.Space
var tex r.RenderTexture2D

// Embeds

//go:embed shaders/blur.fs
var blurFs string

//go:embed shaders/default.vs
var defaultVs string

// Layers
type LayerType int

const (
	LayerTerrain LayerType = 1000
	LayerWater   LayerType = 0
	LayerLava    LayerType = 1
	LayerStone   LayerType = 2
	LayerBomb    LayerType = 3
)

// Layer contains all the data for a layer, which is a material
type Layer struct {
	Type      LayerType
	Friction  float64
	Mass      float64
	Radius    float64
	Name      string
	Color     r.Color
	Extradata map[string]float64
}

var layers = []Layer{
	{
		Type:     LayerWater,
		Friction: 0,
		Mass:     1,
		Radius:   5,
		Color:    r.NewColor(0, 255, 255, 255),
		Name:     "Water",
	},
	{
		Type:     LayerLava,
		Friction: 0,
		Mass:     1,
		Radius:   5,
		Color:    r.NewColor(255, 255/2, 0, 255),
		Name:     "Lava",
	},
	{
		Type:     LayerStone,
		Friction: 10,
		Mass:     10,
		Radius:   5,
		Color:    r.NewColor(255/5, 255/5, 255/5, 255),
		Name:     "Stone",
	},
	{
		Type:     LayerBomb,
		Friction: 10,
		Mass:     1,
		Radius:   5,
		Color:    r.NewColor(77, 196, 109, 255),
		Name:     "Bomb",
		Extradata: map[string]float64{
			"BombRange": 100,
			"BombForce": 1000,
		},
	},
}

// Just a floor (for now)
var terrain = []cp.Vector{
	{X: 0, Y: 0},
	{X: 0, Y: height - terrainWidth}, {X: width, Y: height - terrainWidth},
	{X: width, Y: 0},
}
