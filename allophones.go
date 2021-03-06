package main

// A (somewhat incompletely described) list of allophones from the TI-99/4A
// speech synthesizer ROMs.
// As specified in "Compute!'s guide to TI-99/4A sound and graphics" p162, p163
func allophones() {
	allophone001Addition := []uint8{0xa6, 0xf6, 0x35, 0xc3, 0xac, 0xaa, 0xb4, 0xda, 0xd7, 0xc, 0xb3, 0xaa, 0xd2, 0x6a, 0x5f, 0x33, 0xcc, 0xaa, 0xca, 0x3}
	allophone002Annuity := []uint8{0xa6, 0xce, 0xc3, 0x93, 0x25, 0xe3, 0xb4, 0x3a, 0xf, 0x4f, 0x96, 0x8c, 0xd3, 0xea, 0x3c, 0x3c, 0x59, 0x32, 0xce, 0x3}
	allophone003Delta := []uint8{0xa6, 0x76, 0xbe, 0xa6, 0xc8, 0x9a, 0xb6, 0xda, 0xf9, 0x9a, 0x22, 0x6b, 0xda, 0x6a, 0xe7, 0x6b, 0x8a, 0xac, 0xe9, 0x3}
	allophone004OnTime := []uint8{0xa6, 0xe, 0x66, 0xdb, 0x59, 0x9b, 0xb4, 0x3a, 0x98, 0x6d, 0x67, 0x6d, 0xd2, 0xea, 0x60, 0xb6, 0x9d, 0xb5, 0xc9, 0x3}
	allophone005Autonomy := []uint8{0xaa, 0xf6, 0x2a, 0x3f, 0x58, 0x5b, 0x97, 0xda, 0xab, 0xfc, 0x60, 0x6d, 0x5d, 0x6a, 0xed, 0x6f, 0x1c, 0xb3, 0xce, 0x3}
	allophone006Anonymity := []uint8{0xaa, 0x76, 0xea, 0x5e, 0x29, 0x53, 0x97, 0xda, 0xa9, 0x7b, 0xa5, 0x4c, 0x5d, 0x6a, 0x1d, 0x7f, 0x58, 0x3b, 0xc3, 0x3}
	allophone007Eliminate := []uint8{0xaa, 0x1e, 0x59, 0x9, 0x7c, 0x57, 0xa6, 0x7a, 0x64, 0x25, 0xf0, 0x5d, 0x99, 0xea, 0x91, 0x95, 0xc0, 0x77, 0xe5, 0x3}
	allophone008Enough := []uint8{0xa6, 0xee, 0x2d, 0x98, 0x29, 0x66, 0xb7, 0xba, 0xb7, 0x60, 0xa6, 0x98, 0xdd, 0xea, 0xde, 0x82, 0x99, 0x62, 0xf6, 0x3}
	allophone009Context := []uint8{0xa6, 0x4e, 0xd5, 0x4d, 0x32, 0xab, 0xb6, 0x3a, 0x55, 0x37, 0xc9, 0xac, 0xda, 0xea, 0x54, 0xdd, 0x24, 0xb3, 0xea, 0x3}
	allophone010Ancient := []uint8{0xa6, 0x4e, 0xc5, 0x43, 0x34, 0x12, 0x8f, 0x3a, 0x15, 0xf, 0xd1, 0x48, 0x3c, 0xea, 0x54, 0x3c, 0x44, 0x23, 0xf1, 0x3}
	allophone011Western := []uint8{0xa6, 0x8e, 0x65, 0x98, 0x5b, 0xb5, 0xb7, 0x3a, 0x96, 0x61, 0x6e, 0xd5, 0xde, 0xea, 0x58, 0x86, 0xb9, 0x55, 0xfb, 0x3}
	allophone012Synthesis := []uint8{0xaa, 0x8e, 0x45, 0x5c, 0x7c, 0x96, 0x94, 0x3a, 0x16, 0x71, 0xf1, 0x59, 0x52, 0xea, 0x58, 0xc4, 0xc5, 0x67, 0xc9, 0x3}
	allophone013Mane := []uint8{0xa6, 0x9e, 0x41, 0x15, 0xba, 0x5a, 0x8f, 0x7a, 0x6, 0x55, 0xe8, 0x6a, 0x3d, 0xea, 0x19, 0x54, 0xa1, 0xab, 0xf5, 0x3}
	allophone014TookOn := []uint8{0xa6, 0xf6, 0xc6, 0x33, 0x92, 0xf2, 0xb6, 0xda, 0x1b, 0xcf, 0x48, 0xca, 0xdb, 0xea, 0x60, 0x2d, 0x3c, 0xa9, 0xdf, 0x3}
	allophone015Donation := []uint8{0xaa, 0xf6, 0x22, 0xba, 0x1e, 0xe0, 0x96, 0xda, 0x8b, 0xe8, 0x7a, 0x80, 0x5b, 0x6a, 0x2f, 0xa2, 0xeb, 0x1, 0xee, 0x3}
	allophone016Annual := []uint8{0xaa, 0x96, 0xa5, 0x9c, 0xcc, 0x37, 0xa6, 0x5a, 0x96, 0x72, 0x32, 0xdf, 0x98, 0x6a, 0x59, 0xca, 0xc9, 0x7c, 0xe3, 0x3}
	allophone017Unique := []uint8{0xa6, 0xb6, 0x46, 0xab, 0xc, 0xc3, 0xb6, 0xda, 0x1a, 0xad, 0x32, 0xc, 0xdb, 0x6a, 0x6b, 0xb4, 0xca, 0x30, 0xec, 0x3}
	allophone018Above := []uint8{0xa6, 0x8e, 0xc6, 0x37, 0x9c, 0xbb, 0x8e, 0x3a, 0x1a, 0xdf, 0x70, 0xee, 0xda, 0xea, 0x68, 0x7c, 0xc3, 0xb9, 0xeb, 0x3}
	allophone019Instrument := []uint8{0xa6, 0x4e, 0x3e, 0x4b, 0xc9, 0xe7, 0xb4, 0x3a, 0xf9, 0x2c, 0x25, 0x9f, 0x33, 0xea, 0xe0, 0x2c, 0x95, 0x33, 0xe9, 0x3}
	allophone020Underneath := []uint8{0xa6, 0x4e, 0x3e, 0x4b, 0xc9, 0xe7, 0x8c, 0x3a, 0xf9, 0x2c, 0x25, 0x9f, 0x33, 0xea, 0xe4, 0xb3, 0x94, 0x7c, 0xce, 0x3}
	allophone021Roses := []uint8{0xa6, 0xee, 0x21, 0x54, 0xc, 0xf2, 0xb5, 0xba, 0x87, 0x50, 0x31, 0xc8, 0xd7, 0xea, 0x1e, 0x42, 0xc5, 0x20, 0xdf, 0x3}
	allophone022Basement := []uint8{0xa6, 0xf6, 0xae, 0x82, 0xc3, 0x53, 0x8e, 0xda, 0xbb, 0xa, 0xe, 0x4f, 0x39, 0x6a, 0xef, 0x2a, 0x38, 0x3c, 0xe5, 0x3}
	allophone023Seeker := []uint8{0xa6, 0x76, 0x43, 0xb9, 0xca, 0xbc, 0xb7, 0xda, 0xd, 0xe5, 0x2a, 0xf3, 0xde, 0xeb, 0x54, 0xbb, 0xa1, 0x5c, 0x65, 0xde, 0x1f}
	allophone024Ratio := []uint8{0xa6, 0xf6, 0xb2, 0x3a, 0x4d, 0xbd, 0xb5, 0xda, 0xcb, 0xea, 0x34, 0xf5, 0xd6, 0x6a, 0xab, 0xba, 0xca, 0xd9, 0x6f, 0xa9, 0x8d, 0xa9, 0x6a, 0x47, 0xef, 0xa9, 0x66, 0xb5, 0x36, 0x44, 0x32, 0x3d}
	allophone025Funny := []uint8{0xaa, 0xee, 0x55, 0x99, 0xad, 0x67, 0x97, 0xba, 0x57, 0x65, 0xb6, 0x9e, 0x5d, 0xea, 0x5e, 0x95, 0xd9, 0x7a, 0x76, 0xaa, 0x47, 0x56, 0x2, 0xdf, 0x95, 0xf}
	allophone026Hat := []uint8{0xae, 0x8e, 0x2d, 0xca, 0xb8, 0x9a, 0xac, 0x3a, 0xb6, 0x28, 0xe3, 0x6a, 0x72, 0xeb, 0x5b, 0xaf, 0x3a, 0xb6, 0x28, 0xe3, 0x6a, 0xf2}
	allophone027Hot := []uint8{0xae, 0xf6, 0xb9, 0x47, 0xc5, 0x9b, 0xac, 0xda, 0xe7, 0x1e, 0x15, 0x6f, 0xb2, 0xeb, 0x5d, 0xaf, 0xda, 0xe7, 0x1e, 0x15, 0x6f, 0xf2}
	allophone028Height := []uint8{0xae, 0xb6, 0xb9, 0xc7, 0xb8, 0x13, 0xaf, 0xda, 0xe6, 0x1e, 0xe3, 0x4e, 0x7c, 0xea, 0x58, 0x23, 0x83, 0xab, 0xc9, 0xaa, 0x73, 0x77, 0x35, 0xad, 0xba, 0xab, 0x2e, 0x5d, 0xd4, 0xac, 0xaa, 0x3e}
	allophone029Cart := []uint8{0xa6, 0xb6, 0xbd, 0x22, 0x8d, 0x9b, 0x8e, 0xda, 0xf6, 0x8a, 0x34, 0x6e, 0x3a, 0xeb, 0x51, 0xbb, 0x6d, 0x51, 0xca, 0xb5, 0x5b, 0xed, 0xf, 0xad, 0x49, 0x47, 0x6e, 0xb5, 0x5f, 0xd8, 0x29, 0x99, 0xe0, 0x1}
	allophone030House := []uint8{0xa6, 0xe, 0xbe, 0xdb, 0xd5, 0xeb, 0x8e, 0x3a, 0xf8, 0x6e, 0x57, 0xaf, 0xbb, 0x6a, 0x17, 0x66, 0x54, 0xa2, 0xde, 0xaa, 0x4d, 0xd8, 0x11, 0x89, 0xba, 0xa3, 0xb6, 0x66, 0x9e, 0xd5, 0xe6, 0x3c}
	allophone031Sought := []uint8{0xa6, 0x76, 0xe6, 0xd7, 0xc8, 0xe7, 0x8d, 0xda, 0x99, 0x5f, 0x23, 0x9f, 0x37, 0xeb, 0x5d, 0x8f, 0xda, 0x99, 0x5f, 0x23, 0x9f, 0xf7}
	allophone032Heat := []uint8{0xa6, 0x1e, 0x59, 0x9, 0x7c, 0x57, 0x8e, 0x7a, 0x64, 0x25, 0xf0, 0x5d, 0x39, 0xeb, 0x59, 0x8f, 0x7a, 0x36, 0x65, 0xb0, 0x5e, 0xf5}
	allophone033Pierce := []uint8{0xa6, 0xae, 0x39, 0x4, 0xd2, 0xbb, 0x8f, 0xba, 0xe6, 0x10, 0x48, 0xef, 0xbe, 0xea, 0xd4, 0x3a, 0x20, 0x45, 0x7d, 0xab, 0xfd, 0x34, 0x91, 0x91, 0xfc, 0xad, 0x76, 0x4d, 0x55, 0x57, 0xb3, 0x3d}
	allophone034Set := []uint8{0xae, 0x8e, 0xd9, 0x5d, 0xa3, 0x26, 0xaf, 0x3a, 0x66, 0x77, 0x8d, 0x9a, 0xbc, 0xeb, 0x55, 0xc7, 0x6c, 0x29, 0x91, 0xa3, 0x47, 0xed, 0xaa, 0x86, 0x65, 0x97, 0x7e}
	allophone035Therapy := []uint8{0xae, 0x2e, 0xd5, 0x4d, 0x79, 0xda, 0xac, 0xba, 0x54, 0x37, 0xe5, 0x69, 0xb3, 0xea, 0xb8, 0xd5, 0xcd, 0xbc, 0xcd, 0xa8, 0xc3, 0x95, 0x70, 0xb3, 0x36, 0xad, 0xf6, 0x87, 0x32, 0xcd, 0x9b, 0x3c}
	allophone036Take := []uint8{0xae, 0x6e, 0x29, 0x99, 0x2b, 0xd6, 0xac, 0xba, 0xa5, 0x64, 0xae, 0x58, 0xb3, 0xea, 0x9a, 0x8c, 0xb5, 0x7b, 0xcd, 0xaa, 0x6b, 0x15, 0xd2, 0xc9, 0x26, 0xab, 0xae, 0x45, 0x88, 0xbb, 0x5a, 0x3f}
	allophone037Hurt := []uint8{0xa6, 0x76, 0xc3, 0xc9, 0x47, 0xfc, 0x8f, 0xda, 0xd, 0x27, 0x1f, 0xf1, 0x3f, 0xeb, 0x59, 0x8f, 0xda, 0xd, 0x27, 0x1f, 0xf1, 0xff}
	allophone038Issue := []uint8{0xae, 0x8e, 0x45, 0x5c, 0x7c, 0x96, 0xac, 0x3a, 0x16, 0x71, 0xf1, 0x59, 0xb2, 0xeb, 0x5d, 0xaf, 0x3a, 0x16, 0x71, 0xf1, 0x59, 0xf2}
	allophone039Choice := []uint8{0xae, 0x76, 0xb6, 0x32, 0x9d, 0xb3, 0xae, 0xda, 0xd9, 0xca, 0x74, 0xce, 0x7a, 0x6a, 0x9f, 0xc2, 0x42, 0x6d, 0xf5, 0xa9, 0x63, 0x48, 0xd1, 0xf4, 0x35, 0xa7, 0xe, 0x2e, 0xc4, 0x23, 0xab, 0x3c}
	allophone040Cook := []uint8{0xae, 0xf6, 0xc6, 0x33, 0x92, 0xf2, 0xae, 0xda, 0x1b, 0xcf, 0x48, 0xca, 0xbb, 0xeb, 0x5b, 0xaf, 0x3a, 0x3a, 0x4f, 0xb, 0xdc, 0xf6}
	allophone041Poorly := []uint8{0xaa, 0x16, 0xa5, 0x96, 0xd4, 0xbb, 0x94, 0x5a, 0x94, 0x5a, 0x52, 0xef, 0xd2, 0x6a, 0x51, 0xb3, 0x4c, 0x35, 0xdb, 0xa8, 0x65, 0xf3, 0x4c, 0x91, 0xb6, 0xad, 0x96, 0x53, 0xb2, 0x45, 0xb3, 0x3c}
	allophone042Horse := []uint8{0xaa, 0x56, 0x3d, 0x5b, 0xd4, 0xba, 0x96, 0x5a, 0xf5, 0x6c, 0x51, 0xeb, 0x3a, 0x6a, 0x3d, 0x3d, 0x53, 0xa4, 0x6d, 0xab, 0xcd, 0x91, 0x4a, 0xd1, 0xcc, 0xa5, 0xb6, 0x9b, 0x2a, 0x9d, 0x5b, 0x3c}
	allophone043Boat := []uint8{0xae, 0xf6, 0xb2, 0x3a, 0x4d, 0xbd, 0xad, 0xda, 0xcb, 0xea, 0x34, 0xf5, 0xb6, 0x6a, 0xab, 0xba, 0xca, 0xd9, 0xef, 0xaa, 0x65, 0xac, 0x6a, 0x47, 0xef, 0xa5, 0x66, 0xb5, 0x36, 0x44, 0x32, 0x3d}
	allophone044Shoot := []uint8{0xa6, 0xe, 0xc6, 0x95, 0x4a, 0xf6, 0x8e, 0x3a, 0x18, 0x57, 0x2a, 0xd9, 0xbb, 0x6a, 0x67, 0x3d, 0xb9, 0xb4, 0xeb, 0xaa, 0x8d, 0xf3, 0xe4, 0xd2, 0xae, 0xad, 0xe6, 0xa1, 0xdc, 0x42, 0x17, 0x3d}
	allophone045Hut := []uint8{0xae, 0x8e, 0xc6, 0x37, 0x9c, 0xbb, 0xae, 0x3a, 0x1a, 0xdf, 0x70, 0xee, 0xba, 0xeb, 0x5b, 0xaf, 0x3a, 0x1a, 0xdf, 0x70, 0xee, 0xfa}
	allophone046Boot := []uint8{0xa6, 0x36, 0x22, 0x46, 0x9d, 0xbb, 0x8c, 0xda, 0x88, 0x18, 0x75, 0xee, 0xb2, 0x6a, 0x23, 0x62, 0xcc, 0xc9, 0xdf, 0xa8, 0xd, 0x8f, 0x53, 0x53, 0x7f, 0xa3, 0x36, 0x3c, 0x4e, 0x4d, 0xfd, 0x3d}
	allophone047Had := []uint8{0xae, 0x8e, 0x2d, 0xca, 0xb8, 0x9a, 0xac, 0x3a, 0xb6, 0x28, 0xe3, 0x6a, 0xb2, 0xeb, 0x5b, 0xdf, 0xfa, 0xd4, 0xa1, 0x45, 0xaa, 0x56, 0x9c, 0x55, 0xc7, 0xea, 0x65, 0x5a, 0x71, 0x1e}
	allophone048Odd := []uint8{0xae, 0xe, 0x2e, 0x5f, 0x42, 0xe7, 0xac, 0x3a, 0xb8, 0x7c, 0x9, 0x9d, 0xb3, 0xeb, 0x5d, 0xef, 0x7a, 0xd7, 0xab, 0xe, 0x2e, 0x5f, 0x42, 0xe7, 0x3c}
	allophone049Hide := []uint8{0xae, 0xb6, 0xb9, 0xc7, 0xb8, 0x13, 0xaf, 0xda, 0xe6, 0x1e, 0xe3, 0x4e, 0x7c, 0xea, 0xd8, 0xbd, 0xc4, 0xdb, 0xf1, 0xa9, 0x53, 0xf7, 0xe2, 0x88, 0x3a, 0xa7, 0xce, 0xdd, 0x92, 0x32, 0x9b, 0xac, 0x3a, 0x35, 0x9, 0xc9, 0xaa, 0xb3, 0xea, 0x90, 0x54, 0x7c, 0x63, 0xf2, 0x3}
	allophone050Card := []uint8{0xae, 0xb6, 0xb1, 0xd7, 0x31, 0x9a, 0xae, 0xda, 0xc6, 0x5e, 0xc7, 0x68, 0xba, 0x6a, 0x17, 0xeb, 0x9c, 0x22, 0xdd, 0xa9, 0xbd, 0xcb, 0x4d, 0xb6, 0xe6, 0xa7, 0xf6, 0x21, 0x3b, 0x55, 0x9a, 0x9f, 0xda, 0x57, 0x8f, 0xc, 0x4e, 0x3e, 0x6a, 0xdf, 0xdd, 0x32, 0xb8, 0xe9, 0x3}
	allophone051Loud := []uint8{0xae, 0x76, 0xad, 0xca, 0xd5, 0x6a, 0xaf, 0xda, 0xb5, 0x2a, 0x57, 0xab, 0x7d, 0x6a, 0x57, 0xaa, 0x42, 0xad, 0xf6, 0xa9, 0x6d, 0x9a, 0x36, 0xd3, 0xc6, 0xab, 0x36, 0x69, 0xc6, 0x4c, 0x9a, 0x8c, 0x5a, 0xf5, 0x29, 0x51, 0x6d, 0x5d, 0x6a, 0x71, 0xda, 0xa8, 0xcd, 0xc9, 0x3}
	allophone052Saw := []uint8{0xa6, 0x76, 0xe6, 0xd7, 0xc8, 0xe7, 0x8d, 0xda, 0x99, 0x5f, 0x23, 0x9f, 0x37, 0xeb, 0x59, 0xef, 0x7a, 0xd7, 0xa3, 0x76, 0xe6, 0xd7, 0xc8, 0xe7, 0x3d}
	allophone053Seed := []uint8{0xa6, 0xee, 0x49, 0x19, 0xaa, 0xe7, 0x8c, 0xba, 0x27, 0x65, 0xa8, 0x9e, 0x33, 0xeb, 0x59, 0xef, 0x7a, 0xd6, 0xa3, 0xee, 0x49, 0x19, 0xaa, 0xe7, 0x3c}
	allophone054Heel := []uint8{0xa6, 0xee, 0x49, 0x19, 0xaa, 0xc7, 0x8c, 0xba, 0x27, 0x65, 0xa8, 0x1e, 0x33, 0xeb, 0x55, 0x97, 0x6c, 0x82, 0x51, 0x53, 0x56, 0x9d, 0x8a, 0x1b, 0x59, 0xce, 0x59, 0x75, 0xa8, 0x11, 0x2c, 0xbe, 0xa4, 0xd5, 0xba, 0xf5, 0x8, 0x5b, 0xdb, 0x54, 0xf3, 0x3e, 0xcb, 0xe2, 0x69, 0x5d, 0x4d, 0xef, 0x86, 0xb1, 0x39, 0x4e, 0x35, 0xcb, 0x1f, 0x29, 0x6c, 0xe7, 0x1}
	allophone055Hear := []uint8{0xae, 0xae, 0x39, 0x4, 0xd2, 0xbb, 0xaf, 0xba, 0xe6, 0x10, 0x48, 0xef, 0xbe, 0xeb, 0x55, 0xa7, 0xea, 0x42, 0xa5, 0x5d, 0x57, 0x1d, 0x87, 0x39, 0x96, 0x74, 0x1f, 0xb5, 0x9f, 0x26, 0x32, 0x92, 0x7f, 0xd4, 0xae, 0xa9, 0xea, 0x6a, 0xb6, 0x7}
	allophone056Said := []uint8{0xae, 0x4e, 0x3d, 0x5c, 0xa5, 0xda, 0xac, 0x3a, 0xf5, 0x70, 0x95, 0x6a, 0xb3, 0xea, 0xd4, 0xdd, 0x45, 0xbb, 0xcd, 0xa9, 0x53, 0xf7, 0x10, 0xa9, 0xd6, 0xa7, 0x4e, 0xcd, 0x82, 0x2d, 0x9b, 0xae, 0x3a, 0x76, 0xb, 0xf6, 0x9c, 0xb7, 0xea, 0xd8, 0x2d, 0xd8, 0x73, 0xde, 0x3}
	allophone057There := []uint8{0xae, 0x2e, 0xd5, 0x5c, 0x74, 0xda, 0xac, 0xba, 0x54, 0x73, 0xd1, 0x69, 0x73, 0xea, 0xdc, 0xd4, 0xd4, 0xba, 0xcd, 0xaa, 0xd3, 0x50, 0x35, 0xeb, 0x2e, 0xab, 0x8e, 0x5b, 0xdd, 0xcc, 0xdb, 0x8c, 0x3a, 0x5c, 0x9, 0x37, 0x6b, 0xd3, 0x6a, 0x7f, 0x28, 0xd3, 0xbc, 0xc9, 0x3}
	allophone058Day := []uint8{0xae, 0x6e, 0x29, 0x99, 0x2b, 0xd6, 0xac, 0xba, 0xa5, 0x64, 0xae, 0x58, 0x73, 0xea, 0x9a, 0x8c, 0xb5, 0x7b, 0xcd, 0xad, 0x57, 0x5d, 0xab, 0x90, 0x4e, 0x36, 0x19, 0x75, 0x2d, 0x42, 0xdc, 0xd5, 0x7a, 0xd4, 0xb5, 0x8, 0x71, 0x57, 0xeb, 0x7}
	allophone059Heard := []uint8{0xa6, 0x76, 0xc3, 0xc9, 0x47, 0xfc, 0xb7, 0xda, 0xd, 0x27, 0x1f, 0xf1, 0x3f, 0xeb, 0x59, 0xcf, 0x7a, 0xd6, 0xa3, 0x76, 0xc3, 0xc9, 0x47, 0xfc, 0x3f}
	allophone060Hid := []uint8{0xae, 0x8e, 0x45, 0x5c, 0x7c, 0x96, 0xac, 0x3a, 0x16, 0x71, 0xf1, 0x59, 0xb2, 0xeb, 0x5b, 0xef, 0x7a, 0xd7, 0xab, 0x8e, 0x45, 0x5c, 0x7c, 0x96, 0x3c}
	allophone061Hill := []uint8{0xae, 0x8e, 0x45, 0x5c, 0x7c, 0x96, 0xac, 0x3a, 0x16, 0x71, 0xf1, 0x59, 0xb2, 0xea, 0x58, 0xd4, 0xc5, 0x6a, 0xc9, 0xaa, 0x7d, 0xf1, 0x10, 0x8d, 0x25, 0xab, 0x76, 0x35, 0x93, 0xc5, 0xda, 0xb4, 0x5a, 0xb7, 0x1a, 0x11, 0x6b, 0x9b, 0x6a, 0xde, 0x67, 0x59, 0x3c, 0xad, 0xab, 0xd9, 0xbd, 0x62, 0x9, 0xa7, 0xa1, 0xa6, 0x77, 0xc3, 0xd8, 0x1c, 0x3f}
	allophone062Think := []uint8{0xae, 0xce, 0x45, 0x5c, 0x7c, 0x96, 0xac, 0x3a, 0x17, 0x71, 0xf1, 0x59, 0xb2, 0xeb, 0x5d, 0xef, 0xba, 0xd4, 0xc1, 0xc3, 0x98, 0xa5, 0x2a, 0x57, 0x33, 0x67, 0xb7, 0xa4, 0xaa, 0x43, 0x6d, 0x94, 0x6b, 0x48, 0x38, 0x7d}
	allophone063Boy := []uint8{0xae, 0xb6, 0xa6, 0xc6, 0x4d, 0xfc, 0xae, 0xda, 0x9a, 0x1a, 0x37, 0xf1, 0xbb, 0x6a, 0x6b, 0x6a, 0xdd, 0xd8, 0xdf, 0xa9, 0x93, 0x49, 0xd7, 0xf2, 0xc6, 0xa7, 0xf6, 0xc5, 0x29, 0xbd, 0xa3, 0xac, 0x3a, 0x16, 0x71, 0xf1, 0x59, 0xb2, 0xea, 0x58, 0xc4, 0xc5, 0x67, 0xc9, 0x3}
	allophone064Could := []uint8{0xae, 0xf6, 0xc6, 0x33, 0x92, 0xf2, 0xae, 0xda, 0x1b, 0xcf, 0x48, 0xca, 0xbb, 0xeb, 0x5b, 0xdf, 0xfa, 0xd6, 0xab, 0xf6, 0xc6, 0x33, 0x92, 0xf2, 0x3e}
	allophone065Poor := []uint8{0xa6, 0x16, 0xa5, 0x96, 0xd4, 0xbb, 0xb4, 0x5a, 0x94, 0x5a, 0x52, 0xef, 0x32, 0x6a, 0x51, 0xb3, 0x4c, 0x35, 0xdb, 0xaa, 0x65, 0xf3, 0x4c, 0x91, 0xb6, 0xa3, 0x96, 0x53, 0xaa, 0x44, 0x33, 0xb7, 0x5a, 0xd, 0xaa, 0x74, 0x6e, 0xd1, 0x6a, 0xdd, 0x5c, 0x74, 0xd0, 0xe7, 0x3}
	allophone066Core := []uint8{0xa6, 0x56, 0x3d, 0x5b, 0xd4, 0xba, 0xb6, 0x5a, 0xf5, 0x6c, 0x51, 0xeb, 0xba, 0x6a, 0x3d, 0x3d, 0x53, 0xa4, 0xed, 0xa8, 0xcd, 0x91, 0x4a, 0xd1, 0xcc, 0xa3, 0x36, 0x47, 0xb2, 0x45, 0xb3, 0xb4, 0xda, 0x6e, 0xaa, 0x74, 0x6e, 0xd1, 0x6a, 0x37, 0x5d, 0x74, 0xd0, 0xe7, 0x3}
	allophone067Low := []uint8{0xae, 0x4e, 0x3c, 0xba, 0x4c, 0xfc, 0x9e, 0x3a, 0xf1, 0xe8, 0x32, 0xf1, 0x7b, 0x6a, 0x27, 0xab, 0xd2, 0xc5, 0xdb, 0xa9, 0x8d, 0xae, 0x2a, 0xc7, 0x6c, 0xa3, 0x16, 0xa1, 0x36, 0x8d, 0xb2, 0xb6, 0x9a, 0xd5, 0xda, 0x10, 0xc9, 0x54, 0x6a, 0xba, 0xba, 0x93, 0x2c, 0xd5, 0x3}
	allophone068Shoe := []uint8{0xa6, 0x4e, 0xc6, 0x95, 0x4a, 0xf6, 0x8e, 0x3a, 0x19, 0x57, 0x2a, 0xd9, 0xbb, 0xeb, 0x55, 0x7, 0xee, 0xc9, 0xa5, 0x5d, 0x77, 0x3d, 0x6a, 0x1e, 0xca, 0x2d, 0x74, 0x51, 0xab, 0x59, 0x6c, 0xf, 0xd3, 0x5, 0xf}
	allophone069Mud := []uint8{0xae, 0x8e, 0xde, 0xd6, 0x4d, 0xd6, 0xac, 0x3a, 0x7a, 0x5b, 0x37, 0x59, 0xb3, 0xeb, 0x5b, 0xdf, 0x7a, 0xd7, 0xab, 0x8e, 0xde, 0xd6, 0x4d, 0xd6, 0x3c}
	allophone070Skull := []uint8{0xae, 0x8e, 0xde, 0xd6, 0x4d, 0xd6, 0xac, 0x3a, 0x7a, 0x5b, 0x37, 0x59, 0x73, 0xeb, 0x51, 0x9b, 0x1c, 0x6b, 0xaa, 0x6d, 0x5a, 0x2d, 0x7b, 0xad, 0x89, 0xa5, 0x4d, 0x35, 0x9f, 0xbd, 0x26, 0x9e, 0x36, 0xd4, 0xf4, 0x6e, 0x18, 0x9b, 0xe3, 0x50, 0xd3, 0xbb, 0x61, 0x6c, 0x8e, 0x1f}
	allophone071Pull := []uint8{0xae, 0x36, 0xbe, 0xd2, 0x9d, 0xb3, 0xad, 0xda, 0xf8, 0x4a, 0x77, 0xce, 0xb6, 0xeb, 0x5d, 0xef, 0x7a, 0xd4, 0x3a, 0x56, 0xba, 0x73, 0xb6, 0x52, 0xf3, 0x3a, 0x61, 0x62, 0x69, 0x43, 0x4d, 0xef, 0x86, 0xb1, 0x39, 0x8e, 0xb5, 0xae, 0x55, 0xcd, 0xf2, 0x47, 0xa, 0xdb, 0x79}
	allophone072Moon := []uint8{0xa6, 0x36, 0x22, 0x5b, 0x83, 0xb7, 0x8c, 0xda, 0x88, 0x6c, 0xd, 0xde, 0xb2, 0x6a, 0x23, 0xb3, 0x35, 0x78, 0xcb, 0xa8, 0x8d, 0xaa, 0x12, 0xc7, 0xfe, 0xad, 0x16, 0x7a, 0x52, 0x9c, 0xd7, 0x94, 0x9a, 0x94, 0x71, 0x35, 0x6f, 0x5c, 0x6a, 0x52, 0xc6, 0xd5, 0xbc, 0xf1, 0x3}
	allophone073Like := []uint8{0xa9, 0x56, 0x2e, 0x36, 0x21, 0x93, 0xd6, 0xba, 0xd7, 0xad, 0x56, 0x2e, 0x36, 0x21, 0x93, 0x3e}
	allophone074Bowl := []uint8{0xa1, 0xa6, 0x77, 0xc3, 0xd8, 0x1c, 0xc7, 0x3a, 0xd6, 0xa6, 0x66, 0xf9, 0x23, 0x85, 0xed, 0x3c}
	allophone075Awful := []uint8{0xa3, 0x76, 0xf2, 0x8f, 0xc5, 0xda, 0xb4, 0x5a, 0x9b, 0x3b, 0x61, 0x6b, 0x9b, 0x6a, 0xde, 0x67, 0x59, 0x3c, 0xad, 0xab, 0xd9, 0xbd, 0x62, 0x9, 0xa7, 0xae, 0xa6, 0x77, 0xc3, 0xd8, 0x1c, 0xa7, 0x9a, 0xe5, 0x8f, 0x14, 0xb6, 0xf3}
	allophone076May := []uint8{0xa9, 0x96, 0x5e, 0x52, 0xb9, 0x9d, 0xd6, 0xba, 0xd5, 0xd2, 0x4b, 0x2a, 0xb7, 0xd3, 0x7}
	allophone077Hum := []uint8{0xad, 0x76, 0x5a, 0x92, 0x10, 0x9d, 0xe7, 0x3a, 0xd7, 0xb9, 0xe, 0xb5, 0xd5, 0x52, 0x86, 0x96, 0x3c, 0xd4, 0x56, 0x73, 0x2b, 0xb8, 0xbb, 0x54, 0x6b, 0x53, 0x29, 0xe1, 0x8e, 0x42, 0x1d, 0x72, 0x25, 0x46, 0xcb, 0x7d}
	allophone078Nice := []uint8{0xa5, 0x76, 0x1a, 0x3a, 0xca, 0x6a, 0x95, 0x3a, 0x5a, 0x28, 0x4f, 0x6b, 0xdc, 0xea, 0xe0, 0x39, 0x58, 0xc2, 0xcb, 0x3}
	allophone079Sane := []uint8{0xa9, 0x76, 0xe, 0x53, 0xa4, 0xdd, 0x85, 0x5a, 0x6, 0xd1, 0xe0, 0x6e, 0x1d, 0x6a, 0x69, 0xc4, 0xdb, 0xdb, 0xb2, 0xaf, 0x43, 0x5d, 0x52, 0xa8, 0xa9, 0xa5, 0x79}
	allophone080Think := []uint8{0xa9, 0xe, 0x1e, 0xc6, 0x2c, 0x55, 0xb9, 0x9a, 0x39, 0xbb, 0x25, 0x55, 0x1d, 0x6a, 0xa3, 0x5c, 0x43, 0xc2, 0xe9, 0x3}
	allophone081Thing := []uint8{0xa3, 0xce, 0x41, 0xc1, 0x2f, 0x6c, 0xb5, 0x5a, 0x39, 0xa9, 0xa, 0xb7, 0x99, 0x6a, 0xed, 0x68, 0x4a, 0xdb, 0x62, 0xaa, 0xb5, 0xa3, 0x4d, 0xd, 0xb, 0xa1, 0xd6, 0xce, 0xb8, 0xc3, 0x1d, 0x9b, 0x5a, 0x3b, 0xe3, 0xe, 0x77, 0xfc}
	allophone082Real := []uint8{0xa6, 0xe6, 0x4d, 0xdd, 0x33, 0xe4, 0xbb, 0x5a, 0x75, 0xc, 0x4b, 0xd3, 0x9e, 0x6a, 0x33, 0x30, 0x2c, 0x4d, 0xfb, 0x3}
	allophone083Witch := []uint8{0xa6, 0xa6, 0x36, 0xdf, 0x89, 0xec, 0xa5, 0x9a, 0xdb, 0xf9, 0x42, 0x56, 0xda, 0x6a, 0xe1, 0xfb, 0xb, 0x38, 0xfd, 0x3}
	allophone084Which := []uint8{0x8, 0x50, 0x2b, 0x3, 0x3, 0xa6, 0xa6, 0x36, 0xdf, 0x89, 0xec, 0xa5, 0x9a, 0xdb, 0xf9, 0x42, 0x56, 0x5a, 0x6a, 0xe1, 0xfb, 0xb, 0x38, 0xfd, 0x3}
	allophone085You := []uint8{0xad, 0x1e, 0x5e, 0x14, 0xad, 0xfa, 0xf6, 0x7a, 0xd4, 0x35, 0x38, 0x43, 0xf8, 0xbe, 0x7}
	allophone086Bad := []uint8{0xaa, 0x96, 0x21, 0x3, 0xd5, 0x53, 0xb3, 0x9a, 0x24, 0x37, 0xd6, 0x4c, 0xf4}
	allophone087Dab := []uint8{0xa2, 0x26, 0xbe, 0x5d, 0x25, 0xda, 0xd2, 0x3a, 0xd5, 0x2a, 0xbb, 0xb8, 0x85, 0x62, 0x52, 0x3b, 0xe3, 0x1b, 0xce, 0x5d, 0x1f}
	allophone088Dig := []uint8{0xac, 0x36, 0xd2, 0xd5, 0xc3, 0x52, 0xaf, 0xfa, 0x8, 0xb7, 0x2c, 0x4b, 0xf5}
	allophone089Bid := []uint8{0xac, 0xa6, 0xd6, 0x3b, 0x23, 0x6d, 0xaa, 0xba, 0x44, 0x15, 0x16, 0xc9, 0x3c, 0xea, 0x18, 0x45, 0x94, 0xa2, 0xa6, 0xaa, 0x4b, 0x54, 0x61, 0x91, 0xcc, 0xf}
	allophone090Give := []uint8{0xae, 0x9e, 0x35, 0x44, 0x98, 0x9d, 0xac, 0x7a, 0x14, 0x37, 0xa0, 0x5a, 0xfb}
	allophone091Go := []uint8{0x60, 0xa2, 0x8e, 0x29, 0x58, 0x3c, 0x33, 0xbb, 0x3a, 0xa6, 0x60, 0xf1, 0xcc, 0xbc, 0xea, 0x40, 0x35, 0xa3, 0xd4, 0xf7, 0x3}
	allophone092Bag := []uint8{0xac, 0xce, 0x9d, 0x40, 0xb3, 0xe2, 0x98, 0x3a, 0x77, 0x2, 0xcd, 0x8a, 0x23, 0x6a, 0x6e, 0xd4, 0xd7, 0xcd, 0x35, 0xab, 0x89, 0xb5, 0x9c, 0x30, 0x37, 0xa9, 0x2e, 0xa7, 0xc5, 0xc9, 0x6c, 0x99, 0x3a, 0x47, 0xc6, 0xf0, 0x6c, 0xe8, 0xea, 0xdc, 0x9c, 0x95, 0x34, 0xf6, 0x3}
	allophone093Jug := []uint8{0x0, 0xe, 0xb8, 0x8e, 0x60, 0xd5, 0xc3, 0x9a, 0x89, 0x78, 0x9f, 0x7}
	allophone094Budge := []uint8{0xaa, 0x36, 0x91, 0x54, 0x6f, 0x2c, 0xa7, 0xda, 0x44, 0x52, 0xbd, 0xb1, 0x6c, 0x6a, 0x1a, 0x34, 0x7a, 0x83, 0xaa, 0xaa, 0x99, 0x75, 0xcf, 0x19, 0xc9, 0xae, 0xde, 0xbd, 0x2a, 0x1, 0x18, 0x3b, 0xe0, 0x8c, 0xce, 0x0, 0xb4, 0x34, 0xcd, 0x80, 0xd8, 0x8d, 0x1e}
	allophone095This := []uint8{0xaa, 0x76, 0x4e, 0x52, 0xc4, 0x33, 0xa7, 0xda, 0x39, 0x49, 0x11, 0xcf, 0x9c, 0xea, 0xe0, 0x34, 0x85, 0x79, 0x49, 0xaa, 0x93, 0xf5, 0x10, 0xc6, 0xcc, 0xa9, 0x4e, 0xce, 0x5d, 0x8, 0x3d, 0x97, 0x3a, 0x7b, 0x77, 0x23, 0x72, 0xfa}
	allophone096Clothe := []uint8{0xac, 0xce, 0xde, 0xdd, 0x88, 0x9c, 0x9a, 0x3a, 0x7b, 0x77, 0x23, 0x72, 0x6a, 0xeb, 0x5a, 0x33, 0x80, 0x3a, 0xf, 0x9, 0x30, 0x80, 0x3a, 0x8f, 0x7}
	allophone097Vine := []uint8{0xac, 0x96, 0xca, 0xbd, 0xcc, 0x5d, 0x9b, 0x5a, 0x2a, 0xf7, 0x32, 0x77, 0xed, 0x6a, 0xa9, 0xdc, 0xcb, 0xdc, 0xf5, 0x3}
	allophone098Alive := []uint8{0xa2, 0x16, 0x51, 0xdc, 0x58, 0xdb, 0x84, 0x5a, 0x44, 0x71, 0x63, 0x6d, 0x13, 0x6a, 0x21, 0x4d, 0xc2, 0xcc, 0x8d, 0x3, 0x8a, 0x31, 0xd5, 0x0, 0x1, 0x8a, 0x31, 0x7d}
	allophone099Zoo := []uint8{0xa2, 0xae, 0xdc, 0x4c, 0x18, 0xf3, 0xb8, 0xba, 0x72, 0x33, 0x61, 0xcc, 0x43, 0x0, 0x1, 0xe6, 0x52, 0xfd, 0xad, 0xaa, 0xa, 0xe9, 0xaf, 0x75, 0xab, 0x97, 0x74, 0x13, 0x26, 0x7f, 0xad, 0x5e, 0x3a, 0xc8, 0x18, 0xfd, 0x3e}
	allophone100Does := []uint8{0xa2, 0xfe, 0x9e, 0xd9, 0xa8, 0xd5, 0x86, 0xfa, 0x7b, 0x66, 0xa3, 0x56, 0x1b, 0x6b, 0x1, 0x7c, 0x67, 0x2e, 0x1, 0x2, 0x7c, 0x2d, 0xfa}
	allophone101Azure := []uint8{0xa6, 0xe, 0xc1, 0xca, 0xc3, 0xeb, 0x8c, 0x3a, 0x4, 0x2b, 0xf, 0xaf, 0x53, 0xea, 0x61, 0xcd, 0x44, 0xbc, 0x4f, 0xae, 0x5, 0xb0, 0x1c, 0x81, 0x4, 0x52, 0x3d, 0xac, 0x99, 0x88, 0xf7, 0x69, 0xf5, 0xb0, 0x66, 0x22, 0xde, 0xe7, 0x1}
	allophone102Beige := []uint8{0xaa, 0x1e, 0xd6, 0x4c, 0xc4, 0xfb, 0x94, 0x7a, 0x58, 0x33, 0x11, 0xef, 0x93, 0x6b, 0x3, 0xf4, 0x5e, 0xa9, 0x1, 0xd, 0x48, 0x80, 0x1, 0xbd, 0x57, 0x3e}
	allophone103Skate := []uint8{0x0, 0x60, 0x80, 0x5a, 0xc6, 0x1e}
	allophone104Case := []uint8{0x0, 0xe0, 0x80, 0xbc, 0x4c, 0x1c, 0x10, 0xb3, 0xa2, 0x2, 0x72, 0xe, 0x7c}
	allophone105Make := []uint8{0x0, 0x0, 0xa, 0xc8, 0xcb, 0x44, 0x1, 0x31, 0x2b, 0xa, 0x20, 0xe7, 0xc0, 0x7}
	allophone106Key := []uint8{0x0, 0x90, 0x80, 0xe2, 0xd9, 0x12, 0x90, 0xbd, 0x89, 0x3, 0x72, 0x50, 0x7e}
	allophone107Cough := []uint8{0x0, 0x10, 0x80, 0xd0, 0x9b, 0x3d, 0x20, 0x80, 0xd0, 0x9b, 0x1f}
	allophone108Space := []uint8{0x8, 0x8, 0xde, 0x84, 0x0, 0xc1, 0x9b, 0x0, 0x30, 0x80, 0x7a, 0x93, 0x7}
	allophone109Pie := []uint8{0x0, 0x10, 0x0, 0xe2, 0x4d, 0xc, 0x40, 0xbc, 0xc9, 0x3}
	allophone110Nap := []uint8{0x0, 0x0, 0x60, 0x0, 0xe2, 0x4d, 0x18, 0x40, 0xbc, 0x39, 0x2, 0x88, 0x37, 0x7f}
	allophone111Stake := []uint8{0x8, 0x8, 0x30, 0x92, 0x0, 0x1, 0x46, 0x2, 0x18, 0x60, 0x6b, 0x93, 0x7}
	allophone112Tie := []uint8{0x0, 0x10, 0x80, 0x1d, 0xc5, 0x3, 0x70, 0xac, 0x87, 0x1, 0x52, 0x28, 0x7e}
	allophone113Late := []uint8{0x0, 0x0, 0xe, 0x58, 0xc2, 0x5c, 0x3, 0x4, 0x58, 0xc2, 0xfc, 0x1}
	allophone114Church := []uint8{0x0, 0xa0, 0x80, 0x7b, 0x53, 0x3c, 0x10, 0x80, 0xd9, 0x2b, 0x15, 0x90, 0x83, 0xea, 0x3}
	allophone115Fat := []uint8{0x8, 0x28, 0x26, 0x8c, 0x0, 0xc5, 0x84, 0x71, 0xc0, 0x2, 0x1e, 0x48, 0x40, 0x31, 0x61, 0xf}
	allophone116Laugh := []uint8{0xc, 0xf0, 0xda, 0x45, 0x1, 0x5e, 0xbb, 0x38, 0xa0, 0x38, 0x37, 0x3, 0xe4, 0x58, 0xe6, 0x80, 0x6a, 0x33, 0x1c, 0x90, 0x6d, 0x86, 0x1, 0x72, 0x4c, 0x7b}
	allophone117Hit := []uint8{0x8, 0xc8, 0x96, 0x4, 0x1, 0xd9, 0x92, 0x50, 0x40, 0x3, 0xa, 0xc8, 0x96, 0xe4, 0x1}
	allophone118Home := []uint8{0x8, 0x30, 0xb6, 0x1a, 0x1, 0xc6, 0x56, 0x13, 0xc0, 0x98, 0x2a, 0xa, 0x30, 0xc0, 0x98, 0xaa, 0x7}
	allophone119Hut := []uint8{0x8, 0xf0, 0x21, 0x3, 0x1, 0x3e, 0x64, 0x8, 0x20, 0xda, 0x4c, 0x5, 0x78, 0x1f, 0x65, 0x0, 0x9f, 0xb2, 0x1e}
	allophone120Seem := []uint8{0x8, 0xf8, 0x29, 0x82, 0x0, 0x3f, 0x45, 0x48, 0x40, 0x3, 0x6, 0xf8, 0x29, 0xd3, 0x3, 0x1, 0xf8, 0xc9, 0xed, 0x1}
	allophone121Miss := []uint8{0x2, 0x38, 0x82, 0xcc, 0x1, 0x47, 0x90, 0x19, 0xe0, 0x3b, 0x56, 0x3, 0xfc, 0xac, 0x6a, 0x80, 0x5f, 0x4d, 0xd, 0xf0, 0x8b, 0x85, 0x0, 0x7e, 0x64, 0x41, 0xc0, 0xf, 0x22, 0xf}
	allophone122Shine := []uint8{0xc, 0xe8, 0xbd, 0xd2, 0x0, 0xbd, 0x57, 0x7a, 0x20, 0x2, 0x11, 0x48, 0x40, 0xef, 0x95, 0xf}
	allophone123Wash := []uint8{0xa, 0x98, 0x8e, 0x28, 0x1, 0xd3, 0x11, 0x25, 0x60, 0xf7, 0x94, 0x0, 0xec, 0x3e, 0x11, 0x80, 0x3d, 0xaa, 0x1c, 0x30, 0x57, 0x85, 0x3, 0x7a, 0xaf, 0x7c}
	allophone124Thing := []uint8{0x4, 0x28, 0xce, 0x44, 0x0, 0xc5, 0x99, 0x58, 0x40, 0x3, 0x2, 0x28, 0xce, 0xec, 0x1}
	allophone125With := []uint8{0xc, 0xf0, 0xcc, 0xcd, 0x0, 0x9e, 0xb9, 0x39, 0x20, 0x4b, 0x23, 0x1, 0x84, 0xec, 0x86, 0x0, 0x15, 0x54, 0xe1, 0x1}
}


/*
 1. AE1  as in ".A.ddition"
 2. AE1N as in ".A.nnuity"
 3. AH1  as in "Delt.a."s
 4. AH1N as in ".O.n time"
 5. AW1  as in ".Au.tonomy"
 6. AW1N as in "An.o.nimity"
 7. E1   as in ".E.liminate"
 8. E1N  as in ".E.nough"
 9. EH1  as in "Cont.e.xt"
10. EH1N as in "Anci.e.nt"
11. ER1N as in "West.e.rn"
12. I1   as in "Synth.e.s.i.s"
13. I1N  as in ".I.nane"
14. OO1  as in "T.oo.k on"
15. OW1N as in "D.o.nation"
16. U1   as in "Ann.u.al"
17. U1N  as in ".U.nique"
18. UH1  as in ".A.bove"
19. UH1M as in "Instr.u.ments"
20. UH1N as in ".U.nderneath"
21. Y1   as in "Ros.e.s"
22. Y1N  as in "Basem.e.nt"
23. ER1  as in "Seek.e.r"
24. OW1  as in "Rati.o."
25. Y2   as in "Funn.y."
26. AE2  as in "H.a.t"
27. AH2  as in "H.o.t"
28. AI2  as in "H.ei.ght"
29. AR2  as in "C.a.rt"
30. AU2  as in "H.ou.se"
31. AW2  as in "S.ou.ght"
32. E2   as in "H.ea.t"
33. EER2 as in "P.ie.rce"
34. EH2  as in "S.e.t"
35. EHR2 as in "Th.e.rapy"
36. EI2  as in "T.a.ke"
37. ER2  as in "H.u.rt"
38. I2   as in ".I.ssue"
39. OI2  as in "Ch.oi.ce"
40. OO2  as in "C.oo.k"
41. OOR2 as in "P.oor.ly"
42. OR2  as in "H.or.se"
43. OW2  as in "B.oa.t"
44. U2   as in "Sh.oo.t"
45. UH2  as in "H.u.t"
46. UU2  as in "B.oo.t"
47. AE3  as in "H.a.d"
48. AH3  as in ".O.dd"
49. AI3  as in "H.i.de"
50. AR3  as in "C.a.rd"
51. AU3  as in "L.ou.d"
52. AW3  as in "S.a.w"
53. E3   as in "S.ee.d"
54. EELL as in "H.eel."
55. EER3 as in "H.ear."
56. EH3  as in "S.ai.d"
57. EHR3 as in "Th.ere."
58. EI3  as in "D.ay."
59. ER3  as in "H.ear.d"
60. I3   as in "H.i.d"
61. ILL  as in "H.ill."
62. ING2 as in "Th.in.k"
63. OI3  as in "B.oy."
64. OO3  as in "C.ou.ld"
65. OOR3 as in "P.oo.r"
66. OR3  as in "C.ore."
67. OW3  as in "L.ow."
68. U3   as in "Sh.oe."
69. UH3  as in "M.u.d"
70. ULL  as in "Sk.ull."
71. UHL  as in "P.ull."
72. UU3  as in "M.oo.n"
73. L    as in ".L.ike"
74. L-   as in "Bow.l."
75. LL   as in "Awf.ul."
76. M    as in ".M.ay"
77. MM   as in "Hu.m."
78. N    as in ".N.ice"
79. NN   as in "Sa.ne."
80. NG1  as in "Thi.n.k"
81. NG2  as in "Thi.ng."
82. R    as in ".R.eal"
83. W    as in ".W.itch"
84. WH   as in ".Wh.ich"
85. Y    as in ".Y.ou"
86. B    as in ".B.ad"
87. BB   as in "Da.b."
88. D    as in ".D.ig"
89. DD   as in "Bi.d."
90. G1   as in ".G.ive"
91. G2   as in ".G.o"
92. GG   as in "Ba.g."
93. J    as in ".J.ug"
94. JJ   as in "Bu.dge."
95. THV  as in ".Th.is"
96. THV- as in "Clo.the."
97. V    as in ".V.ine"
98. VV   as in "Li.ve."
99. Z    as in ".Z.oo"
100. ZZ   as in "Doe.s."
101. ZH   as in "A.z.ure"
102. ZH-  as in "Bei.ge."
103. K2   as in "S.k.ate"
104. KH   as in ".C.ase"
105. KH-  as in "Ma.ke."
106. KH1  as in ".K.ey"
107. KH2  as in ".C.ough"
108. P    as in "S.p.ace"
109. PH   as in ".P.ie"
110. PH-  as in "Na.p."
111. T    as in "S.t.ake"
112. TH   as in ".T.ie"
113. TH-  as in "La.te."
114. CH   as in ".Ch.ur.ch."
115. F    as in ".F.at"
116. FF   as in "Lau.gh."
117. HI   as in ".H.it"
118. HO   as in ".H.ome"
119. HUH  as in ".H.ut"
120. S    as in ".S.eem"
121. SS   as in "Mi.ss."
122. SH   as in ".Sh.ine"
123. SH-  as in "Wa.sh."
124. THF  as in ".Th.ing"
125. THF- as in "Wi.th."
126. Pause1 <short pause>
127. Pause2 <long pause>
*/



/*

aa 04 4d 4f 52 45 04 b6 0e b5 00 46 42 51 01 2d 09 a8 00 18 00 48 dc 7d 01 2e 00 00 00 22 00 50
ec 5c 01 30 00 00 00 00 00 13 c3 46 01 31 00 0e 00 4a 00 14 09 53 01 32 00 00 00 40 00 14 5c 3e
01 33 00 00 00 00 00 14 9a 4d 01 34 00 36 00 54 00 14 e7 4a 01 35 00 00 00 5e 00 15 31 77 01 36
00 00 00 00 00 15 a8 40 01 37 00 2c 00 a9 00 15 e8 4f 01 38 00 00 00 7c 00 16 37 2d 01 39 00 00
00 00 00 16 64 80 01 41 00 72 00 90 00 16 e4 1c 02 41 31 00 00 00 9b 00 17 00 14 05 41 42 4f 55
54 00 00 00 00 00 17 14 55 05 41 46 54 45 52 00 86 00 d1 00 17 69 3c 05 41 47 41 49 4e 00 00 00
c5 00 17 a5 62 03 41 4c 4c 00 00 00 00 00 18 07 29 02 41 4d 00 b7 00 dc 00 18 30 46 02 41 4e 00
00 00 e7 00 18 76 36 03 41 4e 44 00 00 00 00 00 18 ac 67 06 41 4e 53 57 45 52 00 68 01 7e 00 19
13 4f 03 41 4e 59 00 00 00 00 00 19 62 45 03 41 52 45 01 02 01 1a 00 55 6e 32 02 41 53 00 00 01
25 00 19 a7 41 06 41 53 53 55 4d 45 00 00 00 00 00 19 e8 3d 02 41 54 01 0e 01 56 00 1a 25 1d 01
42 00 00 01 49 00 1a 42 22 04 42 41 43 4b 00 00 00 00 00 1a 64 2b 04 42 41 53 45 01 3f 01 63 00
1a 8f 4f 02 42 45 00 00 01 6e 00 1a 42 22 07 42 45 54 57 45 45 4e 00 00 00 00 00 1a de 69 05 42
4c 41 43 4b 01 34 01 cd 00 1b 47 43 04 42 4c 55 45 00 00 01 99 00 1b 8a 2c 04 42 4f 54 48 00 00
00 00 00 1b b6 34 06 42 4f 54 54 4f 4d 01 8c 01 b5 00 1b ea 36 03 42 55 54 00 00 01 c1 00 1c 20
28 03 42 55 59 00 00 00 00 00 1c 48 3e 02 42 59 01 a6 01 ee 00 1c 48 3e 03 42 59 45 00 00 01 e4
00 1c 48 3e 01 43 00 00 00 00 00 1c 86 53 03 43 41 4e 01 d8 01 fa 00 1c d9 37 08 43 41 53 53 45
54 54 45 00 00 02 0b 00 1d 10 37 06 43 45 4e 54 45 52 00 00 00 00 00 1d 47 3b 05 43 48 45 43 4b
00 f3 03 72 00 1d 82 20 06 43 48 4f 49 43 45 00 00 00 00 00 1d a2 44 05 43 4c 45 41 52 02 28 02
45 00 1d e6 3a 05 43 4f 4c 4f 52 00 00 02 53 00 1e 20 34 04 43 4f 4d 45 00 00 00 00 00 1e 54 33
05 43 4f 4d 45 53 02 37 02 8c 00 1e 87 57 05 43 4f 4d 4d 41 00 00 02 7c 00 1e de 3c 07 43 4f 4d
4d 41 4e 44 00 00 00 00 00 1f 1a 57 08 43 4f 4d 50 4c 45 54 45 02 6e 02 9d 00 1f 71 5c 09 43 4f
4d 50 4c 45 54 45 44 00 00 02 af 00 1f cd 67 08 43 4f 4d 50 55 54 45 52 00 00 00 00 00 20 34 57
09 43 4f 4e 4e 45 43 54 45 44 02 60 03 18 00 20 8b 68 07 43 4f 4e 53 4f 4c 45 00 00 02 e2 00 20
f3 49 07 43 4f 52 52 45 43 54 00 00 00 00 00 21 3c 46 06 43 4f 55 52 53 45 02 d2 03 01 00 21 82
3e 04 43 59 41 4e 00 00 03 0e 00 21 c0 43 01 44 00 00 00 00 00 22 03 39 04 44 41 54 41 02 f2 03
43 00 22 3c 58 06 44 45 43 49 44 45 00 00 03 34 00 22 94 69 06 44 45 56 49 43 45 00 00 00 00 00
22 fd 69 03 44 49 44 03 25 03 4f 00 23 66 5e 09 44 49 46 46 45 52 45 4e 54 00 00 03 61 00 23 c4
69 08 44 49 53 4b 45 54 54 45 00 00 00 00 00 24 2d 53 02 44 4f 02 c0 04 12 00 24 80 33 04 44 4f
45 53 00 00 03 8a 00 24 b3 37 05 44 4f 49 4e 47 00 00 00 00 00 24 ea 54 04 44 4f 4e 45 03 7d 03
a5 00 25 3e 5b 06 44 4f 55 42 4c 45 00 00 03 b4 00 25 99 3a 04 44 4f 57 4e 00 00 00 00 00 25 d3
3f 04 44 52 41 57 03 98 03 e8 00 26 12 56 07 44 52 41 57 49 4e 47 00 00 03 de 00 26 68 63 01 45
00 00 00 00 00 26 cb 25 04 45 41 43 48 03 ce 03 f5 00 26 f0 33 05 45 49 47 48 54 00 00 04 03 00
16 37 2d 06 45 49 47 48 54 59 00 00 00 00 00 27 23 36 06 45 4c 45 56 45 4e 03 c1 04 63 00 27 59
5d 04 45 4c 53 45 00 00 04 2e 00 27 b6 3f 03 45 4e 44 00 00 00 00 00 27 f5 71 04 45 4e 44 53 04
21 04 47 00 28 66 47 05 45 4e 54 45 52 00 00 04 55 00 28 ad 42 05 45 52 52 4f 52 00 00 00 00 00
28 ef 48 07 45 58 41 43 54 4c 59 04 3a 04 89 00 29 37 68 03 45 59 45 00 00 04 7f 00 37 93 3c 01
46 00 00 00 00 00 29 9f 23 07 46 49 46 54 45 45 4e 04 73 04 99 00 29 c2 5b 05 46 49 46 54 59 00
00 04 a7 00 2a 1d 43 06 46 49 47 55 52 45 00 00 00 00 00 2a 60 77 04 46 49 4e 44 02 1a 07 2a 00
2a d7 47 04 46 49 4e 45 00 00 00 00 00 2b 1e 3d 06 46 49 4e 49 53 48 04 c3 04 df 00 2b 5b 39 08
46 49 4e 49 53 48 45 44 00 00 04 f0 00 2b 94 43 05 46 49 52 53 54 00 00 00 00 00 2b d7 3d 03 46
49 54 04 d0 05 23 00 2c 14 2a 04 46 49 56 45 00 00 05 17 00 15 31 77 03 46 4f 52 00 00 00 00 00
14 e7 4a 05 46 4f 52 54 59 05 0a 05 31 00 2c 3e 41 04 46 4f 55 52 00 00 05 3e 00 14 e7 4a 08 46
4f 55 52 54 45 45 4e 00 00 00 00 00 2c 7f 9a 06 46 4f 55 52 54 48 04 fe 05 9d 00 2d 19 5b 04 46
52 4f 4d 00 00 05 6b 00 2d 74 48 05 46 52 4f 4e 54 00 00 00 00 00 2d bc 2f 01 47 05 5e 05 83 00
2d eb 3d 05 47 41 4d 45 53 00 00 05 91 00 2e 28 64 03 47 45 54 00 00 00 00 00 2e 8c 2e 07 47 45
54 54 49 4e 47 05 79 05 c8 00 2e ba 7e 04 47 49 56 45 00 00 05 ba 00 2f 38 55 05 47 49 56 45 53
00 00 00 00 00 2f 8d 6f 02 47 4f 05 ad 05 d3 00 2f fc 35 04 47 4f 45 53 00 00 05 e0 00 30 31 48
05 47 4f 49 4e 47 00 00 00 00 00 30 79 5d 04 47 4f 4f 44 05 4f 06 8b 00 30 d6 24 09 47 4f 4f 44
20 57 4f 52 4b 00 00 00 00 00 30 fa 4e 07 47 4f 4f 44 42 59 45 05 fb 06 1d 00 31 48 58 03 47 4f
54 00 00 06 29 00 31 a0 31 04 47 52 41 59 00 00 00 00 00 31 d1 4c 05 47 52 45 45 4e 06 0d 06 5c
00 32 1d 61 05 47 55 45 53 53 00 00 06 52 00 32 7e 42 01 48 00 00 00 00 00 32 c0 2f 03 48 41 44
06 44 06 68 00 32 ef 4a 04 48 41 4e 44 00 00 06 75 00 33 39 46 0d 48 41 4e 44 48 45 4c 44 20 55
4e 49 54 00 00 00 00 00 33 7f 86 03 48 41 53 06 36 06 d9 00 34 05 45 04 48 41 56 45 00 00 06 a4
00 34 4a 42 04 48 45 41 44 00 00 00 00 00 34 8c 59 04 48 45 41 52 06 97 06 be 00 34 e5 35 05 48
45 4c 4c 4f 00 00 06 cc 00 35 1a 57 04 48 45 4c 50 00 00 00 00 00 35 71 3d 04 48 45 52 45 06 b1
07 01 00 34 e5 35 06 48 49 47 48 45 52 00 00 06 f5 00 35 ae 5c 03 48 49 54 00 00 00 00 00 36 0a
34 04 48 4f 4d 45 06 e6 07 0e 00 36 3e 4b 03 48 4f 57 00 00 07 1a 00 36 89 66 07 48 55 4e 44 52
45 44 00 00 00 00 00 36 ef 68 05 48 55 52 52 59 05 ee 08 61 00 37 57 3c 01 49 00 00 00 00 00 37
93 3c 05 49 20 57 49 4e 07 38 07 50 00 37 cf 81 02 49 46 00 00 07 5b 00 38 50 22 02 49 4e 00 00
00 00 00 38 72 43 04 49 4e 43 48 07 42 07 96 00 38 b5 45 06 49 4e 43 48 45 53 00 00 07 82 00 38
fa 51 0b 49 4e 53 54 52 55 43 54 49 4f 4e 00 00 00 00 00 39 4b 72 0c 49 4e 53 54 52 55 43 54 49
4f 4e 53 07 73 07 ab 00 39 bd 75 02 49 53 00 00 07 b6 00 3a 32 48 02 49 54 00 00 00 00 00 3a 7a
34 01 4a 07 66 08 10 00 3a ae 3f 08 4a 4f 59 53 54 49 43 4b 00 00 07 dc 00 3a ed 5f 04 4a 55 53
54 00 00 00 00 00 3b 4c 3e 01 4b 07 cb 07 f3 00 3b 8a 2f 03 4b 45 59 00 00 07 ff 00 3b b9 30 08
4b 45 59 42 4f 41 52 44 00 00 00 00 00 3b e9 66 04 4b 4e 4f 57 07 e9 08 35 00 3c 4f 40 01 4c 00
00 08 27 00 3c 8f 41 05 4c 41 52 47 45 00 00 00 00 00 3c d0 49 06 4c 41 52 47 45 52 08 1d 08 44
00 3d 19 4e 07 4c 41 52 47 45 53 54 00 00 08 54 00 3d 67 77 04 4c 41 53 54 00 00 00 00 00 3d de
40 05 4c 45 41 52 4e 07 c1 09 00 00 3e 1e 5a 04 4c 45 46 54 00 00 08 7c 00 3e 78 3a 04 4c 45 53
53 00 00 00 00 00 3e b2 56 03 4c 45 54 08 6f 08 95 00 3f 08 27 04 4c 49 4b 45 00 00 08 a2 00 3f
2f 3b 05 4c 49 4b 45 53 00 00 00 00 00 3f 6a 6b 04 4c 49 4e 45 08 89 08 d7 00 3f d5 76 04 4c 4f
41 44 00 00 08 ca 00 40 4b 88 04 4c 4f 4e 47 00 00 00 00 00 40 d3 6a 04 4c 4f 4f 4b 08 bd 08 e4
00 41 3d 54 05 4c 4f 4f 4b 53 00 00 08 f2 00 41 91 56 05 4c 4f 57 45 52 00 00 00 00 00 41 e7 4c
01 4d 08 b0 09 4c 00 42 33 34 04 4d 41 44 45 00 00 09 17 00 42 67 47 07 4d 41 47 45 4e 54 41 00
00 00 00 00 42 ae 80 04 4d 41 4b 45 09 0a 09 34 00 43 2e 4f 02 4d 45 00 00 09 3f 00 43 7d 40 04
4d 45 41 4e 00 00 00 00 00 43 bd 48 06 4d 45 4d 4f 52 59 09 27 09 7c 00 44 05 67 07 4d 45 53 53
41 47 45 00 00 09 6b 00 44 6c 6b 08 4d 45 53 53 41 47 45 53 00 00 00 00 00 44 d7 7a 06 4d 49 44
44 4c 45 09 5b 09 8b 00 45 51 42 05 4d 49 47 48 54 00 00 09 99 00 45 93 4c 06 4d 4f 44 55 4c 45
00 00 00 00 00 45 df 63 01 2b 00 00 00 00 00 51 b3 7e 04 4d 4f 53 54 00 00 00 00 00 46 93 4c 04
4d 4f 56 45 09 b2 09 cc 00 46 df 5e 04 4d 55 53 54 00 00 09 d9 00 47 3d 49 01 4e 00 00 00 00 00
47 86 3a 04 4e 41 4d 45 09 bf 0a 0a 00 47 c0 73 04 4e 45 41 52 00 00 09 fd 00 48 33 4d 04 4e 45
45 44 00 00 00 00 00 48 80 5c 08 4e 45 47 41 54 49 56 45 09 f0 0a 1b 00 48 dc 7d 04 4e 45 58 54
00 00 0a 28 00 49 59 4c 08 4e 49 43 45 20 54 52 59 00 00 00 00 00 49 a5 a9 04 4e 49 4e 45 09 e3
0a 87 00 16 64 80 06 4e 49 4e 45 54 59 00 00 0a 55 00 4a 4e 5d 02 4e 4f 00 00 00 00 00 3c 4f 40
03 4e 4f 54 0a 46 0a 6c 00 4a ab 2f 03 4e 4f 57 00 00 0a 78 00 4a da 46 06 4e 55 4d 42 45 52 00
00 00 00 00 4b 20 5d 01 4f 0a 60 0a a8 00 4b 7d 3d 02 4f 46 00 00 0a 9c 00 4b ba 59 03 4f 46 46
00 00 00 00 00 4c 13 2e 02 4f 48 0a 91 0a b3 00 4b 7d 3d 02 4f 4e 00 00 0a be 00 4c 41 4a 03 4f
4e 45 00 00 00 00 00 14 09 53 04 4f 4e 4c 59 0a 39 0b 68 00 4c 8b 51 02 4f 52 00 00 0a e2 00 4c
dc 58 05 4f 52 44 45 52 00 00 00 00 00 4d 34 56 05 4f 54 48 45 52 0a d7 0a fe 00 4d 8a 4a 03 4f
55 54 00 00 0b 0a 00 4d d4 36 04 4f 56 45 52 00 00 00 00 00 4e 0a 5c 01 50 0a f0 0b 3e 00 4e 66
39 04 50 41 52 54 00 00 0b 2e 00 4e 9f 41 07 50 41 52 54 4e 45 52 00 00 00 00 00 4e e0 51 05 50
41 52 54 53 0b 21 0b 4c 00 4f 31 50 06 50 45 52 49 4f 44 00 00 0b 5b 00 4f 81 64 04 50 4c 41 59
00 00 00 00 00 4f e5 48 05 50 4c 41 59 53 0b 17 0b c3 00 50 2d 66 06 50 4c 45 41 53 45 00 00 0b
85 00 50 93 59 05 50 4f 49 4e 54 00 00 00 00 00 50 ec 5c 08 50 4f 53 49 54 49 4f 4e 0b 76 0b a4
00 51 48 6b 08 50 4f 53 49 54 49 56 45 00 00 0b b5 00 51 b3 7e 05 50 52 45 53 53 00 00 00 00 00
52 31 3c 05 50 52 49 4e 54 0b 93 0b f1 00 52 6d 3d 07 50 52 49 4e 54 45 52 00 00 0b e1 00 52 aa
4f 07 50 52 4f 42 4c 45 4d 00 00 00 00 00 52 f9 67 08 50 52 4f 42 4c 45 4d 53 0b d1 0c 02 00 53
60 8e 07 50 52 4f 47 52 41 4d 00 00 0c 12 00 53 ee 89 03 50 55 54 00 00 00 00 00 54 77 33 07 50
55 54 54 49 4e 47 0a ca 0d 62 00 54 aa 76 01 51 00 00 00 00 00 55 20 4e 01 52 0c 2e 0c 42 00 55
6e 32 08 52 41 4e 44 4f 4d 4c 59 00 00 0c 53 00 55 a0 b2 04 52 45 41 44 00 00 00 00 00 56 52 61
05 52 45 41 44 31 0c 38 0c 96 00 57 c1 40 0e 52 45 41 44 59 20 54 4f 20 53 54 41 52 54 00 00 0c
85 00 56 b3 92 08 52 45 43 4f 52 44 45 52 00 00 00 00 00 57 45 7c 03 52 45 44 0c 6e 0c a2 00 57
c1 40 05 52 45 46 45 52 00 00 0c b0 00 58 01 60 08 52 45 4d 45 4d 42 45 52 00 00 00 00 00 58 61
6e 06 52 45 54 55 52 4e 0c 60 0d 12 00 58 cf 6b 06 52 45 57 49 4e 44 00 00 0c df 00 59 3a 88 05
52 49 47 48 54 00 00 00 00 00 7c 38 55 05 52 4f 55 4e 44 0c d0 0c fb 00 59 c2 98 01 53 00 00 0d
05 00 5a 5a 47 04 53 41 49 44 00 00 00 00 00 5a a1 4e 04 53 41 56 45 0c ed 0d 38 00 5a ef 76 03
53 41 59 00 00 0d 2b 00 5b 65 3d 04 53 41 59 53 00 00 00 00 00 5b a2 59 06 53 43 52 45 45 4e 0d
1f 0d 47 00 5b fb 60 06 53 45 43 4f 4e 44 00 00 0d 56 00 5c 5b 64 03 53 45 45 00 00 00 00 00 1c
86 53 04 53 45 45 53 0c c1 0e 0c 00 5c bf 5c 03 53 45 54 00 00 0d 7b 00 5d 1b 35 05 53 45 56 45
4e 00 00 00 00 00 15 e8 4f 07 53 45 56 45 4e 54 59 0d 6f 0d 99 00 5d 50 55 05 53 48 41 50 45 00
00 0d a7 00 5d a5 39 06 53 48 41 50 45 53 00 00 00 00 00 5d de 49 05 53 48 49 46 54 0d 89 0d e2
00 5e 27 35 05 53 48 4f 52 54 00 00 0d d2 00 5e 5c 49 07 53 48 4f 52 54 45 52 00 00 00 00 00 5e
a5 7f 06 53 48 4f 55 4c 44 0d c4 0d f1 00 5f 24 49 04 53 49 44 45 00 00 0d fe 00 5f 6d 5b 05 53
49 44 45 53 00 00 00 00 00 5f c8 52 03 53 49 58 0d b6 0e 60 00 15 a8 40 05 53 49 58 54 59 00 00
0e 26 00 60 1a 56 05 53 4d 41 4c 4c 00 00 00 00 00 60 70 3e 07 53 4d 41 4c 4c 45 52 0e 18 0e 44
00 60 ae 43 08 53 4d 41 4c 4c 45 53 54 00 00 0e 55 00 60 f1 62 02 53 4f 00 00 00 00 00 61 53 44
04 53 4f 4d 45 0e 34 0e 89 00 61 97 2f 05 53 4f 52 52 59 00 00 0e 7b 00 61 c6 60 05 53 50 41 43
45 00 00 00 00 00 62 26 37 06 53 50 41 43 45 53 0e 6d 0e 98 00 62 5d 6f 05 53 50 45 4c 4c 00 00
0e a6 00 62 cc 67 06 53 51 55 41 52 45 00 00 00 00 00 63 33 49 05 53 54 41 52 54 0c 1e 11 5b 00
63 7c 49 04 53 54 45 50 00 00 00 00 00 63 c5 32 04 53 54 4f 50 0e c3 0e dd 00 63 f7 2c 03 53 55
4d 00 00 0e e9 00 61 97 2f 08 53 55 50 50 4f 53 45 44 00 00 00 00 00 64 23 66 0b 53 55 50 50 4f
53 45 44 20 54 4f 0e d0 0f 25 00 64 89 6b 04 53 55 52 45 00 00 0f 1b 00 64 f4 5d 01 54 00 00 00
00 00 65 51 3a 04 54 41 4b 45 0f 0e 0f 32 00 65 8b 34 04 54 45 45 4e 00 00 0f 3f 00 65 bf 44 04
54 45 4c 4c 00 00 00 00 00 66 03 4b 03 54 45 4e 0e fa 0f bc 00 66 4e 48 11 54 45 58 41 53 20 49
4e 53 54 52 55 4d 45 4e 54 53 00 00 0f 72 00 66 96 c5 04 54 48 41 4e 00 00 00 00 00 67 5b 5b 04
54 48 41 54 0f 58 0f 8c 00 67 b6 60 11 54 48 41 54 20 49 53 20 49 4e 43 4f 52 52 45 43 54 00 00
0f a6 00 68 16 e8 0d 54 48 41 54 20 49 53 20 52 49 47 48 54 00 00 00 00 00 68 fe 76 03 54 48 45
0f 7f 0f e3 00 69 74 42 04 54 48 45 31 00 00 0f d5 00 69 b6 2b 05 54 48 45 49 52 00 00 00 00 00
6a 72 6c 04 54 48 45 4e 0f c8 0f f0 00 69 e1 91 05 54 48 45 52 45 00 00 0f fe 00 6a 72 6c 05 54
48 45 53 45 00 00 00 00 00 6a de 69 04 54 48 45 59 0f 4c 10 b8 00 6b 47 59 05 54 48 49 4e 47 00
00 10 27 00 6b a0 6f 06 54 48 49 4e 47 53 00 00 00 00 00 6c 0f 64 05 54 48 49 4e 4b 10 19 10 44
00 6c 73 49 05 54 48 49 52 44 00 00 10 52 00 6c bc 55 08 54 48 49 52 54 45 45 4e 00 00 00 00 00
6d 11 91 06 54 48 49 52 54 59 10 36 10 8d 00 6d a2 3c 04 54 48 49 53 00 00 10 7f 00 6d de 48 05
54 48 52 45 45 00 00 00 00 00 14 9a 4d 05 54 48 52 45 57 10 72 10 9b 00 6e 26 43 07 54 48 52 4f
55 47 48 00 00 10 ab 00 6e 26 43 04 54 49 4d 45 00 00 00 00 00 6e 69 47 02 54 4f 10 63 11 05 00
14 5c 3e 08 54 4f 47 45 54 48 45 52 00 00 10 d4 00 6e b0 6f 04 54 4f 4e 45 00 00 00 00 00 6f 1f
6e 03 54 4f 4f 10 c3 10 ed 00 14 5c 3e 03 54 4f 50 00 00 10 f9 00 6f 8d 2e 03 54 52 59 00 00 00
00 00 6f bb 54 09 54 52 59 20 41 47 41 49 4e 10 e1 11 33 00 70 0f 83 04 54 55 52 4e 00 00 11 24
00 70 92 3c 06 54 57 45 4c 56 45 00 00 00 00 00 70 ce 4b 06 54 57 45 4e 54 59 11 17 11 42 00 71
19 57 03 54 57 4f 00 00 11 4e 00 14 5c 3e 04 54 59 50 45 00 00 00 00 00 71 70 4e 01 55 10 0c 12
8d 00 71 be 36 04 55 48 4f 48 00 00 00 00 00 71 f4 51 05 55 4e 44 45 52 11 65 11 80 00 72 45 58
0a 55 4e 44 45 52 53 54 41 4e 44 00 00 11 93 00 72 9d 92 05 55 4e 54 49 4c 00 00 00 00 00 73 2f
70 02 55 50 11 72 11 c6 00 73 9f 24 05 55 50 50 45 52 00 00 11 ba 00 73 c3 40 03 55 53 45 00 00
00 00 00 74 03 46 01 56 11 ac 11 d0 00 74 49 3e 04 56 41 52 59 00 00 11 dd 00 74 87 53 04 56 45
52 59 00 00 00 00 00 74 da 46 01 57 11 a1 12 33 00 75 20 7d 04 57 41 49 54 00 00 12 01 00 75 9d
42 04 57 41 4e 54 00 00 00 00 00 75 df 42 05 57 41 4e 54 53 11 f4 12 1c 00 76 21 5c 03 57 41 59
00 00 12 28 00 76 b0 67 02 57 45 00 00 00 00 00 76 7d 33 05 57 45 49 47 48 12 0e 12 5d 00 76 b0
67 06 57 45 49 47 48 54 00 00 12 50 00 75 9d 42 04 57 45 4c 4c 00 00 00 00 00 77 17 45 04 57 45
52 45 12 41 12 6a 00 77 5c 60 04 57 48 41 54 00 00 12 77 00 77 bc 2d 0d 57 48 41 54 20 57 41 53
20 54 48 41 54 00 00 00 00 00 77 e9 8c 04 57 48 45 4e 11 ea 13 2a 00 78 75 36 05 57 48 45 52 45
00 00 12 a8 00 78 ab 49 05 57 48 49 43 48 00 00 00 00 00 78 f4 30 05 57 48 49 54 45 12 9a 12 c4
00 79 24 45 03 57 48 4f 00 00 12 d0 00 79 69 4b 03 57 48 59 00 00 00 00 00 79 b4 5d 04 57 49 4c
4c 12 b6 13 02 00 7a 11 5a 04 57 49 54 48 00 00 12 f6 00 7a 6b 40 03 57 4f 4e 00 00 00 00 00 14
09 53 04 57 4f 52 44 12 e9 13 0f 00 7a ab 5f 05 57 4f 52 44 53 00 00 13 1d 00 7b 0a 6b 04 57 4f
52 4b 00 00 00 00 00 7b 75 47 07 57 4f 52 4b 49 4e 47 12 dc 13 77 00 7b bc 7c 05 57 52 49 54 45
00 00 13 48 00 7c 38 55 01 58 00 00 00 00 00 7c 8d 25 01 59 13 3a 13 5c 00 7c b2 46 06 59 45 4c
4c 4f 57 00 00 13 6b 00 7c f8 60 03 59 45 53 00 00 00 00 00 7d 58 41 03 59 45 54 13 52 13 9f 00
7d 99 42 03 59 4f 55 00 00 13 8f 00 71 be 36 07 59 4f 55 20 57 49 4e 00 00 00 00 00 7d db 72 04
59 4f 55 52 13 83 13 ac 00 7e 4d 4c 01 5a 00 00 13 b6 00 7e 99 3b 04 5a 45 52 4f 00 00 00 00 00
13 c3 46 b5 9a 5c 13 18 bf 79 45 99 2c c4 97 ee 4d 07 bb 23 55 b3 92 3a 6c cd 35 6c d4 8d a2 4b
cc 4e f7 36 2a 65 ce 9e 92 7b ab 1b 54 97 95 7b e6 f7 60 99 dd 25 b4 ae 98 2a 2d f4 8b 2d 9e 0f
c1 f8 3f 37 e8 dd 3c 00 00 65 a5 6c fb 91 37 a5 61 db 9f 42 6a 59 54 87 df d0 1c be f6 62 c6 3e
b9 46 a4 79 b3 14 3a 35 4c 57 1c dd 58 b2 11 6a 54 ca e6 6e 9e 0e c1 cf 29 36 97 6b 30 73 86 51
2d dc 6c 53 04 58 3b b5 b3 14 b1 1e de da 64 c9 08 56 a9 79 5b 32 a4 a1 59 5a 59 e0 40 1b 8a 3c
00 e3 5e 18 04 a1 47 a7 68 86 a2 a6 9b f6 3d c7 99 2d 2c 8a d6 92 93 7b d5 6a dc 88 95 74 9b ce
d7 5e 96 54 b3 ad d9 58 1a 54 ce cf b3 56 09 46 51 cb cb 55 82 8f 72 ee ae 4f 30 10 53 22 40 58
0d 02 13 b2 a8 ce d3 2d 9c f0 9d 97 28 67 b5 7d 27 7e 0a 0a 95 fb 89 a7 a4 a0 84 9e e2 7a 5e 18
90 bd fa a2 d3 83 15 6f 6e 6a b5 b4 82 19 cd f7 2b 74 30 7a 91 a7 66 a6 fc 1d 83 d6 4f 59 70 c6
62 f5 1c d5 ca 33 c0 03 00 c4 b4 24 05 80 d0 1a 04 a3 3d ef 88 32 59 59 25 21 6f 9c a1 76 e9 c9
19 21 31 af 7a a2 b3 7b 54 89 c6 4b 1a be d6 aa 8d 4f 64 b7 95 e7 db 43 d2 3f 5d 92 74 91 aa 2f
c7 66 ad a9 d5 33 dd d9 b7 0c 6f 88 ac 6e de 00 00 03 01 17 b3 62 77 39 f6 8c 66 9c 9d cf ac a2
6b 2f 26 7b ee e4 9e bb c9 5e fb b8 28 a2 72 67 c2 e9 ce a0 bc a2 10 ba a3 a8 a7 2a 84 6a 60 e6
09 ca e3 2a a8 28 ae f2 d8 d5 e8 c2 a4 9a ba 36 7a 55 9a ae 2f 95 9a 96 ab 2b 8b e3 56 85 7a ca
db 08 d1 61 62 ba b6 c2 14 58 58 ac 2d 2f 74 d2 95 ae 46 cc 19 45 63 5a 90 9b 04 a9 de 99 dd 60
20 1c e4 21 00 a7 ab f8 60 19 6d 23 80 f4 66 62 04 80 cb 69 18 02 20 f5 69 25 06 85 4a d1 f6 4d
49 1a 23 a3 e6 93 72 62 a9 59 1d 22 b4 6e be 12 1b c8 a8 c0 00 08 06 45 d1 00 71 b5 4a 03 f3 79
58 08 03 f1 67 f0 00 00 30 1f 7b 32 40 58 0b 01 a0 34 0d 46 e8 b4 c2 d4 8f b2 76 4d cb 49 98 f8
76 37 8a 3a a7 9d 5a 1a 5c 6e a9 ef 79 78 0c f5 ee a6 4d 94 42 34 65 1d 9b 26 b3 09 51 a7 66 c8
1c 31 cb b9 1a 2d e8 d8 94 ef 46 8c 7a a3 24 cb 11 b3 9e 6b 09 8e ef b4 f4 a3 2a c5 c6 ad 3d 2a
aa b3 78 4c 4d 5a a1 1d 5d 73 73 66 6c d1 6a ab 44 13 aa ce c6 f4 e0 00 00 40 22 b6 53 80 20 18
84 61 81 e0 01 4a 55 4c d4 5b 6d 95 35 a9 96 b3 12 75 8d 7a e3 b5 b2 75 5c 5d 30 8a 55 4e 55 07
dc a4 0d b3 95 36 39 2c b3 54 d4 ee 86 4b 34 d7 31 3b df 92 4b 94 4c 50 e8 e3 b3 44 57 95 46 15
0e 93 5e e5 b2 9c cc 1a d9 71 74 6e b0 ab 15 cb 5f 0b c2 2d ac 6e 98 4a 69 0b 33 13 a6 11 c2 a2
c4 55 5d 8c 30 c4 b3 18 c3 62 98 a4 4a cd 65 58 2d c7 12 95 11 66 28 6a bc 44 d8 61 8b 94 ec 53
31 d0 00 0f e5 51 73 a4 e5 54 bb 46 e7 dc fd 9f b5 6a ab 4e 05 22 be 99 b2 b7 3f 07 e1 e0 00 00
b5 6c a4 62 3e 34 b3 66 6c cd b7 b6 f7 4e e7 dd 1b fc 00 00 85 10 74 f3 52 db ef 46 2a 93 63 9d
93 5b 6a a2 c7 d8 73 ba aa b4 c2 60 55 e5 61 8a a0 a9 b7 64 dd ea a7 b6 ad d5 bd 7b ab 75 72 a7
bf 2d 96 2b b9 aa f5 52 b6 43 71 12 2b ae ad ac cf 35 6c e2 2b 45 ed d2 e1 a6 00 00 40 26 b7 81
01 c5 bb 08 07 b2 18 03 c0 a4 d1 c4 c3 85 c3 39 34 6f 4c e7 55 5e cd d9 fb 47 6a ad 8c 01 a8 63
58 0b 00 01 80 51 4c f2 a1 fa 37 69 59 bd eb ad d3 c1 d8 37 05 11 31 2e 46 2d c0 55 34 48 cd 93
6e 6b 33 0e 4f 10 0b 8b bd aa 75 e5 ce 2a b5 ca a3 18 87 45 d2 f2 b9 42 15 91 6e b6 ad ed 44 84
97 b5 6c cd b8 b2 8d 44 75 e9 25 e7 c4 d4 91 ba c6 49 37 36 e5 c4 b4 ab 42 2e 28 f3 36 27 95 a2
10 b1 c7 ba f7 66 2d 8d e6 4d 66 47 9b f1 7c 31 10 4f 31 d8 60 97 a5 33 62 6d f8 37 04 e0 2c 0b
08 99 5b 49 78 00 00 94 8d 67 ea 9d 67 2b 2d e5 c4 ac 77 6b 1b d9 73 2e 66 cd 59 ab 35 d2 c4 d7
31 46 dc 95 ba 17 23 5b 6a 14 a6 7d eb bd 37 f8 c5 14 b2 bb 26 a8 b9 45 2c d0 ab 52 2e 55 3c b4
22 54 ab b5 76 ae 55 2b ad 9a cc 73 b5 f2 b8 d6 29 59 ea 9b b7 16 b6 95 6a 93 1b b5 ac 71 d2 a3
2a df 73 ee 65 cb 71 52 c3 bb 19 72 5e 92 a7 6e 57 dc bb fc 00 00 c5 53 cb 32 3d 44 fb 46 a8 a6
99 cd 72 39 ba 3b 47 2a a6 1e 4c ee 47 da b7 56 eb c5 b1 b2 c2 9c c8 a5 6b 5e 56 8b 6e c6 dc db
d5 7a 60 70 a9 92 35 de bb fc 00 00 85 13 cb 32 3d 44 f3 46 27 a6 99 cd 72 39 d9 b3 35 a6 a6 1e
4c ee 47 69 e9 1a 54 66 cd da a9 ca 8d 59 b2 96 ae 54 96 45 5c e5 ac 8c e9 89 9b 35 2b a4 fa a1
74 cd c2 f5 c1 59 59 eb 6e bd 50 3a 46 c7 03 af 4c 0e 91 ee 34 cb d2 e2 b6 5c 8d 2b 05 b4 75 9b
22 49 c4 f9 43 2a 29 d1 1e 36 0b c5 18 4a 32 ae b3 5a f0 a5 11 9a d4 3d 54 7b 57 28 96 1e 51 6a
39 d9 f7 3e 26 7c a4 a0 e6 bc e9 5a f3 77 1e b6 ae 69 55 79 d4 b2 b2 8d 80 76 0a cf 02 60 38 0e
4d de d6 46 32 ca b0 ed a5 ae ba 32 da d9 a1 7a ee 53 34 f3 a6 60 bb 8e 5b c5 a9 58 26 e7 3b 30
99 f0 e5 55 b4 ca 2a eb 7b 57 68 ed 14 a2 11 49 35 c5 55 28 a6 9e 10 67 4b da b5 6d 74 8c 83 39
cb 5d 6a a9 ab 5e 4e 18 4a 04 cf 5b 00 66 0f 4a ad 5b b0 d5 8b dc 47 b4 cd 45 63 03 b9 9b d3 ad
19 3d 4a 6e f3 2c f0 e4 f4 bc ca ab 69 79 45 2f 32 aa f9 ee d1 ca 21 79 94 69 b2 f6 ae d5 aa df
47 94 69 c6 ee cb db 7b 67 6c 2d 55 1e 5a 2a 1a eb 55 27 87 8c 89 32 d6 9e 9a 17 25 13 b5 40 68
88 a9 00 e0 14 02 81 e0 b5 8c 94 c2 c4 e4 f3 65 ec a5 8d 7c 4a bc 74 a7 83 80 98 21 9e 03 c0 70
1e 74 3c c0 78 16 85 ab 98 58 9b d6 64 bc 97 96 b3 5e 8a d5 4a d8 e3 a5 46 55 ce d8 dc fb df 83
71 6e 2d cb c0 d4 d1 93 5c bc e5 3b 47 68 ed 1b aa f6 46 c0 00 50 13 53 a3 00 d2 1d 88 03 95 06
01 e0 d5 b7 8b 28 ba dd b5 6d a2 ea 0f 7b 5c 59 9a 29 8c 46 b3 76 6e cd d9 9b 2f 6d 6d 8d ed bc
f7 f8 00 00 e5 b2 9a d3 9c 6b b9 84 2a b6 c9 73 6e 61 2a b5 b3 5c b7 b7 f6 ee db db 7b 54 e8 00
10 0b 43 22 80 84 d1 50 04 9e e4 20 1d 14 3f 65 67 3c c4 bc 48 39 54 a7 10 6c 77 6e 55 39 b2 2c
4f 97 b5 f2 b9 c9 8d 66 74 ba ae 6f 63 4a 9c df 6b b5 93 52 a3 ec 64 ed 6c e6 a8 8b 5a 49 5d 95
b3 19 45 b1 40 74 e8 8a 03 cd 54 60 1f 8b 2b 00 fd 9c d8 07 a6 a4 a0 34 9b 61 00 d9 50 78 85 61
65 cc c4 38 b9 45 5a d0 72 b7 ee 55 56 b3 9b a6 f5 97 24 b1 2d 51 2a 03 01 8a ac 98 09 c2 a8 80
3f 2c c2 00 d7 43 b4 cf 44 cc 38 5c f5 3c 97 4a ac 77 3d 51 59 38 a5 45 cf 35 6a ac 56 0c 1e 2d
7b af 76 6e dd dd 9b a1 73 9e 4d 30 b9 b7 5a d6 86 c3 53 6d 96 25 a7 ce d6 d2 55 87 62 e5 2c c6
11 62 54 91 6b 32 3f 25 48 3d 4c 44 39 29 5b 1e dc 4b 2f 3d 55 06 a6 d0 bc bb 95 35 d1 b2 13 25
e5 4f 93 72 26 3b 7b 5f 2b 85 9f 12 2d d9 db 77 75 2f 85 ce 5a 69 85 5c 1b 80 02 00 8e 1c 90 04
68 a0 60 24 f9 01 00 e8 29 78 b4 d0 81 e3 29 ab f1 54 e4 38 e4 8e fc d5 aa 9a e9 85 44 9e ee ad
d5 9a f3 6d ee 85 e8 6d d2 c0 cd df 79 b0 34 36 52 e6 49 e0 00 00 45 a2 94 b3 35 4b 75 5a df 3c
a5 13 ce d5 da b5 52 9d 5d 02 f7 5d d7 7a e5 6a 29 74 84 fe 79 dc db 93 80 10 08 a5 d1 40 4a 11
4b 01 80 4d d6 c4 01 d6 83 0f 45 c4 45 3c 34 39 f5 53 dd 3e 90 b9 ce d5 ba b5 5b 65 5c 6c f5 dd
d9 7b 14 02 55 2a 90 29 5b 6c fa 4f 0e cb d7 5b 23 6b ae 5b d7 0d e9 58 fb 9b 72 70 3e 1e 00 00
55 a3 65 33 b4 4b 75 54 18 fe d4 9a bd d5 da bb 57 6b cd 72 b1 b5 67 49 c8 5a 6c 00 00 70 13 33
b0 80 5b dc c3 c0 00 00 75 43 65 33 b4 4b 79 43 e5 1e 89 13 2e d5 ca 9d 2c c5 0c 71 f6 ae d5 da
b9 5b a9 76 c9 8d fe d7 db 35 74 32 d2 93 69 3b dd 6b a1 47 86 98 a5 da f8 51 95 ee 22 77 be 77
46 f4 df e0 00 00 10 0b 2d 42 00 d1 11 5c 07 d4 ab 00 3f 18 67 81 10 1e 03 00 f2 d5 38 e5 1f 06
8c be ce 47 58 c9 1b b7 d7 b1 f6 4e cd 8a 2f 59 40 58 f6 ee ad d7 9b 29 6d 2c 44 95 95 c9 dd 0b
a4 e3 29 e9 85 b2 e7 a9 0f 79 18 ac ba 2f 05 cb 94 32 ae 83 41 61 9d cd e0 30 16 42 33 00 cc 54
dc 04 96 ab 49 67 97 65 59 ae f6 ae d5 ca a5 76 74 7a 8a f6 be d9 db 37 66 ed cd 97 b2 55 b5 e5
68 b6 ec 9d 73 ad d5 26 07 0a 99 23 7d 67 b3 c0 70 0f a4 12 c0 58 16 9d 92 b3 87 b8 d5 6a 60 1a
6b b3 40 68 0e 03 ed 98 68 09 02 d5 3b 0c ef 17 5d dd 7b af 36 62 b7 59 a1 9e e3 95 a9 35 02 c0
78 04 01 2a 95 18 1e 40 1f 49 aa 00 ef 58 54 07 e8 a2 c0 3f 68 57 01 f9 2b 18 06 59 57 93 c2 f4
6e af 75 ed 5d ab 15 3d ea d2 11 5c 8d 92 02 ae 78 b5 6e b2 2c 4b 5c fb 66 6c bd 95 b2 35 ae b4
d9 e0 70 19 bd 4c c0 88 06 00 ce 78 c4 70 9b b3 cc e3 37 2e e5 dc f9 aa f6 3e c0 00 68 07 00 60
0c 0f 60 19 95 54 c0 68 02 00 d9 b9 d4 6d 64 dc 49 cf 77 26 e5 d4 d1 5c e3 c1 cd b7 56 ab 8c 71
5e 0a e7 9b ad 84 99 2e b4 f6 3a 60 fa 39 4b dc 52 18 c6 0e 88 91 98 20 1f 7a 12 c0 58 0d 01 60
24 02 81 e0 00 00 40 10 8d 14 00 7b e2 60 05 5c d4 80 1c fa 9d 46 c5 5c a4 25 b7 72 72 59 49 8d
6a 2d 9d 9b b3 6e 7e 55 2a b1 ac bd d7 b5 5a b7 e8 66 d9 b4 75 b6 76 be e6 dc bb d3 87 80 00 00
70 11 e3 ac 00 8c a8 de 02 c0 e4 6f 75 5b b2 4d bb 1e 23 5b be 8d 4f 57 79 27 35 e8 d5 4e 8c c1
d1 db 77 66 6e 95 ed b2 b4 d9 6d 23 82 f0 5e 29 c3 c0 00 00 70 11 b3 2b 80 73 9d 1c 04 5d 2a a0
23 38 58 49 26 55 b4 1c b7 b3 76 6e 4c e8 5d ab 1d 93 b3 ee ab 56 d9 a5 2a 14 b2 36 c6 d8 db 1b
9f 73 6f 4d e5 bf c0 60 10 ba a3 80 7c 58 58 04 a1 68 60 15 56 6e 49 09 d3 b3 ad cf 92 42 17 2e
e9 64 ec bc 95 b2 c9 46 6a 26 e6 c5 4d 9e 41 42 c9 25 63 dc 34 a9 0f 78 d9 1b 63 6c 2d 63 4e 91
29 db 0b 5a b7 86 70 b1 ca ef 30 5e 0d 25 ac 03 42 02 60 1d 41 b2 80 fc 4d 58 07 90 6d e0 80 10
cb 1c 00 8b 21 5c 04 a5 ae 40 15 c6 6d 44 db d4 3b bd 73 b2 76 4d 48 bb 73 a4 bc 72 d4 36 69 4a
85 2c bd 55 ab 16 2f 16 8d 0d 20 bd b6 bd b1 b7 84 1b 15 70 7e 0b bf c0 00 00 60 10 db 1a 80 7b
a4 d8 03 26 a7 ca a4 ca b6 83 96 f6 ae d5 8a a1 02 a7 06 93 e6 ac d5 aa a3 59 5b 29 aa f6 ae d5
ba f7 66 ad 7d de 59 a6 2b 9b ad 73 68 d0 4b 4d 3a de cb d4 d3 b9 3b 91 b6 f5 de bb d9 79 95 11
2d 11 d6 5e 47 4c 3c d3 66 17 b5 8d 0a f9 60 6d fe 60 10 9d a6 27 6d 2a 1d 1e 49 ba b7 55 2b 74
0e d0 d9 59 3b 61 6b de 36 a5 2b 57 5c 87 bb c8 53 71 97 b1 4c f0 ac e4 00 60 0f 6b 3b 40 a4 8f
4d 3c 2b 68 f7 26 a6 72 e1 6a 23 35 8a 20 d4 3b 99 ae 5a a8 30 d6 b6 64 37 b1 d7 b5 83 77 14 d4
4d e3 cc 00 0c 05 8d 2c b0 0e 01 40 0f 40 0b a3 99 c0 b5 4b ac 4a cb 48 f7 5d 2c 95 07 4d 11 69
0b 1b 11 f3 88 7a be c4 de 08 02 b5 47 78 0e 02 8b 59 e4 8f 63 c4 24 4c f9 2b 9e cc f1 56 ad 4b
28 92 bb de 8f 32 c2 78 d8 cd 26 36 42 00 0f a8 40 0e 3b b4 ab a4 59 c5 d7 6d ba fb 6e ee a5 f2
83 b1 b0 7d 11 81 97 13 11 33 13 e0 8c 54 97 69 79 70 b7 0f 50 10 9e 23 40 e5 51 73 cb c0 5b 77
55 eb a5 6d 9a 52 a1 4b 1b 68 f1 40 06 84 4c 00 94 d4 5a 07 24 a4 a1 41 da f9 b9 35 2c 1c d2 8d
13 cc cc e0 2e b8 89 54 fb 2c 95 b5 b7 35 66 dc df 6b dd 65 55 d8 97 c3 06 dc df 28 a5 6e c1 b5
27 97 67 5b 30 84 5a 49 5e 55 e0 70 11 e3 ac 00 8c a8 de 04 24 8d 1e 5a 1b 49 d9 f9 3b e7 0e ca
f9 c9 4e f7 ba b8 b4 b6 f4 72 69 4e 21 c2 ab 9c 96 76 89 61 bb 66 25 96 9a 09 f1 d7 29 a2 b7 55
46 52 1e b0 00 50 09 82 c5 a6 a4 59 c5 d7 6d 9a 37 55 6d 45 e8 64 34 c3 cc 63 7b 6f 55 ea 4b b4
bc b8 5d 65 22 a8 68 93 35 5f 18 a2 29 45 64 03 c0 00 00 40 0f 8c 03 40 58 07 03 51 ba 39 0e b2
dc ec 9d 93 72 4a 55 bd 9d 43 16 8f 99 b4 2e b7 01 b5 3b cc 08 80 d0 2d 52 dd 72 b2 f5 da 56 8d
55 2a 4d da 76 c6 f8 e0 fc 19 81 ab ea cb 4e c5 e0 bc 11 89 6d d0 88 ce 43 e2 5c 7c 70 0e e3 92
00 7c 61 a5 1a b1 34 ca b3 6e d1 6a 19 c4 a9 45 99 d6 a9 d9 5b 35 53 e4 b6 8f b1 bd d5 9a ef 64
ad 7d ce 9e b2 09 3b 9d 84 30 14 f1 89 94 e2 00 00 80 30 b7 55 01 a7 45 a0 0b 2e dc 40 3d d2 5e
00 00 50 0e e3 92 c0 58 12 85 37 ad 25 7b a5 61 ac 55 8b 12 2b 13 4c ab 36 b4 ec e2 54 a2 cc ef
55 6c 9d c6 03 94 76 58 0c 03 f2 77 68 0b 01 00 7e 8c 42 03 fb 71 10 1f d2 98 80 eb 49 bc 00 00
60 1f 6a b2 c0 58 0d 01 20 4a 16 4a 65 d3 1b bb 8f 71 ee 3d 47 0c 44 39 64 73 71 ee 3d 4b 2b ba
b9 4d 5f b3 ee ad d5 bb 31 6b a9 0c b4 f1 2c dd 5b a1 82 a2 8e c7 72 36 e0 bc 11 82 a2 aa eb 32
b6 e1 e0 30 08 4b 46 26 9b d4 d2 67 97 89 a7 14 34 ca c1 6a 69 b1 11 76 b2 58 9a 90 43 4d ac 9e
67 c5 57 92 98 55 e3 2f 55 eb bd b5 b9 f8 3b 2b 49 48 6e f0 8e bf 11 41 4c aa a7 80 35 a1 76 44
54 24 cd 69 15 33 32 ee 0c 59 99 53 99 bc 93 93 c6 4a c9 1b 9d ed 5a 9e 78 34 67 cb 66 e7 dc ec
00 cc 59 70 14 7c 9d ab 8b dc d8 67 31 8b 1d 36 94 8c 92 66 dc 5e e7 cd 3b c7 76 98 ba f0 cb 17
54 8e 4f cd 98 c5 cd 0c 14 d3 35 d0 ed 1a 93 b9 03 c0 00 00 b5 58 34 2a 24 4c f5 55 1b 0e 8a 9b
55 55 18 bc 22 56 75 80 7e 28 b1 02 20 40 0f c9 96 20 44 05 81 88 60 e8 b7 2b 8d e2 57 4a e5 92
63 99 9b 33 de 8b 54 eb 4c a3 db 52 f5 52 b7 ee 24 a6 93 76 b5 fb 69 38 b4 e5 af 7c dd 22 e2 0b
7b eb 7c af 6c cd ea 6a 19 2b dc e7 36 a7 a9 22 f1 c2 5e a9 93 68 bc 6d a7 a8 96 d7 80 03 00 46
bb 72 ac 3d e4 99 2f dc aa 8c 79 28 5b b5 ac ab 1a 8a 09 dd 3b 5d 6b 96 f0 a6 53 ce 55 06 65 a1
bd af b3 76 6e cb c9 66 bc aa 42 b1 f6 6e 51 27 e3 ba 35 6f b6 e3 18 c5 96 5b 29 94 c1 f2 4a 99
e8 54 30 74 18 ae 6e 94 80 16 0b 54 01 fc 9a b8 0f d1 d9 c0 7d a8 ca 03 f1 44 50 1f 91 91 80 fb
0d 08 07 9a 6b e0 45 c8 33 53 cb b6 a5 76 12 c6 aa 6e 4c 4f 88 5b 04 a6 d7 53 d6 34 80 cd ac d4
b2 9b 8a 54 fa f5 24 a2 e2 b7 3d ad 4b 18 b9 b5 57 6f 13 45 ee 6b 53 e4 c5 55 64 31 d5 6d 6d 75
9b 0e 50 de db 5f 47 43 14 37 96 b8 3e f0 65 f3 32 a2 34 bb 1d 75 14 e8 8f 1b 04 63 3a c3 20 b6
57 c0 00 00 35 56 7c 31 b2 6f 5f 57 6b ed 1a 9e 52 76 46 9f 66 a5 75 55 9a 6e 99 0a f1 3a 85 31
72 18 08 25 55 00 2a 2d 2e 02 6e a6 54 6d c8 d3 c8 cf 2f 56 2b 76 4a 1b 09 fd 8b 1f b3 95 b1 ba
f2 d7 d5 a5 9d 6c fc be 11 b8 28 6b 65 ef 86 0a 4a c4 ad 42 9b 82 af 6b 51 64 9e 56 0a 02 ec 66
60 1c 2b 32 00 7a 2c 48 02 22 ec 00 1e b5 fa 6b a9 98 37 79 54 a7 0a 6b 5a 5e 55 18 b9 b5 57 6e
d5 d5 90 c7 55 b5 80 1f 9b b4 c0 90 1f 8b 33 00 a8 f5 41 80 56 8c cc 02 d1 52 e4 72 93 95 b5 67
3b 27 25 9c 9a 99 67 3b d9 fb 56 ab 8c 5a 9c f6 33 1b 40 01 40 45 2e aa 03 89 68 50 18 3b 38 78
74 72 01 cd 4c b6 b1 44 22 ca d2 df 3e 4d 44 35 15 34 df 72 6e 4c cb 99 ad 44 8a d6 92 93 7b d5
6b 1c 88 95 74 9b d6 d8 5e 96 54 b3 ad d9 f7 b6 f5 e1 e0 b5 72 7b b2 b2 7b 3b 3e a4 64 de 9d b2
a9 b9 73 2e 65 d5 0f 84 cb cb 68 f3 56 6c cd 97 ba 97 35 ef 0e ef d4 9d d3 b9 d7 ad ef 2e f1 8b
7d e6 02 a0 37 58 07 00 a0 78 b5 92 83 2a b9 ef 77 26 26 7c 1a 59 97 75 29 11 76 77 7a c4 46 4a
cb 09 91 76 75 99 c8 d2 86 51 9d 9e 39 76 aa d6 a6 17 3d ae 6c dd b9 ba d7 57 52 63 0f 6b 9d ed
bd 77 a2 f6 b8 6e 0a a3 88 b9 ae 2b 82 a8 d9 af 4b 56 28 a6 c2 4b d6 a6 57 47 76 78 00 00 03 4c
65 44 4b c3 6c d5 11 55 52 d1 1b 95 71 83 4b 34 79 7b 57 6a e5 51 7c 54 ac 4a fb 57 2a 8b a2 e5
22 57 ca a0 f8 a7 78 91 f2 b8 42 29 dd e2 ba b1 f0 8a 77 78 ed ad 7c 1e 21 9e 69 0b a0 f0 b6 98
56 42 e8 3c 2d a6 15 6a b9 d0 0c 68 95 5b 31 5b 4b 31 6d 40 bc 8f 91 72 f0 d5 74 43 b4 d3 4a b5
5c 9e d2 ed 9a de 54 c7 bf 62 24 bf b5 76 ad d4 8b 48 fb 75 48 b6 ba c4 be ba 10 e5 77 ae d6 1e
d7 ca 67 6b ad b2 b8 da bc 32 79 5c 70 5e 03 c3 c0 00 00 75 92 4c 33 ac a9 31 6c 9e ec ac 9e cd
50 f7 e5 b2 c5 77 74 6e ad d5 bb 33 6d ad 64 2f 55 1a 5d 7b af 75 6f 85 f0 78 5b 4c 2b 21 84 1e
16 d3 0a b8 e0 ec 56 d2 cc 5b 50 37 15 e2 7c 4b 97 80 75 0e bd 2a 4b 27 6d 5e 23 2c e0 49 5c 55
07 ca 9b bb 97 15 32 8a ef 21 27 d5 4c 7b f2 b9 db 7b 56 aa 63 e7 e4 9a 5f ba b7 56 ea cd 79 b2
d6 aa 97 aa 6e f5 95 c9 c5 61 b3 3c 21 72 37 52 e1 2e 58 5a 9d cb b1 cc 39 d7 29 ef 48 cc 93 55
cf 73 4b 04 27 3c 00 00 65 d0 7c a9 bb b9 5d 73 28 ae f2 12 7b 56 c7 bf 2b 9d b7 35 da a7 27 75
d8 66 f9 a3 33 e6 6c cd 79 ef 46 2a 6b b2 99 12 db ab 21 0b a6 53 96 f2 ea 42 e5 1c e6 bc ba d2
c9 26 29 af 6f de 14 c3 30 55 0b ab 5a b0 5b 8b 29 6d a2 bc 17 82 30 5b 8b 65 61 a0 ac 16 e6 3b
87 71 ab 06 a9 ca ac a2 29 c7 c0 84 54 81 92 ed d4 eb 2d e5 cc db 9b 73 ee 7d cf 8a 2b 44 21 ba
d2 66 ab d5 5a 8c 01 9a 9e ca dc 86 67 e7 ac f8 d5 17 8b 98 4d f9 31 46 24 c6 15 99 cd d1 ba 33
55 ac c5 99 02 2b b8 fb 87 70 01 00 cc 4d 54 06 f7 30 d0 14 07 39 72 80 35 96 95 01 76 b3 00 80
7e 09 bc 04 00 d4 b4 72 8b 4a 4c ea 35 9c 9a b2 63 5b bd 66 a4 b4 d0 93 2e d5 d8 f3 22 6b 50 08
14 54 28 bd 1c 82 6b 4b 9a b7 5e 2e bd 98 c2 6f 4b 3c 1f 92 f2 3e 9e 00 00 e5 78 83 38 c3 6b 7b
5f 2b 95 a5 85 a2 d9 ab 23 08 63 29 97 6e bd 56 f5 d4 9b ec 77 95 42 f2 ae d5 1d ed 1d 9f b3 ea
77 ca 91 dd eb ce b5 b1 b4 78 3b 1c ed dd b9 be 98 3a 92 94 8b 62 76 2d 8a ba 92 59 5d 93 20 6c
e7 32 45 66 c8 23 40 d4 89 19 36 28 cc 93 23 46 4a 83 2c 24 4a fc 84 d2 a3 3b 2c e4 e9 33 e6 ef
0b 7a 5e d1 da bb 56 2a 5c b6 21 5f 3d 6a 94 ff c5 73 96 56 ca da eb 91 0b 8d 02 ae 2a fa b2 09
d1 1b 6a 03 a9 a3 68 0e 02 e0 64 58 06 03 83 56 30 1f 4b 4f 80 c5 13 b2 b2 3c c6 f9 45 2d 2c 8a
fa de cf c9 a7 69 65 59 ce f2 6a 52 95 2a b3 bc 9a 93 a5 5a cc e7 27 a4 69 51 9b 37 ca a7 2a 35
66 ca 72 b9 52 59 15 73 9a b2 33 a6 26 6c d6 2e 93 ea 85 d3 37 0b d7 05 65 67 ad ba f5 40 e9 1b
1c 0e bd 30 3a 47 b8 d3 2f 4b 8a d9 72 34 ac 16 d1 d6 6c 89 27 13 e5 0c a8 a7 44 78 d8 2f 14 61
28 ca ba cd 6b c0 e5 53 9c a3 33 d6 f9 54 eb 0c c9 15 ce 55 3a cb aa 4d 77 b5 76 ae 55 1b bb a3
54 57 b5 76 cd dd 6b 9b 36 56 49 a9 c2 f7 2c 32 ad 4b b1 be 38 3f 16 e4 eb 8e c4 fa 57 32 3e e2
00 34 d4 20 0c a6 21 20 24 03 80 50 3c e5 14 9c 32 c2 f7 39 45 29 2c 8c be 4e 4f 4b 2b b1 bd 93
b3 4a 73 3d ee 17 99 53 a1 c0 52 0e ee 02 c7 47 c5 0f ca 2b b9 eb 73 6d ef 86 2c b8 55 43 68 99
93 2e 6b 6f 11 b5 64 b9 1b 63 ba 8f c0 00 00 d4 f5 93 39 3b e9 79 3d a7 2c 4e fa ee 4f 1a b2 ac
c5 d7 b3 5a 56 e0 de e4 b5 57 9a f3 da a7 6d 91 ec 76 b7 aa f5 de eb df 5c 2f 85 b1 6d 9d 17 cb
6a ec 9d 95 58 c2 da 33 37 68 66 d5 a5 4e d1 d5 66 63 c5 27 80 00 00 e5 54 ac a9 c3 eb 7b 5e ec
c5 93 b3 12 d3 eb 27 5c 6c b5 91 8a 03 db 58 a0 05 c0 2a e0 2e 01 5c 51 93 c3 aa 43 d3 93 c7 71
0c 58 e4 ec fd 9f b3 f6 6d 4f 3d a4 24 5b 68 40 0e 2c c0 04 01 e9 53 80 12 8a bb 40 a5 09 a4 bb
33 69 2f 56 2f 85 99 91 6b ab 5c 21 84 24 17 35 d1 78 62 b6 bc 6e 33 29 98 b1 c9 5d 2c 8a 0f d4
d1 a4 52 35 c8 f7 37 67 ed 1d ab b5 76 bd 5a f8 c4 1b c6 4b 77 5f 17 01 16 b2 3c 06 80 b0 16 01
40 f0 50 17 64 2b 40 e4 34 94 21 b3 7b 3b 27 65 ac d1 9d 80 51 50 2a 02 89 78 68 02 00 4d a8 00
50 08 63 3e 23 bc 5c c2 6f c9 a9 31 26 30 ac ce 66 5c cd 79 ed 56 2c c8 11 5d cb d7 7b 61 74 d6
69 59 30 17 de cb ca e8 a9 88 6d 37 06 e0 bc 11 81 d6 d1 69 2d a3 60 85 2d 6a 43 8b c0 30 15 86
2a 40 58 0f 03 51 ce 90 c6 8d e6 ec 9a 92 54 66 35 9c 66 25 6b 99 aa 36 21 39 d4 06 66 ea b0 12
01 c0 00 50 17 2c b2 80 bc a5 25 0d a0 c8 51 9e cc d1 7a e1 75 5e 86 93 1f 47 de dc 1b 9a f1 f0
30 0d 3c 39 80 9a dd 94 05 5a cc d0 39 0c 20 ea b3 35 ee 47 58 cb 1a 2e b3 91 5a 6c 87 6f ed b4
75 ab 81 db e9 55 22 d2 d0 34 8e 43 4c 47 54 cb 3b 71 14 1c f5 2f 2c db a5 5a b3 13 93 29 2d 56
1e 8a 2d 57 cd 57 58 39 84 3e f7 56 4e 52 80 f5 76 d5 b1 9c b0 c3 cb f9 7c 6f 2c 70 6f 7b 5e e8
c9 1d cc 9d d8 b6 e8 93 4b 37 66 0e a1 34 bb 49 19 8c 2d 08 4a ae af 30 0c 6a 3a 80 8b 15 5a 03
40 e4 8d 84 6b 9c c9 39 3b 6d 70 c9 52 cd d1 ba b7 56 ea d5 71 c3 d2 34 49 37 66 6d a5 b3 b2 a3
bc e5 2b 6d 6e 85 c6 73 4c 25 c6 1f 7b f0 6e 16 c9 ce 65 67 9c 9d 70 95 d1 9b b2 03 54 fc 30 0c
6a 3b 00 8b 15 62 04 40 d4 2d 84 6b 9c c9 39 33 6d 70 c9 52 ce cd d9 bb 3f 68 e5 51 c3 d2 34 49
3b 67 6d e5 d3 b2 a3 bc e5 37 76 ef 85 e6 73 4c 25 c6 23 7b ef 7d ed c1 78 1f 0f 40 12 53 aa c0
68 0f 03 92 c6 4c c7 51 a4 e4 d3 9c 39 43 5b 7b 37 68 95 51 70 44 0c d9 2f 56 6c c5 b4 6b 92 34
5a 33 75 ae 92 cd aa 6a a2 e0 37 7c 87 80 f0 1a 02 40 28 1e 40 14 53 92 c0 68 0d 03 91 ce 50 8c
f1 6d ec 9d 96 d3 52 ae 84 6f 5e d5 14 a3 9b aa 74 f9 5d 60 a4 ad 59 9d d9 9b 69 74 8c c9 57 30
96 01 ee 53 c4 08 81 10 12 00 00 28 0c 35 d9 80 2d ed 04 01 ea 75 f0 10 10 5b 20 c0 18 07 01 e0
42 26 61 96 a8 a6 db 9f 73 ee 7c d1 8a dd 79 25 b6 77 da d7 e1 91 29 9d f5 be b8 30 0f 51 8d 80
79 e8 6b 00 80 7f 6e 22 01 66 a9 00 50 16 83 3b 00 db 51 fc 20 11 74 32 40 58 0b 01 e0 62 39 41
d5 6a 75 9d 93 b2 f2 78 c2 25 56 37 9b aa f6 3e d5 de 00 0a 02 aa 74 40 16 6d a7 80 00 00 10 12
32 b1 40 50 14 7b ba 21 44 33 a4 0a df 59 2d 29 e9 98 c2 b6 eb cb 89 95 47 86 a3 bd de 72 00 b9
1a 38 09 c6 56 52 52 30 84 f9 ed bd 18 b2 d1 73 18 8f 64 ed 7d ce ba cd 74 38 8b 5b 70 4e 3e 00
20 0f 4b 2a 80 92 9d a1 29 a7 56 f0 c9 fb 4a 88 d7 99 9c ce 53 1b 15 c5 0b 73 a4 c7 d4 4a ba df
71 3b 24 ec d8 ae f5 00 58 25 81 00 a9 49 c0 01 c0 51 ec 6d 01 00 51 c4 b4 82 8b 44 7b 9a 56 24
b2 d0 92 6a a7 89 b0 b4 23 99 a9 da 7c 2d 0c ea aa 74 a2 e7 23 5c 98 9d 2b a1 c8 a3 21 99 2b 2a
72 47 e8 65 ca db 2c d1 f6 19 92 ba 6b 34 ac 98 0c 2e 8a d4 df 9e 16 8b cc b5 bb 58 65 a2 f2 2d
6a 52 d3 68 bc cb 5c a4 a6 52 2e 3a d7 31 22 24 8b 92 b5 bd 58 71 f8 00 00 10 0b 44 21 00 8a 58
d0 05 96 c4 c0 2b 0a 72 00 3e d3 d2 54 ba b2 00 24 d7 96 12 97 7b 26 11 55 84 c5 dd e9 83 5b e5
39 b5 58 a2 d3 6b 54 7c 4d ba 2c de d5 ae cd 58 a7 77 a5 ac a9 e4 4a 5b 65 72 e4 b4 d3 31 b2 00
93 b1 0c 01 00 f3 ba 88 07 99 d9 00 4d 2c 46 01 50 55 18 1e 01 01 14 aa 10 0a bd c5 c0 4a 12 f4
95 5a 39 98 29 ae 26 5e 16 9e 06 e5 a9 99 0c c3 59 8a 6e 7d cf 8a 1c eb e3 49 a5 e6 ac d9 9b 69
73 d8 b0 c2 ae f8 dc eb da b4 a6 16 93 be f7 de cb da b3 b5 07 77 3e f7 df e0 00 00 60 15 84 a3
c0 88 11 02 92 af 49 2f 42 f5 ac d7 9f 95 3d f5 70 4f 2f dd 7b b2 d7 be 41 6e 96 ac 8e 0d bd d7
c9 b2 cc a8 98 03 01 65 59 b4 0f 55 7d d5 09 83 2b a0 06 e2 61 8a b7 47 43 36 9b ea ab ca 10 81
76 ba a2 d4 82 26 7d 66 e8 dd 1b ab 75 6e cb 5b 49 a0 a5 4f 6e b6 ca e9 ad 4d 7a 2b 4e ba 2b 33
9e 88 4b ee 5d c9 bf c0 55 a8 24 54 6b 27 6d 6d 9e c8 8d 3a 4d cf c9 e9 25 66 7c ae ee 7c 4f 4a
b1 b3 64 6f 33 de 7b 51 59 28 b4 5d 4b 35 66 bb d7 7b 71 6d 5e e2 b0 f5 a9 5c b6 8d 59 5b 71 b7
36 e6 5c 76 a5 62 42 8d b7 2e e4 5c 46 4c 43 c4 6e 57 fe 44 6c 88 af 60 1c 38 13 00 fb 94 9c 07
de 6a c0 3e b3 53 01 33 c1 78 00 00 c5 75 52 ab 03 2d 35 2d 64 e8 8f 1a 4d 4f 19 b4 a4 4e 2f 74
6e 8d d5 ba b3 5d 2c 75 56 1a 6e 96 fb 40 00 14 05 55 0c 40 12 f6 60 f0 00 00 c5 79 9d 4a 99 38
f5 26 a0 e6 ab 2e ce 49 67 42 1d 4e 8f 93 c6 8e a5 35 9c e4 f2 ab 42 ae c8 f5 3c 6c d2 31 93 60
80 14 9b 34 a1 b3 e4 c5 5e d5 ca a9 67 45 88 d2 72 ca c9 cd 22 74 7c ba f3 63 17 8b 15 ae b3 d4
8e e6 25 2b e4 e1 4b 98 49 b2 f4 b9 36 22 52 6c c1 ad 1c 84 a4 c2 b0 63 07 68 29 30 cc 1c c0 bb
18 72 23 03 35 ae 8e 25 4c ca 75 63 44 27 32 b2 9c 14 9e 2a 19 e0 00 00 56 06 44 5c cb 26 e5 76
68 ec 05 5a da dd 8b b1 6a 90 32 db ee fc 5f 5b 41 a3 97 de f7 e2 e8 49 d9 60 fe d8 b6 51 48 88
4d be ac 7c 16 9d 9b d9 8a d7 17 c7 5a b5 52 a5 41 b5 d2 6c 0e a9 31 6b 43 ab 53 a6 62 19 a8 dd
da 20 18 94 43 01 28 2a 08 05 c2 54 3c 55 a8 2c 64 62 a7 6d 66 62 66 6a f2 5c 51 86 1a 24 d6 ab
53 d9 6c 87 59 a2 d4 d3 43 ab 6d 56 71 2c 1a cd 1b 59 2c 4c c7 b3 ce d5 4b 13 ba 2a cf 39 59 c5
0d 74 33 5d 66 71 53 9b 2e d3 1e 2c d9 7b 6d 74 52 ee ae 52 ba dd 3b 99 7a 99 0a c8 d5 b7 61 14
42 a2 ac ad b8 a0 0f b9 9d 40 7e 0a ea 03 f0 54 30 1f 71 a9 00 fa cc 84 07 9c 4b e0 a5 72 75 1a
3d d6 b9 45 d8 ae 55 13 6e 55 38 2c 34 9c d3 95 41 6f 8d 2b 3c e5 4f 5c ec 39 5d bb 57 6b dd 99
be 59 31 55 d0 e3 2d 8e 4e c9 6a bf 46 49 6a f3 f0 c5 57 62 b9 54 4d b9 3c e0 b0 d2 73 4e 51 05
be 34 ac f3 94 3d 73 b0 e5 76 dd 5d af 16 b1 55 d0 e3 2d bd d3 c2 38 bb 15 aa fd 19 25 af c9 99
3c 8f 32 80 fe 6e 68 03 88 23 20 1f c3 8a 00 fe e0 48 07 aa 42 60 2e 84 8f 25 42 75 dc c5 a8 a1
5b d8 a6 ab 6e 4b 4f 19 41 a3 cc cf 93 c8 ab 31 4a fb d4 ed 75 fb a0 5c f7 35 a7 54 6f 94 83 67
3a 2d 52 a3 b6 60 9a bd 55 05 cc b3 1e d3 56 db 34 86 83 ad dd d6 be cf 63 4a 9a 8c 31 a4 0c 5b
61 92 dc 97 92 31 59 8b 2c 22 36 9c 4f 93 72 56 50 9b 42 ac d4 7c c5 b1 94 1a 3c cc fb 37 28 9a
16 d6 52 6f da bb 57 6a e5 4d 63 d4 bc 5c 37 5d eb 55 64 3d 4b c4 b8 3c 00 00 85 91 63 a9 c3 4f
b3 57 28 72 e5 65 d2 df da 37 47 68 50 00 ea cc 2d 6b 1a 15 f2 c0 08 a1 10 7b 70 49 5b e7 84 84
ec d2 24 4c 65 11 3a ee eb 2f 1a 42 8c 4a 5c 14 ca b1 62 46 5a 4a 38 67 b7 ae f8 4f 15 e4 00 80
23 47 44 80 40 29 37 4f 95 92 87 2a 89 b4 f5 44 e0 aa 94 96 de 4f 18 1c 4c 3c d3 b2 4a 13 52 5a
a6 99 4a 84 56 86 96 74 a2 62 34 c2 95 3c bc 38 e1 54 5c df 0a 76 ce d7 db 37 76 ee c5 ef 8b 72
ac 45 33 7e 2e 7c 63 15 9e a1 8b 61 26 a6 68 a9 da c8 4a 0d a6 29 74 ad b2 56 aa a6 1f 80 00 00
45 04 5d 54 bb c8 f1 54 e0 c8 b4 ba de cf c9 df 36 87 88 ce f2 87 ce 6d d6 2b fd a3 b5 66 ba 5c
f7 5a 52 43 4a 77 ae f0 04 00 c2 43 98 0c 22 20 3c 20 0b eb 44 a5 5d 1e 98 31 5d 2a 13 44 e4 75
9b d2 85 54 fd 65 1f f6 99 76 3f 4c 4b 8e a6 7c d2 d3 6a 45 b9 b7 3f 28 a5 99 16 23 3b da bb 5e
ad 94 c9 de 76 29 7b e5 7c 62 2d 35 70 49 5f 17 92 dd dc 11 d7 c1 88 77 c7 00 f0 00 00 85 50 a2
62 03 15 65 53 66 1c ad 4e ec 50 fa 0e 1c 3b ef 53 c6 e3 47 0f 34 e4 d2 aa a2 3c 69 b9 35 ee c4
4d 3a 3d cf da 37 56 eb d5 77 c3 98 ae 66 f7 66 2d ae 98 cd 2f 47 0b 9c d0 d8 68 6e be d6 5a d6
07 bb c2 99 16 9a 4d 92 14 c2 45 ab 50 de a4 b8 dd 6c 22 e4 71 32 24 5e f8 a2 32 db 6c 3c 84 f8
7b 28 b4 7b 37 3e a7 9c 9a 99 67 3b b9 f7 46 ea d5 71 8b 53 9e c6 73 64 ed 40 19 3c 2a 80 fc 15
96 02 c0 50 1f 7c 2b 00 fb 62 18 07 51 0c a0 26 ea 31 00 b9 2a 00 01 2b 83 21 da 9d c9 e0 00 00
e4 f4 83 aa d5 64 f9 44 e2 aa d5 79 5e 51 49 31 2e de 4f b5 6a ab c1 88 ae bc 75 ad 35 80 38 0c
c4 d5 c0 6f 73 0d 00 c0 73 97 28 03 59 69 f0 10 0c 6b 52 00 7c 16 a1 24 a4 ae f1 71 4e c7 c9 69
47 65 9a ae 76 5e cd c9 e7 59 74 59 aa f6 ae 57 3a cd a3 46 57 b6 76 cd db 7b 59 79 d7 28 9f 15
36 5e 86 b9 2f 4c 4d d7 d1 4e 88 f1 b0 9d cf c5 57 4a 8e c5 1a 9b f0 00 00 10 08 6b c8 80 63 0a
92 02 c0 a4 72 b3 ca c6 38 fb 2f 65 e4 b2 b3 e1 4c c9 7b 37 67 ed 1d ab 76 6a c7 d5 e1 9e a3 97
b2 f6 32 c3 45 aa 93 22 4d b1 76 2e c3 5a 66 35 5a c3 86 17 4a 28 8a 0f 54 6d ab b3 c0 00 00 10
0e 7a 31 c0 58 18 92 58 77 63 e7 3f 24 9e 26 1c 6a 4b d9 7b 36 66 84 ea 72 5b 27 aa 99 39 1b 15
a1 0d 06 d0 da 13 50 0c 03 00 86 25 1c 0c 59 0b 44 23 d5 73 76 6e cc d9 6b 17 6a b5 58 cd d6 c9
58 9c 55 a2 cd 8d f6 32 d2 ba 15 e6 db 50 be 0b 63 5d c6 cb ec 95 94 c4 a8 d1 df 65 5b e0 14 eb
2e e8 d7 1b 1f 73 af 85 5a 0a 23 2b 3c 5d 8b 9c d1 13 11 87 e2 bc 40 00 10 05 de ec b0 08 04 92
ce 30 3c 00 00 10 0d 6a 29 80 94 5c d4 04 ad 0c a0 1f 03 6e 4d 2b 3c ac 63 8f b2 f6 5d 4f 2b 3e
14 cc 97 74 76 ae dd bb b7 7e 30 7d 5e 19 ea 39 7c 27 8b b1 63 d6 99 76 36 a0 2e 02 94 80 60 3f
05 41 01 f9 a9 88 0f 39 54 3c 30 10 8b 22 40 28 08 02 30 72 84 af 9b b4 cc b5 33 2e e5 ec bc 9a
58 78 46 b8 ef 67 ed 5d af 55 c2 b1 0c f1 b5 dd 9a b1 f1 88 5a 4d 66 ed bd b2 b9 d1 56 34 3b 5b
ee 5d c9 bd b7 b6 f5 de 01 e0 00 00 01 01 06 9a 10 08 bc ed 00 46 26 b8 82 71 86 63 7f 67 63 ec
7d 93 93 4e 90 cc 73 b4 ed 5c b2 52 77 75 8d 67 6e d5 f1 8b 3b 34 69 11 79 5d 0a f3 6d a4 e0 8c
0a e9 57 9b 4d 23 03 2a 5e 62 1c 68 c3 09 65 57 77 5c 2e b3 f1 28 e9 c2 cb 62 f5 64 5b b0 c2 f9
be 59 cd 5c 9e 10 10 4b 11 00 8a 94 c8 04 da 49 29 6c f5 42 a8 d6 6e 4c c7 b8 f5 25 67 08 12 f7
7c cd ba b1 5c 6b 4c 32 2f 7c db 7b af 84 b1 7e 59 13 12 7f 1c 57 91 f2 2e 5e 00 03 01 06 9a 10
08 bc ed 00 46 26 bc aa 52 79 66 5b 6f 6a ed 5c aa 6e 5b b7 49 ae ea c5 11 15 d5 40 4d 2f 35 e6
c4 ce 63 f4 30 cf b7 46 eb b5 c8 a5 62 31 ea eb 84 30 44 f2 c9 d2 e4 ec 0f 8a 96 19 b9 2b 13 df
1d 19 9d 6a c4 f4 c9 55 47 62 31 4a ea 66 d4 31 e0 10 0d 72 29 00 84 15 90 04 66 8f 69 a3 66 63
2a 69 e2 37 55 a0 9b 62 dc 8e 29 8a 97 2b 9f 65 ed 18 ad ac ad a4 45 ee 6d 95 e9 8d eb a1 cf 15
81 a5 b6 aa 72 b0 00 03 00 d8 bc 0c 00 f0 01 80 20 1e f8 a2 80 70 29 03 f1 70 ed 31 5d c3 b8 b7
17 62 d4 51 d4 d2 ad 47 b3 16 23 a5 15 99 ee 97 99 2d 3d a8 a8 93 3a bb d5 8a a9 58 32 97 ba e2
b9 5a 4c e1 66 f8 b2 36 94 58 49 bd ad 7e 21 0a ca ed 5b 65 73 b4 89 30 5a b7 de bb d3 81 ef 2e
08 bd f5 52 57 47 8f 80 00 00 10 0d 74 32 00 7b e2 94 04 54 8a b0 16 05 21 9b 0c e3 66 03 c8 e9
37 34 98 d2 76 4e cd d9 ed 64 9e ea 52 fe 40 00 20 11 4b 2a 00 d2 5e 10 06 10 ec 1e 00 00 10 0a
8e c9 00 6c 2a ca 01 00 54 76 49 12 5f fa 4a b2 eb c5 78 ef 1d a4 5a 77 d9 81 5f 99 6f 35 ea b5
6f 44 5d 78 07 6f 65 ec b5 b0 81 c4 c0 b3 6f 75 6e ad d3 ba 77 3a e5 b4 75 d2 1c ac b5 52 b6 85
4b 12 2e 5d 2d 91 ea 19 e0 01 01 08 b2 14 05 80 c0 23 07 26 01 47 dc a8 08 45 5b 51 be d1 89 11
15 c4 8d c4 5a 44 c5 73 26 24 6d e7 11 e6 2b 89 5b 4a c4 58 8e 5a 67 46 f5 51 ac 96 a9 b1 be 44
6d 65 2b 65 2f 4d 5b 49 5b 69 72 e5 9a 8c 76 b9 dd 0b d3 4d d3 73 d5 b2 f4 63 6c a1 2d 4a bc fd
b2 57 85 59 2f 37 6c 98 a1 c4 3b c3 c0 00 00 20 09 d3 b2 40 58 16 91 b3 aa 44 b5 9e e5 dc b8 96
50 7a 43 89 5e 65 74 a6 7a ce 35 b4 db 2a 67 95 51 9a 42 83 cd 25 5b e6 70 ad 16 fb d5 9a b3 5e
ac 7b 9d 9a a3 c7 9b 67 7a 2e 23 65 e2 69 b4 6b 48 ca 88 88 8d 9a d3 a9 5a 1a 55 72 c2 2b 75 47
3f 4a 30 8c a5 14 6a d4 8c 27 08 56 37 70 93 08 aa 5a 1a 5d be 00 00 00 00 28 07 02 d2 37 68 98
a9 1b dc 98 91 73 b1 8f 37 6d a4 5d 64 a3 86 e9 99 2b 2d a7 5d 95 aa 92 7b 7a 35 5d a2 e6 31 1e
cd d7 bb 33 6d ad ad 18 4e 27 b9 5b 5f 72 ed 00 f0 00 00 c4 90 7b e4 41 dd b7 37 68 ed 7d af b6
72 c8 4d e9 da 2d bb bb 57 42 6f 0e f1 9c b5 f1 a1 db 3d d3 29 7c 26 ab 38 d0 a9 e0 ec 1e e4 9d
95 61 b2 f8 3c cb 25 d8 69 ac f5 87 80 00 00 94 90 7b e4 41 dd af 26 64 cc b9 97 12 c2 6f 4e d1
6d ac d4 99 ae 89 75 4d 24 26 43 ee 99 4a d4 c9 d0 bc b7 52 74 b2 a5 23 35 98 5e 72 aa 05 8d 57
19 ed ac b4 19 8e 8a fd 9b 75 6d e3 2c 52 9b dd 5d 78 cb 9c 25 f7 17 56 12 e7 09 b7 b5 f6 75 49
b9 5d 2b 7c b0 8b 99 99 06 c8 ec 22 f2 84 57 97 b3 06 b5 64 96 ac ca c5 2d 76 56 ab 62 31 43 16
d6 2d d6 6c 98 93 a9 8a 52 1b 26 c1 95 4e 1d be d5 14 ab a1 b4 79 3b 47 68 ed 5c ad b0 6a 78 8e
9d ec 5d e6 02 11 75 48 09 00 a0 18 08 b9 98 e0 14 0f d5 54 8b 23 45 f4 b9 45 9e e8 b1 1e 1e 51
49 32 33 5f 67 94 52 6c ca d5 9d ed 5b ab 95 46 0d 0c f1 94 c5 71 7a c4 3a 66 e1 5c a0 18 92 eb
18 d7 1a dd 5b da 1a d2 ce 27 d8 bb 11 62 94 93 4d 2e 32 db e0 d5 96 9a 1a cc 66 b9 55 63 46 73
19 ae 55 69 5a 22 b7 8f b5 76 ad 59 38 d9 33 5e 0f 39 43 37 35 8e a1 e8 6b d1 1a b0 ee b4 ea e5
03 01 62 9b 10 0e ee dd 20 2c 05 80 80 39 28 b4 80 50 08 05 63 50 20 17 06 3f c4 50 82 b3 c5 64
fb 3f 27 9d 19 88 f6 47 b9 b5 34 a2 ca e5 b1 b7 4e 96 ad d1 69 51 15 12 13 36 cc d3 60 1f cc 2c
c0 80 1d 94 a5 aa 84 16 0e 27 c7 9b 2f 74 ae 6c 16 11 ef 55 3b a5 73 22 b0 91 92 95 00 74 6b b0
0d 0d 11 a0 28 0f bc 91 60 07 80 b5 56 7b a9 d4 66 b3 56 ea 85 71 70 49 b2 69 63 5b eb 7d 8f b5
80 7d ea ca 03 f3 68 58 09 00 14 05 21 72 a0 25 18 4a 4a ca b2 bc a4 db 12 ba 70 cf 25 35 c4 b0
95 3b c9 4d 6f 2b 6a 01 00 87 aa 28 0b 47 75 e0 34 05 01 88 df 52 7b b2 a1 14 66 bb b3 17 55 77
46 a8 ed 96 0d 8c 21 32 c8 91 82 58 e5 20 d4 f3 e0 7c 0b a0 b2 60 ee da 21 db e0 00 00 b5 56 7b
a9 d4 66 b3 55 ea 85 71 70 49 b2 69 63 5b eb 6d 88 03 cf 56 48 0d 01 a0 24 00 50 14 85 ca 80 94
61 29 2b 2a ca f2 93 6c 4a e9 c3 3c 94 d7 52 c2 54 ef 25 35 bc ad a8 08 02 1e a8 a0 2d 1d d4 80
f0 18 06 23 7d 4b 2e ca 84 51 9a ee cc 5d 55 dd 1a a3 b6 18 36 30 84 cb 22 46 09 63 94 83 53 cf
81 f0 4e 29 c4 c0 7a 4d 2b 00 a0 0c 01 80 20 33 26 9f c4 d6 73 99 dc 68 f5 35 5b 06 79 3f 8d 4f
37 b9 ad 4f 63 93 42 0a ec f7 62 ec fd a3 75 6a c7 be 21 21 fb 77 ba f7 d6 fa df 6b fe 95 76 26
77 38 0d 00 9d c0 7d 06 6e 03 ef 35 70 1f 59 23 00 fb 0d 84 02 5e e9 e0 00 00 d4 98 83 99 bc 7b
37 47 28 b4 5d 11 a7 db aa ab 39 63 47 f2 ea b9 4d 99 22 fa 92 b5 b4 38 8d b6 5b 30 8b 19 5e 95
66 00 0a 03 4c 78 78 08 02 4e 92 f0 00 00 85 bd d5 09 83 2b 98 06 e2 61 6a b7 47 43 36 9b ea a8
4d 95 16 ac 9d ab 75 6e ad 55 28 32 3c 55 a3 35 66 cc 5b 59 39 b4 45 8e f8 4b 08 c4 8e aa e2 50
c5 8c 66 bc 76 94 31 82 d1 2a e1 97 e0 80 1c 92 9e a8 a3 a5 89 da ed ca 1c ca a9 73 9b 76 7e cd
c9 9e ea 98 63 9e ea 67 ca 1a 15 b5 7a a2 30 94 59 7d 66 2c 83 a1 1e 26 a8 a0 3f 35 36 80 b0 16
00 c0 06 01 b5 b2 62 7a 46 15 55 b5 bd b2 d7 c6 87 0d 19 a3 76 10 b0 cc 4e 64 cd 94 1c a9 73 0e
20 03 00 d6 b9 18 07 49 14 60 07 80 90 1d 71 06 a5 c3 59 91 1e f9 c9 e6 d9 54 69 f6 f2 78 46 dd
d1 ed bd a3 b5 76 ab 54 f4 e3 3c 44 29 40 7d a9 0f 01 c0 7e 4c aa 03 f8 83 38 03 00 14 05 56 ee
a0 3b 04 44 01 25 ca 78 00 00 90 17 8a a4 40 70 13 8b 1e 22 9d 99 91 ea c9 b9 3b 2f 26 a5 55 55
a6 b9 d9 fb 47 6b d5 cf 6a 3d dd 42 ef 84 f0 86 32 7a 24 6c d4 a3 93 72 f0 80 18 79 b4 40 78 16
7f 10 63 08 af 1e 63 dc 9b 97 75 6e bd db bb b1 7e 28 c8 0d 5e 4a 61 37 99 a6 6d 2e 18 4d 28 51
f2 c3 6e 2b c5 b9 78 00 00 60 18 c4 2b 80 bb 50 de 06 1f d4 55 02 27 e9 98 73 0e a1 c4 5c c2 67
eb 78 95 21 5a f5 6c b2 35 ce aa 94 a9 f9 51 5b ce 8a d1 5a a9 5a 59 3f 4a 2b 7b 56 c6 3e d3 12
bf 16 31 cf 90 e6 af c5 ac a5 52 32 2b f1 7b 6a ee b0 8f 7b 5e da 42 ad 2c de d7 be cc a7 4f 2f
95 ea 82 93 d5 38 db 7a ef 5d eb bf c0 00 00 85 4f 70 7c cb 95 29 5c 1e 1d 15 05 c9 56 f7 07 c5
39 8f 55 45 e7 4c cf 24 e5 0d 7b db 3b db 3b 3f 67 ec fa a9 4f 9d 86 2b a6 ea dd 76 b4 f1 cd 56
2b a5 6e ad f0 bc bb d8 55 49 13 ef f0 00 00 e4 f1 a4 3b 25 da f9 3c a7 10 e5 56 ce 4f 08 d3 c1
d5 73 b3 f2 a7 27 b1 95 2e 5d aa 95 1f 7c 88 f3 52 ad 73 ae 75 c2 c3 f3 21 0e 6b 4e b0 db e6 27
88 db ac 3d 3e 49 56 c6 ab 4d 1d 75 45 6d 9e df a5 8d 63 5a 9d 5a ad 5b 56 b9 01 92 bd 54 f5 d5
d9 9e 73 b5 6a 88 32 aa 90 f4 9a a1 d1 a9 94 4b 26 2a 75 e9 26 0a c9 ba b7 56 ab 7e 9d 5e 0a 6b
9a f5 5c 2b 4a 74 92 fb d9 0b 63 78 41 78 af 80 2c 05 80 90 12 01 c0 28 1e 95 ad 59 f5 24 3c 6d
5b d6 7f 29 0f 2e 54 c7 de 42 26 93 94 39 99 8d 07 73 e4 cd 9c ac c1 bd 37 36 a6 84 6d 1e c9 e9
a9 e3 3a 45 a4 7a ee 8d d7 9b 27 7a 80 b1 0d 56 05 2e 7e 49 a2 56 eb 5b e1 83 f4 4d 12 9e b8 e0
fc 17 81 f0 f0 00 00 a5 0b 53 cd 9c 66 71 53 d0 ef a8 b7 2d 55 04 33 f9 3b b3 94 45 8f 94 c9 6c
e5 0f 7c 65 b1 cb 39 43 67 33 88 92 be cf c9 df 96 6a 54 b5 72 78 e5 d5 9d 65 ba aa 5a 65 26 87
64 2a f5 a6 50 08 bb 20 2c 43 5c 57 d9 c3 10 aa da d5 c1 ec a8 d9 9c c5 b0 6b 42 b4 f6 e9 7b 20
ee b1 b5 2a 5f 05 42 c2 54 ad c0 7d 2c 8f 01 e0 14 00 10 02 9d c3 00 e2 dd 90 04 96 e7 e0 85 86
7f a4 9d 44 69 62 1f ad 41 ed 4c 56 7c 4a c4 5b 8b 95 36 b3 0f 19 0d ed 1d 9f b3 f6 8e d5 cb 21
78 94 6b 6a 6e cb db 3b 94 06 d5 2c b0 16 02 40 38 03 00 01 80 48 d3 06 02 8a a4 f0 00 00 a5 cd
5b 5c a1 ed 2d 64 56 71 6c 33 db 56 97 d5 49 4c 93 35 6a 86 2e f1 dd 66 5a 99 ab aa 6a 49 9e a6
73 ed 8f 55 5b 99 f3 46 2a 84 71 4a 9d 6d 9a f7 6e ae 84 e5 86 21 cf 9b b3 7c 2f 4a d3 25 a1 a6
fb df 73 ae 53 0e 9d 6e 34 eb 96 b2 d8 68 6e 42 e6 aa 21 26 5b 7e 00 00 b5 70 0d 69 d4 5a 31 43
c9 92 75 3a 1e 4e f8 c4 a3 7d 6f b3 f6 8e 57 1a 45 23 5c 8f b6 da e7 c6 20 de 32 51 c4 c0 45 ac
8b 01 40 61 ef 00 00 0e 02 e8 67 10 06 94 bf 80 00 00 65 68 5d db c2 4a a1 6b 11 54 6e b6 ca 5a
e5 2b c3 15 d6 d6 39 8c f4 49 75 e5 71 93 da 35 6d bb 5f 6b e5 52 9a e2 1d 4a 7b 57 28 8d 66 c5
66 49 4a 1d 38 73 8c 85 a8 0f c6 1d e0 30 0f c1 e1 e0 40 0f c1 e2 00 79 93 4c 03 8a 69 20 0e 8d
19 00 8c c6 43 c0 00 00 95 92 53 9e 83 5b 69 64 5a c7 60 f7 bd 57 39 a5 3a ce cb 75 e6 ac d5 9a
23 38 00 03 80 b2 21 d0 06 53 0c 20 1f 37 2f 55 69 8d fa 8d cb 23 55 6c cd 7b ab 94 3a b3 6b 16
e4 ed 1c aa 35 8c 56 b7 27 2c 9d a2 98 ed d9 6b e5 a5 92 ad 58 cb 1a 5d 18 df 62 2e ca 96 15 68
aa 50 00 30 09 ce 05 60 24 0f 23 ae 04 b6 24 b8 99 52 d1 4f 41 32 68 5a c2 eb 48 66 76 96 30 f3
36 11 6b a5 4d 4d f5 0a 5e b3 56 a8 64 36 d8 eb 49 b9 f3 37 27 7b f2 9c ea c9 ca a4 ec 88 47 b2
f2 d9 ca a1 d1 ec b8 be 4f a4 79 9b 15 31 8d 11 62 ab 27 1c 5f 88 00 10 07 e9 11 00 3b 4e 6c 03
69 a9 60 1f 7c e4 00 fc dd dc 07 e0 aa c0 3e c5 46 01 f4 aa 78 00 85 47 77 ac a4 45 25 42 19 77
81 11 cb 5a a6 3f c8 63 cf 94 31 b1 ea 33 34 e5 0c 8d 7a 25 4b 39 53 27 7c 93 72 ce 54 c9 de 2b
5b 53 95 32 57 6d 54 d4 e5 4d 95 d3 b4 b5 39 54 21 52 cd 0d 4e 59 29 ac aa cc 6b 56 c6 ab 09 53
1a c5 f3 aa b2 45 46 a9 8c a8 4e 95 71 a9 e3 4c 64 f5 39 6a 94 b3 13 b9 62 92 a5 2c c4 ad 96 a7
a7 13 b1 7c 95 55 6e 53 e0 00 00 84 ec 53 c6 83 3d ad 7b d8 59 62 cf 6b 5b 27 05 e1 4c fb 55 45
a9 bc 53 3e d5 4c 63 fe 04 cf b1 42 a1 7b 40 93 6c 50 b8 de 50 8b fe d5 29 f7 94 42 ff b5 4a 65
e5 10 dd 6d 62 d9 1b 22 33 ea 5c d5 be bb 05 f2 57 b5 31 c6 e1 7b 96 2e 53 64 48 5f 25 93 5a 79
8e 8f 38 64 f6 15 54 8c a8 d9 15 51 36 d3 2a 56 45 4d 53 3e 28 8d 91 53 13 2e ae 48 65 77 43 a2
a4 95 b7 22 d9 c1 9d ce 65 42 be 6d 76 60 25 27 80 00 00 95 b0 68 dd 84 cf a5 74 9c 19 81 0b ec
5a d5 b5 d8 c9 b7 55 aa 37 f2 11 6d c5 49 c6 e3 84 5a f1 42 db fc e4 d3 3d 4c d5 7f b1 34 cf 13
3d 1b ec 31 6b cc f6 a1 71 fc 35 b9 86 2a 55 7a 51 62 45 8b 17 1e b5 57 b5 f2 d6 4f 69 55 6d 76
b9 92 b9 86 2b 9c 2f 7c 03 9d 5a b8 fb dd 7b 5e 1f 0a b1 b7 5c b7 17 ba b2 8d f7 36 ff 65 89 5c
dc 04 f7 25 63 8e f5 40 b7 6b 57 15 1e 48 2c f6 d3 49 45 76 25 3e c4 d0 3c 6c 22 59 71 33 4d 9e
69 3a ab 4c b4 fd 84 dd 8b 14 35 1b aa ab b4 bd 74 b9 c9 bb 77 2d e3 b0 5b 6a 25 52 e6 00 0a 02
7b 65 50 11 8b 12 00 a3 e8 0c 05 23 04 20 1a f4 3f 65 89 5c dc 04 f7 25 63 8e f5 40 b7 6b 57 15
1e 48 2c f6 d3 49 45 76 25 3e c4 d0 3c 6c 22 59 71 33 4d 9e 69 3a ab 4c b4 fd 84 dd 8b 14 35 1b
aa ab b4 bd 74 b9 c9 bb 77 2d e3 b0 5b 6a 25 52 e6 00 04 01 dd a3 30 1b 53 49 80 fc de 4e 01 c0
38 07 00 a0 78 00 00 95 f4 80 6c b9 0f ef 76 6c c5 49 84 fd 02 db b3 56 65 bc b5 96 13 9e 95 e8
0f 7d 95 48 9b 7a 29 cf 29 6a 2e f3 92 b3 c9 e1 0c 8f 97 bd 87 4e bb 44 d9 5b e0 e5 8c d1 76 37
fa 55 63 34 65 15 67 98 b6 ee 1f 66 36 5b 6d 07 80 00 00 b4 53 7a 2c 55 52 b5 2c e4 a8 f3 58 ae
cb a9 a7 24 96 79 ad 72 88 ba 65 d5 65 9b ab 16 32 58 ee 8d 2a 95 ab 69 53 9d cd a7 74 ef 8d ed
bd b8 2f 10 f0 00 00 56 2b 69 53 9d cd 9d 73 5c 36 8b 16 da 58 f8 0d 9a 3c b7 95 53 4f 04 b5 2c
ed 1d 9f 93 d7 08 a7 39 5c dd 19 ab 36 66 d9 5f 4a 12 25 e6 4e 77 c7 00 03 62 55 4c 54 43 91 b8
bb 18 2d 9d a9 f3 28 c6 33 46 34 a9 18 78 45 05 64 4f b3 b9 2d 72 a3 16 e7 52 2d 5d 08 ad c1 bd
ef 97 4e 2d 2a c9 72 a5 d1 79 b3 1b 4e 51 78 db 1b 8e cd 93 60 54 d4 cc 53 69 40 58 82 6d 03 55
56 2c c8 8f ec d4 b0 72 cb c5 54 b5 2b e0 b8 ed 89 9d 4a f8 b4 ab d3 6e f2 b2 73 b1 ba ca 5a 6c
02 8d 54 60 0f 1b b6 ab 84 92 d5 72 57 db b5 93 e4 b9 09 3a 49 66 e8 3c 52 ac 4e 5a c2 6c f4 51
52 96 b1 a1 cc 34 29 59 9b 64 ad 07 4d 4f 45 87 6a 66 1d c8 e1 73 5c 34 ed 2f 49 58 e7 8d 23 4b
f3 95 52 8f 0e a5 b5 e5 14 a3 bb 2b e9 79 35 26 cc cb 79 ce cd c9 ab 56 45 5d cd f2 ab 62 11 0e
bd 98 b6 92 42 6a bc 95 6f 76 54 a9 04 4e c8 c0 00 18 04 aa a6 a0 23 65 14 01 39 c8 78 35 ca 69
4d ca cc d9 6c f0 0c ac 75 fb db 9a f5 55 ef 4a 06 9e dd d5 9a af 55 ed bd b5 b6 96 da d0 60 cf
e4 a5 b5 a2 12 4d 64 e1 7c 5e 0f 37 6c 36 5f 06 83 cd 4a 4c d7 ca 42 66 b6 9c f0 00 00 75 ac 70
55 34 48 e5 6b a0 14 ae d2 bc 5b 29 2a ab 2d b3 95 5a 92 82 ad eb ed 5d ab 75 62 aa b5 0d 37 20
7b aa f5 e6 cc db 9b a1 73 5c 15 4b 69 b8 df 0b da e1 b6 8a 51 aa f4 b5 32 56 da 4c be 6b a0 52
8b 1f 80 00 00 25 6a 69 b4 94 cf 4d 43 1e 2f 09 32 d8 55 29 b4 2a b5 b3 b5 76 ae d1 a9 e9 56 b6
18 93 3a a7 48 a8 8e 24 cc a9 f3 19 44 4b 1c eb 95 8d d9 53 41 bb 27 6c 6e 85 cd d1 53 4a 3d 21
7b f2 52 78 97 29 5f 0b a3 ac 3e 9a 18 4a 87 09 37 9b 76 11 91 33 d6 44 d5 84 5a 6b 37 50 35 61
38 03 44 62 35 98 d2 40 d1 18 95 f0 65 cb 70 55 24 c6 e1 73 a0 16 ab 6e 4b 5b 09 0c 29 a4 d3 95
d2 af 6a 8f 6c ed 5c aa 74 7a 52 c9 5e 2a 7b a1 a4 f9 27 1a 98 07 94 ea c0 3d 16 26 01 e5 ba a0
0a 91 cf 16 4e 8a ea 8f e5 dd d9 c2 78 c3 28 2d 8d dd f4 0e c9 29 74 6a a9 01 b2 3a 65 e9 e1 b7
4c b7 35 04 49 72 28 0d dd 45 40 73 10 4a 03 53 43 20 12 52 a0 07 80 75 e6 7b ac 35 48 65 7b 62
30 a9 12 b9 5a f9 23 a2 8d b7 16 52 53 0a c7 25 ed 58 a2 76 8a 61 99 6e 25 7d 59 d2 31 47 49 64
eb 25 0b 8a 38 0d 05 51 a0 34 08 43 57 42 aa 9d d7 b5 6a bb c2 15 18 27 d0 ba 89 55 37 4b aa 2f
68 da 66 21 42 c0 2e 40 6c 5b 69 3b 33 0d df b7 6e ec df 4b de e4 a6 59 d2 c3 08 16 61 de 27 32
c3 c5 38 87 89 62 80 e0 0d 0c 07 d8 6a 40 3e a4 5f 95 b3 b0 41 aa 49 ad 5c 66 0d 08 b5 7e 51 19
2c 2b f4 6b b3 f6 6b 4e d8 b5 31 6c 56 d5 bc d5 14 c5 6d cd 56 a9 2d fb 63 4b 65 ec ad d2 bd 2c
fd 44 79 a4 31 5a 77 dc ca f9 0c 96 8f c5 43 d1 78 00 00 75 cd 72 d9 29 ca e1 73 9a 96 6a 6d d9
db 3a f9 54 25 73 09 51 ce d5 ca a1 5a 97 4a 8e 72 a9 5e 65 92 6a 9a ae 78 67 54 b8 e6 2b 9d 15
19 ea a7 2a e2 f4 49 b9 a1 b2 a8 a5 95 d1 d4 60 00 60 24 b8 23 01 76 2b 10 08 c1 4c 3c 00 00 75
8a 72 5a 2b ab 21 6b 5c 96 8a 69 d9 5a c7 2d a2 32 97 36 72 b9 42 2e 0a 9b bd ab b5 72 88 c2 6a
0a a3 9c a2 4e c4 71 a9 2b aa 54 54 a2 31 b6 6a 84 fc 79 79 49 b0 09 10 d7 56 4a 0c a9 0b 2e dd
98 b2 0b 68 7b 19 76 6d a5 cd 5c 5b c8 5d 65 7a 25 98 6e d3 b9 66 88 d6 cb 13 f9 99 a2 51 f6 64
f4 f0 85 b2 78 63 29 a7 a1 6c 1a 16 ea 89 eb d7 2a 8d db 63 2c ae 4e a8 54 6e cd 01 fc 92 35 4a
a3 e2 b8 c5 ed b6 ad 13 9c 94 37 7e 2c 5c de a9 cd 4f 4b 95 94 9d 45 9a cf 09 60 cd 23 eb 24 f1
d8 af 08 de 72 b2 86 0d ba 26 b3 aa d9 7c 34 47 0b 0b 3f b5 92 78 ca 31 b7 69 5b 5d 3c c6 ba cd
d5 4a 11 bb 63 ab 88 d6 7a cf 4a 9c 98 d6 70 be d6 b9 da cb d1 0a 94 76 b5 20 0d 2a 15 a0 38 0f
c6 19 a0 3c 06 80 90 0a 01 00 6a 11 98 06 98 4e 80 2a a8 70 00 02 5f 36 bb 3a b2 8f c0 00 00 a5
52 70 ba c8 2f 69 5c de 0c d8 12 6b 54 97 ec 22 d5 9b 73 66 5b 48 88 e3 aa cd 6e 92 1e b6 6d 33
53 a4 c8 b5 21 bd da a9 52 59 78 2e db 2b 57 13 2f 9b 25 de d7 41 2b 8d 21 37 ae 15 c2 78 43 26
aa 2d 5c 2e ae cd 8c 77 86 84 e0 80 2a 25 08 04 1b 0a 80 20 e7 52 00 f9 c2 07 80 00 00 56 28 52
4d 2b bb 5d 64 1a 34 ae 8e 48 57 16 8d ab 0b 77 15 bd b1 ec 87 b3 dd 5a a1 ce ae 65 6b 1e e8 dd
5b af 16 ba 0f 4a d6 94 85 f3 6c b2 0c a9 0c 07 13 2a 90 16 02 c0 20 01 84 78 80 0b a5 c0 04 02
57 2c a0 38 95 7f d5 12 5c aa fc 46 b9 3c a8 ea cd 59 5e 4f 2b 34 1b e4 6f b3 ee 8d d5 cb 25 37
84 57 cd ee db 5e c8 2b aa cd 6d f8 37 06 e2 cc 5e f2 75 39 72 af 23 64 95 9d 5a 4b 4f c0 00 00
75 b0 59 33 f9 c7 25 74 1a 0d 16 d6 4a 5d 28 02 3d b5 5a b7 56 d9 db 4b 2a f9 64 66 f1 66 ad 55
68 53 22 b6 8f 74 6e 8c 51 88 4a 1b c6 8f 35 6e ac 59 b9 43 03 3e 8b 37 5a fc ce 54 1d f5 76 c2
b9 72 37 7c d3 b0 63 83 1c ee 48 dc 19 8b 5e 38 89 72 c6 62 98 24 32 4d ad 58 9d a9 72 b2 eb 65
c9 6b b4 43 c8 95 6b a4 c4 6b 11 64 5b 0d 41 2b 43 4c 3c 45 66 6a 76 23 32 99 51 da 77 2d 14 b6
54 96 9c 53 b6 2e 15 31 86 d5 0d 52 c4 f6 a4 99 ab ea fb 37 66 e4 d4 bc a1 33 df 7b 3f 68 e5 52
c4 a2 42 4d fb 6f 6e dd f6 bd d4 65 5f 67 65 70 9e 0f bd 57 b5 ea 69 4e e9 55 d0 ab 91 1a 38 bc
85 d0 70 3c db 45 27 65 2a 7b 42 2a 9a 27 aa aa a5 42 b8 b1 f2 8b 20 d5 7e e1 57 9b 33 66 6c d1
8a ad 05 22 da c9 ea bd 52 14 1a 3c ba b3 36 75 05 8f 2e 6e cd f4 be f4 54 08 9e e5 af b5 11 0a
23 46 8c 22 f4 24 88 d1 1b 02 a9 ea 25 dc 87 c1 d8 4a 4c a8 8f a2 7e 0f c7 c0 00 00 46 07 5a e6
bb c8 91 71 98 bf 0b 89 19 5b 27 86 cc c2 93 16 52 53 0e 6f ac ed 5d ab 93 ce 72 ea 73 a4 e4 d4
b3 aa ab 69 63 53 2a 9c 1c 52 d6 e7 9a 37 56 2b 83 e1 19 d7 66 ac 0e ca 58 79 91 12 e8 b2 19 d1
dc 88 02 2b a2 68 1a aa 35 77 54 6d 2f 2d 94 e5 91 8f 5b ab 61 29 64 74 d6 5a f7 49 9d 99 b3 ae
bd 74 56 66 b8 da af 63 25 19 62 64 ab e5 0b 63 25 72 3e f1 00 b7 c2 0c 0f 75 6f 68 7c cc a4 a1
5b da 1e f1 2c ba 57 07 06 34 da 6f 95 ce 6b 0a b7 5c e5 73 aa ca 36 57 79 5c ee b2 71 92 4e 59
3b ac 23 ec 93 36 4a ca 74 95 4f 22 20 0c 03 89 8a 60 1e 94 42 c0 48 09 00 10 06 20 aa 40 39 07
41 81 e0 00 00 66 2a 61 66 bc 24 d9 82 58 5d ca c9 b8 60 f7 07 d4 39 53 97 4d 72 ed 2b 2d e5 b3
6c d3 ba c9 f9 64 5f 15 0a f2 6e 54 f8 cc 42 c4 9b 94 3e 72 75 11 54 d4 ef 8b b5 3e 33 35 3c 5e
cb 2f cc bd 4f 42 ba cd 43 42 40 75 06 70 03 f9 62 80 1e ec 90 10 08 a4 e2 40 62 b4 b4 02 d4 a7
70 13 a4 bd 22 3e 62 21 9d b9 89 55 88 5a c0 b6 da 45 5e 9a 2c 74 96 8d 94 97 a9 0e d5 a3 6b b6
16 ca c5 79 31 53 67 15 2e b6 4a 5c da 43 bb 3d 2e 97 b6 52 b7 4d 18 96 0e 73 cc e4 30 a5 84 14
f1 96 f0 97 61 28 aa 33 5c 28 c0 2d a9 5e 85 ce 58 5c 53 56 a7 6c ed d5 70 9d 53 2d 37 3b 57 6a
e5 53 bd 52 24 39 39 45 70 f2 2c f2 3d d1 aa 2d 43 34 9b 69 ea a9 38 56 66 a0 8f a9 76 b0 0a 29
52 20 52 8b d5 45 8d 6d b9 b2 97 bc e7 37 b0 db 9d ef c1 d8 d5 60 ed b4 e4 66 58 9a 18 bb a9 99
9e 2e c8 4e c5 c0 10 09 6a c7 80 75 8d 58 7c c3 a8 a9 64 1a 18 cf 0e bb 56 ea dc b2 d2 73 b5 76
ae d5 da bb 5e 2c 6e 25 5d e0 c9 3b 80 00 0c 03 09 0e e0 30 88 82 00 ea 28 f8 95 b2 59 b3 39 db
6b 6f 2b 93 ee 55 9e a9 da fb 5f 6b ed 9d b7 17 b9 fb 66 cf 5b be 12 c1 39 a8 47 69 1c 70 96 0a
a6 c1 3b c9 1d 8a 61 b8 30 ce 37 62 97 5e b3 1a 69 98 ad 93 2a 86 ec 56 0f 7c 4b a1 26 bc 00 00
95 b0 70 54 bb 48 e5 6b da 17 0e f1 dd 59 35 b4 b3 25 af 95 45 55 52 a7 a3 ed 1c 99 ea cb 65 4b
55 26 53 ea 1d 0a 67 19 9d 39 65 27 0e d2 d7 d4 ca c7 19 57 86 76 ba e6 5c de 5c ed d7 ba f7 5e
fa df 2c 1b b2 59 94 b5 bf 17 64 cb 22 4c ab d5 d9 3a ec 4d 67 35 66 51 d4 23 22 b7 7c e5 4e 5c
fb a9 cf 7b 56 ea c5 4d 3c fc 28 bf f3 55 eb bd 92 b9 4b cf 24 2b b4 6f 7d ea bc d1 ec 45 1b 1b
30 3b eb 69 06 16 ac 10 e9 db 14 49 a2 f4 b1 ea 8e 13 6a ba 72 33 54 77 27 80 e5 51 6c 7c 93 cd
79 54 5b 1f 06 f3 4e 55 17 47 c9 b4 d3 95 c5 f1 ec 8b 74 e5 71 74 7b 22 5b 39 64 5b 1e ca 96 cb
58 f7 d6 9b 36 b2 f6 62 c7 3e b0 99 75 92 b1 cf 9b 35 3d 1a ad 54 e1 51 1e 54 cb 60 c6 54 46 52
32 dc b6 1d 85 93 46 02 66 67 20 0f 44 20 80 5b a1 43 c0 d4 ab af 43 2b d7 79 53 99 d6 ca ba 5e
54 ca 66 29 cc b3 b5 76 ae d5 db 2d 72 df 38 a6 d2 16 01 56 cc 3c 06 80 f0 1c 04 db 30 c0 27 19
6f e5 6c 7e 63 9b e9 79 43 25 b4 a5 72 5e d5 da b7 56 ea dd 79 b3 16 ae f2 ec f4 dd bd d4 b9 76
89 67 a7 1b 2f 4b 92 8d b1 d8 db 99 72 5c 96 8d 8e 35 5c 78 2d 42 e3 6e 17 40 ee ac eb 63 75 d0
94 30 8c ba 44 02 ea 47 e0 00 00 d4 ab 76 74 03 cb 71 42 65 9e 47 92 cc 50 a6 e7 92 ac 9a f3 d6
79 cf 4a 22 e0 98 74 8f 46 a9 d5 6a e6 a6 86 a2 76 6e cd 5b 16 c5 2b 9d b7 76 e2 e8 4a 19 29 fb
94 be 36 28 59 bc ac af 8c 42 62 e5 99 1b dd 7b 9a 71 72 cc c6 5e e6 94 5d 33 13 c0 b5 48 9d f8
34 66 ed 51 f3 9a 0f 76 3c 54 89 df 8a c5 ab 15 26 f5 a4 af 74 d5 6a c4 d3 2a cd b9 62 f2 f3 28
96 dd 5c ce 25 4a 2c cf 17 33 89 36 8b 34 b5 cd d9 54 b9 da 2d 6b b2 e8 b8 2f 9a 5a ec 22 bc ab
ca 56 3e 0e 87 0e ac 85 b3 95 58 92 a5 44 02 98 ad e0 00 00 94 a8 8f da 19 6a eb 36 27 4d 71 dd
0b 69 8a 13 85 b9 35 b2 d2 a5 68 75 99 ae 50 b1 36 33 78 89 6e 2a 74 dd 1d dd cf 8a dd 91 a6 66
d6 5a b7 60 39 5e ad 54 b5 da 0e 56 ab 4d 2e 76 82 29 d2 f7 6b 9b 65 46 e5 77 c2 f6 e0 36 95 e4
2e c0 f4 77 a5 84 e7 80 00 00 d4 d1 7b ec 1c 4d 3b 57 68 ec fc 9e 2c 6d 54 3b 55 eb 95 ae 7b 61
b4 d8 b3 5f 2d 83 d7 11 66 d5 6b 99 44 95 b6 b2 52 f5 d8 ae a1 6d b3 c2 18 37 21 72 d1 6c 75 ef
d8 d2 ab cc 21 73 e8 8a ee b2 36 5f 0b 3b 00 a2 97 c0 00 00 e5 53 75 db 32 69 39 53 df 96 cb 15
de d5 da b5 5b 21 b8 89 15 db 56 d6 67 9a b6 71 d6 a2 f6 e9 70 d3 00 00 50 13 5b c2 00 e2 dd 8c
03 d9 0c 20 18 e8 40 f0 00 00 e5 0b 75 f4 0a 5d b9 3a db 5f 24 73 5c 4e 98 5f a8 a4 ce d3 a2 75
e6 2d 74 b5 08 8d 72 83 bc ed 5c 58 ea a9 0e 2b 59 36 32 a1 b3 b2 f5 5a a5 be 59 d9 dd d8 b1 96
66 97 3b b5 ad 65 21 1b 15 6d 2b 97 74 6a c4 7e 42 f5 dd 13 29 1e 6e bd 99 34 9a 37 d3 af 4d 19
26 65 27 e0 00 00 06 00 d7 32 a0 07 35 07 92 4e 28 8b 33 9d e4 b9 a2 99 3c f9 77 36 e8 cd 76 b2
f1 33 3a ae e4 ae a3 05 5f 26 a8 eb dc 82 bb 98 6e 32 f8 a8 22 ee 11 8c c1 ea 26 b8 66 df 80 04
80 90 12 05 a4 5d 6a dc 9e 3b c9 d9 4b 98 44 d2 f2 86 de a2 58 ec bc a9 d9 87 a7 2b 1e 2c 77 16
29 13 37 6b 9f c2 d6 4b 89 d2 f7 e8 6e 52 64 12 bd f6 35 75 b5 48 00 38 0a ad dd 00 45 ae a1 e0
50 08 9c 49 c0 68 16 8d 92 d7 8a 1b a6 25 74 75 db 06 6d a9 9d 4b 56 d2 7f 6a 68 4e d8 b9 1e f6
9d fa 45 86 67 e0 00 03 28 6b 86 e0 f9 43 4b a1 98 45 91 93 e6 fc e1 4c 5d 86 56 b2 b6 c3 17 f0
a2 1c f4 32 ca 19 25 98 57 93 b3 87 0c dd 9a 47 e0 50 0e 63 b1 40 38 16 91 74 bb 82 5a de a6 64
b2 61 8f 59 89 d9 5d 77 92 d6 5a 86 da d6 30 36 99 ae 96 37 10 6f e6 f7 75 ae d1 2d ca 4c 8d 72
62 a6 d1 49 c1 3f 05 c1 ba c3 8c 0c 07 a6 ae b0 1a 03 00 fd 61 da 02 00 fd 59 cc 07 14 af e0 00
00 60 10 7b 22 c0 c4 73 73 29 c4 da f1 2c 21 08 9a f3 5d cb 89 5d 28 45 d3 7f e6 7b 50 ea 3a 2d
9b fb 15 42 ee 8b 08 ff c5 51 9c 29 3c 5d 71 65 23 08 6c df 3c db ab 67 39 53 65 f6 6e ec df 6b
e0 f7 65 87 d9 b3 06 39 95 61 ad 46 bc 6d b8 88 85 19 af 22 ea 25 a5 46 cb e6 86 45 66 72 90 06
26 9b c0 00 00 60 13 aa 2a 80 92 96 29 23 25 38 c7 51 cb c9 98 f5 24 27 16 87 b2 3d 49 1a 44 a9
e4 8f 52 ca af 08 99 64 cd 16 ae 94 48 47 ca dd 2e 8c 4a 26 f2 1b 5b e5 7c 5c 6f 35 34 c8 5f 05
9c 5d cc 11 d7 b1 49 9d 0e 85 f0 00 00 50 16 8a b2 00 92 5e e9 12 9b 96 a8 ef 2d 49 08 45 39 d5
8f 72 6a 48 52 25 0f 2b 9b 93 12 ca 8a ca b9 5d c5 13 a2 32 b5 d9 31 5d 69 04 6c fe 5c db 5b a5
7c a4 2e d3 95 ba df 5b e5 7c 24 6e af 99 d9 5e e9 a3 24 55 72 97 c9 50 8c b1 72 96 3f 94 b2 12
29 58 06 90 cc 60 3a 65 47 01 d6 32 20 0b 21 97 c0 00 00 70 13 81 2b 80 9a d9 f1 0d 12 96 ea ce
dc 4b 24 b4 bc 24 b3 52 bd cc f1 0f e3 dc d6 9e b2 61 39 7f 26 6a c5 55 9a 91 bc 7d 33 65 ed a5
f3 b9 a1 a5 7b 67 7c af 94 8d 19 a3 c7 4b e0 c8 08 6a b8 d3 0f ac e1 8a 75 ae 03 20 35 60 1f 62
b3 00 fc 58 98 07 de d0 80 3e e5 6f 50 0d 5a d1 00 5a e2 c4 02 99 94 89 5e 97 cb 48 31 52 44 be
fd 41 6d 98 95 8b bf 61 49 66 a7 82 6b d5 1a 6d ca a4 bb d4 63 9a 72 c9 be e1 58 ad 7c be 73 a6
55 3d 5d b0 94 99 5a 23 33 4c 67 08 34 56 b2 3b 29 99 19 21 9c c2 01 a2 38 10 13 19 5c 00 92 65
9e 02 00 44 21 bc 00 00 60 0e 74 1a c0 75 51 85 42 92 3c b5 43 a3 12 6b 1a 4e 55 06 c4 aa a6 b3
75 ee c7 01 50 14 c4 0c 55 16 cb 1a b6 6f 74 76 5d 49 19 bb 99 de 6e f2 28 0e 66 65 e0 3c 04 00
77 bc 6a 8a ba ac d4 a5 dd b3 b7 5b 09 48 a5 05 65 f3 c6 58 c1 e5 0c 6c f7 86 2a 7a b3 1d 59 5d
8b 18 8f 07 52 76 64 c6 9b 3a 4c b1 59 39 a8 ca 76 eb f0 40 09 63 b9 c0 10 08 64 ce 23 5d ee 59
66 29 a9 9b 3c c5 49 ad ee 6d cd a9 de ea a7 57 8e 5a 76 c6 95 dd ea 92 ab e8 35 86 97 1c 2e fd
a1 58 11 3b 0b 9a f8 77 77 b4 d6 a9 56 f6 43 2b b5 e5 17 29 32 6e d1 6a 25 cf 4b ac 2b a8 cc 04
80 b0 ea d9 33 a8 9c 89 c6 c8 de cc 12 e9 98 c4 ac 23 04 3e 22 5d a3 28 c5 6b 77 78 77 1a 31 83
1d 59 15 56 02 01 eb 73 55 d6 95 b4 12 ab 3c 00 00 20 03 6c 32 80 74 95 d4 03 6c f5 68 9d 44 a7
69 93 76 5e 4c fb b4 34 4c 73 73 ee ac db 2b d9 27 a7 5d 6d 37 15 01 f7 a2 2c 06 80 c0 3f 14 56
01 f8 b2 a8 0f a8 99 20 1c 00 f0 00 00 50 0c a3 c2 40 18 16 89 fa 4a 4a 0b 6e a2 8d 1d 56 23 35
b8 b7 1e e7 d5 72 94 aa bb 6a ad 74 2c d2 89 36 28 61 34 22 ad 33 b9 58 c9 6e ce ee a3 00 50 14
5b 42 00 6e 22 08 02 61 0d e0 50 0c a3 c2 c0 58 14 8d fa 4a 4a 0b 6f 23 8d 1d 56 23 35 d8 fb 24
67 a5 54 71 ab 33 b9 18 02 24 d5 c9 e3 26 73 a6 b2 f6 bd 5b 0c b2 a5 25 be d7 ba cc 8b cb 6f 96
0e c1 2c da 6b 1d 8b 2a 8a d8 e9 a7 65 0a b1 93 ab d1 99 43 8a 8a 89 33 0f 30 03 6c 32 80 74 95
d0 03 6c f4 40 1d 70 8d 46 d8 dc 44 14 bb b2 72 57 bf 2a 0e 1d 9d 97 b3 6a 66 c6 aa 0d ac 8c a8
ef 6b 54 57 9a 2b 1c 66 98 ed b8 eb 10 c7 87 78 52 da a7 b9 31 cd 9b f8 ae 2c 9c 94 29 7f 6c ed
bb bb 17 3e 15 50 81 65 86 11 70 da 2a ab 63 7c 6f 7e 0e c1 ee 2c 44 55 a3 30 5b 56 90 e5 59 e0
30 0d b3 d1 00 75 c2 39 1b 63 71 10 52 ee 4a f7 65 bb 25 b7 93 3d f9 6e a9 a4 e4 cf 85 64 1b 69
79 33 63 77 04 fa 4d 4e c7 65 49 ad b5 55 0e 54 f0 6e eb 45 62 8d bc 1c 39 29 53 9a b1 11 49 5a
59 15 1d 5a 1b bb 56 c1 6f 72 a8 a5 d5 ae 6d 63 a2 4d 31 7b 21 74 aa b3 59 5e e6 cd 32 85 b5 d8
41 c1 8c ca b5 76 10 70 63 ac 1d 5d 83 da 17 69 06 68 61 44 1c ba b2 b4 80 3c 09 48 03 ea 03 50
1f 79 9a 00 e9 c1 90 06 88 8e 80 24 56 9f 50 0c 7c aa 00 75 c2 1d 3c ad 2a ee a9 de 4a e6 46 3c
15 b3 b2 f2 57 29 f9 dc e5 fb 96 93 31 93 88 85 2f 64 c7 95 49 23 3d 51 51 db 31 08 ce 47 54 b9
25 c1 33 92 55 32 a3 aa f1 2e d5 50 c9 59 c2 db f9 5c 34 b2 e9 16 4e d9 db 7b 6e ee e5 d0 9c 4c
2a e7 77 7e 2f 84 5e 25 4e b9 0c 5c f3 b5 18 96 bb 06 38 e5 51 66 aa c1 2f 59 53 99 a2 b1 5a d2
65 26 48 ac 5c b5 67 2a 72 23 27 ad 96 0d 9d a0 4a a2 f8 77 37 75 f8 50 0c 7c aa 40 38 04 01 50
a6 e4 4f 5c 4d 23 69 bb 2e a6 64 29 9d af 39 aa 1a f8 77 5b b6 52 a6 bd e1 5e b6 20 00 e0 30 c7
86 01 96 43 18 08 3a 4c 03 c0 60 0d 73 33 40 58 0c 89 f1 ba 73 36 e6 a4 5c 6a 18 e6 e7 c9 18 e8
98 69 99 f2 58 3a 1a 0a ee 48 00 a9 bb 50 04 83 3e 22 8c 1c 95 9e c5 ca e7 78 35 67 d2 6a ca da
8c dd 75 76 ba b9 81 38 7d 0c 2f 94 cc 27 1a 22 cb dc c4 1e 92 8a 22 f4 b1 2b 51 54 48 bd 0b 78
c5 56 92 2f 62 4d 6d 9d c6 8b de c4 49 56 71 9a e8 2d 52 55 1c 40 02 8a 4a 87 96 97 1f 80 00 00
70 0d 8b 1c 40 60 01 7d bb 80 33 1a d6 06 a2 a4 54 ca 1f 59 c8 a8 f7 30 b7 ba 76 2d 47 37 b9 95
2e f3 52 c9 8e a5 49 bd cc f8 aa 07 79 3a 5b ad eb 95 c7 85 3b 42 e8 a3 73 ae 34 6d e1 97 40 cb
90 db 76 62 cc aa d5 35 2f 09 63 9e 00 00 e5 0e 8d 73 94 4b bb 3f 67 ec fd a3 95 3e 13 72 a9 67
ed 5c b1 f5 78 98 39 3f 6e b5 cd c2 4c bb 57 27 73 af 7d 54 d6 ce 44 cb a3 97 34 53 75 78 00 00
55 a5 ac b3 b5 39 21 6a 30 77 34 cb 3d 58 e9 a5 b5 35 9f 95 46 2d 12 d1 6c e5 52 a3 4b a6 c9 39
3c ac d2 c9 8e 4e 4f 2b 3c b2 5b 93 93 c6 8f 0c 96 e4 d4 f0 8b b4 1d 49 25 3c 5e 18 cf 4b 28 50
d6 16 ba da ad d5 31 47 3a a9 5a 95 91 2a ab c2 c9 71 43 d8 f1 4a d6 cc 54 d6 c5 43 a6 92 95 b1
eb 90 a9 9d 85 8c 70 f3 2b 39 e1 63 dc 1d 0a ae 79 5a d7 0e b2 3b 7a 97 39 6d 51 03 67 a5 cd 44
e2 c1 db b1 6b a0 d5 2c fa 6d 5d 08 ac 44 46 93 17 ce 44 f1 73 63 a5 f3 a0 bb e5 d2 a1 7c 64 13
39 70 26 5e f6 8d dc c2 09 57 ca 20 d1 54 cb f0 00 00 85 c8 8c 3a 3c a9 a9 72 6a cc ae c9 fb 58
bc 23 a3 a2 bf 55 ba e6 e7 0d 36 e5 50 79 ca cb e9 39 3d 14 a8 b9 56 2d 4d 54 9a be d5 67 53 d9
88 69 97 a1 d5 17 9b 18 44 fb 3b 57 6a ed 9a b2 d5 84 05 5f ac 2e 52 15 e5 6a 36 fb 99 71 d2 d3
0d 51 b9 5f 56 c2 18 a4 a6 b7 ba f9 4d 95 09 e3 73 cb c0 75 89 a1 cb 4c 1d ad 63 ac 72 90 53 7e
56 fb 9d 33 2d df 95 47 2b 0c ab 77 ed 5a aa 57 5a 08 99 a5 aa 8c d8 d8 72 65 aa a7 75 53 4b b9
5a aa 65 10 92 b4 b1 ad 36 88 0d 2a 59 60 30 0d 31 a2 94 3d 52 ca d3 72 a5 95 43 32 ab 4f 58 05
c0 ac a0 3f 95 36 80 a0 3f b8 30 08 01 65 4b f2 78 c2 22 61 2d bd a3 b5 72 a7 d1 ee 18 ec 9a a1
f9 67 a5 5d 1e e6 cc b6 89 f9 44 b8 67 0b a2 5b 92 62 1d 34 00 c0 26 98 53 01 64 d9 88 07 48 cc
00 03 c0 00 00 25 4e 5c 22 a3 27 25 52 ea 72 d4 ef 6e 59 2c 35 12 a3 9f 95 d2 91 26 e3 77 ed 7a
ae d5 b8 08 47 79 2d 9a 21 8d 21 da 02 01 e6 a7 30 0b cd 90 2a 54 4b 69 c8 ee f5 96 b2 d2 64 ca
ee 50 9f f6 99 71 7a 99 09 ad 26 6d 8e a5 c6 c9 0a 1b d3 77 44 57 32 c7 e9 19 a4 ed d3 b2 f6 5a
e7 ed 55 e4 ed d4 bd da 53 8a 3b b4 b0 76 d4 e2 92 7e cc 9f d1 59 53 b5 33 36 e1 0a 19 dd 9e 00
00 65 a7 b2 bb cc 27 e1 63 30 30 d2 cb 7c 54 fc 9c a3 b4 f7 95 4a ee e8 b7 6c ec fd 9f 94 4a ef
08 95 6b ed 5b af 76 e6 e7 60 d7 aa a3 3d 89 b8 37 02 40 45 4c cb 4b 6a 18 52 36 c8 8a e3 f0 00
00 85 6d c0 ca d2 3f 39 54 6a b4 70 73 ec 54 d8 ae 2c 55 8e 95 2d af 90 d2 e0 40 0f 5a 21 80 ab
61 4a 02 c0 40 11 84 ad 22 6d 95 aa c6 e7 99 6f 2d e6 bc f7 aa f5 da d7 64 e2 64 fe 95 be b8 4b
16 dd 95 f8 ee ce c5 b7 2b 98 69 54 31 76 0a 65 9a c4 cc 9d 62 89 56 91 1b 14 c1 52 16 25 9e 00
00 44 2c 4b b4 23 b9 21 5b 74 12 ce ef ed 55 2a 2d 13 a5 db b5 72 a8 c5 e5 ca 6d 94 ad b0 2a 53
b7 6b eb 8d 52 aa 32 89 81 ab 27 67 ec db 96 93 bd ce d0 87 5d 7c f1 a9 15 9a 0a f4 ad 34 65 ed
aa b3 51 cd 2b 55 ad 85 d9 97 06 6f 2c 1d a1 a7 83 f6 cb 16 5c a2 e9 66 91 c9 99 b6 c6 d3 69 35
66 8c a2 b6 43 b6 d5 93 f2 6c cc ee d0 00 0f 00 06 48 a8 93 e4 39 8f 15 3d c9 2b 4b a6 d5 52 8a
c2 a4 de fb 56 2a 7c 19 99 6e d4 cb 98 98 89 58 72 00 48 0b 32 1d e0 28 07 b9 98 40 1e ad 3a 89
f2 74 a9 57 5e e4 d4 b0 a5 93 d1 3c f7 2e e6 c4 f0 db 1c aa 4c f3 56 2b 86 5c e1 d1 f5 6b 21 d6
66 72 b9 52 d8 f0 dd 61 25 90 ba 1a 09 66 a7 53 ae 34 1a d9 2d 96 fb 9f 77 80 46 27 a3 2d 42 a9
61 6c 3a 70 16 73 5d 59 4c 3c 92 95 bf b6 72 ca 62 58 18 ee fa b2 93 74 19 49 bd 2c 82 95 db 91
de cb 4b 1b 96 44 92 b6 c8 58 77 d6 b2 1b b2 d5 b9 91 f0 85 34 dd 5a a9 ee bc 91 89 26 ea dd 5a
aa 31 99 82 a6 e7 6a d5 72 c4 5a 1d 3b 37 66 ad 9e dd cc ea c9 bb 73 75 ae 9f 19 cd 2a c8 eb d6
e2 a8 69 6e 36 e5 dc 8b 8e f4 b9 69 69 9a f2 b2 66 62 1b 30 c3 0b c5 55 a7 58 80 7a 6d 06 07 80
00 00 55 45 a5 b2 bd 26 a1 5a b0 70 f4 c6 fc 54 d9 ad b5 1f 5f 95 42 0f 4e ad 3b e5 50 94 4b 95
d8 f9 54 67 14 e5 71 ce 55 09 c5 39 6c 6f 95 42 53 10 56 e3 e5 6f 8c d4 0d 45 35 5b a3 55 03 31
cc 58 c9 e4 39 54 53 56 b6 19 2e 2f 54 c5 ac 86 43 14 56 ed 72 df 92 c5 15 bb 5c c5 e4 29 c4 6e
57 a5 92 cc ce db 95 eb 72 bb ab 47 65 7b 1c 73 0c d1 c9 5e a6 9d 49 c4 8e 17 9d ab 54 6d 21 76
05 7b d4 a4 46 29 8a ce ef 30 ed b9 63 75 41 cb 2c ac c0 1e 08 a6 01 6d 65 0f 20 12 92 2e 28 84
52 5e b6 35 ca a5 45 94 8d b2 72 a8 cd 2c b6 e4 5c aa 13 5b 48 d7 5f 6a e5 71 94 43 8f bb 21 6b
dc b6 89 e9 a4 01 c5 c3 34 06 01 d4 ca 2c 06 80 e0 3e d7 77 01 f7 33 38 0f b1 a1 40 7d 0c c3 03
c0 60 1c 5d 33 40 68 0f 01 20 72 29 c6 19 8d fd 9d 8b b1 f6 4e 4f 3a 43 31 ce d3 b5 72 b9 49 dd
d6 35 9d bb 17 c6 2c ec d1 a4 75 e5 74 2b cd b6 8f 81 b0 2b a5 5e 6d 34 6c 0c a9 79 88 71 ba ec
25 95 5d dd 6a ba 38 64 33 9a d8 07 80 00 00 50 1e 62 2b 40 68 09 02 d2 3e 4f 0a f7 64 dc 9c 92
37 88 44 ab 2f 24 8d 1d d1 b2 c9 a9 65 35 65 6d b2 6e 6c 4f 3a 22 ab ed 8f 55 4e a8 aa f9 a4 c5
73 9a 1a cd 76 f1 64 e6 86 b1 3e 3b 59 39 a1 2c ce 6f 16 52 ee 44 cd e5 c5 94 8c 19 43 6b 29 6c
60 aa ce 96 48 5c f8 a3 b1 ad b1 57 31 2c 90 cc ec 66 17 8c bb 1a 96 cc 04 21 28 20 17 17 81 00
b7 3a f8 00 00 10 1e 7b 3a c0 68 0d 01 a0 34 06 80 b0 35 1c a7 0e 8b 59 de c9 d9 39 2c ab 0c 8b
59 de cd c9 eb 38 25 58 f6 76 bd 5d 4a 9a ab ed 6e f8 57 09 e0 ec 27 41 59 8f 8a 3b 19 b5 46 2e
e0 5e 30 1f 30 42 80 fb d1 14 07 d4 ac a0 3a b5 66 01 e8 33 34 05 01 f7 b2 9c 0e 45 18 3b b2 55
8f b2 ea 58 55 a1 23 24 99 97 73 6e 8d d7 bb 35 6b a2 f2 ad 5a 3d db 9b 6f 65 ac 7c 1a 50 e6 d9
4b 1e d7 77 18 b9 ca dd 1d 59 e1 57 a8 03 cc 25 48 0c 03 ab 46 58 08 02 c6 85 f0 20 1f 52 b2 00
ea d5 94 07 a0 cc b0 01 80 3c d4 e8 01 79 b2 48 14 8d 98 18 69 65 bd e5 c4 8f c9 ca 3b 4f 75 24
72 ac a5 11 fd 49 88 c1 0b be b3 72 76 5d cd 8a 31 38 30 59 d1 ee ad d9 8b 6f 59 30 57 d1 d2 ea
d5 c4 d9 bc 6e b9 ce 0a 66 a6 db ee 7d ee bd af 1a 56 a7 1b 6f 5d e9 c3 c0 00 00 30 11 23 43 80
fb 10 dc 07 e8 44 e0 3f 44 38 01 fa 28 38 0f c1 52 11 5c ce cc 8c da dc b8 99 d3 88 76 9a e6 26
8d 1d 61 66 a6 04 01 6f 64 18 02 01 14 58 d5 70 a4 33 2b e4 f7 66 2e 7b e9 56 1a 37 0b e0 f0 a5
94 92 46 f8 df 1b dd 7b 1a 1c f0 c9 46 de 8c 07 09 a9 5a 6c cb 08 cd e1 0e 23 2f b3 c0 00 00 10
0c 34 a2 80 fc e1 de 03 c0 88 11 01 c0 71 90 d8 92 4d 65 66 8b 97 23 b3 e1 45 d3 eb a9 2f 17 30
88 f6 ee 5c cd 9a af 56 2d bd a9 40 d3 6d 4b e9 53 41 8a f2 53 09 50 d1 1a bd b3 be 58 46 46 89
13 e3 96 10 89 aa ad 77 21 7c 16 ca ae fa 85 01 20 2c bc 06 80 90 0a 07 80 00 00 60 1f 72 93 00
eb 94 d4 07 dc a4 d0 1a 02 80 da 55 2d 14 a5 0a cb 39 ee c9 d9 7b 37 2a 8c da d1 f2 47 da ef 69
80 4c 96 c0 01 40 3d 32 cf 00 40 15 b4 a0 01 e0 70 1f 82 1c 40 78 0f 01 e0 34 03 81 c8 dd 37 a5
8c 8f 76 3e c7 69 63 26 b6 25 8d b2 a7 42 1d d5 b4 77 aa d5 3d f2 cc ab 9b a5 8e 74 a3 a9 66 a1
6b 1a 19 29 31 b7 da ab 90 e2 e9 3a 69 22 f4 26 5e 21 da 60 0e 02 8b 96 95 6f 72 1d 6b c2 ab 7c
f0 8e 2d c0 3c 00 30 1a 96 c2 c0 68 11 01 e0 3c 07 80 f0 39 24 e6 ea ab 76 5e cb d9 fb 57 2a b4
60 91 6b 4b 9a e1 5d 60 c2 d1 7d c7 59 18 8a 54 de 2c 76 00 03 00 83 a4 cc 01 dd 0c 01 e0 50 1a
96 c3 40 78 11 01 e0 3c 07 81 a8 a7 37 55 5b b2 f6 3e c7 d9 39 2d a3 04 8b 5a 5e cf 8a ab 06 16
8b ee 52 d8 c4 52 a6 f1 6b b4 00 01 80 7e 6e ed 01 c0 7e 4e ae 03 f4 85 60 1f 9b ab 00 fc 9d 4c
07 e0 68 80 34 97 9f 10 19 b6 ca c0 78 0f 02 40 6a 5b 11 02 20 34 0e 49 55 42 a6 3e cf b2 f6 6e
51 44 c3 2e 36 cb 35 3e d2 00 d4 39 9c 03 80 70 0e 00 00 02 80 99 e1 88 02 1d 0c 1e 70 18 c6 db
80 c6 37 5c 06 2f 98 d0 0a 06 25 53 6a 5d d9 f9 99 71 2a db b1 0e ab 7d cf aa 99 0c 66 94 5f ea
a7 42 d1 6d 57 da a9 f5 a3 4a 75 f4 aa 75 d0 de 2a 3a 6b e4 e6 78 48 76 00 30 0a 3d d5 40 4a 74
82 01 70 68 f0 10 18 c6 da c0 68 0c 03 18 dd 70 18 be 63 40 28 18 95 cc 87 79 55 6e 25 53 aa a6
09 6f 89 53 1d ba 70 77 e2 54 cb 6a 5c 5d fa 95 70 c8 79 37 be a5 6b e5 9f 4d 6d 89 5f 95 58 67
5e 03 00 84 54 e0 05 30 e9 49 dd 24 79 68 92 62 a7 e8 a1 dd e6 d8 b1 f9 37 69 69 bd ae 7e 8d da
56 6f 4b db 82 88 98 97 42 f6 e0 2a 9a 24 6e c1 b7 35 99 87 23 2e 8f 4d 1d 59 c6 00 00 29 64 c9
47 87 56 78 07 01 9b 6c c4 08 81 20 35 2d 89 81 10 35 1c 17 08 b2 f7 ee 48 f6 3c c4 94 f7 72 ea
57 2e 25 e1 27 bd a3 b5 76 ee dd bb e9 7c 60 90 ec 9e d6 de bc 11 7a 0f 0f 4d 0e b9 61 56 32 3c
b3 ad d7 ca 28 ea 8b e2 10 07 5b c7 80 20 1f a3 23 40 68 0d 01 a0 14 0d 54 f7 d5 ba 45 8f b2 f6
5e cb c9 a1 1a a6 48 cd f6 7e 55 09 cd 32 3e 73 b5 f2 d9 51 a1 96 b3 9b bb 17 c6 48 eb 17 9b bd
f4 be 51 36 76 ae e2 30 2a e6 21 de 46 9b cd 79 93 4f 2d 0a 49 5f a7 4b 32 b3 8d 96 c5 ea aa 8d
72 55 f1 aa b1 96 48 7c 10 1f 6a 1a c0 78 0f 01 a0 24 05 81 c9 61 3a a6 49 ae 76 7e cf da 3b 47
2a 85 6a 55 62 b9 cb 23 45 86 6a ca ee dc 5d 29 a3 3b d6 53 37 de fa 5f 07 2a 3c 47 89 17 0d 95
0f 52 99 4d c9 b9 57 f9 a8 2e 8c dd 7d ea 03 67 44 70 1f 72 9c 40 58 1e 00 00 60 1f 51 3a c0 70
1f 73 22 80 f2 19 b9 24 5c ec ed 1e 4e 47 39 b3 33 47 93 51 ce ea 8a b3 dd 00 30 0c a4 14 00 e3
2a 20 07 da ea a0 3a 83 31 00 f0 55 03 80 ac 68 b1 34 57 05 2d 56 3d 55 49 31 9c 47 b2 96 d2 ea
65 15 db 85 f2 82 0b 75 c6 d9 8c 98 47 55 50 bf 30 1f 9b 22 40 48 0b 01 a0 2c 08 4a 96 ac 52 33
7a 33 6a 66 3b b5 8d ec bb 9b 33 e6 8c d5 9b 29 62 6b 98 a3 6e 4a dd 0b 91 ad b5 0a 53 46 f7 5e
7e f4 28 ca 99 97 9e de 4a ae dd f0 00 00 20 1f ab 12 40 58 0b 01 60 1c 00 94 6b 62 5d 9a c9 29
2b d1 79 26 72 5c cb 99 37 26 e5 dc b4 a9 6b cb 92 56 f5 2b 5b df d5 59 67 2b 57 73 cb 46 91 c6
f7 60 ee 0b 63 2d 89 f8 3b 17 74 d5 ea 1e 8d cb c0 50 1f 60 a3 40 78 0f 00 a0 03 a5 2c a4 e1 da
5b 69 9e a9 f8 35 b5 ee 6e cd a9 9a ca f7 27 b1 ea 66 33 e9 55 a5 bb 9e b5 5a c7 1b a9 9c ac 98
b5 ad c8 65 7d 16 2f 6c 22 61 2a c9 4c 1b 67 36 ac 74 cf 08 61 08 53 18 44 c9 80 7d a8 4f 02 00
7e ac 0e 03 f8 50 18 00 0e 02 a7 68 60 1c 72 30 78 00 00 60 1f 8a 1c 80 fb 94 e0 07 e2 a7 20 3e
e5 38 01 f8 b3 44 0e 47 13 c5 cb 2c f3 b1 f6 4d 4a d5 5d 53 9b f7 73 6e ac db 4b d1 0a d9 62 9a
c3 02 ea ee 49 a5 2d c1 58 93 99 0c 8f 2b 46 26 bd c2 a3 bb 3c 00 00 30 1e aa 02 40 48 0b 01 60
34 08 49 26 55 b4 1c b7 b2 76 6e 4c e8 5d ab 1d 93 b3 6e 7b 54 d9 a5 2a 14 b2 35 c6 d8 db 1b 5f
73 6e 4d e5 bf c0 60 1c 6b 2b c0 80 1f 7a ac 00 fc 5d a2 04 c0 d3 ee 6b 5b ac b6 f9 1b e4 fb 08
f2 de c7 c8 e1 36 e7 65 96 76 3e 48 ea 37 33 ad b7 53 32 cd cd 29 34 c5 4d b2 e3 60 eb 71 63 ea
96 98 96 ec df 8b e5 53 66 8c b0 de fb 5f 4c 13 ac 6c 6a 97 4e c0 cd 1c 92 a5 b5 ab 84 2d d2 d9
63 98 67 ef 09 4f 30 1e 62 2a c0 78 11 00 60 00 03 00 55 52 f9 3c 2c ce 8d 52 be cf ca 25 66 65
58 d2 ea a9 c9 4d 62 35 79 ae d6 c6 08 8c f5 a3 ad ca 03 eb 65 88 11 01 e0 44 07 81 e0 03 01 90
35 30 0f 31 61 c0 7e 0a ee 03 f1 68 70 1f 74 cb 00 33 a5 c0 39 24 2c ce 8d 52 be cd c9 e5 66 65
58 d2 f2 79 c9 4d 62 35 7b 9a d3 46 08 8c f5 a3 bc cc 03 ed 25 70 1f 72 24 00 fb 19 d8 05 8c cf
2c 19 07 66 98 ed 4f 09 e5 1c 99 93 20 ce b1 5e 08 65 d5 bc 2a 1a fd 80 6c 66 ab 01 40 74 c4 aa
03 e6 15 40 1f 49 b2 00 c9 19 83 c0 80 1f 7a b4 80 fc 59 e4 07 e4 d0 e0 3a b2 70 00 30 0d 44 37
25 64 25 d5 ea 55 c9 9b 19 87 7b 95 f2 77 46 61 9e a5 7c a1 f2 98 57 8b 67 2a 7c 26 59 62 67 ca
a1 09 a8 48 7d ea b6 4e a6 0b 5e 58 b5 53 ca 72 c7 9d b0 4d f2 60 e5 e7 4c 51 7b 9b 08 79 cb 23
de e2 c1 db 6f c9 99 25 fe d4 45 2a 66 68 9c 34 34 34 bc 50 1f 80 a3 40 88 13 00 60 00 08 04 67
18 40 1e f6 d4 91 31 ba 96 29 ae 65 dc bc 96 14 66 85 7d 67 66 dc fb a3 55 46 ce 8f 2d 75 b5 cf
e1 bc 33 dd 2b 84 f1 96 2f db 94 42 66 e1 94 3a eb 02 72 67 65 0e 1b 49 3c 89 b9 78 50 1f 82 43
00 fc 9a 1e 03 40 00 80 15 73 b7 21 83 a6 19 a2 57 c9 1b 48 c7 39 99 f6 5e 4c e9 c6 41 b5 b3 53
be d0 f4 e3 a5 ed 7a b9 f8 75 98 2d 25 f0 b6 30 b1 3c c3 46 e1 84 2e 0b 0f 12 30 09 01 86 44 38
05 ad d0 60 78 60 1e 83 cb 40 88 0d 00 02 40 55 52 9c 96 11 85 65 ab 5f 25 85 da 11 ae cb d9 fb
46 2d 6d 9a 9d ae 35 0c 19 95 a5 6d 8d 80 01 80 4c 20 c6 07 80 00 00 70 1b 5a bc 40 88 0f 00 02
00 5d f0 fa a1 f2 bd 54 89 27 68 ed 5d ab b6 76 cc 5c db e5 29 e3 75 f7 80 00 20 05 bd 9c 60 07
80 00 00 60 1b 53 a3 c0 88 11 01 e0 62 46 2e 69 9d ed 59 9a b4 0e a0 70 0d ac b6 a5 63 a6 d9 56
d7 c9 9c c9 c6 55 9a 62 65 b6 b1 55 ae 99 9b 34 5a a5 32 ed 59 6e 78 ad 6c aa 46 6b a6 2b 63 2a
55 d6 e9 ab 1a b9 a4 75 ba 62 c7 aa 65 19 6e b6 b5 cd a8 26 6b 62 80 a8 09 de 04 40 88 00 06 01
c8 78 60 16 44 30 78 00 00 20 1e b2 13 40 70 1f a2 0b 80 fd 59 5e 07 28 8a 66 4d e2 d7 da b7 50
30 07 ad a6 d2 2d d9 4a af 6b dc bd 9b 73 62 74 bb 64 e2 ad 59 9e f3 d6 a7 01 da 19 bc 06 01 fc
19 00 48 0d 39 5e 00 69 2e f8 aa 8e 76 39 39 b7 6c ed d8 bd 8f 58 98 5c dd 6f 86 29 83 cc cb 58
9d 8a 24 ed 32 f5 b6 62 a6 d3 b5 25 85 98 b1 8c d6 6b 22 f0 00 00 00 02 01 34 a2 b0 0d 67 2a 20
44 0d 4b 7a 6e 75 4a 8f 90 52 35 02 2a b5 d4 b2 7c 31 44 ad f5 2c 1f 0c 98 ef 7d 4e e7 c2 b6 bb
bf 55 3d d4 8d aa ef c5 4e 9c 1c 61 bd ed 5b 2a e9 38 b2 f9 da eb 9b 92 7a 66 ba 2a f5 51 5e 5a
0c 4a bd d4 45 75 85 8a 2f 7d 99 18 dd d6 00 00 1e 90 1a 33 3c 00 f4 a1 69 0e 63 0a 46 96 4d 49
89 41 83 c7 b7 b2 72 5c 45 4c 62 7c 9b 9f 35 5a bc 34 cc ee 73 72 b2 8a 33 6d 9a 63 ac 62 11 f2
d5 aa ab 1c 64 7c c4 66 f8 00 00 90 17 5b bc 00 b3 ad 18 04 a8 ad 88 e5 65 74 8c 75 f6 4e 4b 4b
ab 14 55 97 b3 6e 8c 57 17 03 6d 4b 2e 37 b3 19 26 56 19 94 60 10 02 2e 72 68 08 02 50 60 0f 80
1a 6c a4 40 68 1c 7e d0 64 27 8c ef 64 ec b8 9b 31 33 4b 9c 96 68 cd 57 b3 37 5a ec 45 0c e6 2a
90 bd 6d 2d 96 a4 a3 70 66 0b 61 74 2d 36 d9 82 56 99 29 2e 25 60 86 26 ca d3 6d 98 3c ac d3 4f
2a 00 0f 09 01 a5 42 38 0b 36 12 91 c6 ae c9 15 1d ec bd 97 b2 f2 68 5a 21 1e ab 9d 9f 94 3a 73
6a 93 64 ed 58 ad 76 ab 55 8c de 6d b5 a9 c5 5a 2c d8 eb 75 2e 46 aa 94 e6 47 0b 8f d9 95 2b 8a
b2 f3 76 1d 52 d3 8e bd 52 68 87 a4 97 80 50 0d 3b 2b 80 c1 61 14 04 a1 08 80 1d 14 4e 47 2c d3
aa 45 77 b2 76 5e cd d9 f9 44 74 ee 6f 11 ce d7 bb 73 74 2e 7b 42 d2 5e c5 1b a1 7b da 18 91 32
97 5e e6 0d a3 d4 a5 97 a9 87 69 17 21 55 e7 72 c3 37 c6 fc 00 00 70 15 53 2b 80 9a 25 75 24 17
2b 0f 11 2e 49 09 d2 b3 c5 7b 92 46 94 8c f1 a5 d4 92 dc 9b 2c 59 1d 3c 24 e0 8b 31 c6 49 a8 ed
30 9a 69 80 7e 8a af 01 c0 7e cc 56 83 fa 94 43 64 fe a7 8b 29 99 1b d7 a9 a2 ba 67 37 f1 ea 68
b6 59 c5 fc 96 9f 47 49 64 69 ab 80 f3 0d 1a 03 c0 58 1a 9a 4e a5 82 3c de e7 dc f9 a2 d4 39 65
36 55 5e 8d 50 03 d7 80 78 0e 03 d5 70 10 0c 22 d2 00 a3 81 a0 06 28 f1 ab 22 e7 74 58 73 6e c9
5c c7 33 a4 2a 9e 57 ad ae e8 cc ed a5 ee 7d 50 3c 38 ef 8d 32 7b aa 0e 1e 52 cc da e5 87 09 8e
b7 24 e6 8d 64 e6 68 17 91 a8 0c a0 d9 80 7e 8a 0f 01 00 79 e2 49 01 40 7e 80 3e a5 93 7b bb 11
39 67 65 2d 44 55 95 6e 47 2b 0e e8 35 6b 76 4e be 57 28 3d 22 b6 97 95 4e f1 26 97 24 ed 5d a3
b4 76 8e d5 da b9 6c ee f4 6b 11 ce dd db b7 7d f0 96 0d 8a 4b 3b b8 e1 83 9e 32 eb 2e b7 60 d7
0c b2 d3 b5 f8 36 f5 dc 6b a3 18 44 1a 6c 80 a7 a3 21 61 5d 39 e0 44 ea 8a aa be 48 99 62 ce ee
ce f6 a6 5c f4 1c ac 9d 89 57 ad ce 84 f3 6a 85 f2 7c 3a 13 4c e9 5c a2 d2 ac de 5e 57 0a bc 24
56 77 95 c7 0f 29 13 5c e5 52 cc d2 34 47 39 54 77 12 8f 2d ce 55 0c c4 2b 5b 6f 95 3e 8f 0c d4
92 d5 52 73 33 a3 38 e9 6b 18 c6 c9 52 50 68 0d 00 e0 07 80 00 00 a5 d4 8a a4 13 3e f9 54 1e b3
0b 5e 5e 50 f8 b4 bb de 77 94 36 12 f3 17 94 e5 0c 84 bc c6 53 50 02 1d 15 a8 e6 f6 57 57 d6 72
c9 42 11 55 fd bc b6 92 74 64 7f 6e ae 9c 59 15 63 eb 8b e9 36 54 56 de 5b 09 c5 94 d1 b6 76 c2
92 65 52 3b f5 af a3 99 1c cf e7 4b e8 98 47 34 bd a8 0e 20 8d 60 30 0f cd 4e 00 60 c6 c6 00 ec
6b 40 0b 81 e6 a8 9d 19 55 6b 47 d9 39 1c e2 cb 07 76 4d 45 26 39 51 46 6e 91 45 e1 8f 2e ac a4
4d 4a 3b e4 36 d5 11 98 f3 ec ee 40 80 11 9b 5c 22 5d 4e da 4a 55 58 a5 12 74 72 f2 57 46 46 6d
24 4c aa d5 91 9f 65 54 e8 ed a4 4b c9 64 2b ce f1 0a e7 2f 49 36 6d 44 f7 53 d2 56 4f 51 be 2f
50 bf 1a e5 8f 8b 4d 36 d6 f7 7d f2 ae 60 00 00 c0 32 e4 54 01 86 2b 18 0a 2d 4f c0 00 00 65 8b
63 a4 16 4a e9 73 da 51 07 17 2e 5d 48 aa 41 33 ef 95 3e 2d 2e f7 9d ed 1d a3 b4 28 04 c1 e0 e0
5a b8 cd 60 5a 73 da b2 11 65 36 7e e5 ad 8a d9 8d 97 67 00 20 55 37 80 d0 00 55 26 ac ef 42 f5
9d 16 9d 58 69 98 0b a6 27 56 6a 9e 03 57 c9 9c fa ca 27 b1 f6 6e d1 ca e9 56 66 39 ad df 1b e5
5c e1 9d 2a 8c 8b 56 20 00 60 13 6c 43 00 dc 0d 03 c0 00 00 85 ad 61 cb 1b d8 9d 6b 58 70 86 fa
18 5b 47 1a a3 05 f3 54 56 8c a8 a7 b4 e4 f6 c4 29 a1 6f 75 3d a4 84 75 3d bc d5 9a af 64 ae bb
90 8b ab b2 ec 22 72 6b f7 86 3e f6 61 27 82 d7 52 2d b8 78 00 00 95 ce 7a c9 93 ca e3 76 2d 93
56 dd 97 6b da b9 54 5c f7 26 f6 5e d5 aa a2 f7 c9 37 92 da b6 4e e5 8b 5a 92 b2 4e 78 71 56 97
80 45 85 74 34 35 c8 d1 61 97 0f 0b 92 b4 56 83 cb cb 53 8d 15 b1 0c ce f0 e3 85 55 3b 2a 32 ca
f9 3c a4 ec a5 76 be 4d 2a b3 9a 75 b3 93 4e ae ca 79 6b e4 d3 bc 32 15 d8 f9 35 2d 0c a5 56 3e
51 3b 43 a1 65 8f 95 4e d0 ca 57 23 e5 72 b3 c2 96 49 39 64 aa f2 a7 72 4e 5b 09 bc 32 54 8e 97
3d a1 aa f2 aa 7d ee bd cd 0d 75 a7 5b 2f 53 07 5d 2d b6 ab d0 d4 b6 4d 89 42 f5 b6 12 17 1b 2e
be 2f 53 52 a7 20 80 5c 51 00 00 93 c9 6f 2c 8e db f0 45 66 7c 33 b6 46 d1 59 9f 0c e9 b1 34 58
56 ca ca 54 90 f6 1a b4 a2 19 da 9c 6e ae 4f 67 80 cb 2e e7 ec fc 9e 75 67 45 cb 67 67 ec fd a3
95 4b 0a ac d7 74 e5 51 ca b3 3b dd 3b 66 ad 87 4a 61 62 c9 9b ad 73 f8 53 2c f6 49 5c fd 8e 3a
5b a2 17 3f 03 50 76 d9 65 d0 e1 33 13 ca 95 6c 32 62 a8 ce c4 5b 19 3a a1 bc 88 3c 00 00 95 af
6b 32 91 cd 23 6c 6c a5 72 63 b2 19 4d 31 54 d3 0c 90 db 2c 4a d7 3a 45 c6 2b 52 41 0e 8d f1 61
c4 97 84 08 54 ec f3 2d e6 ad 14 aa f1 62 18 9d 66 6c b5 d4 92 89 4c e8 ef 7d af ab d8 8e 5f 69
7b e9 7d 1c e6 70 bf 4b 5f 96 34 0a 3c fa d7 fe ac 90 0a f6 80 1d 59 1c 00 fd 54 9c 07 f5 28 c0
3f 33 43 01 14 3a f8 65 88 8a b2 ad c8 d5 69 dd 06 ad 6e c7 59 15 aa 9a bb 92 55 ca 0f 48 ad a5
ed 5d ab 95 4e b3 08 8f a5 ed 5d ab 95 d2 f0 a8 93 a5 ed 9c ba 76 74 57 ac e5 2f 94 c6 1b 36 27
4c 24 c6 48 f7 ad 4b 09 25 56 72 21 70 c2 70 23 ab 98 13 b0 93 50 67 5e 16 cc 60 b3 4b a7 8a 78
02 00 b6 bb 88 07 a1 dc 80 39 2c 68 02 0e 65 e4 91 8b ab b5 68 7b 27 24 9d 1d 15 63 47 c9 6b 79
44 47 d1 f2 6a e6 91 0d 75 7c 9e da 94 34 5d 1f 2a ae a0 8d 9f 45 8a eb 96 05 78 cd 5a ba 58 c2
26 2a 32 b2 10 06 d7 76 64 2c 73 83 e1 9c b9 0b 1e e0 e9 59 32 42 c8 38 3a 16 14 b0 b1 ad 0e 85
94 e4 ac 8a 9c a2 9e 93 0b a5 06 15 56 cd f8 00 10 06 5d c9 80 5b 99 cc 02 da ee c0 1e 87 75 00
e4 b1 b4 0e 47 18 b2 b1 5c 6f 91 d2 6e 6e 97 5a ec 7d 8f 53 53 70 6a 73 62 ed 7d b7 18 53 8c 6a
95 59 86 2c 68 de a4 95 1b 93 72 66 51 78 65 3d 18 db 93 b2 8b c3 29 e4 3a ec aa 85 09 8d 4a 37
26 64 b4 4b 3a a4 b1 00 7e b0 c9 01 20 1c 03 81 e0 00 00 10 0d 2b 28 80 98 dd 0a 01 c0 e4 71 7b
3b a5 e6 f9 14 5b 09 0f 39 9d 49 5c da 31 35 cf b3 76 ad 5b 4d 48 aa 3d a6 98 ad a4 f9 30 e4 9e
2d c8 96 c1 50 ac d4 e3 06 01 3d b2 b8 08 c5 89 40 51 f4 0a 02 91 82 10 0d 7a 1f 80 10 07 3c b9
00 49 6a 0c 02 48 f0 a0 26 d6 4e 3f 27 24 d2 4c cf 91 b9 af 32 d1 73 c4 6e ba 24 cb 4f 33 26 65
cd 1b b7 77 62 f8 65 91 29 27 f9 be d8 47 68 a9 29 77 9e 28 c4 8f 86 a9 a6 d1 b2 33 21 a5 f2 47
0c a2 a5 59 88 b4 3b 2a 51 60 ce e6 22 01 4d 56 f0 00 06 01 65 42 18 06 b2 1b 51 31 b3 14 cb 62
e4 8c 84 3d 5a bf 75 2b 25 09 18 cf ec 4c ca b1 c6 33 f6 93 3f 6a 68 f0 ef 25 04 66 44 4c b4 80
08 02 50 f3 20 2a c5 38 3f 6a e2 a8 24 8f 52 52 4e 84 d3 db c4 92 7a 1b 5d 52 b1 35 98 a6 b5 34
ab 51 67 29 2d 4d 2f 15 da 0a 29 13 52 b5 d6 72 92 4c d4 e9 7d 9e a4 93 34 b8 61 37 12 34 d4 29
d8 c1 84 cf 74 94 76 2b 62 35 cb 25 1d 9b 1a 31 90 ad c8 67 29 84 52 ca 71 da 4e c1 0e f6 a4 6e
81 e0 04 01 46 ba ac 05 80 d0 39 1c 51 4d 2d 16 ae 48 e7 3b cd 34 f3 52 32 8a d7 2b 3a cc 94 91
fa 08 76 9b 90 ea 08 01 67 31 34 0d 51 4c 29 9b 55 b3 16 d7 30 a0 6d ee 7d cb c0 f8 97 1f 25 67
4c 4d 9c a6 cd 59 93 31 49 0e 24 58 a6 42 b9 54 c5 56 b1 6a d0 13 6a 55 ac 53 2c 13 de dd 65 20
c8 aa b6 4b d5 c9 a1 04 87 7b c6 ee 3d 49 17 a4 bb 56 ab b2 ee 7a cf 5a 98 07 96 68 f0 1a 03 c0
78 0d 01 a0 34 0f 10 0d 3b 18 80 82 dd 0a 01 c0 68 12 81 76 39 59 49 a4 e6 ac b6 91 58 4a 48 5b
ee 23 5c d2 d6 95 f5 a9 1b 04 a8 68 7a e2 58 2d 5e 70 6d d9 9f 35 da f9 29 5d f8 66 f0 c1 89 7b
9d 0b 63 72 5e 6b d1 3a 78 00 78 09 01 c0 4e 14 e8 01 b9 86 78 1c 8d f5 ba 67 7a e7 64 ec 9d 97
b2 f2 68 66 69 5a 23 dc aa 5a 98 54 98 ef 6a e5 93 bc aa ac d5 37 6e 6e 95 cd 70 db 23 58 9f 7b
2e 5b 87 18 e9 da db 97 73 2f 4b 8e e0 a6 4b e0 90 16 44 2b 00 7a 65 79 3c a4 ee 8e b7 4e 55 38
c3 14 2d fa d6 4e ae 63 29 36 75 b2 b9 0a ca ea b1 54 f5 06 2e bb de 55 4b d3 9a 9d fb 95 4e 6e
ea cf b4 e5 11 93 ca ac 69 3b 46 28 73 9a e0 e7 46 eb 62 e5 77 46 92 66 ed 5b 37 b3 bb 15 de 98
3e 4b 46 cf a6 a6 2e 8c 39 4c 57 61 92 e8 af 35 32 66 6a ca 9b db c3 91 1a b2 0c d2 f6 d3 f0 90
19 53 4b 80 d2 da 90 04 23 4d c8 60 ca b9 44 9a f6 3d 47 05 5e 39 a4 b7 12 41 39 ac 4b 65 c4 ae
4e 6a 9b 49 69 33 53 ba ac aa 5a 54 c4 ed b4 13 b2 55 b8 f7 8f 02 eb 95 ab 6e 4b b1 46 e1 72 65
70 ea 71 d9 5e e5 cd 41 3b 79 57 ae ab 46 90 dc 55 ed 81 ca 9d 3b 57 7a af 54 15 d8 ee 46 8c 12
e7 76 49 96 1b 03 ae 5e 1a 5d 7e 00 00 00 08 01 26 c2 c0 0b 9a 55 80 51 58 b8 8a 0c bb 83 6b 27
63 ec 9d 97 b3 76 ac 5a d9 e5 40 c5 8c 40 12 12 e0 0a 01 29 89 38 00 01 e0 00 00 70 14 85 45 40
70 14 b6 33 c0 b4 6a c4 2c 59 4b b1 22 e5 8e d6 1b 4e 48 f6 65 35 0d d3 92 41 b9 52 89 64 ec 9d
97 93 c9 b7 50 ab 63 ed 5c b6 33 87 67 7c e6 ee cd f9 be 98 42 8a d1 17 0d a6 11 83 a3 dc 43 25
84 1c ad 33 50 37 61 2a 21 2b dc 2f c0 00 00 70 14 85 45 40 70 14 b6 36 23 6d d5 a5 d1 db a8 d7
77 7a 73 d6 72 47 b7 21 9d 64 7d 8b b0 ea 08 d1 d9 d7 22 9b 83 70 ea 38 d1 99 db 62 9c 96 71 65
77 ac de 26 ac 64 4f 22 e5 29 e7 24 06 d8 a8 aa 73 a2 66 ac dc 94 a2 8b 33 76 87 66 26 ab 9c ca
2e d9 c8 e7 07 35 aa cd f2 3a ba 19 a2 33 7d 8b b1 76 4e cf da b5 6c 62 f2 a9 31 4c df 7b e1 7b
66 ac eb 0e a7 60 f2 ca 33 3d a9 58 25 ab 10 98 e3 f0 50 17 5c ac 40 40 0f 9c 21 00 44 ad 39 14
57 8c 7a ea ee c5 d8 fb 27 66 d5 11 7e a0 e4 1d f3 56 2a 8d 30 ce d5 6f 9a ef 5c ac 94 02 d2 1a
47 3b 2b 6c ad 8b 82 d6 9a 35 3b 7c 00 00 30 0e cc ba 80 75 26 21 01 ab 7a 88 5b 58 c9 a9 59 19
a9 27 b6 f6 5d cb d9 75 3a e5 92 e3 52 4d d1 ba b7 5e 2b 5c 32 d8 9a 67 9a ef 64 ac 5b a6 d0 d2
38 cb 12 c9 84 56 56 d2 d9 b2 19 45 9b 8a 01 cd 95 30 0c 84 89 00 4b e5 7c 50 0c 84 3b c0 50 0b
8a b8 6a 25 cf 29 19 a7 9c 96 30 a9 57 3b 27 25 9c a1 d9 9b 3b c9 69 58 77 19 af 5a 86 bd 5d 93
9a 8c a1 48 75 c7 66 d2 2a 9b d8 e6 1d cb 00 28 c8 37 01 36 c9 52 39 e1 c9 0a 2c 76 ae 55 08 4c
e8 a4 af a6 45 8a b3 37 0c 23 11 3a ca 81 b7 1f 40 13 53 3b c0 60 12 6c 3d 23 9b 25 a5 52 35 c9
5d 4a 98 3a 8a f2 58 ca a6 4d f3 9c 9a 35 89 83 ab 27 28 96 21 9d 66 b9 aa a7 96 55 5b ae e2 ca
5d 15 9a aa b0 b6 12 36 77 8e 98 00 08 02 4e a0 40 26 31 22 01 71 95 10 0a 94 2f 80 00 00 a5 5b
b0 19 c4 f6 2f 56 aa cc 95 41 63 d7 89 af 16 40 88 da e6 6c 4c d4 45 24 ae ae f3 de ab d9 6b 5a
86 a7 92 bb 56 e9 df 1b df 7a ee 45 cd 34 a2 a4 b1 3c 00 00 d4 ad 8c 62 3c b6 b9 1c 1d 56 e7 3a
3e 3e ca 53 b3 eb 57 6e 5d a2 00 b5 d5 00 08 04 56 62 10 3d d3 8c ad 5c d4 ef 5c ec 39 5d bb 47
6a ed 5d ab 55 31 55 d0 e3 2d cd 58 a9 6a bf 46 49 6d aa 71 af d9 94 df 4a 93 1c 78 62 72 22 a3
d9 ed 8d 93 9e e5 13 5c d5 2a 59 7b 47 68 e5 12 7c d2 93 e7 29 43 20 92 4d 2f b8 cc f9 dd 52 5e
94 eb 91 b6 54 66 c3 4b be 86 96 da ad 44 66 9c ed 7c b6 32 87 26 6b 7e 2e 74 dd 13 62 4f 7b eb
8c 31 75 8d 22 da 6a ec 59 81 7a b8 71 bb 17 e1 49 62 5d 6a c6 39 84 14 45 6f 80 00 00 d5 0f 93
f1 c5 38 f9 53 e5 54 ab 15 ce d5 ca a5 0b 85 27 b2 4a a5 c1 60 a2 6c 90 a9 10 5b 67 b7 0b 2a 1d
5e 2d a5 23 4b 22 b5 53 68 8e 5a b7 d5 90 ed e5 f9 af 36 62 c8 65 d4 65 9e d2 02 e7 05 80 1f 92
ac 40 70 1f 8b 1a 00 fb 0c 41 00 1e 8f 5a 8d f1 76 67 c8 ef 27 8c e1 95 b2 cb d9 f7 46 6a cd 58
b1 f4 67 67 b6 9d ef a5 ed b3 2b b5 b5 2b 7c 2f 6b 8a d9 a9 66 fb d9 79 9e d9 49 49 39 de eb c6
fa 99 67 4c c3 0a 29 d1 1e 36 0d c3 c0 00 00 e5 4e 7d 5b a4 d7 39 54 1d 54 e7 55 be d5 6a 98 e5
76 56 b3 4e a9 d5 2a da d1 b6 79 59 b6 d4 5f 06 b2 c2 2b 8e 40 55 52 92 03 49 94 40 09 7c 4e a4
93 9d 19 df c9 d9 3b 2e a6 8c 61 95 9b c9 a9 de f7 a7 57 d6 6a a6 ba 31 d6 2e 52 ac f9 bc 42 b9
54 ec 8d b0 b4 fb bb 34 99 5b ad 37 2e 18 69 c6 ab 4b dc 36 49 8d aa d2 db 65 82 9c 9e 00 00 d4
8d 94 e2 3d 58 fb 36 a6 65 2b 11 72 c7 ba 2d 52 a7 5a e9 36 b7 d4 00 06 02 32 75 10 0a 73 b0 c0
f0 00 00 e4 f3 74 6b aa 4f 39 33 a1 58 6f 32 ab 4c aa ce 15 53 6c 00 80 12 36 ee 02 32 76 94 cd
d1 53 ba 6b 6f 55 ab 6e c6 25 93 47 5b 65 73 74 4b 52 7a 38 5e cc 8d 5a bd 65 d8 3b 06 95 09 29
f0 00 00 b5 5c bc c0 01 ed b5 45 ed 92 06 73 fe 51 79 4b 03 1d ff b4 72 8a b5 94 e4 f7 d9 a3 35
5a aa 99 2d 65 2f d7 ab 35 66 bb d9 4b 1c 87 85 74 f1 d6 c2 00 25 fc 20 0d b0 41 00 7e ca 06 03
ef 34 30 1c 68 b7 80 00 00 a5 52 7c 22 1a b9 b5 55 25 08 6a 96 5e d1 da 39 45 a9 04 2b 1e 5d 51
69 39 13 cf 97 15 5a 6e 23 15 e5 c5 56 8b 88 cc f7 2d 5d e0 c4 55 5d bb 59 68 28 94 d7 72 16 56
28 29 55 94 f0 00 00 65 ab 5b 32 9c 4d 65 74 1d 10 a8 db 3e d5 ca 1b 28 77 7c b1 f6 7e cb d9 79
34 25 2b 0e da bd 54 eb ba c3 b6 b3 b6 5a d6 68 ea 19 2d 97 bb 17 bf 6a eb 27 77 b6 12 b9 4a 44
5b 69 94 6c 86 d5 1e c9 63 25 b1 ad 56 72 38 bb 29 3d 81 6e ea 6b c9 79 2f 1f 75 8c 63 b1 32 6c
a9 64 dc ea ae fb 5e cf d9 fb 3f 26 93 dd 55 e6 eb c9 e5 28 54 86 9f 72 87 e1 cd e0 e7 db ab 75
ea d7 e6 11 60 e6 fa ba 17 74 58 4b be 70 a6 11 92 a2 e3 fa e9 8c 9e a6 d1 3a a9 e5 3c bc 00 00
35 83 55 45 4d 38 f0 e6 19 10 6a f2 ee 50 da ab c5 5d ab 13 2e ed 2f 5b 1a c4 aa b4 cb 35 e6 63
2b 27 43 fe 15 5e d3 7a 31 52 95 b6 ce 97 ac 54 d5 cc b4 0f f3 56 55 8f 27 04 ff b5 b9 79 38 ad
da ef 74 af a3 d5 42 df d7 7b ed 7c a8 86 94 db ca 61 19 2a 27 27 d2 98 35 ae eb a9 f3 86 0b 6b
bb e2 6e dd 8a 61 2b 4c db b6 62 a5 db 5c 2b f5 98 b9 8e 99 2a ec 55 af 5a a5 3a 29 3c 65 86 6d
7a 28 bb 1d 59 9f 7e 68 4f c7 54 6a de a2 16 d6 d4 29 f1 b0 c7 b4 e4 f1 8b 3b c2 6d 39 44 66 ac
cf 39 ce 55 38 29 ad 4e 6f 35 de d7 5e d6 1c 4d 5c 31 77 80 00 38 0a 99 e1 40 59 54 9e 00 00 85
6a 3c 7b 99 b4 dd 59 db be e2 52 3b 56 88 7f b0 25 b7 15 31 7f aa 4b 66 dd 5c aa 2a cc 83 4b 67
6b ed 9a b5 d1 c9 63 7a e5 ee 95 cd a4 2c 25 45 25 73 d5 0c cc ae 58 de 00 e0 1e a7 24 80 1e 00
00 55 43 8d fa 8b a7 65 59 df 5e e0 92 eb 54 97 d7 b0 2c f7 15 31 93 ee 0f 6d b5 0c 7d fa 0d 58
b3 3e 28 66 b1 12 25 9b 7a af 64 ad 66 2d 4d a9 9b 0b 5b d5 65 79 51 b2 e5 46 8e 0e db 46 9d cd
55 82 79 28 01 00 82 1d 50 06 4d 30 80 3f 76 73 01 fc 32 10 0f b1 54 40 7c e8 fe 00 00 35 a3 c6
c4 09 b9 13 6c ad 4b ee 91 0a f9 aa e5 03 39 ba cc ee bc 57 67 b1 86 5e 4f 35 e6 ac d5 9a f1 65
aa c6 0d 5d c9 d8 bb 13 63 2d 93 81 ae a9 99 e0 55 83 96 54 2b 19 1d 59 9d 7e aa 6a 4a 56 96 57
a1 8c d7 55 2d 73 ec 49 3d e5 4d 74 e2 2a ed 39 44 e6 ce cb 1a 5e cf ca 29 37 36 7a d1 f6 ae d5
da fb 67 2e b4 5c 92 27 47 bb b1 7d e7 04 4f 5e ab 5f 5a 21 2c 66 4a 57 d2 26 31 39 43 75 ef 3a
ae 63 c1 15 7a ce ed d6 a8 d5 5e c5 1c 73 3b 55 17 31 2c 56 cc e5 f0 01 48 2b 64 aa b2 91 75 42
a2 da f0 4d e6 76 9d 12 af 54 5b e7 26 63 de ed 62 d7 d9 bb 2f 66 ed 1a a9 72 9d 74 c8 e5 ac 4e
72 5d 62 a9 7b 6f 74 ae 3f 72 4d 96 55 1b d9 79 bb 70 2f 2e a6 5e 6e da 2a 53 8f c0 45 63 86 5c
2b 18 dd 59 9f 7e a6 2e 59 56 86 d7 c2 84 bf 55 b5 53 71 00 e7 d5 6d 73 c5 61 b9 f5 5a a7 11 12
53 7d 54 c9 c9 c5 23 ff 35 62 a7 4e c9 35 d6 f9 aa f5 5e ab 54 eb 59 26 3a df 15 bf 16 69 48 b7
a5 8f c4 12 d9 bc dd 6b dd 04 50 ca c5 5d 08 49 91 ca a9 16 c5 70 ca d0 e3 f0 00 00 64 25 c7 c9
2a c9 25 22 9b fc 88 b2 db 48 b7 7e 32 2d 77 12 b2 3b 4c 93 24 dc d8 a1 f4 c7 74 8b 27 6a ad ab
bc 14 03 cd 50 d0 0a 07 80 00 95 46 b5 f8 1c 5c f9 52 e1 3e 48 b6 ce 54 e7 c7 b2 bc af 94 36 f3
26 f5 23 ec f6 aa 0b 96 49 57 24 00 9a a1 d8 04 21 2c 60 1f 37 41 00 b7 33 4a 83 56 62 56 64 7c
ae 2a 99 46 4b be 2c a2 22 8d 4e 7e e0 16 3a c7 00 f0 36 b0 0a 00 98 80 2c 8e 6e a2 d1 c8 66 44
ee 2a 7b dd d5 4f 57 ca 9f 08 86 56 f6 72 a7 59 e1 63 32 9c a2 31 a6 66 7f 1e a6 7c e5 5e 22 97
a9 a1 3a 57 78 a9 e2 88 d2 99 96 32 80 e0 30 c8 85 81 e0 00 00 a4 a9 87 59 33 6a b1 23 a5 70 a8
bd cd 49 1b 3b aa dd 57 72 6e 5e cf da bb 66 ae 95 a6 90 de d7 0b d9 15 63 2b 95 47 08 e0 fc 19
83 9a 36 8d 6e b6 e2 dc 7c 00 00 00 94 c6 ad ea 1a cb 71 32 e3 36 aa d7 dd 4d 19 ba ab ae f3 b3
ee 8d d5 da f9 64 69 26 b4 d3 ed dd 8b e1 87 47 a4 be e6 fb 60 fb aa 45 ad d2 58 3f 25 15 0b a2
76 10 e8 cc 24 e6 5d 84 36 a8 ab 0e 35 61 1d 23 20 a3 97 c0 30 0a 93 ca 80 64 dd d6 04 a2 3c 2e
91 8e 59 58 b5 1c 1e ea d7 1a 5d c9 d9 f1 54 64 e4 f3 19 b4 d4 00 0c 03 bb b7 58 0b 01 60 24 04
81 e0 00 00 10 0c ab b0 52 54 4b 3d 48 e5 58 95 6f 9f 73 3b 66 a5 6c ab d0 e2 db c9 5d 5a d6 47
b2 f2 58 5e 6d d2 2c 9c 9a 58 99 64 5b 2f 28 ae 21 95 9e 5b da e5 6c a0 69 11 54 d0 00 02 00 6a
da 38 0b 2e 98 c0 1d f2 de 00 10 0a 8c a8 80 5a ac d8 02 9b 69 49 0f 18 98 47 91 6a 27 26 25 dd
67 98 89 88 9a 88 59 dd a3 62 26 d6 0f 69 79 29 33 c8 f8 d4 53 7a d5 5b 2b 6c 2e 61 ea 91 c6 fe
eb d0 bc 94 73 d6 2b 02 53 1d 59 ec 6a c5 2b a7 65 56 e7 80 30 08 84 41 80 4b ee 64 3c d6 55 c8
8d 4e 4e c7 4f b8 a5 b3 93 39 b3 ee 49 ac e4 ee 7d f2 94 59 37 47 6a ed 5c ad f1 ba 66 69 2f 6e
e5 d1 9b d3 b3 59 3b 7e af 95 5d 99 e7 29 9b f1 84 a8 ca b5 35 4b 61 28 9b bb e4 12 58 4d e6 b1
53 02 76 14 b1 13 35 53 40 00 09 74 19 2e ed 2e cf 34 a3 85 d4 bc 26 91 60 a9 97 08 c5 a6 54 67
67 99 8c 96 95 a1 95 ec 83 75 d5 4c 6c d9 b2 ef b9 53 9a ef 29 59 be d1 da 35 53 59 72 e9 1d bd
d5 ba f1 63 1b 5a e1 36 5c db 7b 69 6a 23 9a c5 0e 47 5a 6e dd 18 d3 79 d7 1b 97 22 b0 e3 65 c6
cd d1 1c b8 d5 71 db 94 e8 e5 bf 65 a6 75 f3 28 36 db 6c 6c a5 4c 3c 7b b0 29 f5 53 51 38 88 d2
fd 4e c6 3c c5 35 d3 53 b9 8d 11 d1 99 cd 19 ab 36 4e d1 00 75 ba 10 07 b1 d0 80 41 4e a4 02 2a
76 20 11 5b b1 80 6b 5d bc 00 00 85 86 6e fb 01 c7 e5 51 7b 2e 4d 13 ba 54 7e a3 3b 4d e6 94 1f
4e 93 09 72 a5 09 ac a4 e1 cd 2b 45 6a b5 4d a4 0c 69 cf 2f 66 2c 84 ec 8e d9 7f 9b 75 6c 63 66
52 eb fd 5b 28 5a 0b d3 be 96 c7 30 63 0d 27 55 a5 8a ad c6 26 d7 72 2e 24 1d 62 ed b4 bb d7 79
ef 00 00 25 92 64 b4 32 37 3c 55 03 8e 55 2b 27 1d 59 df be a2 52 7a 54 69 dd b3 1b bf 15 26 ef
0c e4 ff c5 09 9c 34 da 2f ed 42 a9 46 f6 8b fa 54 9b 39 74 15 aa 75 4a b4 d2 0b 6c a5 74 b1 b7
71 7d 47 ed 6c b5 af bc 8a 63 2f ed 6c 2f 06 55 0b 76 5c 29 3b 4c d3 69 36 a6 d4 5a 38 c2 c5 5b
6d 57 9d 12 f0 f4 e4 40 1e 82 b2 c0 68 0b 01 a0 28 0e a5 17 c0 55 83 a6 e3 9a b6 a1 6a 17 5e ca
0a ca 5a a6 57 ba 82 9f 55 b1 93 91 26 97 e5 4e 55 4b da a5 f9 53 1d 52 94 c7 7c 54 cb 49 45 a2
ff 35 5a a6 d2 c5 71 cf f4 a9 d8 91 5c 55 ed 2a 93 65 dc d9 48 01 00 1b 16 37 81 e0 55 a4 95 db
2a 28 e5 72 5d 5b 0c 0e 7c 56 99 55 b3 13 bf 95 2e 14 ef 46 a7 e5 4d 6e 42 5a a9 f9 53 25 88 b6
c7 7d 50 e9 d1 2e 2a fe b4 20 06 a1 1d 20 24 08 00 d3 33 ea ac c6 18 59 b5 7d b3 76 da eb 39 8c
e6 22 54 be 0d 34 a9 76 a3 2f 3b 61 f1 5d 94 cb ce d6 8c 38 6d 2a f3 b1 ab 4e 5b 68 bd 2c 59 e3
86 d2 2f 6b 02 6d 25 c6 6b 9e c2 58 36 6d a2 f9 cd 55 c5 1b 81 e0 00 00 45 a6 84 bb 9c 38 d1 59
71 09 0d 2e 35 56 5c 3a 45 c2 91 96 1f 4a 71 8c e4 95 8b e0 cd 2a 3b f5 5b 23 58 e6 b9 ee d7 ca
a1 39 97 3b 8e 76 ae 55 1a cc b2 5c 77 95 cb 0f 26 f3 5d e5 93 a2 23 45 54 b3 74 af 94 cd 22 ea
86 cc 18 96 49 97 68 00 02 80 ba 9e 3c a5 14 8d 32 34 7b 3b 57 2b 95 21 99 6f 39 da fb 5e 6c 8d
a0 00 00 0c 03 d2 76 50 1f aa ba 40 38 07 00 60 0c 0f 45 c3 7d fa 83 37 5d 61 dd 3e c8 4e ca 58
88 df 9a 1c f3 55 ae bb 4a 91 65 dd 7b af 95 c1 99 ac af 54 ed 7a ad f0 ca 66 8a a6 ab 8c 36 16
56 b5 ba b1 54 2a ac b5 6d bb d5 3a a5 54 a2 8a 6f 51 26 d4 ba bc 00 00 75 db 90 22 3d f4 a3 74
ae cb c1 51 f3 a5 4b 30 b4 52 7d f1 f2 aa b5 d4 e2 f5 9c a2 2f 68 78 9d 1f 28 83 de dd 9e d7 8a
1e e7 c9 38 b2 52 a8 35 f6 90 ed d7 b3 76 5a e5 4e 7d 55 2d 95 be 17 a6 79 c6 c9 6c 66 07 c5 d2
31 da 95 81 e9 b4 8c ab 25 60 d8 24 e3 39 6d 17 be 28 6e 6f 2c f0 00 00 55 59 70 2a 56 64 e3 65
6d d5 b9 ab b0 1b 7b 79 65 22 ea b1 1e ce db cb 69 37 75 67 d2 ea d8 d9 e4 d7 5b b9 ae d5 43 11
22 d6 dc 85 4d 8b aa 35 c4 23 3b 00 fb 11 96 02 c0 48 07 00 a0 14 01 81 e0 75 f9 68 31 ce e2 a3
85 2f bc 19 02 2f b9 ab eb 26 41 8a f1 f2 fb 56 58 94 fe fc b6 95 68 66 7f 2f 6e ed ba b2 33 5a
47 ac e5 2b 8c 53 15 b6 42 69 d2 a7 6b 67 71 00 28 09 ad d9 00 4a 32 61 e0 00 00 64 d4 41 2a 6e
d2 e5 6e 20 28 6d 5d ba 5d 79 12 0c 57 6b 96 56 ce 80 ef ed e5 56 b5 30 30 df f1 34 5a aa b4 bf
4a 46 76 5d 60 a5 9a 71 4e 38 c8 e9 8c ef 86 34 b2 da 76 26 71 98 66 bc aa 68 79 6a 1b af 2b ac
65 56 13 59 db 3b 6f 6d ed b2 b9 2d 66 77 7b 1b ed 55 a5 83 43 4d 44 d5 61 5e f4 b1 89 b4 56 29
c4 bc db 71 15 0d d3 72 d4 9c f0 00 00 65 58 58 32 4d f4 9f 54 6a 9d 58 aa 2f 57 48 6d a5 ea b5
4d 8b 45 39 a9 ef 5d ec 95 ca ca 4e a2 cd 67 84 b0 66 91 f5 92 79 2c 17 84 6f 39 59 53 06 dd 13
59 d5 6e c2 1a 23 85 85 9b 30 54 99 61 dd b4 8c 19 15 38 67 76 07 80 00 00 a5 52 24 33 1c ce f1
5e ce ec a6 6f dd d9 6b 3e f6 56 15 77 6a 8b b2 95 0d 2f 9c 9e 93 94 43 6d 9f 26 ac 9c 4e 27 c9
d9 b5 3d 62 82 93 79 cd d5 9a b3 65 ee 85 d3 71 1c 5e 62 fc 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00


*/