package pixel

type Row uint64

func (row Row) Focus(pos int) Row {
	pix := Pixel(pos)
	return row & pix
}

func Pixel(pos int) Row {
	if pos < 0 {
		return 0
	} else {
		return 1 << pos
	}
}

func Sprite(pos int) Row {
	row := Row(0)
	for i := -1; i < 2; i++ {
		row = row | Pixel(pos+i)
	}
	return row
}
