// Code generated by "stringer -type EMouseCursor -trimprefix EMouseCursor_"; DO NOT EDIT.

package internal

import "strconv"

const _EMouseCursor_name = "usernonearrowibeamhourglasswaitarrowcrosshairupsizenwsizesesizenesizeswsizewsizeesizensizessizewesizenssizeallnohandblankmiddle_pannorth_pannorth_east_panast_pansouth_east_pansouth_pansouth_west_panwest_pannorth_west_panaliascellcolresizecopycurverticaltextrowresizezoominzoomouthelpcustomlast"

var _EMouseCursor_index = [...]uint16{0, 4, 8, 13, 18, 27, 36, 45, 47, 53, 59, 65, 71, 76, 81, 86, 91, 97, 103, 110, 112, 116, 121, 131, 140, 154, 161, 175, 184, 198, 206, 220, 225, 229, 238, 245, 257, 266, 272, 279, 283, 289, 293}

func (i EMouseCursor) String() string {
	if i < 0 || i >= EMouseCursor(len(_EMouseCursor_index)-1) {
		return "EMouseCursor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EMouseCursor_name[_EMouseCursor_index[i]:_EMouseCursor_index[i+1]]
}