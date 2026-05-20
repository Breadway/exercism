package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var output = ""
	if name != "" {
		output = "One for " + name + ", one for me."
	} else {
		output = "One for you, one for me."
	}
	return output
}
