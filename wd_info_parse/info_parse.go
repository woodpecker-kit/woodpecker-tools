package wd_info_parse

import (
	"fmt"
	gitUrls "github.com/chainguard-dev/git-urls"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"strings"
)

func ParseRepositoryInfoByWoodPeckerInfo(info *wd_info.WoodpeckerInfo) error {

	if info.RepositoryInfo.CIRepoCloneURL == "" {
		return fmt.Errorf("RepositoryInfo.CIRepoCloneURL is empty")
	}

	parse, errParse := gitUrls.Parse(info.RepositoryInfo.CIRepoCloneURL)
	if errParse != nil {
		return fmt.Errorf("RepositoryInfo.CIRepoCloneURL parse err: %s", errParse)
	}
	if strings.Contains(parse.Scheme, "http") {
		info.RepositoryInfo.CiRepoHost = parse.Host
		info.RepositoryInfo.CiRepoHostname = parse.Hostname()
		info.RepositoryInfo.CiRepoPort = parse.Port()
	}

	return nil
}

func ParseRepositoryInfoByRepositoryInfo(info *wd_info.RepositoryInfo) error {

	if info.CIRepoCloneURL == "" {
		return fmt.Errorf("RepositoryInfo.CIRepoCloneURL is empty")
	}

	parse, errParse := gitUrls.Parse(info.CIRepoCloneURL)
	if errParse != nil {
		return fmt.Errorf("RepositoryInfo.CIRepoCloneURL parse err: %s", errParse)
	}
	if strings.Contains(parse.Scheme, "http") {
		info.CiRepoHost = parse.Host
		info.CiRepoHostname = parse.Hostname()
		info.CiRepoPort = parse.Port()
	}

	return nil
}
