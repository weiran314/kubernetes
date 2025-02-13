package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "32"
	gitVersion   = "v1.32.2-k3s1"
	gitCommit    = "865a265e0f265275e5e91c468720d3ad5400abfb"
	gitTreeState = "clean"
	buildDate = "2025-02-13T20:05:57Z"
)
