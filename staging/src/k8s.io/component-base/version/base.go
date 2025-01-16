package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "31"
	gitVersion   = "v1.31.5-k3s1"
	gitCommit    = "1663a51d0c577a9c04721bf5017fab5c7f74b04a"
	gitTreeState = "clean"
	buildDate = "2025-01-16T02:39:43Z"
)
