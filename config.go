package main

type Config struct {
	Static string
	Source string
	Target string
	CssDir string
	CssComp struct {
		Sass string
		Sccs string
	}
	Output string
}