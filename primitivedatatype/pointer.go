package primitivedatatype

func playWithPointer() {

	var firstName *string = new(string)
	*firstName = "Arthur"
	println(*firstName)

}
