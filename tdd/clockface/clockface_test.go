package clockface_test

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"

	. "example.com/clockface"
)

// func TestSecondHandAtMidnight(t *testing.T) {
// 	tm := simpleTime(0, 0, 0)

// 	want := Point{X: 150, Y: 150 - 90}
// 	got := SecondHand(tm)

// 	if got != want {
// 		t.Errorf(`got %v, want %v`, got, want)
// 	}
// }

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := simpleTime(0, 0, 30)

// 	want := Point{X: 150, Y: 150 + 90}
// 	got := SecondHand(tm)

// 	if got != want {
// 		t.Errorf(`got %v, want %v`, got, want)
// 	}
// }

func TestSecondsInRadians(t *testing.T) {
	testCases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 30), 30 * (math.Pi / 30)},
		{simpleTime(0, 0, 45), 45 * (math.Pi / 30)},
		{simpleTime(0, 0, 7), 7 * (math.Pi / 30)},
	}

	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			got := SecondsInRadians(tC.time)
			want := tC.angle

			if !roughlyEqualFloat64(got, want) {
				t.Fatalf(`got %v, want %v rad`, got, want)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	testCases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			got := MinutesInRadians(tC.time)
			want := tC.angle

			if !roughlyEqualFloat64(got, want) {
				t.Fatalf(`got %v, want %v rad`, got, want)
			}
		})
	}
}
func TestHoursInRadians(t *testing.T) {
	testCases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			got := HoursInRadians(tC.time)
			want := tC.angle

			if !roughlyEqualFloat64(got, want) {
				t.Fatalf(`got %v, want %v rad`, got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			got := SecondHandPoint(tC.time)
			want := tC.point

			if !roughlyEqualPoint(got, want) {
				t.Errorf(`got %v, want %v`, got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			got := MinuteHandPoint(tC.time)
			want := tC.point

			if !roughlyEqualPoint(got, want) {
				t.Errorf(`got %v, want %v`, got, want)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			got := HourHandPoint(tC.time)
			want := tC.point

			if !roughlyEqualPoint(got, want) {
				t.Errorf(`got %v, want %v`, got, want)
			}
		})
	}
}

// func TestSVGWriterAtMidnight(t *testing.T) {
// 	tm := simpleTime(0, 0, 0)

// 	b := bytes.Buffer{}
// 	SVGWriter(&b, tm)

// 	svg := SVG{}
// 	xml.Unmarshal(b.Bytes(), &svg)

// 	want := Line{150, 150, 150, 60}

// 	for _, line := range svg.Line {
// 		if line == want {
// 			return
// 		}
// 	}
// 	t.Errorf(`want line %v, in the svg %v`, want, svg.Line)
// }

func TestSVGWriterSecondHand(t *testing.T) {
	testCases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}
	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, tC.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tC.line, svg.Line) {
				t.Errorf(`want line %v, in the svg %v`, tC.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	testCases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
	}
	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, tC.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tC.line, svg.Line) {
				t.Errorf(`want line %+v, in the svg %+v`, tC.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	testCases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
	}
	for _, tC := range testCases {
		t.Run(testName(tC.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, tC.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tC.line, svg.Line) {
				t.Errorf(`want line %+v, in the svg %+v`, tC.line, svg.Line)
			}
		})
	}
}

func containsLine(want Line, got []Line) bool {
	for _, line := range got {
		if line == want {
			return true
		}
	}
	return false
}

func testName(tm time.Time) string {
	return tm.Format("15:04:05")
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(1337, time.January, 1, h, m, s, 0, time.UTC)
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e7
	return math.Abs(a-b) < equalityThreshold
}
