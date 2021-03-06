// Copyright 2018 The ZikiChomgo Authors. All rights reserved.  Use of this source
// code is governed by a license that can be found in the License file.

package dct

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"

	"zikichombo.org/codec/wav"
	"zikichombo.org/sound/ops"
)

func TestPlot(t *testing.T) {
	//n := gen.Sin(820 * freq.Hertz)
	var n audio.Source
	n, _ = wav.Load("/Users/scott/Dev/dawg/snd/Happy14/0.wav")
	//n = ops.AfterDur(n, 20*time.Second)
	Ns := []int{128, 256, 512, 1024, 2048}
	b := image.Rect(0, 0, 768, 384)
	for _, N := range Ns {
		d := make([]float64, N)
		ops.Slurp(n, d)
		dct := New(N)
		dct.Do(d)
		p := 0.0
		for i, v := range d {
			d[i] = v //math.Abs(v)
			p += v * v
		}
		fmt.Printf("power %f\n", p)
		img := Plot(b, d)
		w, e := os.Create(fmt.Sprintf("dct-a2-%d.png", N))
		if e != nil {
			t.Fatal(e)
		}
		if e := png.Encode(w, img); e != nil {
			t.Fatal(e)
		}
		w.Close()
		dct.Inv(d)
		img = Plot(b, d)
		w, e = os.Create(fmt.Sprintf("dctfx-a2-%d.png", N))
		if e != nil {
			t.Fatal(e)
		}
		if e := png.Encode(w, img); e != nil {
			t.Fatal(e)
		}
	}
}
