## Go Project Manager (AKA gopm)

![CI](https://github.com/dterbah/gopm/actions/workflows/go-test.yml/badge.svg)
[![codecov](https://codecov.io/gh/dterbah/gopm/branch/main/graph/badge.svg)](https://codecov.io/gh/dterbah/gopm)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gopm&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=dterbah_gopm)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gopm&metric=bugs)](https://sonarcloud.io/summary/new_code?id=dterbah_gopm)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gopm&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=dterbah_gopm)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gopm&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_gopm)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gopm&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=dterbah_gopm)

This CLI is a NPM like for your golang projects. It provides basic commands to manage your dependencies, start scripts or create golang projects from scratch.
Feel free to give me ideas for improvement !
<img src="./assets/logo.webp" width="200" />

### How to install this tool ?

To install this tool, you can use the following command :

```bash
go install github.com/dterbah/gopm@latest
```

### List of available commands

#### Create a new project

To create a new project, you can use this command :

```bash
gopm init
```

You can type the different information asking by the CLI. You could see default values for each information between parenthesis.
This command will create a new folder with these information :

```
<your-project-name>/
├── LICENSE.txt # This file contains the LICENSE information of your project
├── gopm.json # JSON file used to manage your project
├── go.mod # Module file for Go dependencies
├── go.sum # Checksum file for Go dependencies
└── <your-entry-file>.go # Entry point of your project
```

#### Install new dependency

To install a new dependency and add it in your Go project, use this command :

```bash
gopm i name-of-your-dependecy
# or
gopm install name-of-your-dependency
```

#### Install all dependencies

```bash

gopm install
# or
gopm i
```

#### Show current version of CLI

```bash
gopm -v
```

#### Use custom commands

If you want to use custom commands with this CLI, you can add these directly in the `gopm.json` file in the `scripts` section. When a project is created, few default commands are directly created.
