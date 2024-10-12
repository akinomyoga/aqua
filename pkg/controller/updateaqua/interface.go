package updateaqua

import (
	"context"

	"github.com/aquaproj/aqua/v2/pkg/github"
	"github.com/sirupsen/logrus"
)

type AquaInstaller interface {
	InstallAqua(ctx context.Context, logE *logrus.Entry, version string) error
}

type RepositoriesService interface {
	GetLatestRelease(ctx context.Context, logE *logrus.Entry, repoOwner, repoName string) (*github.RepositoryRelease, *github.Response, error)
}

type ConfigFinder interface {
	Finds(wd, configFilePath string) []string
}
