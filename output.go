package main

import "fmt"

func (textOut *Textout) calculate(fsize, psize, isize Item, offset, doffset Offset, dmat bool) {

	// Calculate the paper offset on the backing
	vpaper := (fsize.Size.heighth - psize.Size.heighth) / 2.0
	hpaper := (fsize.Size.width - psize.Size.width) / 2.0
	textOut.pout.Text = (fmt.Sprintf("Paper margin on backing is %.3f vertical, and %.3f horizontal.", vpaper, hpaper))
	textOut.pout.Refresh()
	// // calculate inner mat
	vinner := (fsize.Size.heighth-isize.Size.heighth)/2.0 + offset.value
	hinner := (fsize.Size.width-isize.Size.width)/2.0 + offset.value
	textOut.iout.Text = (fmt.Sprintf("Mat margin is  %.3f vertical, and %.3f horizontal.", vinner, hinner))
	textOut.iout.Refresh()
	// // if double calculate outer mat

	textOut.oout.Text = ("")
	if dmat {
		vouter := vinner - doffset.value
		houter := hinner - doffset.value
		textOut.oout.Text = (fmt.Sprintf("Outer mat margin is  %.3f vertical, and %.3f horizontal.", vouter, houter))
		textOut.oout.Refresh()
	}
}

func refreshStatus() {
	textout.sout.Text = fmt.Sprintf("Frame Size: %.3f x %.3f Paper Size: %.3f x %.3f Image Size: %.3f x %.3f",
		fsize.Size.heighth, fsize.Size.width,
		psize.Size.heighth, psize.Size.width,
		isize.Size.heighth, isize.Size.width)
	textout.sout.Refresh()

}



