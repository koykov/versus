package cbytebuf

// Common testing data.

import "bytes"

var (
	source = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque euismod ante non arcu " +
		"commodo tempor. Praesent quis nulla sed urna dictum iaculis. Pellentesque malesuada lacinia leo, eu " +
		"hendrerit tellus sodales sit amet. Sed ut finibus purus, ac lacinia metus. Nam tortor nunc, gravida " +
		"hendrerit posuere eu, tristique id elit. Proin id blandit purus. Donec aliquam quam nec erat sodales, eu " +
		"aliquet elit vestibulum. Morbi cursus vehicula semper. Sed dolor lorem, mattis et erat a, elementum " +
		"tincidunt purus. Integer sit amet porta mauris. Curabitur eu est sed augue rutrum tristique et a augue. " +
		"Proin dictum cursus quam vel varius. Duis viverra massa sed lacus gravida, a ullamcorper ipsum iaculis. " +
		"Maecenas interdum congue neque, in ultricies erat ornare id. Suspendisse vitae imperdiet eros.")
	space        = []byte(" ")
	expected     = append(source, space...)
	expectedLong = bytes.Repeat(source, 1000)
	parts        = bytes.Split(source, space)
)
