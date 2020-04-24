package main

// Copyright 2019 The Oto Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"flag"
	"fmt"
	"io"
	"math"
	"sync"
	"time"

	"github.com/hajimehoshi/oto"
)

var (
	sampleRate     = flag.Int("samplerate", 44100, "sample rate")
	channelNum     = flag.Int("channelnum", 2, "number of channel")
	herpDerp       = flag.Int("herpderp", 2, "herp derp?")
	someStupidText = flag.String("doh", "", "doh!")
)

// SineWave contains properties for an instance of a sinewave
type SineWave struct {
	freq   float64
	length int64
	pos    int64

	remaining []byte
}

const bitDepthInBytes = 2
const numberOfVoices = 3

func newSineWave(freq float64, duration time.Duration) *SineWave {
	l := int64(*channelNum) * bitDepthInBytes * int64(*sampleRate) * int64(duration) / int64(time.Second)
	l = l / 4 * 4
	return &SineWave{
		freq:   freq,
		length: l,
	}
}

func (s *SineWave) Read(buf []byte) (int, error) {
	if len(s.remaining) > 0 {
		n := copy(buf, s.remaining)
		s.remaining = s.remaining[n:]
		return n, nil
	}

	if s.pos == s.length {
		return 0, io.EOF
	}

	eof := false
	if s.pos+int64(len(buf)) > s.length {
		buf = buf[:s.length-s.pos]
		eof = true
	}

	var origBuf []byte
	if len(buf)%4 > 0 {
		origBuf = buf
		buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	}

	length := float64(*sampleRate) / float64(s.freq)

	num := (bitDepthInBytes) * (*channelNum)
	p := s.pos / int64(num)

	for i := 0; i < len(buf)/num; i++ {
		const max = 32767
		b := int16(math.Sin(2*math.Pi*float64(p)/length) * max / numberOfVoices)
		for ch := 0; ch < *channelNum; ch++ {
			buf[num*i+2*ch] = byte(b)
			buf[num*i+1+2*ch] = byte(b >> 8)
		}
		p++
	}

	s.pos += int64(len(buf))

	n := len(buf)
	if origBuf != nil {
		n = copy(origBuf, buf)
		s.remaining = buf[n:]
	}

	if eof {
		return n, io.EOF
	}
	return n, nil
}

func play(context *oto.Context, freq float64, duration time.Duration) error {
	p := context.NewPlayer()
	s := newSineWave(freq, duration)
	if _, err := io.Copy(p, s); err != nil {
		return err
	}
	if err := p.Close(); err != nil {
		return err
	}
	return nil
}

func run() error {

	freqToPlay := []float64{
		523.3, // C
		659.3, // E
		784.0, // G
	}

	c, err := oto.NewContext(*sampleRate, *channelNum, bitDepthInBytes, 4096)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for index, freq := range freqToPlay {
		fmt.Println(index, freq)
		wg.Add(1)
		go startNote(&wg, c, index, freq)
	}

	wg.Wait()
	c.Close()
	return nil
}

func startNote(wg *sync.WaitGroup, c *oto.Context, index int, freq float64) {
	{

		defer wg.Done()
		time.Sleep(time.Duration(index) * time.Second)
		if err := play(c, freq, 3*time.Second); err != nil {
			panic(err)
		}
	}
}

func main() {
	flag.Parse()
	fmt.Println(*someStupidText, *herpDerp, *channelNum)

	if err := run(); err != nil {
		panic(err)
	}
}
