# Changelog

All notable changes to this project will be documented in this file.

## [unreleased]

### Documentation

- *(changelog)* Regenerate CHANGELOG.md [skip ci]

### Testing

- Remove internal/awx from ignore

## [24.6.1-2] - 2026-04-29

### Bug Fixes

- When host is set to http but it redirects to https it breaks the redirect flow

### CI/CD

- *(workflows)* Introduced changelog generation for automated release notes

### Dependencies

- *(deps)* Updated dependencies in package.json and package-lock.json
- *(deps)* Bump the actions group with 3 updates (#178)

### Documentation

- *(resources)* Documented AWX AWS credential resource details
- *(changelog)* Regenerate CHANGELOG.md [skip ci]
- *(changelog)* Regenerate CHANGELOG.md [skip ci]
- *(changelog)* Regenerate CHANGELOG.md [skip ci]
- *(changelog)* Regenerate CHANGELOG.md [skip ci]

### Features

- *(testdata)* Introduced AWX settings resources for comprehensive tests

### Miscellaneous

- Added .gitignore entry for temporary build files
- *(config)* Organized git-cliff config for changelog generation
- Updated .gitignore to exclude temporary build files
- Updated versions.yaml

### Refactor

- Cleanup and reorganize the code and add credential type
- *(generator)* Consolidate model/resource/data source templates into tf_object.go.tpl

### Testing

- Updated VCR test example to include request latency simulation
- *(testdata/cassettes)* Added awx resources and tests to testdata
- *(testdata/cassettes)* Added instance group integration tests
- *(testdata)* Added integration test for job template survey spec

## [24.6.1-1] - 2026-04-26

### Dependencies

- *(deps)* Bump golangci/golangci-lint-action from 6 to 7  (#142)
- *(deps)* Bump goreleaser/goreleaser-action from 6.2.1 to 6.3.0  (#143)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-testing  (#141)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework-validators  (#146)
- *(deps)* Bump golangci/golangci-lint-action from 7 to 8  (#145)
- *(deps)* Bump github/codeql-action from 3 to 4  (#158)
- *(deps)* Bump actions/setup-go from 5 to 6  (#157)
- *(deps)* Bump actions/stale from 9 to 10  (#156)
- *(deps)* Bump github.com/stretchr/testify from 1.10.0 to 1.11.1  (#154)
- *(deps)* Bump the terraform group across 1 directory with 4 updates  (#159)
- *(deps)* Bump github.com/spf13/cobra from 1.9.1 to 1.10.1  (#155)
- *(deps)* Bump actions/checkout from 4 to 6  (#163)
- *(deps)* Bump golangci/golangci-lint-action from 8 to 9  (#161)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-log  (#162)
- *(deps)* Bump the terraform group with 2 updates  (#164)
- *(deps)* Bump github.com/spf13/cobra from 1.10.1 to 1.10.2  (#165)
- *(deps)* Bump the actions group with 7 updates (#177)
- *(deps)* Bump the terraform group across 1 directory with 3 updates (#171)

### Documentation

- Fix typo in local testing section
- Clarified usage instructions in README
- Removed outdated TODO section from README.md

### Features

- *(testdata)* Added awx_job_template_survey_spec resource
- *(resources/api)* Introduced json support for config properties
- *(config)* Enabled omit_empty for conditional serialization in JSON fields

### Miscellaneous

- Remove deprecated.md from root
- Bump go to 1.24
- Gofmt
- Fix .goreleaser deprecation warnings
- Removed obsolete entries from versions.yaml
- Updated versions.yaml

### Refactor

- Consolidate generator templates into internal/framework package  (#175)

### Testing

- Replace context.Background() with t.Context() in tests
- Migrate golang-ci to version 2

## [24.6.1-0] - 2025-02-28

### Miscellaneous

- Generated version 24.6.1 with tag v24.6.1-0

## [24.2.0-3] - 2025-02-28

### Dependencies

- *(deps)* Bump github.com/spf13/cobra from 1.8.1 to 1.9.1  (#136)
- *(deps)* Bump the terraform group with 3 updates  (#140)

### Miscellaneous

- Updated versions.yaml
- Generated version 24.2.0 with tag v24.2.0-3

### Refactor

- Code generation, improve template and simplify it  (#132)

## [24.2.0-2] - 2025-02-24

### Miscellaneous

- Generated version 24.2.0 with tag v24.2.0-2

## [24.1.0-4] - 2025-02-24

### Miscellaneous

- Generated version 24.1.0 with tag v24.1.0-4

## [24.0.0-4] - 2025-02-24

### Miscellaneous

- Generated version 24.0.0 with tag v24.0.0-4

## [23.9.0-4] - 2025-02-24

### Miscellaneous

- Generated version 23.9.0 with tag v23.9.0-4

## [23.8.1-5] - 2025-02-24

### Miscellaneous

- Generated version 23.8.1 with tag v23.8.1-5

## [23.5.1-5] - 2025-02-24

### Bug Fixes

- Inventory group variable pre wrap

### Build

- Fix goreleaser version to ~> 2

### Dependencies

- *(deps)* Bump goreleaser/goreleaser-action from 6.1.0 to 6.2.1  (#135)

### Miscellaneous

- Update config.json for all versions
- Updated versions.yaml
- Generated version 23.5.1 with tag v23.5.1-5

## [24.2.0-1] - 2024-12-19

### Miscellaneous

- Generated version 24.2.0 with tag v24.2.0-1

## [24.1.0-3] - 2024-12-19

### Miscellaneous

- Generated version 24.1.0 with tag v24.1.0-3

## [24.0.0-3] - 2024-12-19

### Miscellaneous

- Generated version 24.0.0 with tag v24.0.0-3

## [23.9.0-3] - 2024-12-19

### Miscellaneous

- Generated version 23.9.0 with tag v23.9.0-3

## [23.8.1-4] - 2024-12-19

### Miscellaneous

- Generated version 23.8.1 with tag v23.8.1-4

## [23.5.1-4] - 2024-12-19

### Bug Fixes

- Added missing secrets validations for OIDC and Notification Template  (#122)

### Build

- Fix goreleaser deprecated --rm-dist flag

### Dependencies

- *(deps)* Bump golangci/golangci-lint-action from 4 to 6  (#101)
- *(deps)* Bump github.com/spf13/cobra from 1.8.0 to 1.8.1  (#106)
- *(deps)* Bump goreleaser/goreleaser-action from 5.0.0 to 6.0.0  (#105)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework-validators  (#107)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-testing  (#109)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-testing  (#114)
- *(deps)* Bump goreleaser/goreleaser-action from 6.0.0 to 6.1.0  (#124)
- *(deps)* Bump codecov/codecov-action from 4 to 5  (#125)
- *(deps)* Bump github.com/stretchr/testify from 1.9.0 to 1.10.0  (#128)

### Miscellaneous

- Added workflow for stale issues and pr
- Updated versions.yaml
- Updated versions.yaml
- Updated versions.yaml
- Generated version 23.5.1 with tag v23.5.1-4

### Refactor

- Simplify templating  (#123)

## [24.2.0-0] - 2024-04-15

### Build

- We can now specify which versions to be built for the provider

### Features

- Added AWX API 24.2.0

### Miscellaneous

- Updated versions.yaml

## [24.1.0-2] - 2024-04-13

### Miscellaneous

- Generated version 24.1.0 with tag v24.1.0-2

## [24.0.0-2] - 2024-04-13

### Miscellaneous

- Generated version 24.0.0 with tag v24.0.0-2

## [23.9.0-2] - 2024-04-13

### Miscellaneous

- Generated version 23.9.0 with tag v23.9.0-2

## [23.8.1-3] - 2024-04-13

### Miscellaneous

- Generated version 23.8.1 with tag v23.8.1-3

## [23.7.0-3] - 2024-04-13

### Miscellaneous

- Generated version 23.7.0 with tag v23.7.0-3

## [23.6.0-3] - 2024-04-13

### Miscellaneous

- Generated version 23.6.0 with tag v23.6.0-3

## [23.5.1-3] - 2024-04-13

### Bug Fixes

- Missing write-only from data sources, causing credential data look up to break
- Removed some properties that should not be in SettingsMiscSystem

### Build

- Fix the build command so it includes resources api as part of the autobuild

### Miscellaneous

- Cleanup removing plan files
- Update Makefile and README.md
- Updated versions.yaml
- Generated version 23.5.1 with tag v23.5.1-3

## [24.1.0-1] - 2024-04-07

### Miscellaneous

- Generated version 24.1.0 with tag v24.1.0-1

## [24.0.0-1] - 2024-04-07

### Miscellaneous

- Generated version 24.0.0 with tag v24.0.0-1

## [23.9.0-1] - 2024-04-07

### Miscellaneous

- Generated version 23.9.0 with tag v23.9.0-1

## [23.8.1-2] - 2024-04-07

### Miscellaneous

- Generated version 23.8.1 with tag v23.8.1-2

## [23.7.0-2] - 2024-04-07

### Miscellaneous

- Generated version 23.7.0 with tag v23.7.0-2

## [23.6.0-2] - 2024-04-07

### Miscellaneous

- Generated version 23.6.0 with tag v23.6.0-2

## [23.5.1-2] - 2024-04-07

### Bug Fixes

- Removed write only from data sources as they are only for resource creation

### Build

- Sort propertyWriteOnlyKeys so the order is always the same when building the templates

### Dependencies

- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework  (#92)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-testing  (#89)

### Miscellaneous

- Added generated resources for version 23.9.0, 24.0.0, 24.1.0
- Updated versions.yaml
- Make 24.1.0 default version when building
- Updated versions.yaml
- Make generate to sort the write only keys
- Updated versions.yaml
- Generated version 23.5.1 with tag v23.5.1-2

## [24.1.0-0] - 2024-04-07

### Miscellaneous

- Generated version 24.1.0 with tag v24.1.0-0

## [24.0.0-0] - 2024-04-07

### Miscellaneous

- Generated version 24.0.0 with tag v24.0.0-0

## [23.9.0-0] - 2024-04-07

### Features

- Generate AWX API 23.9.0
- Generate AWX API 24.0.0
- Generate AWX API 24.1.0

### Miscellaneous

- Updated versions.yaml
- Updated versions.yaml
- Generated version 23.9.0 with tag v23.9.0-0

## [23.8.1-1] - 2024-04-07

### Miscellaneous

- Generated version 23.8.1 with tag v23.8.1-1

## [23.5.1-1] - 2024-04-07

### Miscellaneous

- Generated version 23.5.1 with tag v23.5.1-1

## [23.6.0-1] - 2024-04-07

### Miscellaneous

- Generated version 23.6.0 with tag v23.6.0-1

## [23.7.0-1] - 2024-04-07

### Bug Fixes

- Added write only fields for credentials and the ability to add custom constraints on fields

### Miscellaneous

- Updated versions.yaml
- Updated versions.yaml
- Generated version 23.7.0 with tag v23.7.0-1

### Tooling

- Added a generate build command to automate the build process for the providers
- Added missing git push command

## [23.8.1-0] - 2024-04-03

### Miscellaneous

- Generated version 23.8.1 with tag v23.8.1-0

## [23.7.0-0] - 2024-04-03

### Miscellaneous

- Generated version 23.7.0 with tag v23.7.0-0

## [23.6.0-0] - 2024-04-03

### Miscellaneous

- Generated version 23.6.0 with tag v23.6.0-0

## [23.5.1-0] - 2024-04-03

### Bug Fixes

- Inventory variables will be converted internally to json from yaml when needed  (#85)

### Dependencies

- *(deps)* Bump github.com/stretchr/testify from 1.8.4 to 1.9.0  (#88)

### Features

- Allow to use token authentication next to basic auth  (#86)
- Added multiple version of the AWX API  (#93)

### Miscellaneous

- Gofmt -s -w -r 'interface{} -> any'
- Goimports -w .
- Upgrade golang.org/x/exp
- Generated version 23.5.1 with tag v23.5.1-0

### Refactor

- Removed 21.8.0 and moved the generic properties to the root config  (#95)

## [23.8.1] - 2024-02-21

### Bug Fixes

- *(build)* Created a provider with the wrong binary name

### Dependencies

- *(deps)* Bump github.com/ilijamt/envwrap from 1.0.0 to 1.1.0  (#79)
- *(deps)* Bump golangci/golangci-lint-action from 3 to 4  (#80)

### Features

- Added API for 23.8.1  (#81)

### Testing

- Added tests for doRequest to cover some edge cases

## [23.7.0] - 2024-02-03

### Dependencies

- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework  (#75)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go  (#76)
- *(deps)* Bump codecov/codecov-action from 3 to 4  (#77)

## [23.6.0] - 2024-01-07

### Documentation

- Update README with built versions

### Miscellaneous

- Generated the API for 23.6.0

## [0.1.5] - 2024-01-07

### Bug Fixes

- AWX Client will not skip verification of SSL certificates if set
- Terraform provider for AWX v23.5.1  (#69)

### Dependencies

- *(deps)* Bump github.com/hashicorp/terraform-plugin-go  (#67)

### Features

- Added the AWX API for v23.6.0
- Associate the InstanceGroup with a JobTemplate  (#71)

### Refactor

- Separate config generation per API version and DL AWX API version for 23.5.1  (#68)

## [0.1.4] - 2023-12-14

### Bug Fixes

- Generated code had wrong StaticString value when we use a choice

### Dependencies

- *(deps)* Bump github/codeql-action from 2 to 3  (#66)

## [0.1.3] - 2023-12-14

### Bug Fixes

- Move the documentation back to the root

## [0.1.2] - 2023-12-14

### Bug Fixes

- Several resources required string wrapping of the json  (#65)

### Dependencies

- *(deps)* Bump goreleaser/goreleaser-action from 3.2.0 to 4.1.0
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go
- *(deps)* Bump github.com/hashicorp/terraform-plugin-log
- *(deps)* Bump goreleaser/goreleaser-action from 4.1.0 to 4.2.0
- *(deps)* Bump github.com/stretchr/testify from 1.8.1 to 1.8.2
- *(deps)* Bump actions/setup-go from 3 to 4
- *(deps)* Bump github.com/hashicorp/terraform-plugin-docs
- *(deps)* Bump github.com/spf13/cobra from 1.6.1 to 1.7.0
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go
- *(deps)* Bump github.com/stretchr/testify from 1.8.2 to 1.8.3
- *(deps)* Bump github.com/stretchr/testify from 1.8.3 to 1.8.4
- *(deps)* Bump goreleaser/goreleaser-action from 4.2.0 to 4.3.0
- *(deps)* Bump github.com/hashicorp/terraform-plugin-docs
- *(deps)* Bump github.com/hashicorp/terraform-plugin-log
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go
- *(deps)* Bump github.com/hashicorp/terraform-plugin-docs
- *(deps)* Bump github.com/iancoleman/strcase from 0.2.0 to 0.3.0
- *(deps)* Bump goreleaser/goreleaser-action from 4.3.0 to 4.4.0
- *(deps)* Bump actions/checkout from 3 to 4
- *(deps)* Bump goreleaser/goreleaser-action from 4.4.0 to 4.6.0
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go
- *(deps)* Bump crazy-max/ghaction-import-gpg from 5 to 6
- *(deps)* Bump goreleaser/goreleaser-action from 4.6.0 to 5.0.0
- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework  (#55)
- *(deps)* Bump golang.org/x/net from 0.14.0 to 0.17.0  (#56)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework  (#59)
- *(deps)* Bump github.com/spf13/cobra from 1.7.0 to 1.8.0  (#60)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go  (#61)
- *(deps)* Bump github.com/hashicorp/terraform-plugin-testing  (#62)
- *(deps)* Bump actions/setup-go from 4 to 5  (#63)

### Documentation

- Added badges to README.md

### Miscellaneous

- Update goreleaser to generate changelogs
- Cleanup of unnecesseary files in examples
- Remove old resources/api/21.8.0.json as it has been split
- Add .terraformrc in .gitignore

### Refactor

- Upgrade terraform plugin framework to 1.4.0 
- Split the API data into pieces and reorganize the structure of the project  (#54)
- Rename provider option from insecure_skip_verify to verify_ssl
- Modify the provider so we can add tests more easily  (#58)
- Rename the package name from awx to awx_21_8_0

### Testing

- Added tests for the general provider  (#57)
- Added tests to check the configuration of the provider

## [0.1.1] - 2022-12-08

### Bug Fixes

- Linting complaining of ireggular iota usage
- Various fixes and liniting issues

### Dependencies

- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework

## [0.1.0] - 2022-11-23

### Bug Fixes

- Makefile had the wrong directory for output of code coverage file
- Hook for credential panics on import of a blank resource and added more details to the log
- Hook for application and settings to prevent panic on import of a blank resource

### CI/CD

- Added dependabot
- Added golangci-lint and codeql-analysis workflows
- Added goreleaser to release the provider

### Dependencies

- *(deps)* Bump github.com/stretchr/testify from 1.7.2 to 1.8.1
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go
- *(deps)* Bump github.com/hashicorp/terraform-plugin-framework-validators
- *(deps)* Bump github.com/hashicorp/terraform-plugin-go

### Features

- Added provider version so it identifies with version also when doing calls to AWX
- Added associate/dissacociate of an Galaxy Credential to an Organization
- Added associate/disassociate of a Galaxy Credential to an Organization (#8)

### Miscellaneous

- Terraform fmt on the examples and remove redundant go:generate line in main.go
- Golangci-lint fixes
- Disabled associate/disassociate with Execution Environment for Organization

### Refactor

- Move all the setting of values to dedicated functions so they can be more easily tested and reused

<!-- generated by git-cliff -->
