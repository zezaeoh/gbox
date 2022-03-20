package common

var Version string

func GetVersion() string {
	if Version != "" {
		return Version
	} else {
		return "dev"
	}
}
