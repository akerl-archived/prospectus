prospectus
=========

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/akerl/prospectus/build.yml?branch=main)](https://github.com/akerl/prospectus/actions)
[![GitHub release](https://img.shields.io/github/release/akerl/prospectus.svg)](https://github.com/akerl/prospectus/releases)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

Tool to check for changes in expected vs. actual state

## Usage

### Check specification

Checks must implement responses for the following commands:

### load

The `load` command accepts a hash with a single key, the directory being checked, and returns an array of checks with optional metadata.

Input: `{"dir": "/path/to/main/dir"}`
Output: `[{"name": "check_N", "metadata": {"foo": "bar"}}, ...]`

### execute

The `execute` command accepts a hash representing the check object. Metadata provided during the `load` call is included. The return value must be a Result object for the given check.

Input: `{"dir": "/path/to/main/dir", "file": "/path/to/main/dir/.prospectus.d/checkfile", "name": "check_N", "metadata": {"foo": "bar"}}`
Output: `{"actual": "unhappy", "expected": {"type": "string", "data": {"expected": "happy"}}}`

### fix

The `fix` command can attempt to fix a failed check automatically. It accepts a hash representing the failed result, which includes the originating check. The return value must be a Result object for the given check.

**Note:** The check must respond to the `fix` command, but if it does not support automatic fixes, it can respond by emiting the same result object it was given.

Input: `{"actual": "unhappy", "expected": {"type": "string", "data": {"expected": "happy"}}, "check": {"dir": "/path/to/main/dir", "file": "/path/to/main/dir/.prospectus.d/checkfile", "name": "check_N", "metadata": {"foo": "bar"}}}`
Output: `{"actual": "happy", "expected": {"type": "string", "data": {"expected": "happy"}}}`

## Installation

## License

prospectus is released under the MIT License. See the bundled LICENSE file for details.
