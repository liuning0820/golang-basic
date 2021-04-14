package datatype

func Pointer() string {

	var firstName *string = new(string)
	*firstName = "Arthur"
	println(*firstName)
	return *firstName

}
