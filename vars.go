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
	fps                = 30
	particlePlaceSpeed = 180
	fontSize           = 24
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
type RenderType int

const (
	LayerTerrain LayerType = 0
	LayerWater   LayerType = 1
	LayerLava    LayerType = 2
	LayerStone   LayerType = 3
	LayerBomb    LayerType = 4

	RenderParticle = 0
	RenderSegment  = 1
)

// Layer contains all the data for a layer, which is a material
type Layer struct {
	Type       LayerType
	RenderType RenderType
	Friction   float64
	Mass       float64
	Size       float64 // Radius for particles, thickness for segments
	Name       string
	Color      r.Color
	Extradata  map[string]float64
}

var layers = []Layer{
	{
		Type:       LayerTerrain,
		RenderType: RenderSegment,
		Friction:   1,
		Size:       5,
		Color:      r.NewColor(0, 0, 0, 255),
		Name:       "Terrain",
	},
	{
		Type:       LayerWater,
		RenderType: RenderParticle,
		Friction:   0,
		Mass:       1,
		Size:       5,
		Color:      r.NewColor(0, 255, 255, 255),
		Name:       "Water",
	},
	{
		Type:       LayerLava,
		RenderType: RenderParticle,
		Friction:   0,
		Mass:       1,
		Size:       5,
		Color:      r.NewColor(255, 255/2, 0, 255),
		Name:       "Lava",
	},
	{
		Type:       LayerStone,
		RenderType: RenderParticle,
		Friction:   10,
		Mass:       10,
		Size:       5,
		Color:      r.NewColor(255/5, 255/5, 255/5, 255),
		Name:       "Stone",
	},
	{
		Type:       LayerBomb,
		RenderType: RenderParticle,
		Friction:   10,
		Mass:       1,
		Size:       5,
		Color:      r.NewColor(77, 196, 109, 255),
		Name:       "Bomb",
		Extradata: map[string]float64{
			"BombRange": 100,
			"BombForce": 1000,
		},
	},
}

// A floor and 2 walls
var terrain = []cp.Vector{
	{X: 0, Y: 0},
	{X: 0, Y: height - layers[LayerTerrain].Size},
	{X: width, Y: height - layers[LayerTerrain].Size},
	{X: width, Y: 0},
}

// Tools
type ToolType int

const (
	PlaceTool   = 0
	UselessTool = 1
)

type Tool struct {
	Type      ToolType
	Name      string
	Tool      func(Tool)
	IntData   map[string]int
	FloatData map[string]float64
	BoolData  map[string]bool
}

var tools = []Tool{
	{
		Name: "Place Tool",
		Tool: placeTool,
		IntData: map[string]int{
			"tool": 0,
		},
		FloatData: map[string]float64{
			"p1x": 0,
			"p1y": 0,
		},
		BoolData: map[string]bool{
			"hasPlaced": false,
		},
	},
	{
		Name: "Useless tool",
		Tool: uselessTool,
	},
}
