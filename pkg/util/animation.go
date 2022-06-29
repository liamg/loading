package util

/*
20 	(space) 	000000 	⠀ (braille pattern blank) 	⠀ 	(space)
21 	! 	011101 	⠮ (braille pattern dots-2346) 	⠮ 	the
22 	" 	000010 	⠐ (braille pattern dots-5) 	⠐ 	(contraction)
23 	# 	001111 	⠼ (braille pattern dots-3456) 	⠼ 	(number prefix)
24 	$ 	110101 	⠫ (braille pattern dots-1246) 	⠫ 	ed
25 	% 	100101 	⠩ (braille pattern dots-146) 	⠩ 	sh
26 	& 	111101 	⠯ (braille pattern dots-12346) 	⠯ 	and
27 	' 	001000 	⠄ (braille pattern dots-3) 	⠄ 	'
28 	( 	111011 	⠷ (braille pattern dots-12356) 	⠷ 	of
29 	) 	011111 	⠾ (braille pattern dots-23456) 	⠾ 	with
2A 	* 	100001 	⠡ (braille pattern dots-16) 	⠡ 	ch
2B 	+ 	001101 	⠬ (braille pattern dots-346) 	⠬ 	ing
2C 	, 	000001 	⠠ (braille pattern dots-6) 	⠠ 	(uppercase prefix)
2D 	- 	001001 	⠤ (braille pattern dots-36) 	⠤ 	-
2E 	. 	000101 	⠨ (braille pattern dots-46) 	⠨ 	(italic prefix)
2F 	/ 	001100 	⠌ (braille pattern dots-34) 	⠌ 	st or /
30 	0 	001011 	⠴ (braille pattern dots-356) 	⠴ 	”
31 	1 	010000 	⠂ (braille pattern dots-2) 	⠂ 	,
32 	2 	011000 	⠆ (braille pattern dots-23) 	⠆ 	;
33 	3 	010010 	⠒ (braille pattern dots-25) 	⠒ 	:
34 	4 	010011 	⠲ (braille pattern dots-256) 	⠲ 	.
35 	5 	010001 	⠢ (braille pattern dots-26) 	⠢ 	en
36 	6 	011010 	⠖ (braille pattern dots-235) 	⠖ 	!
37 	7 	011011 	⠶ (braille pattern dots-2356) 	⠶ 	( or )
38 	8 	011001 	⠦ (braille pattern dots-236) 	⠦ 	“ or ?
39 	9 	001010 	⠔ (braille pattern dots-35) 	⠔ 	in
3A 	: 	100011 	⠱ (braille pattern dots-156) 	⠱ 	wh
3B 	; 	000011 	⠰ (braille pattern dots-56) 	⠰ 	(letter prefix)
3C 	< 	110001 	⠣ (braille pattern dots-126) 	⠣ 	gh
3D 	= 	111111 	⠿ (braille pattern dots-123456) 	⠿ 	for
3E 	> 	001110 	⠜ (braille pattern dots-345) 	⠜ 	ar
3F 	? 	100111 	⠹ (braille pattern dots-1456) 	⠹ 	th

ASCII hex 	ASCII glyph 	Braille dots 	Braille glyph 	Unicode Braille glyph 	Braille meaning
40 	@ 	000100 	⠈ (braille pattern dots-4) 	⠈ 	(accent prefix)
41 	A 	100000 	⠁ (braille pattern dots-1) 	⠁ 	a
42 	B 	110000 	⠃ (braille pattern dots-12) 	⠃ 	b
43 	C 	100100 	⠉ (braille pattern dots-14) 	⠉ 	c
44 	D 	100110 	⠙ (braille pattern dots-145) 	⠙ 	d
45 	E 	100010 	⠑ (braille pattern dots-15) 	⠑ 	e
46 	F 	110100 	⠋ (braille pattern dots-124) 	⠋ 	f
47 	G 	110110 	⠛ (braille pattern dots-1245) 	⠛ 	g
48 	H 	110010 	⠓ (braille pattern dots-125) 	⠓ 	h
49 	I 	010100 	⠊ (braille pattern dots-24) 	⠊ 	i
4A 	J 	010110 	⠚ (braille pattern dots-245) 	⠚ 	j
4B 	K 	101000 	⠅ (braille pattern dots-13) 	⠅ 	k
4C 	L 	111000 	⠇ (braille pattern dots-123) 	⠇ 	l
4D 	M 	101100 	⠍ (braille pattern dots-134) 	⠍ 	m
4E 	N 	101110 	⠝ (braille pattern dots-1345) 	⠝ 	n
4F 	O 	101010 	⠕ (braille pattern dots-135) 	⠕ 	o
50 	P 	111100 	⠏ (braille pattern dots-1234) 	⠏ 	p
51 	Q 	111110 	⠟ (braille pattern dots-12345) 	⠟ 	q
52 	R 	111010 	⠗ (braille pattern dots-1235) 	⠗ 	r
53 	S 	011100 	⠎ (braille pattern dots-234) 	⠎ 	s
54 	T 	011110 	⠞ (braille pattern dots-2345) 	⠞ 	t
55 	U 	101001 	⠥ (braille pattern dots-136) 	⠥ 	u
56 	V 	111001 	⠧ (braille pattern dots-1236) 	⠧ 	v
57 	W 	010111 	⠺ (braille pattern dots-2456) 	⠺ 	w
58 	X 	101101 	⠭ (braille pattern dots-1346) 	⠭ 	x
59 	Y 	101111 	⠽ (braille pattern dots-13456) 	⠽ 	y
5A 	Z 	101011 	⠵ (braille pattern dots-1356) 	⠵ 	z
5B 	[ 	010101 	⠪ (braille pattern dots-246) 	⠪ 	ow
5C 	\ 	110011 	⠳ (braille pattern dots-1256) 	⠳ 	ou
5D 	] 	110111 	⠻ (braille pattern dots-12456) 	⠻ 	er
5E 	^ 	000110 	⠘ (braille pattern dots-45) 	⠘ 	(currency prefix)
5F 	_ 	000111 	⠸ (braille pattern dots-456) 	⠸ 	(contraction)
*/

var loadingAnimationFrames = []rune{
    '⠀',
    '⠈',
    '⠘',
    '⠸',
    '⠼',
    '⠾',
    '⠿',
    '⠷',
    '⠧',
    '⠇',
    '⠃',
    '⠁',
}

func LoadingBlockRune(offset int) rune {
    return loadingAnimationFrames[offset%len(loadingAnimationFrames)]
}
