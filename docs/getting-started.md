---
id: getting-started
title: Getting started
sidebar_position: 2
---

# Getting started

## Prerequisites

- Go 1.23 or newer
- Git (to clone the repository)
- Node.js 20+ only if you want to build the Docusaurus documentation

## Build the current CLI

~~~bash
git clone https://github.com/atstaeff/ai-factory-cli.git
cd ai-factory-cli
go test ./...
go build -o factory ./cmd/factory
./factory help
~~~

On Windows, choose a suitable output name such as `factory.exe`.

## Inspect a local project

~~~bash
./factory scan .
./factory processes .
./factory doctor .
./factory status .
~~~

`scan` is read-only and reports whether familiar files such as `_bmad/`, `openspec/`, `AGENTS.md` or `.factory/PROCESS.md` exist. It is **not** a semantic, security or compatibility audit.

`doctor` currently checks only for `.factory/PROCESS.md`. It does not execute tools or validate model availability.

## Run the documentation locally

~~~bash
npm install
npm run start
~~~

Open `http://localhost:3000/ai-factory-cli/`. Use `npm run build` to check the production bundle.

:::note
The initial website uses npm install because a reproducible package lock could not be generated in the current authoring environment. Before a production release, generate and commit `package-lock.json` with `npm install --package-lock-only` and switch CI to `npm ci`.
:::

## Where to go next

- [CLI commands](./cli-reference.md)
- [BMAD setup](./guides/bmad.md)
- [Examples](./examples.md)
