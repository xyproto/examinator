package examinator

import (
	"github.com/xyproto/distrodetector"
	"testing"
)

func TestFindPackageByName(t *testing.T) {
	distro := distrodetector.New()
	if distro.Name() == "Arch Linux" {
		packages := FindPackageByName(distro, "pacman")
		if !has(packages, "pacman") {
			t.Fatal("pacman should be a package in Arch Linux")
		}
	}
}
