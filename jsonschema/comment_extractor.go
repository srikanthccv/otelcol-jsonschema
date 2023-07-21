package jsonschema

import (
	"fmt"
	"io/fs"
	"os"
	gopath "path"
	"path/filepath"
	"regexp"
	"strings"

	"go/ast"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
)

func modCacheLoc() string {
	return build.Default.GOPATH + "/pkg/mod/"
}

var (
	semVerRegexCompiled = regexp.MustCompile(`\@v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)
	modCacheRegex       = regexp.MustCompile(fmt.Sprintf(`^%s`, modCacheLoc()))
)

// ExtractGoComments will read all the go files contained in the provided path,
// including sub-directories, in order to generate a dictionary of comments
// associated with Types and Fields. The results will be added to the `commentsMap`
// provided in the parameters and expected to be used for Schema "description" fields.
//
// The `go/parser` library is used to extract all the comments and unfortunately doesn't
// have a built-in way to determine the fully qualified name of a package. The `base` paremeter,
// the URL used to import that package, is thus required to be able to match reflected types.
//
// When parsing type comments, we use the `go/doc`'s Synopsis method to extract the first phrase
// only. Field comments, which tend to be much shorter, will include everything.
func ExtractGoComments(base, path string, commentMap map[string]string) error {
	fset := token.NewFileSet()
	dict := make(map[string][]*ast.Package)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
			if err != nil {
				return err
			}
			for _, v := range d {
				// paths may have multiple packages, like for tests
				k := gopath.Join(base, path)
				dict[k] = append(dict[k], v)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	for pkg, p := range dict {
		for _, f := range p {
			gtxt := ""
			typ := ""
			ast.Inspect(f, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.TypeSpec:
					typ = x.Name.String()
					if !ast.IsExported(typ) {
						typ = ""
					} else {
						txt := x.Doc.Text()
						if txt == "" && gtxt != "" {
							txt = gtxt
							gtxt = ""
						}
						txt = doc.Synopsis(txt)
						// check if README.md exists in the directory
						// if so, use the contents as the description
						// for "markdownDescription"
						var markdownDescription string

						readmePath := filepath.Join(pkg, "README.md")
						if _, err := fs.Stat(os.DirFS(pkg), "README.md"); err == nil {
							// read the file
							data, err := os.ReadFile(readmePath)
							if err == nil {
								markdownDescription = string(data)
							}
						}

						// remove semver from the pkg
						pkg = semVerRegexCompiled.ReplaceAllString(pkg, "")
						// remove the mod cache path
						pkg = modCacheRegex.ReplaceAllString(pkg, "")
						commentMap[fmt.Sprintf("%s.%s", pkg, typ)] = strings.TrimSpace(txt)
						if markdownDescription != "" {
							commentMap[fmt.Sprintf("%s.%s.markdownDescription", pkg, typ)] = strings.TrimSpace(markdownDescription)
						}
					}
				case *ast.Field:
					txt := x.Doc.Text()
					if typ != "" && txt != "" {
						for _, n := range x.Names {
							if ast.IsExported(n.String()) {
								k := fmt.Sprintf("%s.%s.%s", pkg, typ, n)
								commentMap[k] = strings.TrimSpace(txt)
							}
						}
					}
				case *ast.GenDecl:
					// remember for the next type
					gtxt = x.Doc.Text()
				}
				return true
			})
		}
	}

	return nil
}
