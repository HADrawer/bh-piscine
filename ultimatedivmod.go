package piscine

func UltimateDivMod(a *int, b *int) {
	var num1 int = *a
	var num2 int = *b

	*a = num1 / num2
	*b = num1 % num2
}
