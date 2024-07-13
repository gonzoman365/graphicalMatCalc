package main


func (i *Item) parseInput(s string) {

	if s == "Custom" {

		i.Size.getSizeDialog(w)

	} else {
		convSize := map[string]Size{
			"18 x 24":  {18, 24},
			"13 x 19":  {13, 19},
			"11 x 17":  {11, 17},
			"11 x 14":  {11, 14},
			"10 x 12":  {10, 12},
			"8.5 x 11": {8.5, 11},
			"8 x 10":   {8, 10},
		}
		i.Size = (convSize[s])
		refreshStatus()
	}
}

func (o *Offset) parseOffset(s string) {
	convOffset := map[string]float64{
		"0.5":   0.5,
		"0.25":  0.25,
		"0":     0,
		"-0.25": -.25,
		"-.5":   -.5,
	}

	o.value = convOffset[s]
}
