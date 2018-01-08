# LpcToy
## Introduction
Simple Go program to mess around with LPC-encoded speech.
Based on the Talkie Speech library for Arduino: https://github.com/going-digital/Talkie.
All the math is being done in floating point to "improve" sound quality.
## Possible to-do items
### More difficult
- Break words into phonemes for resynthesis
- Allow gradual cross-fading of filter coefficients between frames, for smoother sound
- Multisample and interpolate chirp waveform (sinc?), for better sound (maybe it's good enough already, and aliasing doesn't matter in this case?).
- Direct audio output
### Outlandish aspirations
- Create GUI
