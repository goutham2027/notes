package greeting

// exported
const (
	WelcomeText = "Hello, World to Golang"
	MorningText = "Good Morning"
	EveningText = "Good Evening"
)

// not exported, visible only within the greeting package
var loremIpsumText = `Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
commodo consequat.`
