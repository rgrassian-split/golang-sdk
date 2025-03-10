# Changelog

## [0.5.0](https://github.com/open-feature/go-sdk/compare/v0.4.0...v0.5.0) (2022-09-30)


### ⚠ BREAKING CHANGES

* changed client details signatures to return new type (#84)
* spec v0.5.0 compliance (#82)
* defined type for provider interface evaluation context (#74)
* replaced EvaluationOptions with variadic option setter in client functions (#77)
* introduced context.Context to client and provider api (#75)

### Features

* changed client details signatures to return new type ([#84](https://github.com/open-feature/go-sdk/issues/84)) ([25ecdac](https://github.com/open-feature/go-sdk/commit/25ecdacb8303f95ec88656a7f47c8bd2ef0c019a))
* introduced context.Context to client and provider api ([#75](https://github.com/open-feature/go-sdk/issues/75)) ([d850c88](https://github.com/open-feature/go-sdk/commit/d850c8873d617aec7d1013aa1c751aa5bf0dce92))
* replaced EvaluationOptions with variadic option setter in client functions ([#77](https://github.com/open-feature/go-sdk/issues/77)) ([fc4b871](https://github.com/open-feature/go-sdk/commit/fc4b8716f6d3c904b464d34176d0c6ed67f741fc))
* spec v0.5.0 compliance ([#82](https://github.com/open-feature/go-sdk/issues/82)) ([69b8f8e](https://github.com/open-feature/go-sdk/commit/69b8f8e534ad0b99bf3de67cca531720f4bfc2de))


### Bug Fixes

* add reason indicating pseudorandom split ([#76](https://github.com/open-feature/go-sdk/issues/76)) ([e843f5d](https://github.com/open-feature/go-sdk/commit/e843f5d101041e6e3ba785168b8526fcf7f50c8e))


### Code Refactoring

* defined type for provider interface evaluation context ([#74](https://github.com/open-feature/go-sdk/issues/74)) ([69988c0](https://github.com/open-feature/go-sdk/commit/69988c097f16f3aaca9bdae07ea33fbce148872d))

## [0.4.0](https://github.com/open-feature/go-sdk/compare/v0.3.0...v0.4.0) (2022-09-20)


### ⚠ BREAKING CHANGES

* rename module to go-sdk (#66)

### Features

* rename module to go-sdk ([#66](https://github.com/open-feature/go-sdk/issues/66)) ([75a901a](https://github.com/open-feature/go-sdk/commit/75a901a330ab7517e4c92def5f7bf854391203d6))


### Bug Fixes

* ensure default client logger is updated when global logger changes ([#61](https://github.com/open-feature/go-sdk/issues/61)) ([f8e2827](https://github.com/open-feature/go-sdk/commit/f8e2827639d7e7f1206de933d4ed043489eadd7d))
* return error code from client given by provider ([#67](https://github.com/open-feature/go-sdk/issues/67)) ([f0822b6](https://github.com/open-feature/go-sdk/commit/f0822b6ce9522cbbb10ed5168cecad2df6c29e40))

## [0.3.0](https://github.com/open-feature/golang-sdk/compare/v0.2.0...v0.3.0) (2022-09-14)


### ⚠ BREAKING CHANGES

* remove duplicate Value field from ResolutionDetail structs (#58)

### Bug Fixes

* remove duplicate Value field from ResolutionDetail structs ([#58](https://github.com/open-feature/golang-sdk/issues/58)) ([945bd96](https://github.com/open-feature/golang-sdk/commit/945bd96c808246614ad5a8ab846b0b530ff313cc))

## [0.2.0](https://github.com/open-feature/golang-sdk/compare/v0.1.0...v0.2.0) (2022-09-02)


### ⚠ BREAKING CHANGES

* flatten evaluationContext object (#51)

### Features

* implemented structured logging ([#54](https://github.com/open-feature/golang-sdk/issues/54)) ([04649c5](https://github.com/open-feature/golang-sdk/commit/04649c5b954531601dc3e8a474bbff66094d3b1c))
* introduce UnimplementedHook to avoid authors having to define empty functions ([#55](https://github.com/open-feature/golang-sdk/issues/55)) ([0c0bd32](https://github.com/open-feature/golang-sdk/commit/0c0bd32894346babe8d180b086362e95fd3670ef))
* remove EvaluationOptions from FeatureProvider func signatures. ([91aaeb5](https://github.com/open-feature/golang-sdk/commit/91aaeb5893a79ae7ebc9949c7c59aa72b7651e09))


### Code Refactoring

* flatten evaluationContext object ([#51](https://github.com/open-feature/golang-sdk/issues/51)) ([b8383e1](https://github.com/open-feature/golang-sdk/commit/b8383e148184c1d8e58ff74217cffc99e713d29f))
