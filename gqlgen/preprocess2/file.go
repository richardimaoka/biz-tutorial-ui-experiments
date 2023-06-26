package preprocess2

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/google/uuid"
)

type File struct {
	updatedStepID uuid.UUID
	blobHash      plumbing.Hash //blobHash is equivalent to content, assuming every transition is git
	// ref        FileContentRef //when manual edit is allowed, blobHash is not enought to indicate the contnet
}

func (n *File) IsUpdated(currentStep uuid.UUID) bool {
	return currentStep == n.updatedStepID
}

func (n *File) Diff(from SourceCodeStep) string {
	return ""
}
