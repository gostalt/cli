<p align="center"><img src="https://avatars0.githubusercontent.com/u/50806258" width="200"></p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/gostalt/cli">
    <img src="https://goreportcard.com/badge/github.com/gostalt/cli" />
  </a>
</p>

# CLI

CLI is a repository containing all the CLI applications for Gostalt. Currently,
there is the `gostalt` binary for creating new Gostalt projects.

## `gostalt`

Gostalt is a binary for creating new Gostalt applications. Over time, more
functionality may be added, and additional CLI applications may be created.

### Installing

To install `gostalt`, just run
`go install github.com/gostalt/cli/gostalt@latest` (assuming a correctly
configured Go toolchain).

In time, once the ecosystem is more stable, the `@latest` pin will change to a
given version number.

### Commands

#### `new`

The `new` command is used to create a new Gostalt project in the current
directory. For example, to create a new project called `blog`:

```sh
gostalt new blog
```

You'll get feedback as each step of the `new` command is ran, as well as
instructions on how to run the freshly created application.
