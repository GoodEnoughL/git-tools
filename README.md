## Introduction

This is a simple program that can help you manage your git repositories. Now the feature includes:

1. Create merge request automatically when you push a new branch to the remote repository. Attention: This feature just suit for the developer who want to merge the branch from task/subtask to feature/feat/ branch or from feature/feat/bugfix/hotfix branch to dev branch.

## Usage

### Push code to remote repository

use `gitpush` command directly in the terminal;

## Environment

- Go version: 1.16+
- Git version: 2.23+
- OS: Linux/MacOS/Windows

## How to compile

GOOS=darwin GOARCH=amd64 go build -o gitPush-darwin-amd64
GOOS=linux GOARCH=amd64 go build -o gitPush-linux-amd64
GOOS=windows GOARCH=amd64 go build -o gitPush-windows-amd64.exe

gitPush-darwin-amd64 for MacOS
gitPush-linux-amd64 for Linux
gitPush-windows-amd64.exe for Windows
