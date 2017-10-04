// This file is a GPT of a ChromeOS file system, which has
// 12 partitions (more than most) and is hence useful to test
// basic functions out. It includes a lot of data we don't
// care about and some we do.

package gpt

var block = map[int64][]byte{
	0x00000000: []byte{0xde, 0xd7, 0xd2, 0xb2, 0x38, 0xd7, 0x46, 0x98, 0xa2, 0xec, 0x3f, 0x18, 0xc6, 0x01, 0x53, 0x37}, //|....8.F...?...S7|
	0x00000010: []byte{0xfb, 0xfc, 0xbf, 0x00, 0x06, 0xb9, 0x00, 0x01, 0xf3, 0xa5, 0xea, 0x1f, 0x06, 0x00, 0x00, 0x52}, //|...............R|
	0x00000020: []byte{0x89, 0xe5, 0x83, 0xec, 0x1c, 0x6a, 0x1e, 0xc7, 0x46, 0xfa, 0x00, 0x02, 0x52, 0xb4, 0x41, 0xbb}, //|.....j..F...R.A.|
	0x00000030: []byte{0xaa, 0x55, 0x31, 0xc9, 0x30, 0xf6, 0xf9, 0xcd, 0x13, 0x5a, 0xb4, 0x08, 0x72, 0x17, 0x81, 0xfb}, //|.U1.0....Z..r...|
	0x00000040: []byte{0x55, 0xaa, 0x75, 0x11, 0xd1, 0xe9, 0x73, 0x0d, 0x66, 0xc7, 0x06, 0x59, 0x07, 0xb4, 0x42, 0xeb}, //|U.u...s.f..Y..B.|
	0x00000050: []byte{0x13, 0xb4, 0x48, 0x89, 0xe6, 0xcd, 0x13, 0x83, 0xe1, 0x3f, 0x51, 0x0f, 0xb6, 0xc6, 0x40, 0xf7}, //|..H......?Q...@.|
	0x00000060: []byte{0xe1, 0x52, 0x50, 0x66, 0x31, 0xc0, 0x66, 0x99, 0x40, 0xe8, 0xdc, 0x00, 0x8b, 0x4e, 0x56, 0x8b}, //|.RPf1.f.@....NV.|
	0x00000070: []byte{0x46, 0x5a, 0x50, 0x51, 0xf7, 0xe1, 0xf7, 0x76, 0xfa, 0x91, 0x41, 0x66, 0x8b, 0x46, 0x4e, 0x66}, //|FZPQ...v..Af.FNf|
	0x00000080: []byte{0x8b, 0x56, 0x52, 0x53, 0xe8, 0xc4, 0x00, 0xe2, 0xfb, 0x31, 0xf6, 0x5f, 0x59, 0x58, 0x66, 0x8b}, //|.VRS.....1._YXf.|
	0x00000090: []byte{0x15, 0x66, 0x0b, 0x55, 0x04, 0x66, 0x0b, 0x55, 0x08, 0x66, 0x0b, 0x55, 0x0c, 0x74, 0x0c, 0xf6}, //|.f.U.f.U.f.U.t..|
	0x000000a0: []byte{0x45, 0x30, 0x04, 0x74, 0x06, 0x21, 0xf6, 0x75, 0x19, 0x89, 0xfe, 0x01, 0xc7, 0xe2, 0xdf, 0x21}, //|E0.t.!.u.......!|
	0x000000b0: []byte{0xf6, 0x75, 0x2e, 0xe8, 0xe1, 0x00, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x4f, 0x53}, //|.u....Missing OS|
	0x000000c0: []byte{0x0d, 0x0a, 0xe8, 0xd2, 0x00, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x61, 0x63}, //|.....Multiple ac|
	0x000000d0: []byte{0x74, 0x69, 0x76, 0x65, 0x20, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x0d}, //|tive partitions.|
	0x000000e0: []byte{0x0a, 0x91, 0xbf, 0xbe, 0x07, 0x57, 0x66, 0x31, 0xc0, 0xb0, 0x80, 0x66, 0xab, 0xb0, 0xed, 0x66}, //|.....Wf1...f...f|
	0x000000f0: []byte{0xab, 0x66, 0x8b, 0x44, 0x20, 0x66, 0x8b, 0x54, 0x24, 0xe8, 0x40, 0x00, 0x66, 0x8b, 0x44, 0x28}, //|.f.D f.T$.@.f.D(|
	0x00000100: []byte{0x66, 0x8b, 0x54, 0x2c, 0x66, 0x2b, 0x44, 0x20, 0x66, 0x1b, 0x54, 0x24, 0xe8, 0x70, 0x00, 0xe8}, //|f.T,f+D f.T$.p..|
	0x00000110: []byte{0x2a, 0x00, 0x66, 0x0f, 0xb7, 0xc1, 0x66, 0xab, 0xf3, 0xa4, 0x5e, 0x66, 0x8b, 0x44, 0x34, 0x66}, //|*.f...f...^f.D4f|
	0x00000120: []byte{0x8b, 0x54, 0x38, 0xe8, 0x22, 0x00, 0x81, 0x3e, 0xfe, 0x7d, 0x55, 0xaa, 0x75, 0x85, 0x89, 0xec}, //|.T8."..>.}U.u...|
	0x00000130: []byte{0x5a, 0x5f, 0x07, 0x66, 0xb8, 0x21, 0x47, 0x50, 0x54, 0xfa, 0xff, 0xe4, 0x66, 0x21, 0xd2, 0x74}, //|Z_.f.!GPT...f!.t|
	0x00000140: []byte{0x04, 0x66, 0x83, 0xc8, 0xff, 0x66, 0xab, 0xc3, 0xbb, 0x00, 0x7c, 0x66, 0x60, 0x66, 0x52, 0x66}, //|.f...f....|f`fRf|
	0x00000150: []byte{0x50, 0x06, 0x53, 0x6a, 0x01, 0x6a, 0x10, 0x89, 0xe6, 0x66, 0xf7, 0x76, 0xdc, 0xc0, 0xe4, 0x06}, //|P.Sj.j...f.v....|
	0x00000160: []byte{0x88, 0xe1, 0x88, 0xc5, 0x92, 0xf6, 0x76, 0xe0, 0x88, 0xc6, 0x08, 0xe1, 0x41, 0xb8, 0x01, 0x02}, //|......v.....A...|
	0x00000170: []byte{0x8a, 0x56, 0x00, 0xcd, 0x13, 0x8d, 0x64, 0x10, 0x66, 0x61, 0x72, 0x0c, 0x02, 0x7e, 0xfb, 0x66}, //|.V....d.far..~.f|
	0x00000180: []byte{0x83, 0xc0, 0x01, 0x66, 0x83, 0xd2, 0x00, 0xc3, 0xe8, 0x0c, 0x00, 0x44, 0x69, 0x73, 0x6b, 0x20}, //|...f.......Disk |
	0x00000190: []byte{0x65, 0x72, 0x72, 0x6f, 0x72, 0x0d, 0x0a, 0x5e, 0xac, 0xb4, 0x0e, 0x8a, 0x3e, 0x62, 0x04, 0xb3}, //|error..^....>b..|
	0x000001a0: []byte{0x07, 0xcd, 0x10, 0x3c, 0x0a, 0x75, 0xf1, 0xcd, 0x4b, 0xf3, 0x63, 0x86, 0x27, 0x68, 0x49, 0xee}, //|...<.u..K.c.'hI.|
	0x000001b0: []byte{0xa8, 0xb0, 0xea, 0x11, 0x28, 0xf3, 0x17, 0x07, 0x00, 0x00, 0x00, 0x00, 0x1d, 0x9a, 0x00, 0x00}, //|....(...........|
	0x000001c0: []byte{0x02, 0x00, 0xee, 0xff, 0xff, 0xff, 0x01, 0x00, 0x00, 0x00, 0xbf, 0x5f, 0x69, 0x00, 0x00, 0x00}, //|..........._i...|
	0x000001d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000001e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000001f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0xaa}, //|..............U.|
	// first GPT is here.
	0x00000200: []byte{0x45, 0x46, 0x49, 0x20, 0x50, 0x41, 0x52, 0x54, 0x00, 0x00, 0x01, 0x00, 0x5c, 0x00, 0x00, 0x00}, //|EFI PART....\...|
	0x00000210: []byte{0x92, 0x92, 0x51, 0x22, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|..Q"............|
	0x00000220: []byte{0xbf, 0xcf, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|..C.....".......|
	0x00000230: []byte{0x9e, 0xcf, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2d, 0x1e, 0xd4, 0xba, 0xef, 0x93, 0x4a, 0xb0}, //|..C.....-.....J.|
	0x00000240: []byte{0x84, 0x6e, 0x2e, 0x3a, 0x6a, 0x2d, 0x73, 0xbf, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.n.:j-s.........|
	0x00000250: []byte{0x80, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x57, 0x8e, 0x72, 0x8d, 0x00, 0x00, 0x00, 0x00}, //|........W.r.....|
	0x00000260: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000270: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000280: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000290: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000002a0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000002b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000002c0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000002d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000002e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000002f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000300: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000310: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000320: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000330: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000340: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000350: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000360: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000370: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000380: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000390: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000003a0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000003b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000003c0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000003d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000003e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000003f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	// Partitions start here.
	0x00000400: []byte{0xa2, 0xa0, 0xd0, 0xeb, 0xe5, 0xb9, 0x33, 0x44, 0x87, 0xc0, 0x68, 0xb6, 0xb7, 0x26, 0x99, 0xc7}, //|......3D..h..&..|
	0x00000410: []byte{0x93, 0x6b, 0x59, 0xaa, 0xa8, 0xeb, 0x42, 0xe0, 0xbb, 0x1c, 0x50, 0x0b, 0x28, 0x46, 0xdb, 0x57}, //|.kY...B...P.(F.W|
	0x00000420: []byte{0x00, 0xe0, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x4f, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00}, //|..B......OC.....|
	0x00000430: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x53, 0x00, 0x54, 0x00, 0x41, 0x00, 0x54, 0x00}, //|........S.T.A.T.|
	0x00000440: []byte{0x45, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|E...............|
	0x00000450: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000460: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000470: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000480: []byte{0x5d, 0x2a, 0x3a, 0xfe, 0x32, 0x4f, 0xa7, 0x41, 0xb7, 0x25, 0xac, 0xcc, 0x32, 0x85, 0xa3, 0x09}, //|]*:.2O.A.%..2...|
	0x00000490: []byte{0x48, 0x95, 0xa3, 0x18, 0xf9, 0x63, 0x42, 0xd2, 0xa3, 0xcb, 0xb8, 0x90, 0xbc, 0x13, 0x7b, 0x0a}, //|H....cB.......{.|
	0x000004a0: []byte{0x00, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xcf, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.P..............|
	0x000004b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x01, 0x4b, 0x00, 0x45, 0x00, 0x52, 0x00, 0x4e, 0x00}, //|........K.E.R.N.|
	0x000004c0: []byte{0x2d, 0x00, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.A.............|
	0x000004d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000004e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000004f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000500: []byte{0x02, 0xe2, 0xb8, 0x3c, 0x7e, 0x3b, 0xdd, 0x47, 0x8a, 0x3c, 0x7f, 0xf2, 0xa1, 0x3c, 0xfc, 0xec}, //|...<~;.G.<...<..|
	0x00000510: []byte{0xbe, 0xc5, 0x8b, 0x8b, 0xa6, 0x6f, 0x41, 0x47, 0xb7, 0x30, 0x41, 0xa4, 0xc5, 0xc7, 0xf6, 0x0e}, //|.....oAG.0A.....|
	0x00000520: []byte{0x00, 0x60, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xdf, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.`........B.....|
	0x00000530: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x4f, 0x00, 0x4f, 0x00, 0x54, 0x00}, //|........R.O.O.T.|
	0x00000540: []byte{0x2d, 0x00, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.A.............|
	0x00000550: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000560: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000570: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000580: []byte{0x5d, 0x2a, 0x3a, 0xfe, 0x32, 0x4f, 0xa7, 0x41, 0xb7, 0x25, 0xac, 0xcc, 0x32, 0x85, 0xa3, 0x09}, //|]*:.2O.A.%..2...|
	0x00000590: []byte{0xb3, 0xc4, 0x3d, 0x6a, 0x77, 0x36, 0x4b, 0xde, 0x85, 0xd9, 0xd6, 0x74, 0x20, 0xca, 0x70, 0xc3}, //|..=jw6K....t .p.|
	0x000005a0: []byte{0x00, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.........O......|
	0x000005b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x45, 0x00, 0x52, 0x00, 0x4e, 0x00}, //|........K.E.R.N.|
	0x000005c0: []byte{0x2d, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.B.............|
	0x000005d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000005e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000005f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000600: []byte{0x02, 0xe2, 0xb8, 0x3c, 0x7e, 0x3b, 0xdd, 0x47, 0x8a, 0x3c, 0x7f, 0xf2, 0xa1, 0x3c, 0xfc, 0xec}, //|...<~;.G.<...<..|
	0x00000610: []byte{0x66, 0x27, 0xcf, 0x6e, 0xf8, 0xbc, 0x41, 0x47, 0x93, 0x6b, 0x1e, 0x24, 0xc6, 0xa1, 0x26, 0x97}, //|f'.n..AG.k.$..&.|
	0x00000620: []byte{0x00, 0x50, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x5f, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.P......._......|
	0x00000630: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x4f, 0x00, 0x4f, 0x00, 0x54, 0x00}, //|........R.O.O.T.|
	0x00000640: []byte{0x2d, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.B.............|
	0x00000650: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000660: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000670: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000680: []byte{0x5d, 0x2a, 0x3a, 0xfe, 0x32, 0x4f, 0xa7, 0x41, 0xb7, 0x25, 0xac, 0xcc, 0x32, 0x85, 0xa3, 0x09}, //|]*:.2O.A.%..2...|
	0x00000690: []byte{0xfc, 0x72, 0x5f, 0x08, 0xd1, 0xe9, 0x47, 0xd8, 0x8c, 0x21, 0x51, 0x18, 0xe4, 0xa5, 0x1c, 0x36}, //|.r_...G..!Q....6|
	0x000006a0: []byte{0x40, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|@@......@@......|
	0x000006b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x45, 0x00, 0x52, 0x00, 0x4e, 0x00}, //|........K.E.R.N.|
	0x000006c0: []byte{0x2d, 0x00, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.C.............|
	0x000006d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000006e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000006f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000700: []byte{0x02, 0xe2, 0xb8, 0x3c, 0x7e, 0x3b, 0xdd, 0x47, 0x8a, 0x3c, 0x7f, 0xf2, 0xa1, 0x3c, 0xfc, 0xec}, //|...<~;.G.<...<..|
	0x00000710: []byte{0x9c, 0xf5, 0xa0, 0xce, 0xea, 0x6a, 0x4f, 0x00, 0xbd, 0x17, 0xe3, 0x37, 0xf8, 0xdd, 0xac, 0xf0}, //|.....jO....7....|
	0x00000720: []byte{0x41, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|A@......A@......|
	0x00000730: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x4f, 0x00, 0x4f, 0x00, 0x54, 0x00}, //|........R.O.O.T.|
	0x00000740: []byte{0x2d, 0x00, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.C.............|
	0x00000750: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000760: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000770: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000780: []byte{0xa2, 0xa0, 0xd0, 0xeb, 0xe5, 0xb9, 0x33, 0x44, 0x87, 0xc0, 0x68, 0xb6, 0xb7, 0x26, 0x99, 0xc7}, //|......3D..h..&..|
	0x00000790: []byte{0xca, 0xd7, 0x7c, 0x50, 0x95, 0x32, 0x45, 0xd3, 0xbc, 0x96, 0x26, 0x1e, 0xef, 0x63, 0xef, 0x07}, //|..|P.2E...&..c..|
	0x000007a0: []byte{0x00, 0x50, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xcf, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.P..............|
	0x000007b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x00, 0x45, 0x00, 0x4d, 0x00, 0x00, 0x00}, //|........O.E.M...|
	0x000007c0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000007d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000007e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000007f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000800: []byte{0x3d, 0x75, 0x0a, 0x2e, 0x48, 0x9e, 0xb0, 0x43, 0x83, 0x37, 0xb1, 0x51, 0x92, 0xcb, 0x1b, 0x5e}, //|=u..H..C.7.Q...^|
	0x00000810: []byte{0xed, 0x98, 0x3a, 0x51, 0x3e, 0xc4, 0x4a, 0x14, 0x83, 0x98, 0xa4, 0x7a, 0x7e, 0x6a, 0xe4, 0x2c}, //|..:Q>.J....z~j.,|
	0x00000820: []byte{0x42, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|B@......B@......|
	0x00000830: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x72, 0x00, 0x65, 0x00, 0x73, 0x00, 0x65, 0x00}, //|........r.e.s.e.|
	0x00000840: []byte{0x72, 0x00, 0x76, 0x00, 0x65, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|r.v.e.d.........|
	0x00000850: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000860: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000870: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000880: []byte{0x3d, 0x75, 0x0a, 0x2e, 0x48, 0x9e, 0xb0, 0x43, 0x83, 0x37, 0xb1, 0x51, 0x92, 0xcb, 0x1b, 0x5e}, //|=u..H..C.7.Q...^|
	0x00000890: []byte{0x4e, 0x1b, 0x87, 0x7f, 0x7a, 0x17, 0x4a, 0x85, 0xac, 0xc4, 0x10, 0x4f, 0x45, 0xdf, 0x34, 0x5a}, //|N...z.J....OE.4Z|
	0x000008a0: []byte{0x43, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x43, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|C@......C@......|
	0x000008b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x72, 0x00, 0x65, 0x00, 0x73, 0x00, 0x65, 0x00}, //|........r.e.s.e.|
	0x000008c0: []byte{0x72, 0x00, 0x76, 0x00, 0x65, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|r.v.e.d.........|
	0x000008d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000008e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000008f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000900: []byte{0x8e, 0xe8, 0xb6, 0xca, 0xf3, 0xab, 0x02, 0x41, 0xa0, 0x7a, 0xd4, 0xbb, 0x9b, 0xe3, 0xc1, 0xd3}, //|.......A.z......|
	0x00000910: []byte{0x63, 0x89, 0xfd, 0x96, 0x80, 0x7c, 0x47, 0x9e, 0x8f, 0x85, 0x12, 0xdc, 0xcb, 0xf7, 0x75, 0xe0}, //|c....|G.......u.|
	0x00000920: []byte{0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|@.......?@......|
	0x00000930: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x57, 0x00, 0x46, 0x00, 0x57, 0x00}, //|........R.W.F.W.|
	0x00000940: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000950: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000960: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000970: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000980: []byte{0x28, 0x73, 0x2a, 0xc1, 0x1f, 0xf8, 0xd2, 0x11, 0xba, 0x4b, 0x00, 0xa0, 0xc9, 0x3e, 0xc9, 0x3b}, //|(s*......K...>.;|
	0x00000990: []byte{0x4b, 0xf3, 0x63, 0x86, 0x27, 0x68, 0x49, 0xee, 0xa8, 0xb0, 0xea, 0x11, 0x28, 0xf3, 0x17, 0x07}, //|K.c.'hI.....(...|
	0x000009a0: []byte{0x00, 0xd0, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x4f, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.........O......|
	0x000009b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x45, 0x00, 0x46, 0x00, 0x49, 0x00, 0x2d, 0x00}, //|........E.F.I.-.|
	0x000009c0: []byte{0x53, 0x00, 0x59, 0x00, 0x53, 0x00, 0x54, 0x00, 0x45, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x00, 0x00}, //|S.Y.S.T.E.M.....|
	0x000009d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000009e0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x000009f0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|
	0x00000a00: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	// Backup GPT is here.
	0x879f7e00: []byte{0x45, 0x46, 0x49, 0x20, 0x50, 0x41, 0x52, 0x54, 0x00, 0x00, 0x01, 0x00, 0x5c, 0x00, 0x00, 0x00}, //|EFI PART....\...|
	0x879f7e10: []byte{0xb7, 0x08, 0x9b, 0xe7, 0x00, 0x00, 0x00, 0x00, 0xbf, 0xcf, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00}, //|..........C.....|
	0x879f7e20: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|........".......|
	0x879f7e30: []byte{0x9e, 0xcf, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2d, 0x1e, 0xd4, 0xba, 0xef, 0x93, 0x4a, 0xb0}, //|..C.....-.....J.|
	0x879f7e40: []byte{0x84, 0x6e, 0x2e, 0x3a, 0x6a, 0x2d, 0x73, 0xbf, 0x9f, 0xcf, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.n.:j-s...C.....|
	0x879f7e50: []byte{0x80, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x57, 0x8e, 0x72, 0x8d, 0x00, 0x00, 0x00, 0x00}, //|........W.r.....|
	0x879f7e60: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	// Backup partitions are here. Note they are earlier in the image.
	0x879f3e00: []byte{0xa2, 0xa0, 0xd0, 0xeb, 0xe5, 0xb9, 0x33, 0x44, 0x87, 0xc0, 0x68, 0xb6, 0xb7, 0x26, 0x99, 0xc7}, //|......3D..h..&..|
	0x879f3e10: []byte{0x93, 0x6b, 0x59, 0xaa, 0xa8, 0xeb, 0x42, 0xe0, 0xbb, 0x1c, 0x50, 0x0b, 0x28, 0x46, 0xdb, 0x57}, //|.kY...B...P.(F.W|
	0x879f3e20: []byte{0x00, 0xe0, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x4f, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00}, //|..B......OC.....|
	0x879f3e30: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x53, 0x00, 0x54, 0x00, 0x41, 0x00, 0x54, 0x00}, //|........S.T.A.T.|
	0x879f3e40: []byte{0x45, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|E...............|
	0x879f3e50: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f3e80: []byte{0x5d, 0x2a, 0x3a, 0xfe, 0x32, 0x4f, 0xa7, 0x41, 0xb7, 0x25, 0xac, 0xcc, 0x32, 0x85, 0xa3, 0x09}, //|]*:.2O.A.%..2...|
	0x879f3e90: []byte{0x48, 0x95, 0xa3, 0x18, 0xf9, 0x63, 0x42, 0xd2, 0xa3, 0xcb, 0xb8, 0x90, 0xbc, 0x13, 0x7b, 0x0a}, //|H....cB.......{.|
	0x879f3ea0: []byte{0x00, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xcf, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.P..............|
	0x879f3eb0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x01, 0x4b, 0x00, 0x45, 0x00, 0x52, 0x00, 0x4e, 0x00}, //|........K.E.R.N.|
	0x879f3ec0: []byte{0x2d, 0x00, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.A.............|
	0x879f3ed0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f3f00: []byte{0x02, 0xe2, 0xb8, 0x3c, 0x7e, 0x3b, 0xdd, 0x47, 0x8a, 0x3c, 0x7f, 0xf2, 0xa1, 0x3c, 0xfc, 0xec}, //|...<~;.G.<...<..|
	0x879f3f10: []byte{0xbe, 0xc5, 0x8b, 0x8b, 0xa6, 0x6f, 0x41, 0x47, 0xb7, 0x30, 0x41, 0xa4, 0xc5, 0xc7, 0xf6, 0x0e}, //|.....oAG.0A.....|
	0x879f3f20: []byte{0x00, 0x60, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xdf, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.`........B.....|
	0x879f3f30: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x4f, 0x00, 0x4f, 0x00, 0x54, 0x00}, //|........R.O.O.T.|
	0x879f3f40: []byte{0x2d, 0x00, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.A.............|
	0x879f3f50: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f3f80: []byte{0x5d, 0x2a, 0x3a, 0xfe, 0x32, 0x4f, 0xa7, 0x41, 0xb7, 0x25, 0xac, 0xcc, 0x32, 0x85, 0xa3, 0x09}, //|]*:.2O.A.%..2...|
	0x879f3f90: []byte{0xb3, 0xc4, 0x3d, 0x6a, 0x77, 0x36, 0x4b, 0xde, 0x85, 0xd9, 0xd6, 0x74, 0x20, 0xca, 0x70, 0xc3}, //|..=jw6K....t,0x.p.|
	0x879f3fa0: []byte{0x00, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x4f, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.........O......|
	0x879f3fb0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x45, 0x00, 0x52, 0x00, 0x4e, 0x00}, //|........K.E.R.N.|
	0x879f3fc0: []byte{0x2d, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.B.............|
	0x879f3fd0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4000: []byte{0x02, 0xe2, 0xb8, 0x3c, 0x7e, 0x3b, 0xdd, 0x47, 0x8a, 0x3c, 0x7f, 0xf2, 0xa1, 0x3c, 0xfc, 0xec}, //|...<~;.G.<...<..|
	0x879f4010: []byte{0x66, 0x27, 0xcf, 0x6e, 0xf8, 0xbc, 0x41, 0x47, 0x93, 0x6b, 0x1e, 0x24, 0xc6, 0xa1, 0x26, 0x97}, //|f'.n..AG.k.$..&.|
	0x879f4020: []byte{0x00, 0x50, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x5f, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.P......._......|
	0x879f4030: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x4f, 0x00, 0x4f, 0x00, 0x54, 0x00}, //|........R.O.O.T.|
	0x879f4040: []byte{0x2d, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.B.............|
	0x879f4050: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4080: []byte{0x5d, 0x2a, 0x3a, 0xfe, 0x32, 0x4f, 0xa7, 0x41, 0xb7, 0x25, 0xac, 0xcc, 0x32, 0x85, 0xa3, 0x09}, //|]*:.2O.A.%..2...|
	0x879f4090: []byte{0xfc, 0x72, 0x5f, 0x08, 0xd1, 0xe9, 0x47, 0xd8, 0x8c, 0x21, 0x51, 0x18, 0xe4, 0xa5, 0x1c, 0x36}, //|.r_...G..!Q....6|
	0x879f40a0: []byte{0x40, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|@@......@@......|
	0x879f40b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x45, 0x00, 0x52, 0x00, 0x4e, 0x00}, //|........K.E.R.N.|
	0x879f40c0: []byte{0x2d, 0x00, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.C.............|
	0x879f40d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4100: []byte{0x02, 0xe2, 0xb8, 0x3c, 0x7e, 0x3b, 0xdd, 0x47, 0x8a, 0x3c, 0x7f, 0xf2, 0xa1, 0x3c, 0xfc, 0xec}, //|...<~;.G.<...<..|
	0x879f4110: []byte{0x9c, 0xf5, 0xa0, 0xce, 0xea, 0x6a, 0x4f, 0x00, 0xbd, 0x17, 0xe3, 0x37, 0xf8, 0xdd, 0xac, 0xf0}, //|.....jO....7....|
	0x879f4120: []byte{0x41, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|A@......A@......|
	0x879f4130: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x4f, 0x00, 0x4f, 0x00, 0x54, 0x00}, //|........R.O.O.T.|
	0x879f4140: []byte{0x2d, 0x00, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|-.C.............|
	0x879f4150: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4180: []byte{0xa2, 0xa0, 0xd0, 0xeb, 0xe5, 0xb9, 0x33, 0x44, 0x87, 0xc0, 0x68, 0xb6, 0xb7, 0x26, 0x99, 0xc7}, //|......3D..h..&..|
	0x879f4190: []byte{0xca, 0xd7, 0x7c, 0x50, 0x95, 0x32, 0x45, 0xd3, 0xbc, 0x96, 0x26, 0x1e, 0xef, 0x63, 0xef, 0x07}, //|..|P.2E...&..c..|
	0x879f41a0: []byte{0x00, 0x50, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xcf, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.P..............|
	0x879f41b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x00, 0x45, 0x00, 0x4d, 0x00, 0x00, 0x00}, //|........O.E.M...|
	0x879f41c0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4200: []byte{0x3d, 0x75, 0x0a, 0x2e, 0x48, 0x9e, 0xb0, 0x43, 0x83, 0x37, 0xb1, 0x51, 0x92, 0xcb, 0x1b, 0x5e}, //|=u..H..C.7.Q...^|
	0x879f4210: []byte{0xed, 0x98, 0x3a, 0x51, 0x3e, 0xc4, 0x4a, 0x14, 0x83, 0x98, 0xa4, 0x7a, 0x7e, 0x6a, 0xe4, 0x2c}, //|..:Q>.J....z~j.,|
	0x879f4220: []byte{0x42, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|B@......B@......|
	0x879f4230: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x72, 0x00, 0x65, 0x00, 0x73, 0x00, 0x65, 0x00}, //|........r.e.s.e.|
	0x879f4240: []byte{0x72, 0x00, 0x76, 0x00, 0x65, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|r.v.e.d.........|
	0x879f4250: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4280: []byte{0x3d, 0x75, 0x0a, 0x2e, 0x48, 0x9e, 0xb0, 0x43, 0x83, 0x37, 0xb1, 0x51, 0x92, 0xcb, 0x1b, 0x5e}, //|=u..H..C.7.Q...^|
	0x879f4290: []byte{0x4e, 0x1b, 0x87, 0x7f, 0x7a, 0x17, 0x4a, 0x85, 0xac, 0xc4, 0x10, 0x4f, 0x45, 0xdf, 0x34, 0x5a}, //|N...z.J....OE.4Z|
	0x879f42a0: []byte{0x43, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x43, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|C@......C@......|
	0x879f42b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x72, 0x00, 0x65, 0x00, 0x73, 0x00, 0x65, 0x00}, //|........r.e.s.e.|
	0x879f42c0: []byte{0x72, 0x00, 0x76, 0x00, 0x65, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|r.v.e.d.........|
	0x879f42d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4300: []byte{0x8e, 0xe8, 0xb6, 0xca, 0xf3, 0xab, 0x02, 0x41, 0xa0, 0x7a, 0xd4, 0xbb, 0x9b, 0xe3, 0xc1, 0xd3}, //|.......A.z......|
	0x879f4310: []byte{0x63, 0x89, 0xfd, 0x96, 0x80, 0x7c, 0x47, 0x9e, 0x8f, 0x85, 0x12, 0xdc, 0xcb, 0xf7, 0x75, 0xe0}, //|c....|G.......u.|
	0x879f4320: []byte{0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|@.......?@......|
	0x879f4330: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x00, 0x57, 0x00, 0x46, 0x00, 0x57, 0x00}, //|........R.W.F.W.|
	0x879f4340: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

	0x879f4380: []byte{0x28, 0x73, 0x2a, 0xc1, 0x1f, 0xf8, 0xd2, 0x11, 0xba, 0x4b, 0x00, 0xa0, 0xc9, 0x3e, 0xc9, 0x3b}, //|(s*......K...>.;|
	0x879f4390: []byte{0x4b, 0xf3, 0x63, 0x86, 0x27, 0x68, 0x49, 0xee, 0xa8, 0xb0, 0xea, 0x11, 0x28, 0xf3, 0x17, 0x07}, //|K.c.'hI.....(...|
	0x879f43a0: []byte{0x00, 0xd0, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x4f, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00}, //|.........O......|
	0x879f43b0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x45, 0x00, 0x46, 0x00, 0x49, 0x00, 0x2d, 0x00}, //|........E.F.I.-.|
	0x879f43c0: []byte{0x53, 0x00, 0x59, 0x00, 0x53, 0x00, 0x54, 0x00, 0x45, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x00, 0x00}, //|S.Y.S.T.E.M.....|
	0x879f43d0: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, //|................|

}
