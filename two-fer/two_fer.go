// Package twofer provides a tool for fpr provoding a message with a name
package twofer

// ShareWith returns a message with the name,
// else provides a message by replacing the name with "you" .
func ShareWith(name string) string {

	if name != "" {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}

}
