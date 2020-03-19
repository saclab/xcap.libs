package spacegroup

// Get : return the spacegroup form a spacegroup finger print
// http://img.chem.ucl.ac.uk/sgp/large/sgp.htm
func Get(fingerprint string) string {
	var sg map[string]string
	sg = make(map[string]string)

	// Triclinic
	sg["P1"] = "P 1"
	sg["P-1"] = "P -1"

	// Monolithic
	sg["P121"] = "P 1 2 1"
	sg["P1211"] = "P 1 21 1"
	sg["C121"] = "C 1 2 1"
	sg["P1m1"] = "P 1 m 1"
	sg["P1c1"] = "P 1 c 1"
	sg["C1m1"] = "C 1 m 1"
	sg["C1c1"] = "C 1 c 1"
	sg["P12/m1"] = "P 1 2 / m 1"
	sg["P121/m1"] = " P 1 21 / m 1"
	sg["C12/m1"] = "C 1 2 / m 1"
	sg["P12/c1"] = "P 1 2 / c 1"
	sg["P121/c1"] = "P 1 21 / c 1"
	sg["C12/c1"] = "C 1 2 / c 1"

	// Orthorombic
	sg["P222"] = "P 2 2 2"
	sg["P2221"] = "P 2 2 21"
	sg["P21212"] = "P 21 21 2"
	sg["P212121"] = "P 21 21 2 1"
	sg["C2221"] = "C 2 2 21"
	sg["C222"] = "C 2 2 2"
	sg["F222"] = "F 2 2 2"
	sg["I222"] = "I 2 2 2"
	sg["I212121"] = "I 21 21 2 1"
	sg["Pmm2"] = "P m m 2"
	sg["Pmc21"] = "P m c 21"
	sg["Pcc2"] = "P c c 2"
	sg["Pma2"] = "P m a 2"
	sg["Pca21"] = "P c a 21"
	sg["Pnc2"] = "P n c 2"
	sg["Pmn21"] = "P m n 21"
	sg["Pba2"] = "P b a 2"
	sg["Pna21"] = "P n a 21"
	sg["Pnn2"] = "P n n 2"
	sg["Cmm2"] = "C m m 2"
	sg["Cmc21"] = "C m c 21"
	sg["Ccc2"] = "C c c 2"
	sg["Amm2"] = "A m m 2"
	sg["Abm2"] = "A b m 2"
	sg["Ama2"] = "A m a 2"
	sg["Aba2"] = "A b a 2"
	sg["Fmm2"] = "F m m 2"
	sg["Fdd2"] = "F d d 2"
	sg["Imm2"] = "I m m 2"
	sg["Iba2"] = "I b a 2"
	sg["Ima2"] = "I m a 2"
	sg["Pmmm"] = "P m m m"
	sg["Pnnn"] = "P n n n"
	sg["Pccm"] = "P c c m"
	sg["Pban"] = "P b a n"
	sg["Pmma"] = "P m m a"
	sg["Pnna"] = "P n n a"
	sg["Pmna"] = "P m n a"
	sg["Pcca"] = "P c c a"
	sg["Pbam"] = "P b a m"
	sg["Pccn"] = "P c c n"
	sg["Pbcm"] = "P b c m"
	sg["Pnnm"] = "P n n m"
	sg["Pmmn"] = "P m m n"
	sg["Pbcn"] = "P b c n"
	sg["Pbca"] = "P b c a"
	sg["Pnma"] = "P n m a"
	sg["Cmcm"] = "C m c m"
	sg["Cmca"] = "C m c a"
	sg["Cmmm"] = "C m m m"
	sg["Cccm"] = "C c c m"
	sg["Cmma"] = "C m m a"
	sg["Ccca"] = "C c c a"
	sg["Fmmm"] = "F m m m"
	sg["Fddd"] = "F d d d"
	sg["Immm"] = "I m m m"
	sg["Ibam"] = "I b a m"
	sg["Ibca"] = "I b c a"
	sg["Imma"] = " I m m a"

	// Tetragonal
	sg["P4"] = "P 4"
	sg["P41"] = "P 41"
	sg["P42"] = "P 42"
	sg["P43"] = "P 43"
	sg["I4"] = "I 4"
	sg["I41"] = "I 41"
	sg["P-4"] = "P -4"
	sg["I-4"] = "I -4"
	sg["P4/m"] = "P 4 / m"
	sg["P42/m"] = "P 42 / m"
	sg["P4/n"] = "P 4 / n"
	sg["P42/n"] = "P 42 / n"
	sg["I4/m"] = "I 4 / m"
	sg["I41/a"] = "I 41 / a"
	sg["P422"] = "P 4 2 2"
	sg["P4212"] = "P 4 21 2"
	sg["P4122"] = "P 41 2 2"
	sg["P41212"] = "P 41 21 2"
	sg["P4222"] = "P 42 2 2"
	sg["P42212"] = "P 42 21 2"
	sg["P4322"] = "P 43 2 2"
	sg["P43212"] = "P 43 21 2"
	sg["I422"] = "I 4 2 2"
	sg["I4122"] = "I 41 2 2"
	sg["P4mm"] = "P 4 m m"
	sg["P4bm"] = "P 4 b m"
	sg["P42cm"] = "P 42 c m"
	sg["P42nm"] = "P 42 n m"
	sg["P4cc"] = "P 4 c c"
	sg["P4nc"] = "P 4 n c"
	sg["P42mc"] = "P 42 m c"
	sg["P42bc"] = "P 42 b c"
	sg["I4mm"] = "I 4 m m"
	sg["I4cm"] = "I 4 c m"
	sg["I41md"] = "I 41 m d"
	sg["I41cd"] = "I 41 c d"
	sg["P-42m"] = "P -4 2 m"
	sg["P-42c"] = "P -4 2 c"
	sg["P-421m"] = "P -4 21 m"
	sg["P-421c"] = "P -4 21 c"
	sg["P-4m2"] = "P -4 m 2"
	sg["P-4c2"] = "P -4 c 2"
	sg["P-4b2"] = "P -4 b 2	"
	sg["P-4n2"] = "P -4 n 2"
	sg["I-4m2"] = "I-4 m 2"
	sg["I-4c2"] = "I -4 c 2	"
	sg["I-42m"] = "I -4 2 m"
	sg["I-42d"] = "I -4 2 d"
	sg["P4/mmm"] = "P 4 / m m m"
	sg["P4/mcc"] = "P 4 / m c c"
	sg["P4/nbm"] = "P 4 / n b m"
	sg["P4/nnc"] = "P 4 / n n c"
	sg["P4/mbm"] = "P 4 / m b m"
	sg["P4/mnc"] = "P 4 / m n c"
	sg["P4/nmm"] = "P 4 / n m m"
	sg["P4/ncc"] = "P 4 / n c c"
	sg["P42/mmc"] = "P 42 / m m c"
	sg["P42/mcm"] = "P 42 / m c m"
	sg["P42/nbc"] = "P 42 / n b c	"
	sg["P42/nnm"] = "P 42 / n n m"
	sg["P42/mbc"] = "P 42 / m b c"
	sg["P42/mnm"] = "P 42 / m n m"
	sg["P42/nmc"] = "P 42 / n m c"
	sg["P42/ncm"] = "P 42 / n c m"
	sg["I4/mmm"] = "I 4 / m m m"
	sg["I4/mcm"] = "I 4 / m c m"
	sg["I41/amd"] = "I 41 / a m d"
	sg["I41/acd"] = "I 41 / a c d"

	// Trigonal
	sg["P3"] = "P 3"
	sg["P31"] = "P 31"
	sg["P32"] = "P 32"
	sg["R3"] = "R 3"
	sg["R-3"] = "R -3"
	sg["P312"] = "P 3 1 2"
	sg["P321"] = "P 3 2 1"
	sg["P3112"] = "P 31 1 2"
	sg["P3121"] = "P 31 2 1"
	sg["R3212"] = "P 32 1 2"
	sg["R3221"] = "P 32 2 1"
	sg["R32"] = "R 3 2"
	sg["P3m1"] = "P 3 m 1"
	sg["P31m"] = "P 3 1 m"
	sg["P3c1"] = "P 3 c 1"
	sg["P31c"] = "P 3 1 c"
	sg["R3m"] = "R 3 m"
	sg["R3c"] = "R 3 c"
	sg["P-31m"] = "P -3 1 m"
	sg["P-31c"] = "P -3 1 c"
	sg["P-3m1"] = "P -3 m 1"
	sg["P-3c1"] = "P -3 c 1"
	sg["R-3m"] = "R -3 m"
	sg["R-3c"] = "R -3 c"

	// Hexagonal
	sg["P6"] = "P 6"
	sg["P61"] = "P 61"
	sg["P65"] = "P 65"
	sg["P62"] = "P 62"
	sg["P64"] = "P 64"
	sg["P63"] = "P 63"
	sg["P-6"] = "P -6"
	sg["P6/m"] = "P 6 / m"
	sg["P63/m"] = "P 63 / m"
	sg["P622"] = "P 6 2 2"
	sg["P6122"] = "P 61 2 2"
	sg["P6522"] = "P 65 2 2"
	sg["P6222"] = "P 62 2 2"
	sg["P6422"] = "P 64 2 2"
	sg["P6322"] = "P 63 2 2"
	sg["P6mm"] = "P 6 m m"
	sg["P6cc"] = "P 6 c c"
	sg["P63cm"] = "P 63 c m"
	sg["P63mc"] = "P 63 m c"
	sg["P-6m2"] = "P -6 m 2"
	sg["P-6c2"] = "P -6 c 2"
	sg["P-62m"] = "P -6 2 m"
	sg["P-62c"] = "P -6 2 c"
	sg["P6/mmm"] = "P 6 / m m m"
	sg["P6/mcc"] = "P 6 / m c c"
	sg["P63/mcm"] = "P 63 / m c m"
	sg["P63/mmc"] = "P 63 / m m c"

	// Cubic
	sg["P23"] = "P 2 3"
	sg["F23"] = "F 2 3"
	sg["I23"] = "I 2 3"
	sg["P213"] = "P 21 3"
	sg["I213"] = "I 21 3"
	sg["Pm-3"] = "P m -3"
	sg["Pn-3"] = "P n -3"
	sg["Fm-3"] = "F m -3"
	sg["Fd-3"] = "F d -3"
	sg["Im-3"] = "I m -3"
	sg["Pa-3"] = "P a -3"
	sg["Ia-3"] = "I a -3"
	sg["P432"] = "P 4 3 2"
	sg["P4232"] = "P 42 3 2"
	sg["F432"] = "F 4 3 2"
	sg["F4132"] = "F 41 3 2"
	sg["I432"] = "I 4 3 2"
	sg["P4332"] = "P 43 3 2"
	sg["P4132"] = "P 41 3 2"
	sg["I4132"] = "I 41 3 2"
	sg["P-43m"] = "P -4 3 m"
	sg["F-43m"] = "F -4 3 m"
	sg["I-43m"] = "I -4 3 m"
	sg["P-43n"] = "P -4 3 n"
	sg["F-43c"] = "F -4 3 c"
	sg["I-43d"] = "I -4 3 d"
	sg["Pm-3m"] = "P m -3 m"
	sg["Pn-3n"] = "P n -3 n"
	sg["Pm-3n"] = "P m -3 n"
	sg["Pn-3m"] = "P n -3 m"
	sg["Fm-3m"] = "F m -3 m"
	sg["Fm-3c"] = "F m -3 c"
	sg["Fd-3m"] = "F d -3 m"
	sg["Fd-3c"] = "F d -3 c"
	sg["Im-3m"] = "I m -3 m"
	sg["Ia-3d"] = "I a -3 d"

	return sg[fingerprint]

}
