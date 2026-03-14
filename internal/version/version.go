package version

// Version is set at build time via -ldflags.
var Version = "v5.0.2"

// BinaryName identifies which binary was built (floyd or superfloyd).
// Set at build time via -ldflags "-X github.com/legacy-ai/floyd/internal/version.BinaryName=superfloyd"
var BinaryName = "floyd"

// BuildDate is set at build time via -ldflags.
var BuildDate = "unknown"
