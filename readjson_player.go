package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const Coefficients = 10
const ChirpSize = 41
const SampleRate = 8000
const voicedLength = 3

type LpcFrame struct {
	Energy float32   `json:"Energy"`
	Period float32   `json:"Period"`
	K      []float32 `json:"K"`
}

type PhoneDataFormat map[string][]LpcFrame

func main() {
	phone := loadJsonData("phones.json")

	// Some printing for debugging purposes
	/*	for phoneSymbol, phoneData := range phone {
		fmt.Println(phoneSymbol)
		fmt.Println("  Period", phoneData[0].Period)
		fmt.Println("  Energy", phoneData[0].Energy)
		for i, k := range phoneData[0].K {
			fmt.Println("  K", i, k)
		}
	} */

	// Check which phones are voiced or unvoiced
	var voicedPhone map[string]bool
	voicedPhone = make(map[string]bool)
	for phoneSymbol, phoneData := range phone {
		if phoneData[0].Period > 0 {
			voicedPhone[phoneSymbol] = true
		} else {
			voicedPhone[phoneSymbol] = false
		}
	}
	voicedPhone["L"] = false
	voicedPhone["N"] = false
	voicedPhone["NG"] = false
	voicedPhone["R"] = false
	voicedPhone["W"] = false

	f, _ := os.Create("gotestfile.raw")
	defer f.Close()

	/*for phoneSymbol, phoneData := range phone {
		fmt.Println(phoneSymbol)
		play(phoneData, f)

	}*/

	// phoneRepresentation := []string{"K", "AO","HH", "AH0", "L", "OW1", "M", "AY1", "N", "EY1", "M", "IH1", "Z", "JH", "IH1", "M", "IY0", "P", "AA1", "P", "AH0", "N", "D", "AY1", "M", "AH0", "D", "AH1", "M", "W", "AY1", "T", "G", "AY1", "AY1", "M", "N", "AA1", "T", "OW1", "L", "D", "AO1", "R", "N", "UW1", "B", "AH1", "T", "M", "IH1", "D", "AH0", "L", "S", "K", "UW1", "L", "F", "IH1", "F", "TH", "G", "R", "EY1", "D", "L", "AY1", "K", "JH", "UW1", "N", "Y", "ER0", "HH", "AY1", "AY1", "D", "OW1", "N", "T", "N", "OW1", "IH1", "F", "Y", "AO2", "L", "P", "IY1", "P", "S", "B", "IY1", "G", "IH1", "V", "P", "R", "AA1", "P", "S", "T", "UW1", "M", "AY1", "HH", "OW1"
 
	// phoneRepresentation := []string{"S", "UW", "P", "ER", "K", "AE", "L", "AH", "F", "R", "AE", "JH", "AH", "L", "IH", "S", "T", "IH", "K", "EH", "K", "S", "P", "IY", "AE", "L", "AH", "D", "OW", "SH", "AH", "S"}
	// phoneRepresentation := []string{"L", "AO", "NG", "G", "ER"}
	// phoneRepresentation := []string{"EH", "L", "EH", "S", "D", "IY"}
	// phoneRepresentation := []string{"AH","B","JH","EH","K","SH","AH","N","AH","B","AH","L"}
	// phoneRepresentation := []string{"IH", "K", "S", "T", "IH", "NG", "G", "W", "IH", "SH", "IH", "NG"}
	// phoneRepresentation := []string{"AE", "F", "G", "AE", "N", "AH", "S", "T", "AE", "N", "Z"}
	// phoneRepresentation := []string{"AE0", "F", "G", "AE1", "N", "AH0", "S", "T", "AE2", "N", "Z"}
	// phoneRepresentation := []string{"S","UW2","P","ER0","K","AE2","L","AH0","F","R","AE1","JH","AH0","L","IH2","S","T","IH0","K","EH2","K","S","P","IY0","AE2","L","AH0","D","OW1","SH","AH0","S"}
	 phoneRepresentation := []string{"HH", "AH0", "L", "OW1", "M", "AY1", "N", "EY1", "M", "IH1", "Z", "JH", "IH1", "M", "IY0", "P", "AA1", "P", "AH0", "N", "D", "AY1", "M", "AH0", "D", "AH1", "M", "W", "AY1", "T", "G", "AY1", "AY1", "M", "N", "AA1", "T", "OW1", "L", "D", "AO1", "R", "N", "UW1", "B", "AH1", "T", "M", "IH1", "D", "AH0", "L", "S", "K", "UW1", "L", "F", "IH1", "F", "TH", "G", "R", "EY1", "D", "L", "AY1", "K", "JH", "UW1", "N", "Y", "ER0", "HH", "AY1", "AY1", "D", "OW1", "N", "T", "N", "OW1", "IH1", "F", "Y", "AO2", "L", "P", "IY1", "P", "S", "B", "IY1", "G", "IH1", "V", "P", "R", "AA1", "P", "S", "T", "UW1", "M", "AY1", "HH", "OW1"}

	// phoneRepresentation := []string{"AA0", "L", "K", "AA1", "R", "AA0", "Z"}
	// phoneRepresentation := []string{"K", "AA0", "L", "AA0", "M", "AA1", "R", "IY0"}



	var outputFrame []LpcFrame

	for _, phoneSymbol := range phoneRepresentation {
		baseSymbol, stress := breakOutPhoneData(phoneSymbol)
		fmt.Println(baseSymbol, stress)
		framesInPhone := len(phone[baseSymbol])
		fmt.Println(framesInPhone)
		if voicedPhone[baseSymbol] {
			fmt.Println("voiced!")
			for i := 0; i < (voicedLength * (stress + 1)); i++ {
				outputFrame = append(outputFrame, phone[baseSymbol][i%framesInPhone])
			}
		} else {
			fmt.Println("unvoiced!")
			for i := 0; i < framesInPhone; i++ {
				outputFrame = append(outputFrame, phone[baseSymbol][i])
			}
		}
	}

	play(outputFrame, f)

	// for index, phoneSymbol := range phoneRepresentation {
	// 	fmt.Println(index, phoneSymbol)
	// 	play(phone[phoneSymbol], f)
	// }

}

func play(synth []LpcFrame, fp *os.File) {

	var periodCounter float32
	periodCounter = 0
	x := make([]float32, Coefficients+1)
	u := make([]float32, Coefficients+1)
	samplesPerFrame := 256
	var lfsr uint16 = 1 // Initialize linear feedback shift register

	for frameNumber := 0; frameNumber < len(synth); frameNumber++ {

		// Some debugging for helping with decoding phonemes
		if synth[frameNumber].Energy <= 0 {
			fmt.Print("x") // Silent. Debug
		} else if synth[frameNumber].Period <= 0 {
			fmt.Print("|") // Non-voiced. Debug
		} else {
			fmt.Print("-") // Voiced. Debug
		}

		for sampleInFrame := 0; sampleInFrame < samplesPerFrame; sampleInFrame++ {
			interpolated := makeInterpolatedFrame(frameNumber, synth, sampleInFrame, samplesPerFrame)
			// interpolated := synth[frameNumber] // for debugging only

			interpolated.Period = noteToPeriod((12 * 1)) // Test 3, fixed pitch

			// Generate source signal
			if synth[frameNumber].Period > 0 {
				// Voiced source
				u[10] = playChirp(interpolated, &periodCounter)
			} else {
				// Unvoiced source
				u[10] = playNoise(interpolated, &lfsr)
			}
			// Filter the signal
			filter(interpolated, u, x)

			// Write the sample to file
			writeInt16ToFile(int16(u[0]*128), fp)
		}
	}
	fmt.Println("") // Debug
}

// Check if a phone symbol exists in the phone database. if so, return it
func findSymbol(phone PhoneDataFormat, toFind string) (bool, []LpcFrame) {
	if val, ok := phone[toFind]; ok {
		return true, val
	} else {
		return false, nil
	}
}

func loadJsonData(filename string) PhoneDataFormat {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(strings.NewReader(string(content)))
	var phone PhoneDataFormat
	for {
		if err := dec.Decode(&phone); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return phone
}

func playChirp(synth LpcFrame, periodCounter *float32) float32 {
	// Chirp waveform table
	chirp := []float32{0, 42, -44, 50, -78, 18, 37, 20, 2, -31, -59, 2, 95, 90, 5, 15, 38, -4, -91, -91, -42, -35, -36, -4, 37, 43, 34, 33, 15, -1, -8, -18, -19, -17, -9, -10, -6, 0, 3, 2, 1}

	// The "period" value decides how many samples should be in the "wavelength" of the chirp signal.
	// For example, a value of "41" will give a looping waveform, 42 samples in length.

	// Reset the the chirp waveform if the periodCounter is equal to, or exceeds the period.
	if *periodCounter < synth.Period {
		*periodCounter++
	} else {
		*periodCounter -= synth.Period
	}

	// If periodCounter is larger than the number of samples in the chirp waveform table, pad with zeroes.
	if *periodCounter < ChirpSize {
		return ((chirp[int(*periodCounter)]) * synth.Energy) / 256
	} else {
		return 0
	}
}

func makeInterpolatedFrame(frameNumber int, synth []LpcFrame, sampleInFrame int, samplesPerFrame int) LpcFrame {
	ratio := 2 * float32(sampleInFrame) / float32(samplesPerFrame)
	var newFrame LpcFrame
	newFrame.K = make([]float32, Coefficients+1)
	var otherIndex int
	var currRatio float32
	var otherRatio float32
	if ratio < 1 {
		otherRatio = 1 - ratio
		currRatio = ratio
		otherIndex = frameNumber - 1
	} else {
		otherRatio = ratio - 1
		currRatio = 2 - ratio
		otherIndex = frameNumber + 1
	}
	// If first entry, don't interpolate with previous frame
	if otherIndex < 0 {
		otherIndex++
	}
	// If last entry, don't interpolate with next frame
	if otherIndex >= len(synth) {
		otherIndex--
	}
	// Don't interpolate between unvoiced frames
	if synth[otherIndex].Period == 0 {
		otherIndex = frameNumber
	}
	if synth[frameNumber].Period == 0 {
		otherIndex = frameNumber
	}

	newFrame.Energy = currRatio*synth[frameNumber].Energy + otherRatio*synth[otherIndex].Energy
	newFrame.Period = currRatio*synth[frameNumber].Period + otherRatio*synth[otherIndex].Period
	for coefficient := 0; coefficient < 10; coefficient++ {
		newFrame.K[coefficient] = currRatio*synth[frameNumber].K[coefficient] + otherRatio*synth[otherIndex].K[coefficient]
	}

	return newFrame
}

func playNoise(synth LpcFrame, lfsr *uint16) float32 {
	var output float32
	// Normal Galois configuration linear feedback shift register
	var lsb uint16 = *lfsr & 1 // Get LSB (i.e., the output bit)
	*lfsr >>= 1                // Shift register
	if lsb > 0 {               // If the output bit is 1, apply toggle mask
		*lfsr ^= 0xB800
	}
	if *lfsr&1 > 0 {
		output = synth.Energy
	} else {
		output = -synth.Energy
	}
	output *= float32(*lfsr & 1)
	return output
}

// This is needed to save as a 16 bit raw file, since golang only writes bytes
func writeInt16ToFile(input int16, fp *os.File) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.LittleEndian, input)
	if err != nil {
		fmt.Println(err)
	}
	intByteArray := buff.Bytes()
	fp.Write(intByteArray)
}

func filter(synth LpcFrame, u []float32, x []float32) {
	// Lattice filter forward path
	u[9] = u[10] - ((synth.K[9] * x[9]) / 128)
	u[8] = u[9] - ((synth.K[8] * x[8]) / 128)
	u[7] = u[8] - ((synth.K[7] * x[7]) / 128)
	u[6] = u[7] - ((synth.K[6] * x[6]) / 128)
	u[5] = u[6] - ((synth.K[5] * x[5]) / 128)
	u[4] = u[5] - ((synth.K[4] * x[4]) / 128)
	u[3] = u[4] - ((synth.K[3] * x[3]) / 128)
	u[2] = u[3] - ((synth.K[2] * x[2]) / 128)
	u[1] = u[2] - ((synth.K[1] * x[1]) / 32768)
	u[0] = u[1] - ((synth.K[0] * x[0]) / 32768)

	// Output clamp
	if u[0] > 32767 {
		u[0] = 32767
	}
	if u[0] < -32768 {
		u[0] = -32768
	}

	// Lattice filter reverse path
	x[9] = x[8] + ((synth.K[8] * u[8]) / 128)
	x[8] = x[7] + ((synth.K[7] * u[7]) / 128)
	x[7] = x[6] + ((synth.K[6] * u[6]) / 128)
	x[6] = x[5] + ((synth.K[5] * u[5]) / 128)
	x[5] = x[4] + ((synth.K[4] * u[4]) / 128)
	x[4] = x[3] + ((synth.K[3] * u[3]) / 128)
	x[3] = x[2] + ((synth.K[2] * u[2]) / 128)
	x[2] = x[1] + ((synth.K[1] * u[1]) / 32768)
	x[1] = x[0] + ((synth.K[0] * u[0]) / 32768)

	x[0] = u[0]
}

func noteToPeriod(note float32) float32 {
	// Convert a note number to the period numbers that playChirp() expects
	return (SampleRate / float32(55*math.Pow(2, (float64(note)/12)))) - 1
}

func breakOutPhoneData(inputString string) (string, int) {
	phonePattern := regexp.MustCompile(`([A-Z]+)(\d*)`)
	phoneBreakout := phonePattern.FindStringSubmatch(inputString)
	phoneSymbol := phoneBreakout[1]
	stress, _ := strconv.Atoi(phoneBreakout[2])
	return phoneSymbol, stress
}
