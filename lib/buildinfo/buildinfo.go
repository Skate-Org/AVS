package buildinfo

import (
	"context"
	"log"
	"runtime/debug"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// Instrument logs the version, git commit hash, and timestamp from the runtime build info.
func Instrument(ctx context.Context, version string) {
	_, commit, timestamp := getBuildInfo()

	versionGauge.WithLabelValues(version).Set(1)
	commitGauge.WithLabelValues(commit).Set(1)

	ts, _ := time.Parse(time.RFC3339, timestamp)
	timestampGauge.Set(float64(ts.Unix()))
}

func BuildInfoCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information of this binary",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			goversion, commit, timestamp := getBuildInfo()

			var sb strings.Builder
			_, _ = sb.WriteString("Package version: " + version)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Go version:      " + goversion)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Git Commit:      " + commit)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Git Timestamp:   " + timestamp)
			_, _ = sb.WriteString("\n")

			cmd.Printf(sb.String())
		},
	}
}

// getBuildInfo returns the git commit hash and timestamp from the runtime build info.
func getBuildInfo() (goversion, hash string, timestamp string) {
	goversion, hash, timestamp = "█████", "█████", "█████"
	hashLen := 7

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return goversion, hash, timestamp
	}
	goversion = info.GoVersion

	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			hash = s.Value
			if len(hash) > hashLen {
				hash = hash[:hashLen]
			}
		case "vcs.time":
			timestamp = s.Value
			_, err := time.Parse(time.RFC3339, timestamp)
			if err != nil {
				log.Printf("Failed to parse timestamp: %v", err)
			}
		}
	}

	return goversion, hash, timestamp
}
