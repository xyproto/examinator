package examinator

import (
	"github.com/xyproto/distrodetector"
	"os/exec"
	"strings"
)

// which returns the full path to the given executable, or the original string
func which(executable string) string {
	path, err := exec.LookPath(executable)
	if err != nil {
		return executable
	}
	return path
}

// run a shell command and return the output, or an empty string
func run(shellCommand string) string {
	cmd := exec.Command("sh", "-c", shellCommand)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(stdoutStderr)
}

func has(sl []string, e string) bool {
	for _, s := range sl {
		if e == s {
			return true
		}
	}
	return false
}

func FindPackageByName(d *distrodetector.Distro, name string) []string {
	packages := []string{}
	switch d.Name() {
	case "Arch Linux":
		lines := strings.Split(run("pacman -Ss "+name), "\n")
		for _, line := range lines {
			if !strings.HasPrefix(line, " ") && strings.Contains(line, " ") {
				fields := strings.Fields(line)
				if len(fields) > 0 {
					packageName := fields[0]
					if strings.Contains(packageName, "/") {
						parts := strings.SplitN(packageName, "/", 2)
						packageName = parts[1]
					}
					packages = append(packages, packageName)
				}
			}
		}
		return packages
	// TODO: Fill in the rest
	default:
		return packages
	}
}
