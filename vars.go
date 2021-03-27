package main

import (
	_ "embed"

	"github.com/jakecoffman/cp"
	r "github.com/lachee/raylib-goplus/raylib"
)

const (
	width                = 800 * 2
	height               = 450 * 2
	gravity              = -100
	terrainWidth float64 = 5
)

// Quality settings
var (
	size       = 100
	quality    = 3
	directions = 32
	threshold  = 0.2
)

// Autocalculated
var shader r.Shader
var textures []r.RenderTexture2D
var space *cp.Space

// Embeds

//go:embed shaders/blur.fs
var blurFs string

//go:embed shaders/default.vs
var defaultVs string

// Layers
type LayerType int

const (
	LayerTerrain = 0
	LayerWater   = 1
)

// Layer contains all the data for a layer, which is a material
type Layer struct {
	Type     LayerType
	Friction float64
	Mass     float64
	Radius   float64
	Color    r.Vector4
}

var layers = []Layer{
	{
		Type:     LayerWater,
		Friction: 0,
		Mass:     1,
		Radius:   5,
		Color:    r.NewVector4(0, 1, 1, 1),
	},
}

// Just a floor (for now)
var terrain = []cp.Vector{
	{X: 0, Y: height - terrainWidth}, {X: width, Y: height - terrainWidth},
}
