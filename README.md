# dotenv 
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

The easiest way to use `.env` values in your next project. A flexible CLI that reads your source code project and generates a new file with your `.env` values inside. 

## How does it work? 

Let's say you have a `.env` file that looks like this:

```
STRIPE_API_KEY=4837394857859494
DEBUG=true
```

And you have the source code file:

```kotlin
class Example {
    val name = "dotenv CLI"
}
```

If you were to run the `dotenv` CLI on your source code, the CLI would generate a source code file that looks like this:

```kotlin
object Env {

}
```

That's right. It's empty. This is to prevent bloat but also security if you have information in your `.env` you don't want to expose to your project. If you want to use an environment variable in your source code, you need to first use it in your code:

```kotlin
class Example {
    val key = Env.stripeApiKey
}
```

Your compiler will complain at this time because you don't have a file `Env` with the variable `stripeApiKey`. But you need to define these variables to request from the `dotenv` CLI that you want to use an environment variable value. Now if you run the `dotenv` CLI, it will see `Env.stripeApiKey`, look for a environment variable `STRIPE_API_KEY` in the `.env` file, and generate:

```kotlin
object Env {
    val stripeApiKey = "4837394857859494"
}
```

## Getting started

This CLI is not meant to be used directly in your project (although you surely can do that). It's meant to be a CLI that does all the hard work while you are supposed to then create a wrapper around the CLI that's specific to your language of choice. That means that I recommend that instead of using the `dotenv` CLI directly in your project, instead use one of the following:

* [dotenv-android](https://github.com/levibostian/dotenv-android)
* [dotenv-ios](https://github.com/levibostian/dotenv-ios)

...these are all the wrappers for now, but there may be more in the future! See [goals of this project](#goals-of-this-project) to understand how more languages and build system wrappers may be supported in the future.

If you are still wondering how to get started using the `dotenv` CLI then here you go:

1. [Install the CLI](#install)
2. Run command `dotenv generate` to see the help docs printed to you on how to use the CLI. 

# Install 

There are many ways to download the CLI. 

1. (recommended) [Homebrew](https://brew.sh/) --> `brew install levibostian/tap/dotenv`
2. Download the binary directly from [GitHub releases](https://github.com/levibostian/dotenv/releases). Download the file on your computer anywhere you wish. Make sure to add the binary to your PATH. 

You now have the CLI installed on your machine. If you run `dotenv version` or `dotenv --help` and see information printed out instead of an error message, dotenv is successfully installed. 

## Goals of this project

* **Be flexible and powerful.** This CLI is not meant to be convenient as a CLI. Instead, this CLI is meant to do the heavy work for you and then we build language specific tools to be convenient. Like [dotenv-android](https://github.com/levibostian/dotenv-android) to make using `.env` files in your *Android app* very quick and easy. 
* **Binary form not relying on a specific programming language to execute.** We didn't want this project to be written in lets say Ruby to force the developer using it to have Ruby installed on their machine. Instead, have pre-built binaries created so developers could simply install the CLI and run it. Some developer eco-systems like iOS/Swift are used to tools being written in Ruby so that wouldn't be a pain but for Typescript/node or Android eco-systems Ruby is not common. Binary is universal. 
* **Generate a source code file in X language where the input language can be X or Y.** Rather then assume that your project is written in Java and you want to generate Java code, why not give you the flexibility in outputting Kotlin from your Java source code? 
* **Not support every programming language** This CLI is currently built and maintained by [myself](https://github.com/levibostian/) so I can use this tool in my projects. This means that this CLI only supports languages that I personally know so I can keep the tool maintained. This project's vision is not to support every programming language built-in with the CLI. Instead, I would recommend we refactor this CLI to be abstract enough to allow different programming languages via plugins or config files. This would keep the CLI code-base simple while also allowing others to support whatever language they need without having to become a core contributor of this code base. 

# Development 

* Compile code: 

```
go build -ldflags "-X main.version=1.0.1"
```

*Note: the `1.0.1` does not matter during development. You only want to run this command to allow tests to pass that check the version gets printed correctly.*

* Run tests: 

To make life easier, we recommend using [gotestsum](https://github.com/gotestyourself/gotestsum#install) to run tests. 

```
gotestsum --format pkgname
```

### Add a new language 

At this time, this CLI tool supports the following programming languages:

* Java
* Kotlin

With the plan to add:

* Swift

At this time, I am not looking for this tool to get bloated with dozens of languages. See [goals of this project](#goals-of-this-project) to understand how more languages and build system wrappers may be supported in the future.

**How do we add a new language?**

The codebase has been designed to make adding new languages as easy as possible. This is done by encapsulating all programming language related code to 1-2 files. From there, the codebase takes care of the rest of the work for you. 

1. Your friend is the `lang/` directory. This is where all programming language related code exists. When you want to add a new language, create 2 new files: 
* `<lang-name>.go` (example: `typescript.go`)
* `<lang-name>_test.go` (example: `typescript_test.go`)

This is where the Go source code exists for the new language. 

2. In each of the 2 new files that you created, write all of the code needed for your language. Check out `./lang/langs.go` for comments documenting the interfaces and view the existing languages like `./lang/kotlin.go` for examples. 

3. Go into `./lang/langs.go` and follow the instructions in that file. It will tell you a couple of things that you need to do in that file to add your new lang to the CLI. 

4. Make a e2e test in directory `./e2e/`. See the examples already there. Think of a use case like generating Typescript from Typescript source and make a e2e test for that scenario. 

# Deployment 

* Make sure to set version number when compiling: `go build -ldflags "-X main.version=1.0.1"`
* Set environment variables:
1. `REPO_PUSH_TOKEN` - github personal access token with `repos` permission access. 
2. `DOCKERHUB_USERNAME` - username for dockerhub to push docker images. 
3. `DOCKERHUB_TOKEN` - docker hub token to authenticate with account. 

## Contributors

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key))
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://slhck.info"><img src="https://avatars.githubusercontent.com/u/582444?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Werner Robitza</b></sub></a><br /><a href="#ideas-slhck" title="Ideas, Planning, & Feedback">🤔</a> <a href="https://github.com/levibostian/dotenv/commits?author=slhck" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/levibostian"><img src="https://avatars1.githubusercontent.com/u/2041082?v=4" width="100px;" alt=""/><br /><sub><b>Levi Bostian</b></sub></a><br /><a href="https://github.com/levibostian/dotenv/commits?author=levibostian" title="Code">💻</a> <a href="https://github.com/levibostian/dotenv/commits?author=levibostian" title="Documentation">📖</a> <a href="#maintenance-levibostian" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->
