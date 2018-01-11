# LpcToy
## Introduction
Simple Go program to mess around with LPC-encoded speech. It currently outputs a file in the format raw 8kHz little endian signed 16 bit integer.
Based on the Talkie Speech library for Arduino: https://github.com/going-digital/Talkie.
All the math is being done in floating point to "improve" sound quality.
## Possible to-do items
### More difficult
- Tempo-synchronization of voiced and unvoiced frames (voiced and unvoiced may be treated a bit differently)
- Break words into phonemes for resynthesis
- Gradual cross-fading of FIR filter coefficients between frames, for smoother sound
- Interpolation of frame amplitude data
- Multisample and interpolate chirp waveform (sinc?), for better sound (maybe it's good enough already, and aliasing doesn't matter in this case?).
- Direct audio output
### Outlandish aspirations
- Create GUI
