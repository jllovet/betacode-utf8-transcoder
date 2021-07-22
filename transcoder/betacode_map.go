package transcoder

var BetacodeMap = map[string]string{
	// No marks

	`a`: "\u03b1", // GREEK SMALL LETTER ALPHA
	`A`: "\u03b1",

	`b`: "\u03b2", // GREEK SMALL LETTER BETA
	`B`: "\u03b2",

	`g`: "\u03b3", // GREEK SMALL LETTER GAMMA
	`G`: "\u03b3",

	`d`: "\u03b4", // GREEK SMALL LETTER DELTA
	`D`: "\u03b4",

	`e`: "\u03b5", // GREEK SMALL LETTER EPSILON
	`E`: "\u03b5",

	`z`: "\u03b6", // GREEK SMALL LETTER ZETA
	`Z`: "\u03b6",

	`h`: "\u03b7", // GREEK SMALL LETTER ETA
	`H`: "\u03b7",

	`q`: "\u03b8", // GREEK SMALL LETTER THETA
	`Q`: "\u03b8",

	`i`: "\u03b9", // GREEK SMALL LETTER IOTA
	`I`: "\u03b9",

	`k`: "\u03ba", // GREEK SMALL LETTER KAPPA
	`K`: "\u03ba",

	`l`: "\u03bb", // GREEK SMALL LETTER LAMDA
	`L`: "\u03bb",

	`m`: "\u03bc", // GREEK SMALL LETTER MU
	`M`: "\u03bc",

	`n`: "\u03bd", // GREEK SMALL LETTER NU
	`N`: "\u03bd",

	`c`: "\u03be", // GREEK SMALL LETTER XI
	`C`: "\u03be",

	`o`: "\u03bf", // GREEK SMALL LETTER OMICRON
	`O`: "\u03bf",

	`p`: "\u03c0", // GREEK SMALL LETTER PI
	`P`: "\u03c0",

	`r`: "\u03c1", // GREEK SMALL LETTER RHO
	`R`: "\u03c1",

	`s`:  "\u03c3", // GREEK SMALL LETTER SIGMA
	`s1`: "\u03c3",
	`S`:  "\u03c3",
	`S1`: "\u03c3",

	`s2`: "\u03c2", // GREEK SMALL LETTER FINAL SIGMA
	`S2`: "\u03c2",

	`s3`: "\u03f2", // GREEK LUNATE SIGMA SYMBOL
	`S3`: "\u03f2",

	`t`: "\u03c4", // GREEK SMALL LETTER TAU
	`T`: "\u03c4",

	`u`: "\u03c5", // GREEK SMALL LETTER UPSILON
	`U`: "\u03c5",

	`f`: "\u03c6", // GREEK SMALL LETTER PHI
	`F`: "\u03c6",

	`x`: "\u03c7", // GREEK SMALL LETTER CHI
	`X`: "\u03c7",

	`y`: "\u03c8", // GREEK SMALL LETTER PSI
	`Y`: "\u03c8",

	`w`: "\u03c9", // GREEK SMALL LETTER OMEGA
	`W`: "\u03c9",

	`*a`: "\u0391", // GREEK CAPITAL LETTER ALPHA
	`*A`: "\u0391",

	`*b`: "\u0392", // GREEK CAPITAL LETTER BETA
	`*B`: "\u0392",

	`*g`: "\u0393", // GREEK CAPITAL LETTER GAMMA
	`*G`: "\u0393",

	`*d`: "\u0394", // GREEK CAPITAL LETTER DELTA
	`*D`: "\u0394",

	`*e`: "\u0395", // GREEK CAPITAL LETTER EPSILON
	`*E`: "\u0395",

	`*z`: "\u0396", // GREEK CAPITAL LETTER ZETA
	`*Z`: "\u0396",

	`*h`: "\u0397", // GREEK CAPITAL LETTER ETA
	`*H`: "\u0397",

	`*q`: "\u0398", // GREEK CAPITAL LETTER THETA
	`*Q`: "\u0398",

	`*i`: "\u0399", // GREEK CAPITAL LETTER IOTA
	`*I`: "\u0399",

	`*k`: "\u039a", // GREEK CAPITAL LETTER KAPPA
	`*K`: "\u039a",

	`*l`: "\u039b", // GREEK CAPITAL LETTER LAMDA
	`*L`: "\u039b",

	`*m`: "\u039c", // GREEK CAPITAL LETTER MU
	`*M`: "\u039c",

	`*n`: "\u039d", // GREEK CAPITAL LETTER NU
	`*N`: "\u039d",

	`*c`: "\u039e", // GREEK CAPITAL LETTER XI
	`*C`: "\u039e",

	`*o`: "\u039f", // GREEK CAPITAL LETTER OMICRON
	`*O`: "\u039f",

	`*p`: "\u03a0", // GREEK CAPITAL LETTER PI
	`*P`: "\u03a0",

	`*r`: "\u03a1", // GREEK CAPITAL LETTER RHO
	`*R`: "\u03a1",

	`*s`: "\u03a3", // GREEK CAPITAL LETTER SIGMA
	`*S`: "\u03a3",

	`*s3`: "\u03f9", // GREEK CAPITAL LUNATE SIGMA SYMBOL
	`*S3`: "\u03f9",
	`*3s`: "\u03f9",
	`*3S`: "\u03f9",

	`*t`: "\u03a4", // GREEK CAPITAL LETTER TAU
	`*T`: "\u03a4",

	`*u`: "\u03a5", // GREEK CAPITAL LETTER UPSILON
	`*U`: "\u03a5",

	`*f`: "\u03a6", // GREEK CAPITAL LETTER PHI
	`*F`: "\u03a6",

	`*x`: "\u03a7", // GREEK CAPITAL LETTER CHI
	`*X`: "\u03a7",

	`*y`: "\u03a8", // GREEK CAPITAL LETTER PSI
	`*Y`: "\u03a8",

	`*w`: "\u03a9", // GREEK CAPITAL LETTER OMEGA
	`*W`: "\u03a9",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing

	`a)`: "\u1f00", // GREEK SMALL LETTER ALPHA WITH PSILI
	`A)`: "\u1f00",

	`e)`: "\u1f10", // GREEK SMALL LETTER EPSILON WITH PSILI
	`E)`: "\u1f10",

	`h)`: "\u1f20", // GREEK SMALL LETTER ETA WITH PSILI
	`H)`: "\u1f20",

	`i)`: "\u1f30", // GREEK SMALL LETTER IOTA WITH PSILI
	`I)`: "\u1f30",

	`o)`: "\u1f40", // GREEK SMALL LETTER OMICRON WITH PSILI
	`O)`: "\u1f40",

	`u)`: "\u1f50", // GREEK SMALL LETTER UPSILON WITH PSILI
	`U)`: "\u1f50",

	`w)`: "\u1f60", // GREEK SMALL LETTER OMEGA WITH PSILI
	`W)`: "\u1f60",

	`r)`: "\u1fe4", // GREEK SMALL LETTER RHO WITH PSILI
	`R)`: "\u1fe4",

	`*)a`: "\u1f08", // GREEK CAPITAL LETTER ALPHA WITH PSILI
	`*a)`: "\u1f08",
	`*A)`: "\u1f08",
	`*)A`: "\u1f08",

	`*)e`: "\u1f18", // GREEK CAPITAL LETTER EPSILON WITH PSILI
	`*e)`: "\u1f18",
	`*E)`: "\u1f18",
	`*)E`: "\u1f18",

	`*)h`: "\u1f28", // GREEK CAPITAL LETTER ETA WITH PSILI
	`*h)`: "\u1f28",
	`*H)`: "\u1f28",
	`*)H`: "\u1f28",

	`*)i`: "\u1f38", // GREEK CAPITAL LETTER IOTA WITH PSILI
	`*i)`: "\u1f38",
	`*I)`: "\u1f38",
	`*)I`: "\u1f38",

	`*)o`: "\u1f48", // GREEK CAPITAL LETTER OMICRON WITH PSILI
	`*o)`: "\u1f48",
	`*O)`: "\u1f48",
	`*)O`: "\u1f48",

	`*)w`: "\u1f68", // GREEK CAPITAL LETTER OMEGA WITH PSILI
	`*w)`: "\u1f68",
	`*W)`: "\u1f68",
	`*)W`: "\u1f68",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breathing

	`a(`: "\u1f01", // GREEK SMALL LETTER ALPHA WITH DASIA
	`A(`: "\u1f01",

	`e(`: "\u1f11", // GREEK SMALL LETTER EPSILON WITH DASIA
	`E(`: "\u1f11",

	`h(`: "\u1f21", // GREEK SMALL LETTER ETA WITH DASIA
	`H(`: "\u1f21",

	`i(`: "\u1f31", // GREEK SMALL LETTER IOTA WITH DASIA
	`I(`: "\u1f31",

	`o(`: "\u1f41", // GREEK SMALL LETTER OMICRON WITH DASIA
	`O(`: "\u1f41",

	`u(`: "\u1f51", // GREEK SMALL LETTER UPSILON WITH DASIA
	`U(`: "\u1f51",

	`w(`: "\u1f61", // GREEK SMALL LETTER OMEGA WITH DASIA
	`W(`: "\u1f61",

	`r(`: "\u1fe5", // GREEK SMALL LETTER RHO WITH DASIA
	`R(`: "\u1fe5",

	`*(a`: "\u1f09", // GREEK CAPITAL LETTER ALPHA WITH DASIA
	`*A(`: "\u1f09",
	`*(A`: "\u1f09",

	`*(e`: "\u1f19", // GREEK CAPITAL LETTER EPSILON WITH DASIA
	`*e(`: "\u1f19",
	`*E(`: "\u1f19",
	`*(E`: "\u1f19",

	`*(h`: "\u1f29", // GREEK CAPITAL LETTER ETA WITH DASIA
	`*h(`: "\u1f29",
	`*H(`: "\u1f29",
	`*(H`: "\u1f29",

	`*(i`: "\u1f39", // GREEK CAPITAL LETTER IOTA WITH DASIA
	`*i(`: "\u1f39",
	`*I(`: "\u1f39",
	`*(I`: "\u1f39",

	`*(o`: "\u1f49", // GREEK CAPITAL LETTER OMICRON WITH DASIA
	`*o(`: "\u1f49",
	`*O(`: "\u1f49",
	`*(O`: "\u1f49",

	`*(u`: "\u1f59", // GREEK CAPITAL LETTER UPSILON WITH DASIA
	`*u(`: "\u1f59",
	`*U(`: "\u1f59",
	`*(U`: "\u1f59",

	`*(w`: "\u1f69", // GREEK CAPITAL LETTER OMEGA WITH DASIA
	`*w(`: "\u1f69",
	`*W(`: "\u1f69",
	`*(W`: "\u1f69",

	`*(r`: "\u1fec", // GREEK CAPITAL LETTER RHO WITH DASIA
	`*r(`: "\u1fec",
	`*R(`: "\u1fec",
	`*(R`: "\u1fec",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Acute accent and grave accent

	`a\`: "\u1f70", // GREEK SMALL LETTER ALPHA WITH VARIA
	`A\`: "\u1f70",

	`a/`: "\u1f71", // GREEK SMALL LETTER ALPHA WITH OXIA
	`A/`: "\u1f71",

	`e\`: "\u1f72", // GREEK SMALL LETTER EPSILON WITH VARIA
	`E\`: "\u1f72",

	`e/`: "\u1f73", // GREEK SMALL LETTER EPSILON WITH OXIA
	`E/`: "\u1f73",

	`h\`: "\u1f74", // GREEK SMALL LETTER ETA WITH VARIA
	`H\`: "\u1f74",

	`h/`: "\u1f75", // GREEK SMALL LETTER ETA WITH OXIA
	`H/`: "\u1f75",

	`i\`: "\u1f76", // GREEK SMALL LETTER IOTA WITH VARIA
	`I\`: "\u1f76",

	`i/`: "\u1f77", // GREEK SMALL LETTER IOTA WITH OXIA
	`I/`: "\u1f77",

	`o\`: "\u1f78", // GREEK SMALL LETTER OMICRON WITH VARIA
	`O\`: "\u1f78",

	`o/`: "\u1f79", // GREEK SMALL LETTER OMICRON WITH OXIA
	`O/`: "\u1f79",

	`u\`: "\u1f7a", // GREEK SMALL LETTER UPSILON WITH VARIA
	`U\`: "\u1f7a",

	`u/`: "\u1f7b", // GREEK SMALL LETTER UPSILON WITH OXIA
	`U/`: "\u1f7b",

	`w\`: "\u1f7c", // GREEK SMALL LETTER OMEGA WITH VARIA
	`W\`: "\u1f7c",

	`w/`: "\u1f7d", // GREEK SMALL LETTER OMEGA WITH OXIA
	`W/`: "\u1f7d",

	`*\a`: "\u1fba", // GREEK CAPITAL LETTER ALPHA WITH VARIA
	`*a\`: "\u1fba",
	`*A\`: "\u1fba",
	`*\A`: "\u1fba",

	`*/a`: "\u1fbb", // GREEK CAPITAL LETTER ALPHA WITH OXIA
	`*a/`: "\u1fbb",
	`*A/`: "\u1fbb",
	`*/A`: "\u1fbb",

	`*\e`: "\u1fce", // GREEK PSILI AND OXIA
	`*e\`: "\u1fce",
	`*E\`: "\u1fce",
	`*\E`: "\u1fce",

	`*/e`: "\u1fc9", // GREEK CAPITAL LETTER EPSILON WITH OXIA
	`*e/`: "\u1fc9",
	`*E/`: "\u1fc9",
	`*/E`: "\u1fc9",

	`*\h`: "\u1fca", // GREEK CAPITAL LETTER ETA WITH VARIA
	`*h\`: "\u1fca",
	`*H\`: "\u1fca",
	`*\H`: "\u1fca",

	`*/h`: "\u1fcb", // GREEK CAPITAL LETTER ETA WITH OXIA
	`*h/`: "\u1fcb",
	`*H/`: "\u1fcb",
	`*/H`: "\u1fcb",

	`*\i`: "\u1fda", // GREEK CAPITAL LETTER IOTA WITH VARIA
	`*i\`: "\u1fda",
	`*I\`: "\u1fda",
	`*\I`: "\u1fda",

	`*/i`: "\u1fdb", // GREEK CAPITAL LETTER IOTA WITH OXIA
	`*i/`: "\u1fdb",
	`*I/`: "\u1fdb",
	`*/I`: "\u1fdb",

	`*\o`: "\u1ff8", // GREEK CAPITAL LETTER OMICRON WITH VARIA
	`*o\`: "\u1ff8",
	`*O\`: "\u1ff8",
	`*\O`: "\u1ff8",

	`*/o`: "\u1ff9", // GREEK CAPITAL LETTER OMICRON WITH OXIA
	`*o/`: "\u1ff9",
	`*O/`: "\u1ff9",
	`*/O`: "\u1ff9",

	`*\u`: "\u1fea", // GREEK CAPITAL LETTER UPSILON WITH VARIA
	`*u\`: "\u1fea",
	`*U\`: "\u1fea",
	`*\U`: "\u1fea",

	`*/u`: "\u1feb", // GREEK CAPITAL LETTER UPSILON WITH OXIA
	`*u/`: "\u1feb",
	`*U/`: "\u1feb",
	`*/U`: "\u1feb",

	`*\w`: "\u1ffa", // GREEK CAPITAL LETTER OMEGA WITH VARIA
	`*w\`: "\u1ffa",
	`*W\`: "\u1ffa",
	`*\W`: "\u1ffa",

	`*/w`: "\u1ffb", // GREEK CAPITAL LETTER OMEGA WITH OXIA
	`*w/`: "\u1ffb",
	`*W/`: "\u1ffb",
	`*/W`: "\u1ffb",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing and acute accent

	`a)/`: "\u1f04", // GREEK SMALL LETTER ALPHA WITH PSILI AND OXIA
	`a/)`: "\u1f04",
	`A)/`: "\u1f04",
	`A/)`: "\u1f04",

	`e)/`: "\u1f14", // GREEK SMALL LETTER EPSILON WITH PSILI AND OXIA
	`e/)`: "\u1f14",
	`E)/`: "\u1f14",
	`E/)`: "\u1f14",

	`h)/`: "\u1f24", // GREEK SMALL LETTER ETA WITH PSILI AND OXIA
	`h/)`: "\u1f24",
	`H)/`: "\u1f24",
	`H/)`: "\u1f24",

	`i)/`: "\u1f34", // GREEK SMALL LETTER IOTA WITH PSILI AND OXIA
	`i/)`: "\u1f34",
	`I)/`: "\u1f34",
	`I/)`: "\u1f34",

	`o)/`: "\u1f44", // GREEK SMALL LETTER OMICRON WITH PSILI AND OXIA
	`o/)`: "\u1f44",
	`O)/`: "\u1f44",
	`O/)`: "\u1f44",

	`u)/`: "\u1f54", // GREEK SMALL LETTER UPSILON WITH PSILI AND OXIA
	`u/)`: "\u1f54",
	`U)/`: "\u1f54",
	`U/)`: "\u1f54",

	`w)/`: "\u1f64", // GREEK SMALL LETTER OMEGA WITH PSILI AND OXIA
	`w/)`: "\u1f64",
	`W)/`: "\u1f64",
	`W/)`: "\u1f64",

	`*)/a`: "\u1f0c", // GREEK CAPITAL LETTER ALPHA WITH PSILI AND OXIA
	`*a)/`: "\u1f0c",
	`*a/)`: "\u1f0c",
	`*A)/`: "\u1f0c",
	`*A/)`: "\u1f0c",
	`*)a/`: "\u1f0c",
	`*)A/`: "\u1f0c",
	`*)/A`: "\u1f0c",
	`*/a)`: "\u1f0c",
	`*/A)`: "\u1f0c",
	`*/)a`: "\u1f0c",
	`*/)A`: "\u1f0c",

	`*)/e`: "\u1f1c", // GREEK CAPITAL LETTER EPSILON WITH PSILI AND OXIA
	`*e)/`: "\u1f1c",
	`*e/)`: "\u1f1c",
	`*E)/`: "\u1f1c",
	`*E/)`: "\u1f1c",
	`*)e/`: "\u1f1c",
	`*)E/`: "\u1f1c",
	`*)/E`: "\u1f1c",
	`*/e)`: "\u1f1c",
	`*/E)`: "\u1f1c",
	`*/)e`: "\u1f1c",
	`*/)E`: "\u1f1c",

	`*)/h`: "\u1f2c", // GREEK CAPITAL LETTER ETA WITH PSILI AND OXIA
	`*h)/`: "\u1f2c",
	`*h/)`: "\u1f2c",
	`*H)/`: "\u1f2c",
	`*H/)`: "\u1f2c",
	`*)h/`: "\u1f2c",
	`*)H/`: "\u1f2c",
	`*)/H`: "\u1f2c",
	`*/h)`: "\u1f2c",
	`*/H)`: "\u1f2c",
	`*/)h`: "\u1f2c",
	`*/)H`: "\u1f2c",

	`*)/i`: "\u1f3c", // GREEK CAPITAL LETTER IOTA WITH PSILI AND OXIA
	`*i)/`: "\u1f3c",
	`*i/)`: "\u1f3c",
	`*I)/`: "\u1f3c",
	`*I/)`: "\u1f3c",
	`*)i/`: "\u1f3c",
	`*)I/`: "\u1f3c",
	`*)/I`: "\u1f3c",
	`*/i)`: "\u1f3c",
	`*/I)`: "\u1f3c",
	`*/)i`: "\u1f3c",
	`*/)I`: "\u1f3c",

	`*)/o`: "\u1f4c", // GREEK CAPITAL LETTER OMICRON WITH PSILI AND OXIA
	`*o)/`: "\u1f4c",
	`*o/)`: "\u1f4c",
	`*O)/`: "\u1f4c",
	`*O/)`: "\u1f4c",
	`*)o/`: "\u1f4c",
	`*)O/`: "\u1f4c",
	`*)/O`: "\u1f4c",
	`*/o)`: "\u1f4c",
	`*/O)`: "\u1f4c",
	`*/)o`: "\u1f4c",
	`*/)O`: "\u1f4c",

	`*)/u`: "\u1f5c", // Invalid
	`*u/)`: "\u1f5c",
	`*u)/`: "\u1f5c",
	`*U/)`: "\u1f5c",
	`*U)/`: "\u1f5c",
	`*)/U`: "\u1f5c",
	`*)u/`: "\u1f5c",
	`*)U/`: "\u1f5c",
	`*/u)`: "\u1f5c",
	`*/U)`: "\u1f5c",
	`*/)u`: "\u1f5c",
	`*/)U`: "\u1f5c",

	`*)/w`: "\u1f6c", // GREEK CAPITAL LETTER OMEGA WITH PSILI AND OXIA
	`*w)/`: "\u1f6c",
	`*w/)`: "\u1f6c",
	`*W)/`: "\u1f6c",
	`*W/)`: "\u1f6c",
	`*)w/`: "\u1f6c",
	`*)W/`: "\u1f6c",
	`*)/W`: "\u1f6c",
	`*/w)`: "\u1f6c",
	`*/W)`: "\u1f6c",
	`*/)w`: "\u1f6c",
	`*/)W`: "\u1f6c",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing and grave accent

	`a)\`: "\u1f02", // GREEK SMALL LETTER ALPHA WITH PSILI AND VARIA
	`a\)`: "\u1f02",
	`A)\`: "\u1f02",
	`A\)`: "\u1f02",

	`e)\`: "\u1f12", // GREEK SMALL LETTER EPSILON WITH PSILI AND VARIA
	`e\)`: "\u1f12",
	`E)\`: "\u1f12",
	`E\)`: "\u1f12",

	`h)\`: "\u1f22", // GREEK SMALL LETTER ETA WITH PSILI AND VARIA
	`h\)`: "\u1f22",
	`H)\`: "\u1f22",
	`H\)`: "\u1f22",

	`i)\`: "\u1f32", // GREEK SMALL LETTER IOTA WITH PSILI AND VARIA
	`i\)`: "\u1f32",
	`I)\`: "\u1f32",
	`I\)`: "\u1f32",

	`o)\`: "\u1f42", // GREEK SMALL LETTER OMICRON WITH PSILI AND VARIA
	`o\)`: "\u1f42",
	`O)\`: "\u1f42",
	`O\)`: "\u1f42",

	`u)\`: "\u1f52", // GREEK SMALL LETTER UPSILON WITH PSILI AND VARIA
	`u\)`: "\u1f52",
	`U)\`: "\u1f52",
	`U\)`: "\u1f52",

	`w)\`: "\u1f62", // GREEK SMALL LETTER OMEGA WITH PSILI AND VARIA
	`w\)`: "\u1f62",
	`W)\`: "\u1f62",
	`W\)`: "\u1f62",

	`*)\a`: "\u1f0a", // GREEK CAPITAL LETTER ALPHA WITH PSILI AND VARIA
	`*a)\`: "\u1f0a",
	`*a\)`: "\u1f0a",
	`*A)\`: "\u1f0a",
	`*A\)`: "\u1f0a",
	`*)a\`: "\u1f0a",
	`*)A\`: "\u1f0a",
	`*)\A`: "\u1f0a",
	`*\a)`: "\u1f0a",
	`*\A)`: "\u1f0a",
	`*\)a`: "\u1f0a",
	`*\)A`: "\u1f0a",

	`*)\e`: "\u1f1a", // GREEK CAPITAL LETTER EPSILON WITH PSILI AND VARIA
	`*e)\`: "\u1f1a",
	`*e\)`: "\u1f1a",
	`*E)\`: "\u1f1a",
	`*E\)`: "\u1f1a",
	`*)e\`: "\u1f1a",
	`*)E\`: "\u1f1a",
	`*)\E`: "\u1f1a",
	`*\e)`: "\u1f1a",
	`*\E)`: "\u1f1a",
	`*\)e`: "\u1f1a",
	`*\)E`: "\u1f1a",

	`*)\h`: "\u1f2a", // GREEK CAPITAL LETTER ETA WITH PSILI AND VARIA
	`*h)\`: "\u1f2a",
	`*h\)`: "\u1f2a",
	`*H)\`: "\u1f2a",
	`*H\)`: "\u1f2a",
	`*)h\`: "\u1f2a",
	`*)H\`: "\u1f2a",
	`*)\H`: "\u1f2a",
	`*\h)`: "\u1f2a",
	`*\H)`: "\u1f2a",
	`*\)h`: "\u1f2a",
	`*\)H`: "\u1f2a",

	`*)\i`: "\u1f3a", // GREEK CAPITAL LETTER IOTA WITH PSILI AND VARIA
	`*i)\`: "\u1f3a",
	`*i\)`: "\u1f3a",
	`*I)\`: "\u1f3a",
	`*I\)`: "\u1f3a",
	`*)i\`: "\u1f3a",
	`*)I\`: "\u1f3a",
	`*)\I`: "\u1f3a",
	`*\i)`: "\u1f3a",
	`*\I)`: "\u1f3a",
	`*\)i`: "\u1f3a",
	`*\)I`: "\u1f3a",

	`*)\o`: "\u1f4a", // GREEK CAPITAL LETTER OMICRON WITH PSILI AND VARIA
	`*o)\`: "\u1f4a",
	`*o\)`: "\u1f4a",
	`*O)\`: "\u1f4a",
	`*O\)`: "\u1f4a",
	`*)o\`: "\u1f4a",
	`*)O\`: "\u1f4a",
	`*)\O`: "\u1f4a",
	`*\o)`: "\u1f4a",
	`*\O)`: "\u1f4a",
	`*\)o`: "\u1f4a",
	`*\)O`: "\u1f4a",

	`*)\u`: "\u1f5a", // Invalid
	`*u\)`: "\u1f5a",
	`*u)\`: "\u1f5a",
	`*U\)`: "\u1f5a",
	`*U)\`: "\u1f5a",
	`*)u\`: "\u1f5a",
	`*)U\`: "\u1f5a",
	`*)\U`: "\u1f5a",
	`*\u)`: "\u1f5a",
	`*\U)`: "\u1f5a",
	`*\)u`: "\u1f5a",
	`*\)U`: "\u1f5a",

	`*)\w`: "\u1f6a", // GREEK CAPITAL LETTER OMEGA WITH PSILI AND VARIA
	`*w)\`: "\u1f6a",
	`*w\)`: "\u1f6a",
	`*W)\`: "\u1f6a",
	`*W\)`: "\u1f6a",
	`*)w\`: "\u1f6a",
	`*)W\`: "\u1f6a",
	`*)\W`: "\u1f6a",
	`*\w)`: "\u1f6a",
	`*\W)`: "\u1f6a",
	`*\)w`: "\u1f6a",
	`*\)W`: "\u1f6a",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breathing and acute accent

	`a(/`: "\u1f05", // GREEK SMALL LETTER ALPHA WITH DASIA AND OXIA
	`a/(`: "\u1f05",
	`A(/`: "\u1f05",
	`A/(`: "\u1f05",

	`e(/`: "\u1f15", // GREEK SMALL LETTER EPSILON WITH DASIA AND OXIA
	`e/(`: "\u1f15",
	`E(/`: "\u1f15",
	`E/(`: "\u1f15",

	`h(/`: "\u1f25", // GREEK SMALL LETTER ETA WITH DASIA AND OXIA
	`h/(`: "\u1f25",
	`H(/`: "\u1f25",
	`H/(`: "\u1f25",

	`i(/`: "\u1f35", // GREEK SMALL LETTER IOTA WITH DASIA AND OXIA
	`i/(`: "\u1f35",
	`I(/`: "\u1f35",
	`I/(`: "\u1f35",

	`o(/`: "\u1f45", // GREEK SMALL LETTER OMICRON WITH DASIA AND OXIA
	`o/(`: "\u1f45",
	`O(/`: "\u1f45",
	`O/(`: "\u1f45",

	`u(/`: "\u1f55", // GREEK SMALL LETTER UPSILON WITH DASIA AND OXIA
	`u/(`: "\u1f55",
	`U(/`: "\u1f55",
	`U/(`: "\u1f55",

	`w(/`: "\u1f65", // GREEK SMALL LETTER OMEGA WITH DASIA AND OXIA
	`w/(`: "\u1f65",
	`W(/`: "\u1f65",
	`W/(`: "\u1f65",

	`*(/a`: "\u1f0d", // GREEK CAPITAL LETTER ALPHA WITH DASIA AND OXIA
	`*a(/`: "\u1f0d",
	`*a/(`: "\u1f0d",
	`*A(/`: "\u1f0d",
	`*A/(`: "\u1f0d",
	`*(a/`: "\u1f0d",
	`*(A/`: "\u1f0d",
	`*(/A`: "\u1f0d",
	`*/a(`: "\u1f0d",
	`*/A(`: "\u1f0d",
	`*/(a`: "\u1f0d",
	`*/(A`: "\u1f0d",

	`*(/e`: "\u1f1d", // GREEK CAPITAL LETTER EPSILON WITH DASIA AND OXIA
	`*e(/`: "\u1f1d",
	`*e/(`: "\u1f1d",
	`*E(/`: "\u1f1d",
	`*E/(`: "\u1f1d",
	`*(e/`: "\u1f1d",
	`*(E/`: "\u1f1d",
	`*(/E`: "\u1f1d",
	`*/e(`: "\u1f1d",
	`*/E(`: "\u1f1d",
	`*/(e`: "\u1f1d",
	`*/(E`: "\u1f1d",

	`*(/h`: "\u1f2d", // GREEK CAPITAL LETTER ETA WITH DASIA AND OXIA
	`*h(/`: "\u1f2d",
	`*h/(`: "\u1f2d",
	`*H(/`: "\u1f2d",
	`*H/(`: "\u1f2d",
	`*(h/`: "\u1f2d",
	`*(H/`: "\u1f2d",
	`*(/H`: "\u1f2d",
	`*/h(`: "\u1f2d",
	`*/H(`: "\u1f2d",
	`*/(h`: "\u1f2d",
	`*/(H`: "\u1f2d",

	`*(/i`: "\u1f3d", // GREEK CAPITAL LETTER IOTA WITH DASIA AND OXIA
	`*i(/`: "\u1f3d",
	`*i/(`: "\u1f3d",
	`*I(/`: "\u1f3d",
	`*I/(`: "\u1f3d",
	`*(i/`: "\u1f3d",
	`*(I/`: "\u1f3d",
	`*(/I`: "\u1f3d",
	`*/i(`: "\u1f3d",
	`*/I(`: "\u1f3d",
	`*/(i`: "\u1f3d",
	`*/(I`: "\u1f3d",

	`*(/o`: "\u1f4d", // GREEK CAPITAL LETTER OMICRON WITH DASIA AND OXIA
	`*o(/`: "\u1f4d",
	`*o/(`: "\u1f4d",
	`*O(/`: "\u1f4d",
	`*O/(`: "\u1f4d",
	`*(o/`: "\u1f4d",
	`*(O/`: "\u1f4d",
	`*(/O`: "\u1f4d",
	`*/o(`: "\u1f4d",
	`*/O(`: "\u1f4d",
	`*/(o`: "\u1f4d",
	`*/(O`: "\u1f4d",

	`*(/u`: "\u1f5d", // GREEK CAPITAL LETTER UPSILON WITH DASIA AND OXIA
	`*u(/`: "\u1f5d",
	`*u/(`: "\u1f5d",
	`*U(/`: "\u1f5d",
	`*U/(`: "\u1f5d",
	`*(u/`: "\u1f5d",
	`*(U/`: "\u1f5d",
	`*(/U`: "\u1f5d",
	`*/u(`: "\u1f5d",
	`*/U(`: "\u1f5d",
	`*/(u`: "\u1f5d",
	`*/(U`: "\u1f5d",

	`*(/w`: "\u1f6d", // GREEK CAPITAL LETTER OMEGA WITH DASIA AND OXIA
	`*w(/`: "\u1f6d",
	`*w/(`: "\u1f6d",
	`*W(/`: "\u1f6d",
	`*W/(`: "\u1f6d",
	`*(w/`: "\u1f6d",
	`*(W/`: "\u1f6d",
	`*(/W`: "\u1f6d",
	`*/w(`: "\u1f6d",
	`*/W(`: "\u1f6d",
	`*/(w`: "\u1f6d",
	`*/(W`: "\u1f6d",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breathing and grave accent

	`a(\`: "\u1f03", // GREEK SMALL LETTER ALPHA WITH DASIA AND VARIA
	`a\(`: "\u1f03",
	`A(\`: "\u1f03",
	`A\(`: "\u1f03",

	`e(\`: "\u1f13", // GREEK SMALL LETTER EPSILON WITH DASIA AND VARIA
	`e\(`: "\u1f13",
	`E(\`: "\u1f13",
	`E\(`: "\u1f13",

	`h(\`: "\u1f23", // GREEK SMALL LETTER ETA WITH DASIA AND VARIA
	`h\(`: "\u1f23",
	`H(\`: "\u1f23",
	`H\(`: "\u1f23",

	`i(\`: "\u1f33", // GREEK SMALL LETTER IOTA WITH DASIA AND VARIA
	`i\(`: "\u1f33",
	`I(\`: "\u1f33",
	`I\(`: "\u1f33",

	`o(\`: "\u1f43", // GREEK SMALL LETTER OMICRON WITH DASIA AND VARIA
	`o\(`: "\u1f43",
	`O(\`: "\u1f43",
	`O\(`: "\u1f43",

	`u(\`: "\u1f53", // GREEK SMALL LETTER UPSILON WITH DASIA AND VARIA
	`u\(`: "\u1f53",
	`U(\`: "\u1f53",
	`U\(`: "\u1f53",

	`w(\`: "\u1f63", // GREEK SMALL LETTER OMEGA WITH DASIA AND VARIA
	`w\(`: "\u1f63",
	`W(\`: "\u1f63",
	`W\(`: "\u1f63",

	`*(\a`: "\u1f0b", // GREEK CAPITAL LETTER ALPHA WITH DASIA AND VARIA
	`*a(\`: "\u1f0b",
	`*a\(`: "\u1f0b",
	`*A(\`: "\u1f0b",
	`*A\(`: "\u1f0b",
	`*(a\`: "\u1f0b",
	`*(A\`: "\u1f0b",
	`*(\A`: "\u1f0b",
	`*\a(`: "\u1f0b",
	`*\A(`: "\u1f0b",
	`*\(a`: "\u1f0b",
	`*\(A`: "\u1f0b",

	`*(\e`: "\u1f1b", // GREEK CAPITAL LETTER EPSILON WITH DASIA AND VARIA
	`*e(\`: "\u1f1b",
	`*e\(`: "\u1f1b",
	`*E(\`: "\u1f1b",
	`*E\(`: "\u1f1b",
	`*(e\`: "\u1f1b",
	`*(E\`: "\u1f1b",
	`*(\E`: "\u1f1b",
	`*\e(`: "\u1f1b",
	`*\E(`: "\u1f1b",
	`*\(e`: "\u1f1b",
	`*\(E`: "\u1f1b",

	`*(\h`: "\u1f2b", // GREEK CAPITAL LETTER ETA WITH DASIA AND VARIA
	`*h(\`: "\u1f2b",
	`*h\(`: "\u1f2b",
	`*H(\`: "\u1f2b",
	`*H\(`: "\u1f2b",
	`*(h\`: "\u1f2b",
	`*(H\`: "\u1f2b",
	`*(\H`: "\u1f2b",
	`*\h(`: "\u1f2b",
	`*\H(`: "\u1f2b",
	`*\(h`: "\u1f2b",
	`*\(H`: "\u1f2b",

	`*(\i`: "\u1f3b", // GREEK CAPITAL LETTER IOTA WITH DASIA AND VARIA
	`*i(\`: "\u1f3b",
	`*i\(`: "\u1f3b",
	`*I(\`: "\u1f3b",
	`*I\(`: "\u1f3b",
	`*(i\`: "\u1f3b",
	`*(I\`: "\u1f3b",
	`*(\I`: "\u1f3b",
	`*\i(`: "\u1f3b",
	`*\I(`: "\u1f3b",
	`*\(i`: "\u1f3b",
	`*\(I`: "\u1f3b",

	`*(\o`: "\u1f4b", // GREEK CAPITAL LETTER OMICRON WITH DASIA AND VARIA
	`*o(\`: "\u1f4b",
	`*o\(`: "\u1f4b",
	`*O(\`: "\u1f4b",
	`*O\(`: "\u1f4b",
	`*(o\`: "\u1f4b",
	`*(O\`: "\u1f4b",
	`*(\O`: "\u1f4b",
	`*\o(`: "\u1f4b",
	`*\O(`: "\u1f4b",
	`*\(o`: "\u1f4b",
	`*\(O`: "\u1f4b",

	`*(\u`: "\u1f5b", // GREEK CAPITAL LETTER UPSILON WITH DASIA AND VARIA
	`*u(\`: "\u1f5b",
	`*u\(`: "\u1f5b",
	`*U(\`: "\u1f5b",
	`*U\(`: "\u1f5b",
	`*(u\`: "\u1f5b",
	`*(U\`: "\u1f5b",
	`*(\U`: "\u1f5b",
	`*\u(`: "\u1f5b",
	`*\U(`: "\u1f5b",
	`*\(u`: "\u1f5b",
	`*\(U`: "\u1f5b",

	`*(\w`: "\u1f6b", // GREEK CAPITAL LETTER OMEGA WITH DASIA AND VARIA
	`*w(\`: "\u1f6b",
	`*w\(`: "\u1f6b",
	`*W(\`: "\u1f6b",
	`*W\(`: "\u1f6b",
	`*(w\`: "\u1f6b",
	`*(W\`: "\u1f6b",
	`*(\W`: "\u1f6b",
	`*\w(`: "\u1f6b",
	`*\W(`: "\u1f6b",
	`*\(w`: "\u1f6b",
	`*\(W`: "\u1f6b",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Perispomeni

	`a=`: "\u1fb6", // GREEK SMALL LETTER ALPHA WITH PERISPOMENI
	`A=`: "\u1fb6",

	`h=`: "\u1fc6", // GREEK SMALL LETTER ETA WITH PERISPOMENI
	`H=`: "\u1fc6",

	`i=`: "\u1fd6", // GREEK SMALL LETTER IOTA WITH PERISPOMENI
	`I=`: "\u1fd6",

	`u=`: "\u1fe6", // GREEK SMALL LETTER UPSILON WITH PERISPOMENI
	`U=`: "\u1fe6",

	`w=`: "\u1ff6", // GREEK SMALL LETTER OMEGA WITH PERISPOMENI
	`W=`: "\u1ff6",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing and perispomeni

	`a)=`: "\u1f06", // GREEK SMALL LETTER ALPHA WITH PSILI AND PERISPOMENI
	`a=)`: "\u1f06",
	`A)=`: "\u1f06",
	`A=)`: "\u1f06",

	`h)=`: "\u1f26", // GREEK SMALL LETTER ETA WITH PSILI AND PERISPOMENI
	`h=)`: "\u1f26",
	`H)=`: "\u1f26",
	`H=)`: "\u1f26",

	`i)=`: "\u1f36", // GREEK SMALL LETTER IOTA WITH PSILI AND PERISPOMENI
	`i=)`: "\u1f36",
	`I)=`: "\u1f36",
	`I=)`: "\u1f36",

	`u)=`: "\u1f56", // GREEK SMALL LETTER UPSILON WITH PSILI AND PERISPOMENI
	`u=)`: "\u1f56",
	`U)=`: "\u1f56",
	`U=)`: "\u1f56",

	`w)=`: "\u1f66", // GREEK SMALL LETTER OMEGA WITH PSILI AND PERISPOMENI
	`w=)`: "\u1f66",
	`W)=`: "\u1f66",
	`W=)`: "\u1f66",

	`*)=a`: "\u1f0e", // GREEK CAPITAL LETTER ALPHA WITH PSILI AND PERISPOMENI
	`*a)=`: "\u1f0e",
	`*a=)`: "\u1f0e",
	`*A)=`: "\u1f0e",
	`*A=)`: "\u1f0e",
	`*)a=`: "\u1f0e",
	`*)A=`: "\u1f0e",
	`*)=A`: "\u1f0e",
	`*=)a`: "\u1f0e",
	`*=)A`: "\u1f0e",
	`*=a)`: "\u1f0e",
	`*=A)`: "\u1f0e",

	`*)=h`: "\u1f2e", // GREEK CAPITAL LETTER ETA WITH PSILI AND PERISPOMENI
	`*h)=`: "\u1f2e",
	`*h=)`: "\u1f2e",
	`*H)=`: "\u1f2e",
	`*H=)`: "\u1f2e",
	`*)h=`: "\u1f2e",
	`*)H=`: "\u1f2e",
	`*)=H`: "\u1f2e",
	`*=)h`: "\u1f2e",
	`*=)H`: "\u1f2e",
	`*=h)`: "\u1f2e",
	`*=H)`: "\u1f2e",

	`*)=i`: "\u1f3e", // GREEK CAPITAL LETTER IOTA WITH PSILI AND PERISPOMENI
	`*i)=`: "\u1f3e",
	`*i=)`: "\u1f3e",
	`*I)=`: "\u1f3e",
	`*I=)`: "\u1f3e",
	`*)i=`: "\u1f3e",
	`*)I=`: "\u1f3e",
	`*)=I`: "\u1f3e",
	`*=)i`: "\u1f3e",
	`*=)I`: "\u1f3e",
	`*=i)`: "\u1f3e",
	`*=I)`: "\u1f3e",

	`*)=w`: "\u1f6e", // GREEK CAPITAL LETTER OMEGA WITH PSILI AND PERISPOMENI
	`*w)=`: "\u1f6e",
	`*w=)`: "\u1f6e",
	`*W)=`: "\u1f6e",
	`*W=)`: "\u1f6e",
	`*)w=`: "\u1f6e",
	`*)W=`: "\u1f6e",
	`*)=W`: "\u1f6e",
	`*=)w`: "\u1f6e",
	`*=)W`: "\u1f6e",
	`*=w)`: "\u1f6e",
	`*=W)`: "\u1f6e",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breathing and perispomeni

	`a(=`: "\u1f07", // GREEK SMALL LETTER ALPHA WITH DASIA AND PERISPOMENI
	`a=(`: "\u1f07",
	`A(=`: "\u1f07",
	`A=(`: "\u1f07",

	`h(=`: "\u1f27", // GREEK SMALL LETTER ETA WITH DASIA AND PERISPOMENI
	`h=(`: "\u1f27",
	`H(=`: "\u1f27",
	`H=(`: "\u1f27",

	`i(=`: "\u1f37", // GREEK SMALL LETTER IOTA WITH DASIA AND PERISPOMENI
	`i=(`: "\u1f37",
	`I(=`: "\u1f37",
	`I=(`: "\u1f37",

	`u(=`: "\u1f57", // GREEK SMALL LETTER UPSILON WITH DASIA AND PERISPOMENI
	`u=(`: "\u1f57",
	`U(=`: "\u1f57",
	`U=(`: "\u1f57",

	`w(=`: "\u1f67", // GREEK SMALL LETTER OMEGA WITH DASIA AND PERISPOMENI
	`w=(`: "\u1f67",
	`W(=`: "\u1f67",
	`W=(`: "\u1f67",

	`*(=a`: "\u1f0f", // GREEK CAPITAL LETTER ALPHA WITH DASIA AND PERISPOMENI
	`*a(=`: "\u1f0f",
	`*a=(`: "\u1f0f",
	`*A(=`: "\u1f0f",
	`*A=(`: "\u1f0f",
	`*(a=`: "\u1f0f",
	`*(A=`: "\u1f0f",
	`*(=A`: "\u1f0f",
	`*=a(`: "\u1f0f",
	`*=A(`: "\u1f0f",
	`*=(a`: "\u1f0f",
	`*=(A`: "\u1f0f",

	`*(=h`: "\u1f2f", // GREEK CAPITAL LETTER ETA WITH DASIA AND PERISPOMENI
	`*h(=`: "\u1f2f",
	`*h=(`: "\u1f2f",
	`*H(=`: "\u1f2f",
	`*H=(`: "\u1f2f",
	`*(h=`: "\u1f2f",
	`*(H=`: "\u1f2f",
	`*(=H`: "\u1f2f",
	`*=h(`: "\u1f2f",
	`*=H(`: "\u1f2f",
	`*=(h`: "\u1f2f",
	`*=(H`: "\u1f2f",

	`*(=i`: "\u1f3f", // GREEK CAPITAL LETTER IOTA WITH DASIA AND PERISPOMENI
	`*i(=`: "\u1f3f",
	`*i=(`: "\u1f3f",
	`*I(=`: "\u1f3f",
	`*I=(`: "\u1f3f",
	`*(i=`: "\u1f3f",
	`*(I=`: "\u1f3f",
	`*(=I`: "\u1f3f",
	`*=i(`: "\u1f3f",
	`*=I(`: "\u1f3f",
	`*=(i`: "\u1f3f",
	`*=(I`: "\u1f3f",

	`*(=u`: "\u1f5f", // GREEK CAPITAL LETTER UPSILON WITH DASIA AND PERISPOMENI
	`*u(=`: "\u1f5f",
	`*u=(`: "\u1f5f",
	`*U(=`: "\u1f5f",
	`*U=(`: "\u1f5f",
	`*(u=`: "\u1f5f",
	`*(U=`: "\u1f5f",
	`*(=U`: "\u1f5f",
	`*=(u`: "\u1f5f",
	`*=(U`: "\u1f5f",
	`*=u(`: "\u1f5f",
	`*=U(`: "\u1f5f",

	`*(=w`: "\u1f6f", // GREEK CAPITAL LETTER OMEGA WITH DASIA AND PERISPOMENI
	`*w(=`: "\u1f6f",
	`*w=(`: "\u1f6f",
	`*W(=`: "\u1f6f",
	`*W=(`: "\u1f6f",
	`*(w=`: "\u1f6f",
	`*(W=`: "\u1f6f",
	`*(=W`: "\u1f6f",
	`*=w(`: "\u1f6f",
	`*=W(`: "\u1f6f",
	`*=(w`: "\u1f6f",
	`*=(W`: "\u1f6f",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Perispomeni and ypogegrammeni

	`a=|`: "\u1fb7", // GREEK SMALL LETTER ALPHA WITH PERISPOMENI AND YPOGEGRAMMENI
	`a|=`: "\u1fb7",
	`A=|`: "\u1fb7",
	`A|=`: "\u1fb7",

	`h=|`: "\u1fc7", // GREEK SMALL LETTER ETA WITH PERISPOMENI AND YPOGEGRAMMENI
	`h|=`: "\u1fc7",
	`H=|`: "\u1fc7",
	`H|=`: "\u1fc7",

	`w=|`: "\u1ff7", // GREEK SMALL LETTER OMEGA WITH PERISPOMENI AND YPOGEGRAMMENI
	`w|=`: "\u1ff7",
	`W=|`: "\u1ff7",
	`W|=`: "\u1ff7",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Ypogegrammeni

	`a|`: "\u1fb3", // GREEK SMALL LETTER ALPHA WITH YPOGEGRAMMENI
	`A|`: "\u1fb3",

	`h|`: "\u1fc3", // GREEK SMALL LETTER ETA WITH YPOGEGRAMMENI
	`H|`: "\u1fc3",

	`w|`: "\u1ff3", // GREEK SMALL LETTER OMEGA WITH YPOGEGRAMMENI
	`W|`: "\u1ff3",

	`*a|`: "\u1fbc", // GREEK CAPITAL LETTER ALPHA WITH PROSGEGRAMMENI
	`*A|`: "\u1fbc",
	`*|a`: "\u1fbc",
	`*|A`: "\u1fbc",

	`*h|`: "\u1fcc", // GREEK CAPITAL LETTER ETA WITH PROSGEGRAMMENI
	`*H|`: "\u1fcc",
	`*|h`: "\u1fcc",
	`*|H`: "\u1fcc",

	`*w|`: "\u1ffc", // GREEK CAPITAL LETTER OMEGA WITH PROSGEGRAMMENI
	`*W|`: "\u1ffc",
	`*|w`: "\u1ffc",
	`*|W`: "\u1ffc",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Acute accent and ypogegrammeni

	`a/|`: "\u1fb4", // GREEK SMALL LETTER ALPHA WITH OXIA AND YPOGEGRAMMENI
	`a|/`: "\u1fb4",
	`A/|`: "\u1fb4",
	`A|/`: "\u1fb4",

	`h/|`: "\u1fc4", // GREEK SMALL LETTER ETA WITH OXIA AND YPOGEGRAMMENI
	`h|/`: "\u1fc4",
	`H/|`: "\u1fc4",
	`H|/`: "\u1fc4",

	`w/|`: "\u1ff4", // GREEK SMALL LETTER OMEGA WITH OXIA AND YPOGEGRAMMENI
	`w|/`: "\u1ff4",
	`W/|`: "\u1ff4",
	`W|/`: "\u1ff4",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing and ypogegrammeni

	`a)|`: "\u1f80", // GREEK SMALL LETTER ALPHA WITH PSILI AND YPOGEGRAMMENI
	`a|)`: "\u1f80",
	`A)|`: "\u1f80",
	`A|)`: "\u1f80",

	`h)|`: "\u1f90", // GREEK SMALL LETTER ETA WITH PSILI AND YPOGEGRAMMENI
	`h|)`: "\u1f90",
	`H)|`: "\u1f90",
	`H|)`: "\u1f90",

	`w)|`: "\u1fa0", // GREEK SMALL LETTER OMEGA WITH PSILI AND YPOGEGRAMMENI
	`w|)`: "\u1fa0",
	`W)|`: "\u1fa0",
	`W|)`: "\u1fa0",

	`*)a|`: "\u1f88", // GREEK CAPITAL LETTER ALPHA WITH PSILI AND PROSGEGRAMMENI
	`*a)|`: "\u1f88",
	`*a|)`: "\u1f88",
	`*A)|`: "\u1f88",
	`*A|)`: "\u1f88",
	`*)A|`: "\u1f88",
	`*)|a`: "\u1f88",
	`*)|A`: "\u1f88",
	`*|a)`: "\u1f88",
	`*|A)`: "\u1f88",
	`*|)a`: "\u1f88",
	`*|)A`: "\u1f88",

	`*)h|`: "\u1f98", // GREEK CAPITAL LETTER ETA WITH PSILI AND PROSGEGRAMMENI
	`*h)|`: "\u1f98",
	`*h|)`: "\u1f98",
	`*H)|`: "\u1f98",
	`*H|)`: "\u1f98",
	`*)H|`: "\u1f98",
	`*)|h`: "\u1f98",
	`*)|H`: "\u1f98",
	`*|h)`: "\u1f98",
	`*|H)`: "\u1f98",
	`*|)h`: "\u1f98",
	`*|)H`: "\u1f98",

	`*)w|`: "\u1fa8", // GREEK CAPITAL LETTER OMEGA WITH PSILI AND PROSGEGRAMMENI
	`*w)|`: "\u1fa8",
	`*w|)`: "\u1fa8",
	`*W)|`: "\u1fa8",
	`*W|)`: "\u1fa8",
	`*)W|`: "\u1fa8",
	`*)|w`: "\u1fa8",
	`*)|W`: "\u1fa8",
	`*|w)`: "\u1fa8",
	`*|W)`: "\u1fa8",
	`*|)w`: "\u1fa8",
	`*|)W`: "\u1fa8",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breathing and ypogegrammeni

	`a(|`: "\u1f81", // GREEK SMALL LETTER ALPHA WITH DASIA AND YPOGEGRAMMENI
	`a|(`: "\u1f81",
	`A(|`: "\u1f81",
	`A|(`: "\u1f81",

	`h(|`: "\u1f91", // GREEK SMALL LETTER ETA WITH DASIA AND YPOGEGRAMMENI
	`h|(`: "\u1f91",
	`H(|`: "\u1f91",
	`H|(`: "\u1f91",

	`w(|`: "\u1fa1", // GREEK SMALL LETTER OMEGA WITH DASIA AND YPOGEGRAMMENI
	`w|(`: "\u1fa1",
	`W(|`: "\u1fa1",
	`W|(`: "\u1fa1",

	`*(a|`: "\u1f89", // GREEK CAPITAL LETTER ALPHA WITH DASIA AND PROSGEGRAMMENI
	`*a(|`: "\u1f89",
	`*a|(`: "\u1f89",
	`*A(|`: "\u1f89",
	`*A|(`: "\u1f89",
	`*(A|`: "\u1f89",
	`*(|a`: "\u1f89",
	`*(|A`: "\u1f89",
	`*|a(`: "\u1f89",
	`*|A(`: "\u1f89",
	`*|(a`: "\u1f89",
	`*|(A`: "\u1f89",

	`*(h|`: "\u1f99", // GREEK CAPITAL LETTER ETA WITH DASIA AND PROSGEGRAMMENI
	`*h(|`: "\u1f99",
	`*h|(`: "\u1f99",
	`*H(|`: "\u1f99",
	`*H|(`: "\u1f99",
	`*(H|`: "\u1f99",
	`*(|h`: "\u1f99",
	`*(|H`: "\u1f99",
	`*|h(`: "\u1f99",
	`*|H(`: "\u1f99",
	`*|(h`: "\u1f99",
	`*|(H`: "\u1f99",

	`*(w|`: "\u1fa9", // GREEK CAPITAL LETTER OMEGA WITH DASIA AND PROSGEGRAMMENI
	`*w(|`: "\u1fa9",
	`*w|(`: "\u1fa9",
	`*W(|`: "\u1fa9",
	`*W|(`: "\u1fa9",
	`*(W|`: "\u1fa9",
	`*(|w`: "\u1fa9",
	`*(|W`: "\u1fa9",
	`*|w(`: "\u1fa9",
	`*|W(`: "\u1fa9",
	`*|(w`: "\u1fa9",
	`*|(W`: "\u1fa9",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing, acute accent, and ypogegrammeni

	`a)\|`: "\u1f82", // GREEK SMALL LETTER ALPHA WITH PSILI AND VARIA AND YPOGEGRAMMENI
	`a)|\`: "\u1f82",
	`a\)|`: "\u1f82",
	`a\|)`: "\u1f82",
	`a|)\`: "\u1f82",
	`a|\)`: "\u1f82",
	`A)\|`: "\u1f82",
	`A)|\`: "\u1f82",
	`A\)|`: "\u1f82",
	`A\|)`: "\u1f82",
	`A|)\`: "\u1f82",
	`A|\)`: "\u1f82",

	`*)\a|`: "\u1f8a", // GREEK CAPITAL LETTER ALPHA WITH PSILI AND VARIA AND PROSGEGRAMMENI
	`*a)|\`: "\u1f8a",
	`*a\)|`: "\u1f8a",
	`*a\|)`: "\u1f8a",
	`*a|)\`: "\u1f8a",
	`*a|\)`: "\u1f8a",
	`*A)\|`: "\u1f8a",
	`*A)|\`: "\u1f8a",
	`*A\)|`: "\u1f8a",
	`*A\|)`: "\u1f8a",
	`*A|)\`: "\u1f8a",
	`*A|\)`: "\u1f8a",
	`*)a\|`: "\u1f8a",
	`*)a|\`: "\u1f8a",
	`*)A\|`: "\u1f8a",
	`*)A|\`: "\u1f8a",
	`*)\A|`: "\u1f8a",
	`*)\|a`: "\u1f8a",
	`*)\|A`: "\u1f8a",
	`*)|a\`: "\u1f8a",
	`*)|A\`: "\u1f8a",
	`*)|\a`: "\u1f8a",
	`*)|\A`: "\u1f8a",
	`*\a)|`: "\u1f8a",
	`*\a|)`: "\u1f8a",
	`*\A)|`: "\u1f8a",
	`*\A|)`: "\u1f8a",
	`*\)a|`: "\u1f8a",
	`*\)A|`: "\u1f8a",
	`*\)|a`: "\u1f8a",
	`*\)|A`: "\u1f8a",
	`*\|)a`: "\u1f8a",
	`*\|)A`: "\u1f8a",
	`*\|a)`: "\u1f8a",
	`*\|A)`: "\u1f8a",
	`*|a)\`: "\u1f8a",
	`*|a\)`: "\u1f8a",
	`*|A)\`: "\u1f8a",
	`*|A\)`: "\u1f8a",
	`*|)a\`: "\u1f8a",
	`*|)A\`: "\u1f8a",
	`*|)\a`: "\u1f8a",
	`*|)\A`: "\u1f8a",
	`*|\)a`: "\u1f8a",
	`*|\)A`: "\u1f8a",
	`*|\a)`: "\u1f8a",
	`*|\A)`: "\u1f8a",

	// `*)\w|`: "\u1faa", // GREEK CAPITAL LETTER OMEGA WITH PSILI AND VARIA AND PROSGEGRAMMENI

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breathing, grave accent, and ypogegrammeni

	`a(\|`: "\u1f83", // GREEK SMALL LETTER ALPHA WITH DASIA AND VARIA AND YPOGEGRAMMENI
	`a(|\`: "\u1f83",
	`a\(|`: "\u1f83",
	`a\|(`: "\u1f83",
	`a|(\`: "\u1f83",
	`a|\(`: "\u1f83",
	`A(\|`: "\u1f83",
	`A(|\`: "\u1f83",
	`A\(|`: "\u1f83",
	`A\|(`: "\u1f83",
	`A|(\`: "\u1f83",
	`A|\(`: "\u1f83",

	`h)\|`: "\u1f93", // GREEK SMALL LETTER ETA WITH DASIA AND VARIA AND YPOGEGRAMMENI
	`h)|\`: "\u1f93",
	`h\)|`: "\u1f93",
	`h\|)`: "\u1f93",
	`h|)\`: "\u1f93",
	`h|\)`: "\u1f93",
	`H)\|`: "\u1f93",
	`H)|\`: "\u1f93",
	`H\)|`: "\u1f93",
	`H\|)`: "\u1f93",
	`H|)\`: "\u1f93",
	`H|\)`: "\u1f93",

	`w)\|`: "\u1fa3", // GREEK SMALL LETTER OMEGA WITH DASIA AND VARIA AND YPOGEGRAMMENI
	`w)|\`: "\u1fa3",
	`w\)|`: "\u1fa3",
	`w\|)`: "\u1fa3",
	`w|)\`: "\u1fa3",
	`w|\)`: "\u1fa3",
	`W)\|`: "\u1fa3",
	`W)|\`: "\u1fa3",
	`W\)|`: "\u1fa3",
	`W\|)`: "\u1fa3",
	`W|)\`: "\u1fa3",
	`W|\)`: "\u1fa3",

	`*(\\a|`: "\u1f8b", // GREEK CAPITAL LETTER ALPHA WITH DASIA AND VARIA AND PROSGEGRAMMENI
	`*a(\|`:  "\u1f8b",
	`*a(|\`:  "\u1f8b",
	`*a\(|`:  "\u1f8b",
	`*a\|(`:  "\u1f8b",
	`*a|(\`:  "\u1f8b",
	`*a|\(`:  "\u1f8b",
	`*A(\|`:  "\u1f8b",
	`*A(|\`:  "\u1f8b",
	`*A\(|`:  "\u1f8b",
	`*A\|(`:  "\u1f8b",
	`*A|(\`:  "\u1f8b",
	`*A|\(`:  "\u1f8b",
	`*(a\|`:  "\u1f8b",
	`*(a|\`:  "\u1f8b",
	`*(A\|`:  "\u1f8b",
	`*(A|\`:  "\u1f8b",
	`*(\a|`:  "\u1f8b",
	`*(\A|`:  "\u1f8b",
	`*(\|a`:  "\u1f8b",
	`*(\|A`:  "\u1f8b",
	`*(|a\`:  "\u1f8b",
	`*(|A\`:  "\u1f8b",
	`*(|\a`:  "\u1f8b",
	`*(|\A`:  "\u1f8b",
	`*\a(|`:  "\u1f8b",
	`*\a|(`:  "\u1f8b",
	`*\A(|`:  "\u1f8b",
	`*\A|(`:  "\u1f8b",
	`*\(a|`:  "\u1f8b",
	`*\(A|`:  "\u1f8b",
	`*\(|a`:  "\u1f8b",
	`*\(|A`:  "\u1f8b",
	`*\|a(`:  "\u1f8b",
	`*\|A(`:  "\u1f8b",
	`*\|(a`:  "\u1f8b",
	`*\|(A`:  "\u1f8b",
	`*|a(\`:  "\u1f8b",
	`*|a\(`:  "\u1f8b",
	`*|A(\`:  "\u1f8b",
	`*|A\(`:  "\u1f8b",
	`*|(a\`:  "\u1f8b",
	`*|(A\`:  "\u1f8b",
	`*|(\a`:  "\u1f8b",
	`*|(\A`:  "\u1f8b",
	`*|\a(`:  "\u1f8b",
	`*|\A(`:  "\u1f8b",
	`*|\(a`:  "\u1f8b",
	`*|\(A`:  "\u1f8b",

	`*)\h|`: "\u1f9b", // GREEK CAPITAL LETTER ETA WITH DASIA AND VARIA AND PROSGEGRAMMENI
	`*h)|\`: "\u1f9b",
	`*h\)|`: "\u1f9b",
	`*h\|)`: "\u1f9b",
	`*h|)\`: "\u1f9b",
	`*h|\)`: "\u1f9b",
	`*H)\|`: "\u1f9b",
	`*H)|\`: "\u1f9b",
	`*H\)|`: "\u1f9b",
	`*H\|)`: "\u1f9b",
	`*H|)\`: "\u1f9b",
	`*H|\)`: "\u1f9b",
	`*)h\|`: "\u1f9b",
	`*)h|\`: "\u1f9b",
	`*)H\|`: "\u1f9b",
	`*)H|\`: "\u1f9b",
	`*)\H|`: "\u1f9b",
	`*)\|h`: "\u1f9b",
	`*)\|H`: "\u1f9b",
	`*)|h\`: "\u1f9b",
	`*)|H\`: "\u1f9b",
	`*)|\h`: "\u1f9b",
	`*)|\H`: "\u1f9b",
	`*\h)|`: "\u1f9b",
	`*\h|)`: "\u1f9b",
	`*\H)|`: "\u1f9b",
	`*\H|)`: "\u1f9b",
	`*\)h|`: "\u1f9b",
	`*\)H|`: "\u1f9b",
	`*\)|h`: "\u1f9b",
	`*\)|H`: "\u1f9b",
	`*\|)h`: "\u1f9b",
	`*\|)H`: "\u1f9b",
	`*\|h)`: "\u1f9b",
	`*\|H)`: "\u1f9b",
	`*|h)\`: "\u1f9b",
	`*|h\)`: "\u1f9b",
	`*|H)\`: "\u1f9b",
	`*|H\)`: "\u1f9b",
	`*|)h\`: "\u1f9b",
	`*|)H\`: "\u1f9b",
	`*|)\h`: "\u1f9b",
	`*|)\H`: "\u1f9b",
	`*|\)h`: "\u1f9b",
	`*|\)H`: "\u1f9b",
	`*|\h)`: "\u1f9b",
	`*|\H)`: "\u1f9b",

	`*)\w|`: "\u1fab", // GREEK CAPITAL LETTER OMEGA WITH DASIA AND VARIA AND PROSGEGRAMMENI
	`*w)\|`: "\u1fab",
	`*w)|\`: "\u1fab",
	`*w\)|`: "\u1fab",
	`*w\|)`: "\u1fab",
	`*w|)\`: "\u1fab",
	`*w|\)`: "\u1fab",
	`*W)\|`: "\u1fab",
	`*W)|\`: "\u1fab",
	`*W\)|`: "\u1fab",
	`*W\|)`: "\u1fab",
	`*W|)\`: "\u1fab",
	`*W|\)`: "\u1fab",
	`*)w\|`: "\u1fab",
	`*)w|\`: "\u1fab",
	`*)W\|`: "\u1fab",
	`*)W|\`: "\u1fab",
	`*)\W|`: "\u1fab",
	`*)\|w`: "\u1fab",
	`*)\|W`: "\u1fab",
	`*)|w\`: "\u1fab",
	`*)|W\`: "\u1fab",
	`*)|\w`: "\u1fab",
	`*)|\W`: "\u1fab",
	`*\w)|`: "\u1fab",
	`*\w|)`: "\u1fab",
	`*\W)|`: "\u1fab",
	`*\W|)`: "\u1fab",
	`*\)w|`: "\u1fab",
	`*\)W|`: "\u1fab",
	`*\)|w`: "\u1fab",
	`*\)|W`: "\u1fab",
	`*\|)w`: "\u1fab",
	`*\|)W`: "\u1fab",
	`*\|w)`: "\u1fab",
	`*\|W)`: "\u1fab",
	`*|w)\`: "\u1fab",
	`*|w\)`: "\u1fab",
	`*|W)\`: "\u1fab",
	`*|W\)`: "\u1fab",
	`*|)w\`: "\u1fab",
	`*|)W\`: "\u1fab",
	`*|)\w`: "\u1fab",
	`*|)\W`: "\u1fab",
	`*|\)w`: "\u1fab",
	`*|\)W`: "\u1fab",
	`*|\w)`: "\u1fab",
	`*|\W)`: "\u1fab",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing, accute accent, and ypogegrammeni

	`a)/|`: "\u1f84", // GREEK SMALL LETTER ALPHA WITH PSILI AND OXIA AND YPOGEGRAMMENI
	`a)|/`: "\u1f84",
	`a/)|`: "\u1f84",
	`a/|)`: "\u1f84",
	`a|/)`: "\u1f84",
	`a|)/`: "\u1f84",
	`A)/|`: "\u1f84",
	`A)|/`: "\u1f84",
	`A/)|`: "\u1f84",
	`A/|)`: "\u1f84",
	`A|/)`: "\u1f84",
	`A|)/`: "\u1f84",

	`h)/|`: "\u1f94", // GREEK SMALL LETTER ETA WITH PSILI AND OXIA AND YPOGEGRAMMENI
	`h)|/`: "\u1f94",
	`h/)|`: "\u1f94",
	`h/|)`: "\u1f94",
	`h|/)`: "\u1f94",
	`h|)/`: "\u1f94",
	`H)/|`: "\u1f94",
	`H)|/`: "\u1f94",
	`H/)|`: "\u1f94",
	`H/|)`: "\u1f94",
	`H|/)`: "\u1f94",
	`H|)/`: "\u1f94",

	`w)/|`: "\u1fa4", // GREEK SMALL LETTER OMEGA WITH PSILI AND OXIA AND YPOGEGRAMMENI
	`w)|/`: "\u1fa4",
	`w/)|`: "\u1fa4",
	`w/|)`: "\u1fa4",
	`w|/)`: "\u1fa4",
	`w|)/`: "\u1fa4",
	`W)/|`: "\u1fa4",
	`W)|/`: "\u1fa4",
	`W/)|`: "\u1fa4",
	`W/|)`: "\u1fa4",
	`W|/)`: "\u1fa4",
	`W|)/`: "\u1fa4",

	`*)/a|`: "\u1f8c", // GREEK CAPITAL LETTER ALPHA WITH PSILI AND OXIA AND PROSGEGRAMMENI
	`*a)/|`: "\u1f8c",
	`*a)|/`: "\u1f8c",
	`*a/)|`: "\u1f8c",
	`*a/|)`: "\u1f8c",
	`*a|)/`: "\u1f8c",
	`*a|/)`: "\u1f8c",
	`*A)/|`: "\u1f8c",
	`*A)|/`: "\u1f8c",
	`*A/)|`: "\u1f8c",
	`*A/|)`: "\u1f8c",
	`*A|)/`: "\u1f8c",
	`*A|/)`: "\u1f8c",
	`*)a/|`: "\u1f8c",
	`*)a|/`: "\u1f8c",
	`*)A/|`: "\u1f8c",
	`*)A|/`: "\u1f8c",
	`*)/A|`: "\u1f8c",
	`*)/|a`: "\u1f8c",
	`*)/|A`: "\u1f8c",
	`*)|a/`: "\u1f8c",
	`*)|A/`: "\u1f8c",
	`*)|/a`: "\u1f8c",
	`*)|/A`: "\u1f8c",
	`*/a)|`: "\u1f8c",
	`*/a|)`: "\u1f8c",
	`*/A)|`: "\u1f8c",
	`*/A|)`: "\u1f8c",
	`*/)a|`: "\u1f8c",
	`*/)A|`: "\u1f8c",
	`*/)|a`: "\u1f8c",
	`*/)|A`: "\u1f8c",
	`*/|)a`: "\u1f8c",
	`*/|)A`: "\u1f8c",
	`*/|a)`: "\u1f8c",
	`*/|A)`: "\u1f8c",
	`*|a)/`: "\u1f8c",
	`*|a/)`: "\u1f8c",
	`*|A)/`: "\u1f8c",
	`*|A/)`: "\u1f8c",
	`*|)a/`: "\u1f8c",
	`*|)A/`: "\u1f8c",
	`*|)/a`: "\u1f8c",
	`*|)/A`: "\u1f8c",
	`*|/)a`: "\u1f8c",
	`*|/)A`: "\u1f8c",
	`*|/a)`: "\u1f8c",
	`*|/A)`: "\u1f8c",

	`*)/h|`: "\u1f9c", // GREEK CAPITAL LETTER ETA WITH PSILI AND OXIA AND PROSGEGRAMMENI
	`*h)/|`: "\u1f9c",
	`*h)|/`: "\u1f9c",
	`*h/)|`: "\u1f9c",
	`*h/|)`: "\u1f9c",
	`*h|)/`: "\u1f9c",
	`*h|/)`: "\u1f9c",
	`*H)/|`: "\u1f9c",
	`*H)|/`: "\u1f9c",
	`*H/)|`: "\u1f9c",
	`*H/|)`: "\u1f9c",
	`*H|)/`: "\u1f9c",
	`*H|/)`: "\u1f9c",
	`*)h/|`: "\u1f9c",
	`*)h|/`: "\u1f9c",
	`*)H/|`: "\u1f9c",
	`*)H|/`: "\u1f9c",
	`*)/H|`: "\u1f9c",
	`*)/|h`: "\u1f9c",
	`*)/|H`: "\u1f9c",
	`*)|h/`: "\u1f9c",
	`*)|H/`: "\u1f9c",
	`*)|/h`: "\u1f9c",
	`*)|/H`: "\u1f9c",
	`*/h)|`: "\u1f9c",
	`*/h|)`: "\u1f9c",
	`*/H)|`: "\u1f9c",
	`*/H|)`: "\u1f9c",
	`*/)h|`: "\u1f9c",
	`*/)H|`: "\u1f9c",
	`*/)|h`: "\u1f9c",
	`*/)|H`: "\u1f9c",
	`*/|)h`: "\u1f9c",
	`*/|)H`: "\u1f9c",
	`*/|h)`: "\u1f9c",
	`*/|H)`: "\u1f9c",
	`*|h)/`: "\u1f9c",
	`*|h/)`: "\u1f9c",
	`*|H)/`: "\u1f9c",
	`*|H/)`: "\u1f9c",
	`*|)h/`: "\u1f9c",
	`*|)H/`: "\u1f9c",
	`*|)/h`: "\u1f9c",
	`*|)/H`: "\u1f9c",
	`*|/)h`: "\u1f9c",
	`*|/)H`: "\u1f9c",
	`*|/h)`: "\u1f9c",
	`*|/H)`: "\u1f9c",

	`*)/w|`: "\u1fac", // GREEK CAPITAL LETTER OMEGA WITH PSILI AND OXIA AND PROSGEGRAMMENI
	`*w)/|`: "\u1fac",
	`*w)|/`: "\u1fac",
	`*w/)|`: "\u1fac",
	`*w/|)`: "\u1fac",
	`*w|)/`: "\u1fac",
	`*w|/)`: "\u1fac",
	`*W)/|`: "\u1fac",
	`*W)|/`: "\u1fac",
	`*W/)|`: "\u1fac",
	`*W/|)`: "\u1fac",
	`*W|)/`: "\u1fac",
	`*W|/)`: "\u1fac",
	`*)w/|`: "\u1fac",
	`*)w|/`: "\u1fac",
	`*)W/|`: "\u1fac",
	`*)W|/`: "\u1fac",
	`*)/W|`: "\u1fac",
	`*)/|w`: "\u1fac",
	`*)/|W`: "\u1fac",
	`*)|w/`: "\u1fac",
	`*)|W/`: "\u1fac",
	`*)|/w`: "\u1fac",
	`*)|/W`: "\u1fac",
	`*/w)|`: "\u1fac",
	`*/w|)`: "\u1fac",
	`*/W)|`: "\u1fac",
	`*/W|)`: "\u1fac",
	`*/)w|`: "\u1fac",
	`*/)W|`: "\u1fac",
	`*/)|w`: "\u1fac",
	`*/)|W`: "\u1fac",
	`*/|)w`: "\u1fac",
	`*/|)W`: "\u1fac",
	`*/|w)`: "\u1fac",
	`*/|W)`: "\u1fac",
	`*|w)/`: "\u1fac",
	`*|w/)`: "\u1fac",
	`*|W)/`: "\u1fac",
	`*|W/)`: "\u1fac",
	`*|)w/`: "\u1fac",
	`*|)W/`: "\u1fac",
	`*|)/w`: "\u1fac",
	`*|)/W`: "\u1fac",
	`*|/)w`: "\u1fac",
	`*|/)W`: "\u1fac",
	`*|/w)`: "\u1fac",
	`*|/W)`: "\u1fac",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breating, acute accent, and ypogegrammeni

	`a(/|`: "\u1f85", // GREEK SMALL LETTER ALPHA WITH DASIA AND OXIA AND YPOGEGRAMMENI
	`a(|/`: "\u1f85",
	`a/(|`: "\u1f85",
	`a/|(`: "\u1f85",
	`a|/(`: "\u1f85",
	`a|(/`: "\u1f85",
	`A(/|`: "\u1f85",
	`A(|/`: "\u1f85",
	`A/(|`: "\u1f85",
	`A/|(`: "\u1f85",
	`A|/(`: "\u1f85",
	`A|(/`: "\u1f85",

	`h(/|`: "\u1f95", // GREEK SMALL LETTER ETA WITH DASIA AND OXIA AND YPOGEGRAMMENI
	`h(|/`: "\u1f95",
	`h/(|`: "\u1f95",
	`h/|(`: "\u1f95",
	`h|/(`: "\u1f95",
	`h|(/`: "\u1f95",
	`H(/|`: "\u1f95",
	`H(|/`: "\u1f95",
	`H/(|`: "\u1f95",
	`H/|(`: "\u1f95",
	`H|/(`: "\u1f95",
	`H|(/`: "\u1f95",

	`w(/|`: "\u1fa5", // GREEK SMALL LETTER OMEGA WITH DASIA AND OXIA AND YPOGEGRAMMENI
	`w(|/`: "\u1fa5",
	`w/(|`: "\u1fa5",
	`w/|(`: "\u1fa5",
	`w|/(`: "\u1fa5",
	`w|(/`: "\u1fa5",
	`W(/|`: "\u1fa5",
	`W(|/`: "\u1fa5",
	`W/(|`: "\u1fa5",
	`W/|(`: "\u1fa5",
	`W|/(`: "\u1fa5",
	`W|(/`: "\u1fa5",

	`*(/a|`: "\u1f8d", // GREEK CAPITAL LETTER ALPHA WITH DASIA AND OXIA AND PROSGEGRAMMENI
	`*a(/|`: "\u1f8d",
	`*a(|/`: "\u1f8d",
	`*a/(|`: "\u1f8d",
	`*a/|(`: "\u1f8d",
	`*a|(/`: "\u1f8d",
	`*a|/(`: "\u1f8d",
	`*A(/|`: "\u1f8d",
	`*A(|/`: "\u1f8d",
	`*A/(|`: "\u1f8d",
	`*A/|(`: "\u1f8d",
	`*A|(/`: "\u1f8d",
	`*A|/(`: "\u1f8d",
	`*(a/|`: "\u1f8d",
	`*(a|/`: "\u1f8d",
	`*(A/|`: "\u1f8d",
	`*(A|/`: "\u1f8d",
	`*(/A|`: "\u1f8d",
	`*(/|a`: "\u1f8d",
	`*(/|A`: "\u1f8d",
	`*(|a/`: "\u1f8d",
	`*(|A/`: "\u1f8d",
	`*(|/a`: "\u1f8d",
	`*(|/A`: "\u1f8d",
	`*/a(|`: "\u1f8d",
	`*/a|(`: "\u1f8d",
	`*/A(|`: "\u1f8d",
	`*/A|(`: "\u1f8d",
	`*/(a|`: "\u1f8d",
	`*/(A|`: "\u1f8d",
	`*/(|a`: "\u1f8d",
	`*/(|A`: "\u1f8d",
	`*/|a(`: "\u1f8d",
	`*/|A(`: "\u1f8d",
	`*/|(a`: "\u1f8d",
	`*/|(A`: "\u1f8d",
	`*|a(/`: "\u1f8d",
	`*|a/(`: "\u1f8d",
	`*|A(/`: "\u1f8d",
	`*|A/(`: "\u1f8d",
	`*|(a/`: "\u1f8d",
	`*|(A/`: "\u1f8d",
	`*|(/a`: "\u1f8d",
	`*|(/A`: "\u1f8d",
	`*|/a(`: "\u1f8d",
	`*|/A(`: "\u1f8d",
	`*|/(a`: "\u1f8d",
	`*|/(A`: "\u1f8d",

	`*(/h|`: "\u1f9d", // GREEK CAPITAL LETTER ETA WITH DASIA AND OXIA AND PROSGEGRAMMENI
	`*h(/|`: "\u1f9d",
	`*h(|/`: "\u1f9d",
	`*h/(|`: "\u1f9d",
	`*h/|(`: "\u1f9d",
	`*h|(/`: "\u1f9d",
	`*h|/(`: "\u1f9d",
	`*H(/|`: "\u1f9d",
	`*H(|/`: "\u1f9d",
	`*H/(|`: "\u1f9d",
	`*H/|(`: "\u1f9d",
	`*H|(/`: "\u1f9d",
	`*H|/(`: "\u1f9d",
	`*(h/|`: "\u1f9d",
	`*(h|/`: "\u1f9d",
	`*(H/|`: "\u1f9d",
	`*(H|/`: "\u1f9d",
	`*(/H|`: "\u1f9d",
	`*(/|h`: "\u1f9d",
	`*(/|H`: "\u1f9d",
	`*(|h/`: "\u1f9d",
	`*(|H/`: "\u1f9d",
	`*(|/h`: "\u1f9d",
	`*(|/H`: "\u1f9d",
	`*/h(|`: "\u1f9d",
	`*/h|(`: "\u1f9d",
	`*/H(|`: "\u1f9d",
	`*/H|(`: "\u1f9d",
	`*/(h|`: "\u1f9d",
	`*/(H|`: "\u1f9d",
	`*/(|h`: "\u1f9d",
	`*/(|H`: "\u1f9d",
	`*/|h(`: "\u1f9d",
	`*/|H(`: "\u1f9d",
	`*/|(h`: "\u1f9d",
	`*/|(H`: "\u1f9d",
	`*|h(/`: "\u1f9d",
	`*|h/(`: "\u1f9d",
	`*|H(/`: "\u1f9d",
	`*|H/(`: "\u1f9d",
	`*|(h/`: "\u1f9d",
	`*|(H/`: "\u1f9d",
	`*|(/h`: "\u1f9d",
	`*|(/H`: "\u1f9d",
	`*|/h(`: "\u1f9d",
	`*|/H(`: "\u1f9d",
	`*|/(h`: "\u1f9d",
	`*|/(H`: "\u1f9d",

	`*(/w|`: "\u1fad", // GREEK CAPITAL LETTER OMEGA WITH DASIA AND OXIA AND PROSGEGRAMMENI
	`*w(/|`: "\u1fad",
	`*w(|/`: "\u1fad",
	`*w/(|`: "\u1fad",
	`*w/|(`: "\u1fad",
	`*w|(/`: "\u1fad",
	`*w|/(`: "\u1fad",
	`*W(/|`: "\u1fad",
	`*W(|/`: "\u1fad",
	`*W/(|`: "\u1fad",
	`*W/|(`: "\u1fad",
	`*W|(/`: "\u1fad",
	`*W|/(`: "\u1fad",
	`*(w/|`: "\u1fad",
	`*(w|/`: "\u1fad",
	`*(W/|`: "\u1fad",
	`*(W|/`: "\u1fad",
	`*(/W|`: "\u1fad",
	`*(/|w`: "\u1fad",
	`*(/|W`: "\u1fad",
	`*(|w/`: "\u1fad",
	`*(|W/`: "\u1fad",
	`*(|/w`: "\u1fad",
	`*(|/W`: "\u1fad",
	`*/w(|`: "\u1fad",
	`*/w|(`: "\u1fad",
	`*/W(|`: "\u1fad",
	`*/W|(`: "\u1fad",
	`*/(w|`: "\u1fad",
	`*/(W|`: "\u1fad",
	`*/(|w`: "\u1fad",
	`*/(|W`: "\u1fad",
	`*/|w(`: "\u1fad",
	`*/|W(`: "\u1fad",
	`*/|(w`: "\u1fad",
	`*/|(W`: "\u1fad",
	`*|w(/`: "\u1fad",
	`*|w/(`: "\u1fad",
	`*|W(/`: "\u1fad",
	`*|W/(`: "\u1fad",
	`*|(w/`: "\u1fad",
	`*|(W/`: "\u1fad",
	`*|(/w`: "\u1fad",
	`*|(/W`: "\u1fad",
	`*|/w(`: "\u1fad",
	`*|/W(`: "\u1fad",
	`*|/(w`: "\u1fad",
	`*|/(W`: "\u1fad",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Smooth breathing, ypogegrammeni, and perispomeni

	`a)=|`: "\u1f86", // GREEK SMALL LETTER ALPHA WITH PSILI AND PERISPOMENI AND YPOGEGRAMMENI
	`a)|=`: "\u1f86",
	`a=)|`: "\u1f86",
	`a=|)`: "\u1f86",
	`a|=)`: "\u1f86",
	`a|)=`: "\u1f86",
	`A)=|`: "\u1f86",
	`A)|=`: "\u1f86",
	`A=)|`: "\u1f86",
	`A=|)`: "\u1f86",
	`A|=)`: "\u1f86",
	`A|)=`: "\u1f86",

	`h)=|`: "\u1f96", // GREEK SMALL LETTER ETA WITH PSILI AND PERISPOMENI AND YPOGEGRAMMENI
	`h)|=`: "\u1f96",
	`h=)|`: "\u1f96",
	`h=|)`: "\u1f96",
	`h|=)`: "\u1f96",
	`h|)=`: "\u1f96",
	`H)=|`: "\u1f96",
	`H)|=`: "\u1f96",
	`H=)|`: "\u1f96",
	`H=|)`: "\u1f96",
	`H|=)`: "\u1f96",
	`H|)=`: "\u1f96",

	`w)=|`: "\u1fa6", // GREEK SMALL LETTER OMEGA WITH PSILI AND PERISPOMENI AND YPOGEGRAMMENI
	`w)|=`: "\u1fa6",
	`w=)|`: "\u1fa6",
	`w=|)`: "\u1fa6",
	`w|=)`: "\u1fa6",
	`w|)=`: "\u1fa6",
	`W)=|`: "\u1fa6",
	`W)|=`: "\u1fa6",
	`W=)|`: "\u1fa6",
	`W=|)`: "\u1fa6",
	`W|=)`: "\u1fa6",
	`W|)=`: "\u1fa6",

	`*)=a|`: "\u1f8e", // GREEK CAPITAL LETTER ALPHA WITH PSILI AND PERISPOMENI AND PROSGEGRAMMENI
	`*a)=|`: "\u1f8e",
	`*a)|=`: "\u1f8e",
	`*a=)|`: "\u1f8e",
	`*a=|)`: "\u1f8e",
	`*a|)=`: "\u1f8e",
	`*a|=)`: "\u1f8e",
	`*A)=|`: "\u1f8e",
	`*A)|=`: "\u1f8e",
	`*A=)|`: "\u1f8e",
	`*A=|)`: "\u1f8e",
	`*A|)=`: "\u1f8e",
	`*A|=)`: "\u1f8e",
	`*)a=|`: "\u1f8e",
	`*)a|=`: "\u1f8e",
	`*)A=|`: "\u1f8e",
	`*)A|=`: "\u1f8e",
	`*)=A|`: "\u1f8e",
	`*)=|a`: "\u1f8e",
	`*)=|A`: "\u1f8e",
	`*)|a=`: "\u1f8e",
	`*)|A=`: "\u1f8e",
	`*)|=a`: "\u1f8e",
	`*)|=A`: "\u1f8e",
	`*=)a|`: "\u1f8e",
	`*=)A|`: "\u1f8e",
	`*=)|a`: "\u1f8e",
	`*=)|A`: "\u1f8e",
	`*=a)|`: "\u1f8e",
	`*=a|)`: "\u1f8e",
	`*=A)|`: "\u1f8e",
	`*=A|)`: "\u1f8e",
	`*=|)a`: "\u1f8e",
	`*=|)A`: "\u1f8e",
	`*=|a)`: "\u1f8e",
	`*=|A)`: "\u1f8e",
	`*|a)=`: "\u1f8e",
	`*|a=)`: "\u1f8e",
	`*|A)=`: "\u1f8e",
	`*|A=)`: "\u1f8e",
	`*|)a=`: "\u1f8e",
	`*|)A=`: "\u1f8e",
	`*|)=a`: "\u1f8e",
	`*|)=A`: "\u1f8e",
	`*|=)a`: "\u1f8e",
	`*|=)A`: "\u1f8e",
	`*|=a)`: "\u1f8e",
	`*|=A)`: "\u1f8e",

	`*)=h|`: "\u1f9e", // GREEK CAPITAL LETTER ETA WITH PSILI AND PERISPOMENI AND PROSGEGRAMMENI
	`*h)=|`: "\u1f9e",
	`*h)|=`: "\u1f9e",
	`*h=)|`: "\u1f9e",
	`*h=|)`: "\u1f9e",
	`*h|)=`: "\u1f9e",
	`*h|=)`: "\u1f9e",
	`*H)=|`: "\u1f9e",
	`*H)|=`: "\u1f9e",
	`*H=)|`: "\u1f9e",
	`*H=|)`: "\u1f9e",
	`*H|)=`: "\u1f9e",
	`*H|=)`: "\u1f9e",
	`*)h=|`: "\u1f9e",
	`*)h|=`: "\u1f9e",
	`*)H=|`: "\u1f9e",
	`*)H|=`: "\u1f9e",
	`*)=H|`: "\u1f9e",
	`*)=|h`: "\u1f9e",
	`*)=|H`: "\u1f9e",
	`*)|h=`: "\u1f9e",
	`*)|H=`: "\u1f9e",
	`*)|=h`: "\u1f9e",
	`*)|=H`: "\u1f9e",
	`*=)h|`: "\u1f9e",
	`*=)H|`: "\u1f9e",
	`*=)|h`: "\u1f9e",
	`*=)|H`: "\u1f9e",
	`*=h)|`: "\u1f9e",
	`*=h|)`: "\u1f9e",
	`*=H)|`: "\u1f9e",
	`*=H|)`: "\u1f9e",
	`*=|)h`: "\u1f9e",
	`*=|)H`: "\u1f9e",
	`*=|h)`: "\u1f9e",
	`*=|H)`: "\u1f9e",
	`*|h)=`: "\u1f9e",
	`*|h=)`: "\u1f9e",
	`*|H)=`: "\u1f9e",
	`*|H=)`: "\u1f9e",
	`*|)h=`: "\u1f9e",
	`*|)H=`: "\u1f9e",
	`*|)=h`: "\u1f9e",
	`*|)=H`: "\u1f9e",
	`*|=)h`: "\u1f9e",
	`*|=)H`: "\u1f9e",
	`*|=h)`: "\u1f9e",
	`*|=H)`: "\u1f9e",

	`*)=w|`: "\u1fae", // GREEK CAPITAL LETTER OMEGA WITH PSILI AND PERISPOMENI AND PROSGEGRAMMENI
	`*w)=|`: "\u1fae",
	`*w)|=`: "\u1fae",
	`*w=)|`: "\u1fae",
	`*w=|)`: "\u1fae",
	`*w|)=`: "\u1fae",
	`*w|=)`: "\u1fae",
	`*W)=|`: "\u1fae",
	`*W)|=`: "\u1fae",
	`*W=)|`: "\u1fae",
	`*W=|)`: "\u1fae",
	`*W|)=`: "\u1fae",
	`*W|=)`: "\u1fae",
	`*)w=|`: "\u1fae",
	`*)w|=`: "\u1fae",
	`*)W=|`: "\u1fae",
	`*)W|=`: "\u1fae",
	`*)=W|`: "\u1fae",
	`*)=|w`: "\u1fae",
	`*)=|W`: "\u1fae",
	`*)|w=`: "\u1fae",
	`*)|W=`: "\u1fae",
	`*)|=w`: "\u1fae",
	`*)|=W`: "\u1fae",
	`*=)w|`: "\u1fae",
	`*=)W|`: "\u1fae",
	`*=)|w`: "\u1fae",
	`*=)|W`: "\u1fae",
	`*=w)|`: "\u1fae",
	`*=w|)`: "\u1fae",
	`*=W)|`: "\u1fae",
	`*=W|)`: "\u1fae",
	`*=|)w`: "\u1fae",
	`*=|)W`: "\u1fae",
	`*=|w)`: "\u1fae",
	`*=|W)`: "\u1fae",
	`*|w)=`: "\u1fae",
	`*|w=)`: "\u1fae",
	`*|W)=`: "\u1fae",
	`*|W=)`: "\u1fae",
	`*|)w=`: "\u1fae",
	`*|)W=`: "\u1fae",
	`*|)=w`: "\u1fae",
	`*|)=W`: "\u1fae",
	`*|=)w`: "\u1fae",
	`*|=)W`: "\u1fae",
	`*|=w)`: "\u1fae",
	`*|=W)`: "\u1fae",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Rough breathing, ypogegrammeni, and perispomeni

	`a(=|`: "\u1f87", // GREEK SMALL LETTER ALPHA WITH DASIA AND PERISPOMENI AND YPOGEGRAMMENI
	`a(|=`: "\u1f87",
	`a=(|`: "\u1f87",
	`a=|(`: "\u1f87",
	`a|=(`: "\u1f87",
	`a|(=`: "\u1f87",
	`A(=|`: "\u1f87",
	`A(|=`: "\u1f87",
	`A=(|`: "\u1f87",
	`A=|(`: "\u1f87",
	`A|=(`: "\u1f87",
	`A|(=`: "\u1f87",

	`h(=|`: "\u1f97", // GREEK SMALL LETTER ETA WITH DASIA AND PERISPOMENI AND YPOGEGRAMMENI
	`h(|=`: "\u1f97",
	`h=(|`: "\u1f97",
	`h=|(`: "\u1f97",
	`h|=(`: "\u1f97",
	`h|(=`: "\u1f97",
	`H(=|`: "\u1f97",
	`H(|=`: "\u1f97",
	`H=(|`: "\u1f97",
	`H=|(`: "\u1f97",
	`H|=(`: "\u1f97",
	`H|(=`: "\u1f97",

	`w(=|`: "\u1fa7", // GREEK SMALL LETTER OMEGA WITH DASIA AND PERISPOMENI AND YPOGEGRAMMENI
	`w(|=`: "\u1fa7",
	`w=(|`: "\u1fa7",
	`w=|(`: "\u1fa7",
	`w|=(`: "\u1fa7",
	`w|(=`: "\u1fa7",
	`W(=|`: "\u1fa7",
	`W(|=`: "\u1fa7",
	`W=(|`: "\u1fa7",
	`W=|(`: "\u1fa7",
	`W|=(`: "\u1fa7",
	`W|(=`: "\u1fa7",

	`*(=a|`: "\u1f8f", // GREEK CAPITAL LETTER ALPHA WITH DASIA AND PERISPOMENI AND PROSGEGRAMMENI
	`*a(=|`: "\u1f8f",
	`*a(|=`: "\u1f8f",
	`*a=(|`: "\u1f8f",
	`*a=|(`: "\u1f8f",
	`*a|(=`: "\u1f8f",
	`*a|=(`: "\u1f8f",
	`*A(=|`: "\u1f8f",
	`*A(|=`: "\u1f8f",
	`*A=(|`: "\u1f8f",
	`*A=|(`: "\u1f8f",
	`*A|(=`: "\u1f8f",
	`*A|=(`: "\u1f8f",
	`*(a=|`: "\u1f8f",
	`*(a|=`: "\u1f8f",
	`*(A=|`: "\u1f8f",
	`*(A|=`: "\u1f8f",
	`*(=A|`: "\u1f8f",
	`*(=|a`: "\u1f8f",
	`*(=|A`: "\u1f8f",
	`*(|a=`: "\u1f8f",
	`*(|A=`: "\u1f8f",
	`*(|=a`: "\u1f8f",
	`*(|=A`: "\u1f8f",
	`*=a(|`: "\u1f8f",
	`*=a|(`: "\u1f8f",
	`*=A(|`: "\u1f8f",
	`*=A|(`: "\u1f8f",
	`*=(a|`: "\u1f8f",
	`*=(A|`: "\u1f8f",
	`*=(|a`: "\u1f8f",
	`*=(|A`: "\u1f8f",
	`*=|a(`: "\u1f8f",
	`*=|A(`: "\u1f8f",
	`*=|(a`: "\u1f8f",
	`*=|(A`: "\u1f8f",
	`*|a(=`: "\u1f8f",
	`*|a=(`: "\u1f8f",
	`*|A(=`: "\u1f8f",
	`*|A=(`: "\u1f8f",
	`*|(a=`: "\u1f8f",
	`*|(A=`: "\u1f8f",
	`*|(=a`: "\u1f8f",
	`*|(=A`: "\u1f8f",
	`*|=a(`: "\u1f8f",
	`*|=A(`: "\u1f8f",
	`*|=(a`: "\u1f8f",
	`*|=(A`: "\u1f8f",

	`*(=h|`: "\u1f9f", // GREEK CAPITAL LETTER ETA WITH DASIA AND PERISPOMENI AND PROSGEGRAMMENI
	`*h(=|`: "\u1f9f",
	`*h(|=`: "\u1f9f",
	`*h=(|`: "\u1f9f",
	`*h=|(`: "\u1f9f",
	`*h|(=`: "\u1f9f",
	`*h|=(`: "\u1f9f",
	`*H(=|`: "\u1f9f",
	`*H(|=`: "\u1f9f",
	`*H=(|`: "\u1f9f",
	`*H=|(`: "\u1f9f",
	`*H|(=`: "\u1f9f",
	`*H|=(`: "\u1f9f",
	`*(h=|`: "\u1f9f",
	`*(h|=`: "\u1f9f",
	`*(H=|`: "\u1f9f",
	`*(H|=`: "\u1f9f",
	`*(=H|`: "\u1f9f",
	`*(=|h`: "\u1f9f",
	`*(=|H`: "\u1f9f",
	`*(|h=`: "\u1f9f",
	`*(|H=`: "\u1f9f",
	`*(|=h`: "\u1f9f",
	`*(|=H`: "\u1f9f",
	`*=h(|`: "\u1f9f",
	`*=h|(`: "\u1f9f",
	`*=H(|`: "\u1f9f",
	`*=H|(`: "\u1f9f",
	`*=(h|`: "\u1f9f",
	`*=(H|`: "\u1f9f",
	`*=(|h`: "\u1f9f",
	`*=(|H`: "\u1f9f",
	`*=|h(`: "\u1f9f",
	`*=|H(`: "\u1f9f",
	`*=|(h`: "\u1f9f",
	`*=|(H`: "\u1f9f",
	`*|h(=`: "\u1f9f",
	`*|h=(`: "\u1f9f",
	`*|H(=`: "\u1f9f",
	`*|H=(`: "\u1f9f",
	`*|(h=`: "\u1f9f",
	`*|(H=`: "\u1f9f",
	`*|(=h`: "\u1f9f",
	`*|(=H`: "\u1f9f",
	`*|=h(`: "\u1f9f",
	`*|=H(`: "\u1f9f",
	`*|=(h`: "\u1f9f",
	`*|=(H`: "\u1f9f",

	`*(=w|`: "\u1faf", // GREEK CAPITAL LETTER OMEGA WITH DASIA AND PERISPOMENI AND PROSGEGRAMMENI
	`*w(=|`: "\u1faf",
	`*w(|=`: "\u1faf",
	`*w=(|`: "\u1faf",
	`*w=|(`: "\u1faf",
	`*w|(=`: "\u1faf",
	`*w|=(`: "\u1faf",
	`*W(=|`: "\u1faf",
	`*W(|=`: "\u1faf",
	`*W=(|`: "\u1faf",
	`*W=|(`: "\u1faf",
	`*W|(=`: "\u1faf",
	`*W|=(`: "\u1faf",
	`*(w=|`: "\u1faf",
	`*(w|=`: "\u1faf",
	`*(W=|`: "\u1faf",
	`*(W|=`: "\u1faf",
	`*(=W|`: "\u1faf",
	`*(=|w`: "\u1faf",
	`*(=|W`: "\u1faf",
	`*(|w=`: "\u1faf",
	`*(|W=`: "\u1faf",
	`*(|=w`: "\u1faf",
	`*(|=W`: "\u1faf",
	`*=w(|`: "\u1faf",
	`*=w|(`: "\u1faf",
	`*=W(|`: "\u1faf",
	`*=W|(`: "\u1faf",
	`*=(w|`: "\u1faf",
	`*=(W|`: "\u1faf",
	`*=(|w`: "\u1faf",
	`*=(|W`: "\u1faf",
	`*=|w(`: "\u1faf",
	`*=|W(`: "\u1faf",
	`*=|(w`: "\u1faf",
	`*=|(W`: "\u1faf",
	`*|w(=`: "\u1faf",
	`*|w=(`: "\u1faf",
	`*|W(=`: "\u1faf",
	`*|W=(`: "\u1faf",
	`*|(w=`: "\u1faf",
	`*|(W=`: "\u1faf",
	`*|(=w`: "\u1faf",
	`*|(=W`: "\u1faf",
	`*|=w(`: "\u1faf",
	`*|=W(`: "\u1faf",
	`*|=(w`: "\u1faf",
	`*|=(W`: "\u1faf",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Diaeresis

	`i+`: "\u03ca", // GREEK SMALL LETTER IOTA WITH DIALYTIKA
	`I+`: "\u03ca",

	`*+i`: "\u03aa", // GREEK CAPITAL LETTER IOTA WITH DIALYTIKA
	`*i+`: "\u03aa",
	`*I+`: "\u03aa",
	`*+I`: "\u03aa",

	`i\+`: "\u1fd2", // GREEK SMALL LETTER IOTA WITH DIALYTIKA AND VARIA
	`i+\`: "\u1fd2",
	`I\+`: "\u1fd2",
	`I+\`: "\u1fd2",

	`i/+`: "\u1fd3", // GREEK SMALL LETTER IOTA WITH DIALYTIKA AND OXIA
	`i+/`: "\u1fd3",
	`I/+`: "\u1fd3",
	`I+/`: "\u1fd3",

	`i=+`: "\u1fd7", // GREEK SMALL LETTER IOTA WITH DIALYTIKA AND PERISPOMENI
	`i+=`: "\u1fd7",
	`I=+`: "\u1fd7",
	`I+=`: "\u1fd7",

	`u+`: "\u03cb", // GREEK SMALL LETTER UPSILON WITH DIALYTIKA
	`U+`: "\u03cb",

	`*+u`: "\u03ab", // GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA
	`*u+`: "\u03ab",
	`*U+`: "\u03ab",
	`*+U`: "\u03ab",

	`u\+`: "\u1fe2", // GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND VARIA
	`u+\`: "\u1fe2",
	`U\+`: "\u1fe2",
	`U+\`: "\u1fe2",

	`u/+`: "\u1fe3", // GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND OXIA
	`u+/`: "\u1fe3",
	`U/+`: "\u1fe3",
	`U+/`: "\u1fe3",

	`u=+`: "\u1fe7", // GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND PERISPOMENI
	`u+=`: "\u1fe7",
	`U=+`: "\u1fe7",
	`U+=`: "\u1fe7",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Macron

	`a&`: "\u1fb0", // GREEK SMALL LETTER ALPHA WITH VRACHY
	`A&`: "\u1fb0",

	`i&`: "\u1fd0", // GREEK SMALL LETTER IOTA WITH VRACHY
	`I&`: "\u1fd0",

	`u&`: "\u1fe0", // GREEK SMALL LETTER UPSILON WITH VRACHY
	`U&`: "\u1fe0",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Breve

	`a\'`: "\u1fb1", // GREEK SMALL LETTER ALPHA WITH MACRON
	`a'`:  "\u1fb1",
	`A'`:  "\u1fb1",

	`i\'`: "\u1fd1", // GREEK SMALL LETTER IOTA WITH MACRON
	`i'`:  "\u1fd1",
	`I'`:  "\u1fd1",

	`u\'`: "\u1fe1", // GREEK SMALL LETTER UPSILON WITH MACRON
	`u'`:  "\u1fe1",
	`U'`:  "\u1fe1",

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Basic punctuation

	":": "\u00b7", // MIDDLE DOT
	`'`: "\u2019", // RIGHT SINGLE QUOTATION MARK
	"-": "\u2010", // HYPHEN
	"_": "\u2014", // EM DASH
	" ": " ",
}
