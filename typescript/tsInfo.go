package typescript

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TSInfoPakage struct {
	structs map[string]TSStruct
	types   map[string]TSType
}

type TSInfo struct {
	Packages map[string]TSInfoPakage
}

/// exporter
type TSModuleInfo struct {
	structs map[string]string
	types   map[string]string
}

type TSSourceLine struct {
	Pos    int
	End    int
	Line   int
	Source string
}

type TSSourceFile struct {
	Source string
	Name   string
	Lines  []TSSourceLine
	Len    int
}

func (ts *TSInfo) findStruct(p string, n string) bool {
	if _, ok := ts.Packages[p]; ok {
		if _, ok := ts.Packages[p].structs[n]; ok {
			return true
		}
	}
	return false
}

func (ts *TSInfo) findType(p string, n string) bool {
	if _, ok := ts.Packages[p]; ok {
		if _, ok := ts.Packages[p].types[n]; ok {
			return true
		}
	}
	return false
}

func (ts TSInfo) find(p string, n string) bool {
	return ts.findType(p, n) || ts.findStruct(p, n)
}

// popola TsInfo con tutte le definizioni dei tipi

func (i *TSInfo) Populate() {
	i.Packages = make(map[string]TSInfoPakage)
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				fset := token.NewFileSet()
				packages, err := parser.ParseDir(fset, path, nil, parser.ParseComments)

				if err != nil {
					panic(err)
				}

				for k, f := range packages {
					if _, ok := i.Packages[k]; !ok {
						i.Packages[k] = TSInfoPakage{structs: make(map[string]TSStruct), types: make(map[string]TSType)}
					}

					var src = []TSSourceFile{}
					for n, _ := range f.Files {
						dat, err := os.ReadFile(n)
						if err == nil {
							lines := []TSSourceLine{}
							pos := 0
							line := 1
							for p, k := range dat {
								if string(k) == "\n" {
									l := TSSourceLine{
										Pos:    pos,
										End:    p,
										Line:   line,
										Source: string(dat[pos:p]),
									}
									lines = append(lines, l)
									pos = p + 1
									line++
								}
							}
							s := TSSourceFile{
								Name:   n,
								Source: string(dat),
								Len:    len(dat),
								Lines:  lines,
							}
							src = append(src, s)
						}
					}

					p := doc.New(f, "./", 0)

					for _, t := range p.Types {
						var isTypescript = strings.HasPrefix(t.Doc, "Typescript:")
						for _, spec := range t.Decl.Specs {
							switch spec.(type) {
							case *ast.TypeSpec:
								typeSpec := spec.(*ast.TypeSpec)

								switch typeSpec.Type.(type) {
								case *ast.StructType:
									v := TSStruct{
										Name:       typeSpec.Name.Name,
										Typescript: isTypescript,
										Fields:     []TSSField{},
										SourceInfo: GetSourceInfo(int(typeSpec.Name.NamePos), src),
									}
									v.getStruct(typeSpec, src)
									i.Packages[k].structs[typeSpec.Name.Name] = v
								default:
									//dependOn:ToBeImported(field.Type.(ast.Expr)),
									t := TSType{
										Name:       typeSpec.Name.Name,
										Typescript: isTypescript,
										Type:       getFieldInfo(typeSpec.Type),
										TsType:     getFieldTsInfo(typeSpec.Type),
										dependOn:   toBeImported(typeSpec.Type.(ast.Expr)),
										SourceInfo: GetSourceInfo(int(typeSpec.Name.NamePos), src),
									}
									i.Packages[k].types[typeSpec.Name.Name] = t
								}
							}
						}

					}
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
