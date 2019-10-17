package piscine

func UltimateDivMod(a *int, b *int) {

	f := *a
	*a = *a / *b
	*b = f % *b

}
