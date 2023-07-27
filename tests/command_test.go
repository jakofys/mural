package mural_test

import (
	"testing"

	"github.com/jakofys/mural"
)

func TestCreateCLI(t *testing.T) {
	cli, err := mural.NewCLI("myapp", "1.12.3-alpha+6f8ca4bfb2", "A simple application")
	if err != nil {
		t.Fatal(err)
	}
	if cli.Version().Major != 1 {
		t.Fatal("major version must be 1")
	}
	if cli.Version().Minor != 12 {
		t.Fatal("minor version must be 12")
	}
	if cli.Version().Patch != 3 {
		t.Fatal("patch version must be 3 found ", cli.Version().Patch)
	}
	if cli.Version().PreRelease != "alpha" {
		t.Fatal("release version must be alpha")
	}
	if cli.Version().Build != "6f8ca4bfb2" {
		t.Fatal("build version must be 6f8ca4bfb2")
	}
}
