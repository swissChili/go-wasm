package main

type Config struct {
	Static string
	Source string
	Target string
	CssDir string
	CssComp struct {
		Sass string
		Scss string
	}
	Output string
}