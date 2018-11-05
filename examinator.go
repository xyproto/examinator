package examinator

import (
	"github.com/xyproto/distrodetector"
	"os/exec"
)

// Which returns the full path to the given executable, or the original string
func Which(executable string) string {
	path, err := exec.LookPath(executable)
	if err != nil {
		return executable
	}
	return path
}

// Run a shell command and return the output, or an empty string
func Run(shellCommand string) string {
	cmd := exec.Command("sh", "-c", shellCommand)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(stdoutStderr)
}

func FindPackageByName(d *distrodetector.Distro, name string) []string {
	switch d.name {
	case "Arch Linux":
		return strings.Split(run("pacman -Ss "+name), "\n")
	// TODO: Fill in the rest
	default:
		return []string{}
	}
}
