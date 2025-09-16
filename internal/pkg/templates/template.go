package templates

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed *.tmpl
var templates embed.FS

// //go:embed .gitignore.tmpl
// var gitignore embed.FS

// //go:embed .golangci.yaml.tmpl
// var golangci embed.FS

// //go:embed .pre-commit-config.yaml.tmpl
// var preCommitConfig embed.FS

// //go:embed go.mod.tmpl
// var goMod embed.FS

// //go:embed main.go.tmpl
// var mainGo embed.FS

// //go:embed service.go.tmpl
// var serviceGo embed.FS

// //go:embed Makefile.tmpl
// var makefile embed.FS

type Template struct {
	ProjectName         string
	ModuleName          string
	GoVersion           string
	GolangciLintVersion string
}

func (t Template) Gitignore() []byte {
	return t.parse(".gitignore.tmpl")
}

func (t Template) Golangci() []byte {
	return t.parse(".golangci.yaml.tmpl")
}

func (t Template) PreCommitConfig() []byte {
	return t.parse(".pre-commit-config.yaml.tmpl")
}

func (t Template) GoMod() []byte {
	return t.parse("go.mod.tmpl")
}

func (t Template) MainGo() []byte {
	return t.parse("main.go.tmpl")
}

func (t Template) ServiceGo() []byte {
	return t.parse("service.go.tmpl")
}

func (t Template) Makefile() []byte {
	return t.parse("Makefile.tmpl")
}

func (t Template) parse(name string) []byte {
	tmpl, err := template.ParseFS(templates, name)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	err = tmpl.ExecuteTemplate(&buf, name, t)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
