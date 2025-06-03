[![ci](https://github.com/woodpecker-kit/woodpecker-tools/actions/workflows/ci.yml/badge.svg)](https://github.com/woodpecker-kit/woodpecker-tools/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-tools?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-tools)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-tools?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-tools)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-tools)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-tools)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-tools)](https://github.com/woodpecker-kit/woodpecker-tools)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-tools/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-tools)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-tools)](https://github.com/woodpecker-kit/woodpecker-tools/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-tools)](https://github.com/woodpecker-kit/woodpecker-tools/releases)

## for what

- this project used to github golang lib project

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-tools)](https://github.com/woodpecker-kit/woodpecker-tools/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## depends

in go mod project

```bash
# before above global settings
# test version info
$ git ls-remote -q https://github.com/woodpecker-kit/woodpecker-tools.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/woodpecker-kit/woodpecker-tools
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/woodpecker-kit/woodpecker-tools | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [x] `wd_log` package support debug at line number and gid
- [x] `wd_flag` package is flag for woodpecker plugin
    - `wd_flag.SetTimeFormat(format string)` will change time format like `2006-01-02 15:04:05`, must
      before `wd_urfave_cli_v2.UrfaveCliBindInfo()`
- [x] `wd_info.WoodpeckerInfo` is plugin most use env args
  from [woodpecker-ci/woodpecker](https://github.com/woodpecker-ci/woodpecker)
    - Deprecated `CI_PIPELINE_FINISHED` and `CI_PIPELINE_FILES` for support woodpecker server 3.+ (v1.22+)
- [x] `wd_info.CiSystemVersionMinimumSupport` and `wd_info.CiSystemVersionConstraint` can check plugin support ci system
  version
- [x] `wd_urfave_cli_v2.WoodpeckerUrfaveCliFlags()` bind cli
  as [github.com/urfave/cli/v2](https://github.com/urfave/cli/)
- [x] `wd_info.WoodpeckerInfoSupportVersion` support version begin `2.0.0`
- [x] `env_mock.MockEnvByStruct` support struct tag `mock_env_key` or `mock_env_default` for unit test of plugin
    - `1.19.+` add `wd_mock.Fast*()` method for fast mock
    - `wd_mock.FastWorkSpace` for fast mock workspace
    - `wd_mock.FastCurrentStatus` for fast mock current status most use `wd_info.BuildStatusSuccess` or `wd_info.BuildStatusFailure`
    - `wd_mock.FastTag` for fast mock event `tag`
    - `wd_mock.FastPullRequest` for fast mock event `pull_request`
    - `wd_mock.FastPullRequestClose` for fast mock event `pull_request_closed`
    - `wd_mock.FastPushCommitBranch` for fast mock event `push` commit branch
- [x] `env_transfer.AddOrCoverEnvByKey` `env_transfer.RemoveEnvByKey` and `env_transfer.SaveEnv2File` for transfer env
  between steps
    - please add `.env.woodpecker_transfer.local` at git ignore, to use `env_transfer.DefaultWriterFileName` to transfer
      env
    - `env_transfer.AddOrCoverEnvByKey` or `env_transfer.RemoveEnvByKey` will append `env_transfer.PrefixTransfer` and
      upper case at runner env
    - load env can use `env_transfer.OverloadEnvFromFile`, most use file name as `env_transfer.DefaultWriterFileName`
- [x] `wd_steps_transfer.Out` and `wd_steps_transfer.In` for transfer data between steps with same workflow
    - please add `.woodpecker_kit.steps.transfer` at git ignore, to use `wd_steps_transfer.DefaultWriterFileName` to
      transfer data
- [x] `env_kit` package use [github.com/sinlov-go/unittest-kit](https://github.com/sinlov-go/unittest-kit)
    - `env_kit.FetchOsEnv*` and `env_kit.SetEnv*` for env get or set
    - `env_kit.FindAllEnv4Print`, `env_kit.FindAllEnvByPrefix`, `env_kit.FindAllEnv4PrintAsSortJust` for find print env
      string
- [x] `wd_template` for support Handlebars.js for golang
    - use `wd_template.RegisterSettings(DefaultFunctions)` once to register default functions then use
- [x] `wd_short_info.ParseWoodpeckerInfo2Short` can parse `wd_info.WoodpeckerInfo` to `wd_short_info.WoodpeckerInfoShot` (v1.18+)
  for template more clear
- [x] code check
    - [x] full check by golang version
    - [x] full check for docker build

## usage

### `wd_info.WoodpeckerInfo`

- see example of [app.go](https://github.com/woodpecker-kit/woodpecker-tools/blob/main/cmd/cli/app.go)

### `env_mock.MockEnvByStruct`

- see test case
  of [wd_env_mock.go](https://github.com/woodpecker-kit/woodpecker-tools/blob/main/wd_mock_test/wd_info_mock_test.go)

# dev

see [doc-dev/dev.md](doc-dev/dev.md)