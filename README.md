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

The following modules are included with Prospectus.

Most of the examples below will use either an expected or actual block, based on the most common use case, but any module can be used for either state.

### git_tag

This checks the git tag of the local repo. Supports the Regex helper

```
# This would use the current git tag directly
actual do
  git_tag
end
```

```
# This would convert v1.0.0 into 1.0.0
actual do
  git_tag
  regex /^v([\d.]+)$/
end
```

```
# And this would convert v1_0_0 into 1.0.0
actual do
  git_tag
  regex /^v(\d+)_(\d+)_(\d+)$/, '\1.\2.\3'
end
```

### git_hash

Checks the git hash of a local repository. Supports the chdir helper.

Will return the short hash unless the `long` argument is provided.

Primarily used for checking git submodules.

```
# Returns the short hash
actual do
  git_hash
  dir 'submodules/my-important-other-repo'
end

# Returns the full hash
actual do
  git_hash
  long
  dir 'submodules/other-repo'
end
```

### github_release

This checks the latest GitHub release for a repo (must be a real Release, not just a tag. Use github_tag if there isn't a Release). Supports the Regex helper and uses the GitHub API helper for API access.

```
expected do
  github_release
  repo 'amylum/s6'
end
```

### github_tag

This checks the latest GitHub tag for a repo. Supports the Regex helper and uses the GitHub API helper for API access.

```
expected do
  github_tag
  repo 'reubenhwk/radvd'
end
```

### github_hash

This checks the latest commit hash on GitHub. Uses the github_api helper, which requires octoauth. Designed to be used alongside the git_hash module for comparing local submodules with upstream commits.

Will give the 7 character short hash unless "long" is specified.

```
expected do
  github_hash
  repo 'akerl/keys'
end

expected do
  github_hash
  repo 'akerl/keys'
  long
end
```

### grep

This checks a local file's contents. Supports the Regex helper, and uses the provided regex pattern to match which line of the file to use. If no regex is specified, it will use the full first line of the file.

```
# Searches file for OPENSSL_VERSION = 1.0.1e and returns 1.0.1e
actual do
  grep
  file 'Makefile'
  regex /^OPENSSL_VERSION = ([\w.-]+)$/
end
```

### url_xpath

Used to parse an xpath inside a web page. Requires the nokogiri gem, and supports the Regex helper.

The easiest way to get an xpath is usually to use Chrome's Inspector to find the element you want and right click it -> Copy -> As XPath. There are some quirks, notably that nokogiri doesn't parse the tbody tag (just remove it from the xpath that Chrome provides).

```
# Parses the latest tag for procps-ng
expected do
  url_xpath
  url 'https://gitlab.com/procps-ng/procps/tags'
  xpath '/html/body/div[1]/div[2]/div[2]/div/div/div[2]/ul/li[1]/div[1]/a/strong/text()'
  regex /v([\d.]+)$/
end
```

### gemnasium

Parses the dep status on Gemnasium for a project. Will return the worst color from their dashboard (red, yellow, or green). Requires the netrc gem.

Due to the nature of the response, this is best used with a `static` expected block, like this:

```
item do
  expected do
    static
    set 'green'
  end

  actual do
    gemnasium
    slug 'amylum/server'
  end
end
```

### static

Basic module for staticly defining a value. Useful for testing, and also for comparisons against a known state (like with the gemnasium module).

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

## Included Helpers

### regex

Allows modification of result using regex. Supported by most modules, per the above modules list.

The first argument should be a regex pattern to match against the value. Note that an error will be raised if the value does not match the provided regex. An optional second value specifies the replacement string to use; the default is '\1', which will use the first capture group from your regex.

```
# This would convert v1.0.0 into 1.0.0
actual do
  git_tag
  regex /^v([\d.]+)$/
end
```

```
# And this would convert v1_0_0 into 1.0.0
actual do
  git_tag
  regex /^v(\d+)_(\d+)_(\d+)$/, '\1.\2.\3'
end
```

### chdir

Used to chdir to a different directory before loading the state.

```
actual do
  git_hash
  dir 'submodules/important_repo'
end
```

### github_api

Used by modules to provide authenticated access to the GitHub API. Uses the [octoauth gem](https://github.com/akerl/octoauth)

## Installation

    gem install prospectus

## License

prospectus is released under the MIT License. See the bundled LICENSE file for details.

