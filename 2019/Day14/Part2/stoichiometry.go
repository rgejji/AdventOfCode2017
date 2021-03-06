package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type reaction struct {
	numProduced    int64
	reagentsNeeded []reagent
}

type reagent struct {
	numRequired int64
	name        string
}

var requirements map[string]reaction
var stash map[string]int64

func readInput() {
	rowSlice := strings.Split(inputStr, "\n")
	for _, row := range rowSlice {
		sides := strings.Split(row, " => ")
		lhsParts := strings.Split(sides[0], ", ")
		rhsPart := strings.Split(sides[1], " ")

		//construct reagents
		reagents := []reagent{}
		for _, p := range lhsParts {

			vals := strings.Split(p, " ")
			v, _ := strconv.Atoi(vals[0])
			r := reagent{
				numRequired: int64(v),
				name:        vals[1],
			}
			reagents = append(reagents, r)
		}
		//construct reaction
		v2, _ := strconv.Atoi(rhsPart[0])
		react := reaction{
			numProduced:    int64(v2),
			reagentsNeeded: reagents,
		}
		//Add requirement
		name := rhsPart[1]
		requirements[name] = react
		fmt.Printf("Added requirement: %s - %+v\n", name, react)
	}

	return
}

var oreCount int64

//Subtract what is needed to get new requirement, then return with >= num of requirement in stash
func aquireReagents(name string, num int64) {
	//fmt.Printf("Aquiring %d %s\n", num, name)
	//If we have it already, spend it
	if stash[name] >= num {
		return
	}

	//If its ore, make more ore
	if name == "ORE" {
		oreNeeded := num - stash[name]
		oreCount += oreNeeded
		stash[name] += oreNeeded
		return
	}

	//If not, resolve requirements
	//While we don't have enough of the current item we want, make more
	req := requirements[name]
	for stash[name] < num {
		//Perform reaction one time
		//Aquire base components and spend them
		for _, p := range req.reagentsNeeded {
			aquireReagents(p.name, p.numRequired)
			stash[p.name] -= p.numRequired
		}
		//Collect result
		stash[name] += req.numProduced
	}
	return
}

type double struct {
	numOre int64
	i      int64
}

const fuel string = "FUEL"

func main() {
	//patterns := make(map[string]double)
	requirements = make(map[string]reaction)
	readInput()
	stash = make(map[string]int64)
	aquireReagents(fuel, 1)
	stash[fuel] = 0

	itemList := []string{}
	for key := range stash {
		if key == fuel {
			continue
		}
		itemList = append(itemList, key)
	}

	i := int64(1)
	for oreCount < int64(1000000000000) {

		/*
			//attempt to look for repititions
			hash := ""
			for _, key := range itemList {
				hash += fmt.Sprintf("%s-%d", key, stash[key])
			}

			if v, ok := patterns[hash]; ok {
				fmt.Printf("Found a cycle starting at (i,ore)=(%d,%d) and ending (%d,%d)\n", v.i, v.numOre, i, oreCount)
				break
			}
			patterns[hash] = double{numOre: oreCount, i: i}
		*/

		if i%1000 == 0 {
			fmt.Printf("%v: i=%d %d\n", time.Now(), i, oreCount)
			//fmt.Printf("%d HASH IS %s\n", i, hash)
		}
		//Here we aquire fuel 1 at a time and run over night
		//An alternate solution would be to track stash state needed to get 1K,10K,100K,... fuel, and alternate between scaling the stash state for fast calculation, followed by a relaxation period of regular aquisition to use up the scaled trash state.
		aquireReagents(fuel, 1)
		stash[fuel] = 0
		i++
	}

	fmt.Printf("Ended with failure to find fuel for i = %d\n", i)
	fmt.Printf("Answer should be i=%d\n", i-1)

}

/*const inputStr = `2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`
*/
/*
const inputStr = `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`
*/
/*const inputStr = `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`
*/

const inputStr = `6 GTGRP, 1 VPGRV, 1 KGQR => 6 HSHQR
1 RZXL => 1 MJTV
2 MJTV, 1 NZFM => 6 MGVLC
6 PFWG, 2 NVQG => 5 DCQP
6 MQDF, 1 NTHXM, 10 NZFM => 3 JRKQ
13 KFZXS => 8 MQDF
2 CMBFH => 9 KCXVQ
13 QVTVR, 4 HXTVZ, 2 TGFZK => 9 FCLQJ
2 ZBXVW => 5 WQVSD
20 DXKGN, 10 LWNB, 1 KCGRN, 1 QLZWT, 2 CTKD => 3 LSWQ
10 TGFZK => 8 CMBFH
149 ORE => 4 NTHXM
145 ORE => 5 ZVCW
1 LSFHG => 4 PFWG
1 NTHXM, 1 THSD => 6 LSFHG
1 KFZXS, 4 VTMK => 4 LWNB
20 HXTVZ, 1 LWNB => 7 QNDT
3 FHVXH, 6 NBGZ => 8 MLBKD
9 MQDF, 1 VJLNZ => 9 FHVXH
2 CWLD => 3 HLXNV
7 PFWG, 1 NCRG => 6 JLPG
2 XCTGC, 10 VZDF, 5 JRKQ => 8 NVQG
2 MJCR => 7 VPGRV
18 XTNK, 1 THSD => 3 VJLNZ
3 CWLD => 3 NMKZN
3 LSFHG, 1 PFWG, 6 DXKGN => 1 WVLN
12 NMKZN => 8 VZDF
1 MJTV => 5 NZFM
31 MGVLC => 5 THSD
11 PFWG => 8 JTHQ
2 KGQR, 1 TGFZK, 2 FPZHG => 4 XTXKL
30 GTGRP => 3 NBGZ
17 NVQG => 8 HDWSV
1 THSD, 18 XTNK => 2 FPZHG
5 QNDT, 13 WDGM, 13 NTHXM, 10 NBGZ, 14 GTGRP, 14 KFWM, 3 HDWSV, 5 LSWQ => 1 FUEL
3 VJLNZ, 5 VTMK => 9 DXKGN
1 LWNB, 2 HSHQR, 10 WQVSD => 9 QLZWT
42 VZDF, 3 RZXL, 1 NTHXM => 7 XTNK
3 WVLN => 7 NCRG
14 NZFM => 8 XCTGC
4 NVQG, 2 LSFHG => 7 KGQR
26 HSHQR, 3 BVMKL => 2 QVTVR
1 VJLNZ, 7 XTNK => 1 KCGRN
167 ORE => 3 KCLR
2 ZVCW, 3 RZXL, 1 KCLR => 9 CWLD
5 FCLQJ, 19 MLBKD, 4 SRPRL, 5 RMQRL, 11 WQVSD, 3 QLZWT => 6 KFWM
3 KCLR, 1 VZDF => 5 TGFZK
17 NVQG, 1 VPGRV => 5 BVMKL
4 WQVSD => 4 RMQRL
4 KCGRN, 4 DCQP => 4 SRPRL
2 KCGRN => 4 CTKD
1 HLXNV, 1 KFZXS => 7 MJCR
116 ORE => 6 RZXL
181 ORE => 7 KFZXS
1 FHVXH, 1 NVQG => 5 GTGRP
5 JTHQ, 8 FCLQJ, 1 XTXKL, 1 QVTVR, 1 WQVSD, 10 JLPG => 3 WDGM
1 NZFM, 1 RZXL, 17 MGVLC => 4 VTMK
1 KFZXS => 7 ZBXVW
7 KCXVQ, 29 BVMKL => 6 HXTVZ`
