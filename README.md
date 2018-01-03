# LpcToy
## Introduction
Simple Go program to mess around with LPC-encoded speech.
Based on the Talkie Speech library for Arduino: https://github.com/going-digital/Talkie.
All the math is being done in floating point to "improve" sound quality.
## Possible to-do items
### Low-hanging fruit
- Translate oscillator frequency to Hz
### More difficult
- Break words into phonemes for resynthesis
- Allow gradual cross-fading of filter coefficients between frames, for smoother sound
- Multisample and interpolate chirp waveform, for better sound
- Direct audio output
### Outlandish aspirations
- Create GUI