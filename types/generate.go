package types

type GenerateOptions struct {
	InputLangs     string // string: `java` or list string: `java, kotlin`
	OutputLang     string
	OutputPath     string
	SourceCodePath string
	PackageName    string
	DotenvPath     string
}
