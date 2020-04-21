package primitivedatatype

func play_with_pointer() {

	var firstName *string = new(string)
	*firstName = "Arthur"
	println(*firstName)

}
