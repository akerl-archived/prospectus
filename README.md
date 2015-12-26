prospectus
=========

[![Gem Version](https://img.shields.io/gem/v/prospectus.svg)](https://rubygems.org/gems/prospectus)
[![Dependency Status](https://img.shields.io/gemnasium/akerl/prospectus.svg)](https://gemnasium.com/akerl/prospectus)
[![Build Status](https://img.shields.io/circleci/project/akerl/prospectus.svg)](https://circleci.com/gh/akerl/prospectus)
[![Coverage Status](https://img.shields.io/codecov/c/github/akerl/prospectus.svg)](https://codecov.io/github/akerl/prospectus)
[![Code Quality](https://img.shields.io/codacy/36b84b3bc7b24cd4991c4753f7788850.svg)](https://www.codacy.com/app/akerl/prospectus)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

Write short scripts in a simple DSL and use the prospectus tool to check for changes in expected vs. actual state.

I use this for checking my [homebrew tap](https://github.com/halyard/homebrew-formulae) and [ArchLinux packages](https://github.com/amylum) for outdated package versions: it compares the version I'm packaging now against the upstream latest version.

## Usage

This gem reads a "./.prospectus" file to determine expected/actual state. A prospectus file can be pretty lightweight:

```
item do
  name 'zlib'

  expected do
    github_release
    repo 'madler/zlib'
    regex /^v(.*)$/
  end

  actual do
    git_tag
    regex /^(.*)-\d+$/
  end
end
```

Prospectus works by letting you define "items", each of which have an "expected" and "actual" block. You can specify a "name", as above, otherwise it will infer the name from the directory containing the prospectus file.

The expected/actual blocks first define the module to be used, and then define any configuration for that module.

To run the check, just run `prospectus` in the directory with the .prospectus file, or use `prospectus -d /path/to/directory`.

## Included Modules

### static_test

Used for testing, this allows you to manually declare a state:

```
item do
  expected do
    static_test
    set '0.0.3'
  end
  actual do
    static_test
    set '0.0.1'
  end
end
```

## Installation

    gem install prospectus

## License

prospectus is released under the MIT License. See the bundled LICENSE file for details.

