package main

func podeComprarComida(stock map[string]int, dogs, cats int) bool {
	sD := stock["dog"]       //5
	sC := stock["cat"]       //5
	sU := stock["universal"] //0
	fomeD := dogs - sD       //-1
	fomeC := cats - sD       //1

	if fomeD > 0 && fomeC > 0 {
		if sU >= fomeD+fomeC {
			return true
		}
	} else if sD >= dogs && sC >= cats {
		return true
	} else if fomeD <= 0 && fomeC >= 0 && sU >= fomeC {
		return true
	} else if fomeC <= 0 && fomeD >= 0 && sU >= fomeC {
		return true
	}
	return false
}
