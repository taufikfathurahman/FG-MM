package main

import (
	"fmt"
	"math"
)

type HSL struct {
	H, S, L float64
}

type HSV struct {
	H, S, V float64
}

type RGB struct {
	R, G, B float64
}

func (c *RGB) HSL() *HSL {
	r := (c.R / 255) //RGB from 0 to 255
	g := (c.G / 255)
	b := (c.B / 255)

	min := math.Min(r, math.Min(g, b)) //Min. value of RGB
	max := math.Max(r, math.Max(g, b)) //Max. value of RGB
	del := max - min                   //Delta RGB value

	l := (max + min) / 2

	var h, s float64

	if del == 0 { //This is a gray, no chroma...
		h = 0 //HSL results from 0 to 1
		s = 0
	} else { //Chromatic data...
		if l < 0.5 {
			s = del / (max + min)
		} else {
			s = del / (2 - max - min)
		}

		delR := (((max - r) / 6) + (del / 2)) / del
		delG := (((max - g) / 6) + (del / 2)) / del
		delB := (((max - b) / 6) + (del / 2)) / del

		if r == max {
			h = delB - delG
		} else if g == max {
			h = (1 / 3) + delR - delB
		} else if b == max {
			h = (2 / 3) + delG - delR
		}

		if h < 0 {
			h += 1
		}
		if h > 1 {
			h -= 1
		}
	}
	hsl := &HSL{h, s, l}
	return hsl
}

func (c *RGB) HSV() *HSV {
	r := (c.R / 255) //RGB from 0 to 255
	g := (c.G / 255)
	b := (c.B / 255)

	min := math.Min(r, math.Min(g, b)) //Min. value of RGB
	max := math.Max(r, math.Max(g, b)) //Max. value of RGB
	del := max - min                   //Delta RGB value

	v := max

	var h, s float64

	if del == 0 { //This is a gray, no chroma...
		h = 0 //HSV results from 0 to 1
		s = 0
	} else { //Chromatic data...
		s = del / max

		delR := (((max - r) / 6) + (del / 2)) / del
		delG := (((max - g) / 6) + (del / 2)) / del
		delB := (((max - b) / 6) + (del / 2)) / del

		if r == max {
			h = delB - delG
		} else if g == max {
			h = (1 / 3) + delR - delB
		} else if b == max {
			h = (2 / 3) + delG - delR
		}

		if h < 0 {
			h += 1
		}
		if h > 1 {
			h -= 1
		}
	}
	hsv := &HSV{h, s, v}
	return hsv
}

func (c *HSV) RGB() *RGB {
	var r, g, b float64
	if c.S == 0 { //HSV from 0 to 1
		r = c.V * 255
		g = c.V * 255
		b = c.V * 255
	} else {
		h := c.H * 6
		if h == 6 {
			h = 0
		} //H must be < 1
		i := math.Floor(h) //Or ... var_i = floor( var_h )
		v1 := c.V * (1 - c.S)
		v2 := c.V * (1 - c.S*(h-i))
		v3 := c.V * (1 - c.S*(1-(h-i)))

		if i == 0 {
			r = c.V
			g = v3
			b = v1
		} else if i == 1 {
			r = v2
			g = c.V
			b = v1
		} else if i == 2 {
			r = v1
			g = c.V
			b = v3
		} else if i == 3 {
			r = v1
			g = v2
			b = c.V
		} else if i == 4 {
			r = v3
			g = v1
			b = c.V
		} else {
			r = c.V
			g = v1
			b = v2
		}

		r = r * 255 //RGB results from 0 to 255
		g = g * 255
		b = b * 255
	}
	rgb := &RGB{r, g, b}
	return rgb

}

func hue_2_rgb(v1, v2, vH float64) float64 { //Function Hue_2_RGB
	if vH < 0 {
		vH += 1
	}
	if vH > 1 {
		vH -= 1
	}
	if (6 * vH) < 1 {
		return (v1 + (v2-v1)*6*vH)
	}
	if (2 * vH) < 1 {
		return v2
	}
	if (3 * vH) < 2 {
		return (v1 + (v2-v1)*((2/3)-vH)*6)
	}
	return v1
}

func (c *HSL) RGB() *RGB {
	var r, g, b float64
	if c.S == 0 { //HSL from 0 to 1
		r = c.L * 255 //RGB results from 0 to 255
		g = c.L * 255
		b = c.L * 255
	} else {
		var v1, v2 float64
		if c.L < 0.5 {
			v2 = c.L * (1 + c.S)
		} else {
			v2 = (c.L + c.S) - (c.S * c.L)
		}

		v1 = 2*c.L - v2

		r = 255 * hue_2_rgb(v1, v2, c.H+(1/3))
		g = 255 * hue_2_rgb(v1, v2, c.H)
		b = 255 * hue_2_rgb(v1, v2, c.H-(1/3))
	}
	rgb := &RGB{r, g, b}
	return rgb
}

func main() {
	rgb := &RGB{1, 2, 255}
	hsl := rgb.HSL()
	hsv := rgb.HSV()
	fmt.Println("RGB", *rgb)
	fmt.Println("HSL", *hsl)
	fmt.Println("HSV", *hsv)
	hsvRgb := hsv.RGB()
	fmt.Println("HSV-RGB", *hsvRgb)
	hslRgb := hsl.RGB()
	fmt.Println("HSL-RGB", *hslRgb)
}
