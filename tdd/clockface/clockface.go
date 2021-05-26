package clockface

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)
const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
)
const (
	clockCenterX = 150
	clockCenterY = 150
)
const (
	colorRed   = "F00"
	colorBlack = "000"
)

func makeHand(p Point, len float64) Point {
	p = Point{p.X * len, p.Y * len}                   // scale
	p = Point{p.X, -p.Y}                              // flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY} // translate
	return p
}

func SecondsInRadians(tm time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(tm.Second())))
}

func MinutesInRadians(tm time.Time) float64 {
	return (math.Pi / (minutesInHalfClock / float64(tm.Minute()))) +
		(SecondsInRadians(tm) / minutesInClock)
}

func HoursInRadians(tm time.Time) float64 {
	return (math.Pi / (hoursInHalfClock / float64(tm.Hour()%12))) +
		(MinutesInRadians(tm) / hoursInClock)
}

func SecondHandPoint(tm time.Time) Point {
	return angleToPoint(SecondsInRadians(tm))
}

func MinuteHandPoint(tm time.Time) Point {
	return angleToPoint(MinutesInRadians(tm))
}

func HourHandPoint(tm time.Time) Point {
	return angleToPoint(HoursInRadians(tm))
}

func angleToPoint(angle float64) Point {
	x, y := math.Sin(angle), math.Cos(angle)
	return Point{x, y}
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(SecondHandPoint(t), secondHandLength)
	drawLine(w, p, colorRed)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(MinuteHandPoint(t), minuteHandLength)
	drawLine(w, p, colorBlack)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(HourHandPoint(t), hourHandLength)
	drawLine(w, p, colorBlack)
}

func drawLine(w io.Writer, p Point, color string) {
	fmt.Fprintf(w,
		`<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" style="fill:none;stroke:#%s;stroke-width:3px;"/>`,
		clockCenterX, clockCenterY, p.X, p.Y, color,
	)
}

const (
	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)
