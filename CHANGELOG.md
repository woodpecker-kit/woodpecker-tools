# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

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
