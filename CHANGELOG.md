# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.1.0](https://github.com/woodpecker-kit/woodpecker-tools/compare/1.0.2...v1.1.0) (2024-02-22)

### BREAKING CHANGE:

* wd_urfave_cli_v2/exit_cli/exit.go -> wd_urfave_cli_v2/cli_exit_urfave/exit.go

### ✨ Features

* wd_info.CiSystemVersionMinimumSupport and wd_info.CiSystemVersionConstraint can check version ([a05a6f47](https://github.com/woodpecker-kit/woodpecker-tools/commit/a05a6f47e1c8dd1a9cf2c33d3da0b8f2bb25bd05)), fe [#18](https://github.com/woodpecker-kit/woodpecker-tools/issues/18)

* add wd_info.CiSystemVersionCheck for check woodpecker-tools version support ([607c495a](https://github.com/woodpecker-kit/woodpecker-tools/commit/607c495ad58d3c01d60e89eb563e3a92cebe2348))

* let wd_log can show method info and gid by runtimeStack ([1b3b1f73](https://github.com/woodpecker-kit/woodpecker-tools/commit/1b3b1f738fb2882ddbb871674e4da420fb43b699)), fe [#15](https://github.com/woodpecker-kit/woodpecker-tools/issues/15)

* add flag CI_COMMIT_PRERELEASE and load by cli or mock by unit test ([9f290452](https://github.com/woodpecker-kit/woodpecker-tools/commit/9f290452119eb06dacb1867675881cb10c957216)), fe [#11](https://github.com/woodpecker-kit/woodpecker-tools/issues/11)

* add wd_steps_transfer and test case ([73d2ee5d](https://github.com/woodpecker-kit/woodpecker-tools/commit/73d2ee5d875c637b71fb271ee466e0699152cce6)), fe [#12](https://github.com/woodpecker-kit/woodpecker-tools/issues/12)

* add env_transfer AddEnvByKey RemoveEnvByKey SaveEnv2File ([3577ac4d](https://github.com/woodpecker-kit/woodpecker-tools/commit/3577ac4d60b7c68c58e6d7b9811eefab339ff001))

* add env_transfer AddEnvByKey RemoveEnvByKey SaveEnv2File ([a08be278](https://github.com/woodpecker-kit/woodpecker-tools/commit/a08be2781d98fea797db7d1961d5a1b2b13a18b3))

* add env_transfer package to writer or read as file and load to env ([dce1b994](https://github.com/woodpecker-kit/woodpecker-tools/commit/dce1b99407e00aa48a67fa6a11bddad17ec0ed21)), fe [#9](https://github.com/woodpecker-kit/woodpecker-tools/issues/9)

* add env_transfer package to writer or read as file and load to env ([5dbf22c2](https://github.com/woodpecker-kit/woodpecker-tools/commit/5dbf22c29b72da5a68d31a190ad782ebe93f63b6)), fe [#9](https://github.com/woodpecker-kit/woodpecker-tools/issues/9)

### 👷‍ Build System

* bump codecov/codecov-action from 3.1.4 to 4.0.0 ([bacccfd0](https://github.com/woodpecker-kit/woodpecker-tools/commit/bacccfd0f85c6b2e351e87ae00ba9fa1b7d5924a))

* bump actions/setup-go from 4 to 5 ([25f1a82d](https://github.com/woodpecker-kit/woodpecker-tools/commit/25f1a82daf55e8aebb73956990d03ec1fb0236ca))

* bump actions/checkout from 3 to 4 ([b1515299](https://github.com/woodpecker-kit/woodpecker-tools/commit/b1515299aaeb6277e4a828e11f1b1d28b2bebed9))

## [1.0.2](https://github.com/woodpecker-kit/woodpecker-tools/compare/1.0.1...v1.0.2) (2024-01-22)

### 👷‍ Build System

* check relesae by actions/upload-artifact/tree/v4 ([0f7d883b](https://github.com/woodpecker-kit/woodpecker-tools/commit/0f7d883b05ce14db58717455344445f3e0b16ebe))

* change deploy-tag.yml to support dry_run flag ([89e2f0b5](https://github.com/woodpecker-kit/woodpecker-tools/commit/89e2f0b5c802b4d02c45e8e9c485382f73ed3a57))

* support actions/upload-artifact/tree/v4 overwrite upload ([b36b8f39](https://github.com/woodpecker-kit/woodpecker-tools/commit/b36b8f3948f3ec84570af38e3171f16c8d1df898))

## [1.0.1](https://github.com/woodpecker-kit/woodpecker-tools/compare/1.0.0...v1.0.1) (2024-01-22)

## 1.0.0 (2024-01-21)

### 🐛 Bug Fixes

* mockEnvByStruct can check input to support ptr, but not suppot double ptr ([5f1651a1](https://github.com/woodpecker-kit/woodpecker-tools/commit/5f1651a163846fb429b2ca3a855967f1f0b79ee3))

* fix currentCommitFlag use wrong type to parse ([e79e01b2](https://github.com/woodpecker-kit/woodpecker-tools/commit/e79e01b2dfabc929f8e78363ed319baf8b43ba43))

### ✨ Features

* add WoodpeckerInfoSupportVersion show version support ([73bcd73e](https://github.com/woodpecker-kit/woodpecker-tools/commit/73bcd73e9662b069a9fe33d6548e31b7b83d2386))

* add env_mock kit for sample mock env ([370d318c](https://github.com/woodpecker-kit/woodpecker-tools/commit/370d318c058e20bd526c5d69fe26319b0f64ed26))

* add wd_mock kit for woodpecker plugin develop ([72e73bf0](https://github.com/woodpecker-kit/woodpecker-tools/commit/72e73bf010b0f21401f97379db6db71327f4b630))

* basic flag for woodpecker tools support urfave/cli/v2 to dev ([ab3b9a49](https://github.com/woodpecker-kit/woodpecker-tools/commit/ab3b9a49c42b2d065ae4f2731c0b465146da67b3))

### ♻ Refactor

* let MockEnvByStruct less Tag Lookup when type not String ([64244dd3](https://github.com/woodpecker-kit/woodpecker-tools/commit/64244dd38e51fdf1adf54ac4a92d313d430b8abe))

### 👷‍ Build System

* bump actions/upload-artifact from 3 to 4 (#3) ([3b7d6977](https://github.com/woodpecker-kit/woodpecker-tools/commit/3b7d697734e29396d2f872e458dc7fdbeb576089))

* bump actions/setup-go from 4 to 5 (#2) ([ea37c10d](https://github.com/woodpecker-kit/woodpecker-tools/commit/ea37c10d84b899cbb0ac721d003739d2d5e636e0))

* bump actions/download-artifact from 3 to 4 (#1) ([5586c6e1](https://github.com/woodpecker-kit/woodpecker-tools/commit/5586c6e11e4d234591589f362f4011a27ec5089d))
