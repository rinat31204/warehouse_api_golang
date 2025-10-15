package entities

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

var branch = Branch{
	uuid.New(),
	"test name",
	uuid.New(),
	time.Now(),
}

func Test_BranchName_MustNotBeNullOrEmpty(test *testing.T) {
	if branch.name == "" {
		test.Errorf("Branch name must not be empty")
	}
}

func Test_BranchId_MustNotBeNullOrEmpty(test *testing.T) {
	if branch.id == uuid.Nil {
		test.Errorf("Branch id must not be empty")
	}
}

func Test_BranchUserId_MustNotBeNullOrEmpty(test *testing.T) {
	if branch.userId == uuid.Nil {
		test.Errorf("Branch userId must not be empty")
	}
}
