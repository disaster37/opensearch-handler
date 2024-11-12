// A generated module for OpensearchHandler functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"

	"dagger/opensearch-handler/internal/dagger"

	"emperror.dev/errors"
	"github.com/disaster37/dagger-library-go/lib/helper"
)

const (
	mockgenVersion          = "v0.3.0"
	gitUsername      string = "ci"
	gitEmail         string = "ci@localhost"
	defaultGitBranch string = "release-branch.v2"
)

type OpensearchHandler struct {
	// Src is a directory that contains the projects source code
	// +private
	Src *dagger.Directory

	// The golang base image
	// +private
	BaseImage *dagger.Container

	// +private
	GolangModule *dagger.Golang
}

func New(
	ctx context.Context,
	// a path to a directory containing the source code
	// +required
	src *dagger.Directory,
) (*OpensearchHandler, error) {
	// Compute image because of base is not optional
	version, err := inspectModVersion(context.Background(), src)
	if err != nil {
		return nil, err
	}
	base := defaultImage(version)
	base = mountCaches(ctx, base).
		WithDirectory(goWorkDir, src).
		WithWorkdir(goWorkDir).
		WithoutEntrypoint()

	return &OpensearchHandler{
		Src:          src,
		GolangModule: dag.Golang(base, src),
		BaseImage:    base,
	}, nil
}

func (h *OpensearchHandler) Ci(
	ctx context.Context,

	// Set tru if you are on CI
	// +default=false
	ci bool,

	// The codeCov token
	// +optional
	codeCoveToken *dagger.Secret,

	// The git token
	// +optional
	gitToken *dagger.Secret,
) (dir *dagger.Directory, err error) {
	var stdout string

	// Build
	h.Build(ctx)

	// Lint code
	stdout, err = h.Lint(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Error when lint project: %s", stdout)
	}

	// Format code
	dir = h.Format(ctx)

	// Test code
	reportFile := h.Test(ctx)
	dir = dir.WithFile("coverage.out", reportFile)

	// Generate mock
	dir = dir.WithDirectory(".", h.GenerateMock(ctx))

	if ci {
		if codeCoveToken == nil {
			return nil, errors.New("You need to provide CodeCov token")
		}
		stdout, err = h.CodeCov(ctx, dir, codeCoveToken)
		if err != nil {
			return nil, errors.Wrapf(err, "Error when upload report on CodeCov: %s", stdout)
		}

		if _, err = dag.Git().SetConfig(gitUsername, gitEmail, dagger.GitSetConfigOpts{BaseRepoURL: "github.com", Token: gitToken}).SetRepo(dir, dagger.GitSetRepoOpts{Branch: defaultGitBranch}).CommitAndPush(ctx, "Commit from CI. skip ci"); err != nil {
			return nil, errors.Wrap(err, "Error when commit and push files change")
		}
	}

	return dir, nil
}

// Lint permit to lint code
func (h *OpensearchHandler) Lint(
	ctx context.Context,
) (string, error) {
	return h.GolangModule.Lint(ctx)
}

// Format permit to format the golang code
func (h *OpensearchHandler) Format(
	ctx context.Context,
) *dagger.Directory {
	return h.GolangModule.Format()
}

// Test permit to run tests
func (h *OpensearchHandler) Test(
	ctx context.Context,
) *dagger.File {
	return h.BaseImage.
		WithExec(helper.ForgeCommand("go test ./... -v -count 1 -parallel 1 -race -coverprofile=coverage.out -covermode=atomic -timeout 120m")).
		File("coverage.out")
}

// Build permit to build project
func (h *OpensearchHandler) Build(
	ctx context.Context,
) *dagger.Directory {
	return h.GolangModule.Build()
}

func (h *OpensearchHandler) CodeCov(
	ctx context.Context,

	// Optional directory
	// +optional
	src *dagger.Directory,

	// The Codecov token
	// +required
	token *dagger.Secret,
) (stdout string, err error) {
	if src == nil {
		src = h.Src
	}

	return dag.Codecov().Upload(
		ctx,
		src,
		token,
		dagger.CodecovUploadOpts{
			Files:   []string{"coverage.out"},
			Verbose: true,
		},
	)
}

// Test permit to run tests
func (h *OpensearchHandler) GenerateMock(
	ctx context.Context,
) *dagger.Directory {
	return h.BaseImage.WithExec(helper.ForgeScript(`
go install go.uber.org/mock/mockgen@%s
mockgen --build_flags=--mod=mod -destination=mocks/opensearch_handler.go -package=mocks github.com/disaster37/opensearch-handler/v2 OpensearchHandler
	`, mockgenVersion)).Directory(goWorkDir)
}
