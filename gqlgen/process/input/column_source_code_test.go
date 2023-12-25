package input

import (
	"testing"
)

func TestToSourceCommitRow(t *testing.T) {
	// commit is fine
	// tooltip is fine
}

func TestToSourceCommitRowError(t *testing.T) {
	// commit is missing
	// commit is invalid
	// tooltip is missing line number
}

func TestToSourceOpenRow(t *testing.T) {
	// filepath is fine
	// tooltip is fine
}

func TestToSourceOpenRowError(t *testing.T) {
	// filepath is missing
	// filepath is invalid
	// tooltip is missing line number
}

func TestToSourceErrorRow(t *testing.T) {
	// filepath and tooltip are fine
}

func TestToSourceErrorRowError(t *testing.T) {
	// filepath is missing
	// filepath is invalid
	// tooltip is missing line number
	// tooltip is missing - source error needs tooltip
}
