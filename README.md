# dotenv 

WIP. Docs will come. 

# Development 

* Compile code: 

```
go build
```

* Run tests: 

To make life easier, we recommend using [gotestsum](https://github.com/gotestyourself/gotestsum#install) to run tests. 

```
gotestsum --format pkgname
```

### Add a new language 

At this time, I am not looking for this tool to get bloated with dozens of languages. The more we add, the more we have to maintain! If this repo becomes popular then I advice that we create a plugin system or something like it so people can create go modules that can plugin to the CLI to support X language. 

With that being said, this main repo is open to adding new languages:
* Javascript
* Typescript
* Swift
* Go

Why these languages? 
1. I ([levibostian](https://github.com/levibostian/)) the current maintainer knows these languages and can therefore maintain the code. 
2. I have been toying around with the idea of adding these languages to the project at some point anyway so I can use this tool for those projects. 

**How do we add a new language?**

The codebase has been designed to make adding new languages as easy as possible. This is done by encapsulating all programming language related code to 1-2 files. From there, the codebase takes care of the rest of the work for you. 

1. Your friend is the `lang/` directory. This is where all programming language related code exists. When you want to add a new language, create 2 new files: 
* `<lang-name>.go` (example: `typescript.go`)
* `<lang-name>_test.go` (example: `typescript_test.go`)

This is where the Go source code exists for the new language. 

2. In each of the 2 new files that you created, write all of the code needed for your language. Check out `./lang/langs.go` for comments documenting the interfaces and view the existing languages like `./lang/kotlin.go` for examples. 

3. Go into `./lang/langs.go` and follow the instructions in that file. It will tell you a couple of things that you need to do in that file to add your new lang to the CLI. 

