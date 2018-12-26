package main

type CssCompiler struct {
	Less string
	Sass string
	Scss string
}

type Config struct {
	Static string
	Source []string
	Target string
	CssDir string
	CssComp CssCompiler
	Output string
}