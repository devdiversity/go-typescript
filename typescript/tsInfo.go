package typescript

import (
	"errors"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/northwood-labs/golang-utils/exiterrorf"
)

type TSInfoPakage struct {
	structs map[string]TSStruct
	types   map[string]TSType
	enums   map[string]TSEnum
}

type TSInfo struct {
	Packages map[string]TSInfoPakage
}

type TSEnum struct {
	Name string
	Info []TSEnumInfo
}

type TSEnumInfo struct {
	Key   string
	Value string
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

func (i *TSInfo) getConst(p string, c *doc.Value, src []TSSourceFile) {
	var isTypescript = strings.HasPrefix(c.Doc, "Typescript:")
	if isTypescript {
		command := strings.TrimPrefix(c.Doc, "Typescript:")
		command = strings.TrimSpace(command)
		command = strings.Trim(command, "\n")

		enumName := ""
		if strings.Contains(command, "enum=") {
			enumName = strings.TrimPrefix(command, "enum=")
			enum := TSEnum{
				Name: enumName,
				Info: []TSEnumInfo{},
			}
			d := c.Decl
			iota := false
			iotaValue := 0
			for _, s := range d.Specs {
				v := s.(*ast.ValueSpec) // safe because decl.Tok == token.CONST
				if len(v.Values) > 0 {
					be, ok := v.Values[0].(*ast.BinaryExpr)
					if ok {
						x := be.X.(*ast.BasicLit)
						exiterrorf.ExitErrorf(errors.New(fmt.Sprintf("Enum Binary Expression Not implemented %s %s %s AT: %s\n", x.Value, be.Op.String(), be.Y, getSourceInfo(int(x.ValuePos), src))))
					}
					ident, ok := v.Values[0].(*ast.Ident)
					if ok {
						if ident.Name == "iota" {
							iota = true
							iotaValue = v.Names[0].Obj.Data.(int)
							enum.Info = append(enum.Info, TSEnumInfo{Key: v.Names[0].Name, Value: fmt.Sprintf("%d", iotaValue)})
						}
					}
					list, ok := v.Values[0].(*ast.BasicLit)
					if ok {
						enum.Info = append(enum.Info, TSEnumInfo{Key: v.Names[0].Name, Value: fmt.Sprintf("%s", list.Value)})
					}
				} else {
					for _, name := range v.Names {
						if iota {
							iotaValue++
							enum.Info = append(enum.Info, TSEnumInfo{Key: name.Name, Value: fmt.Sprintf("%d", iotaValue)})
						}

					}
				}
			}
			i.Packages[p].enums[enumName] = enum
			t1 := TSType{
				Name:       enumName,
				Typescript: true,
				Type:       "",
				TsType:     fmt.Sprintf("typeof Enum%s[keyof typeof Enum%s] ", enumName, enumName), //getFieldTsInfo(expr.Type),
				dependOn:   false,
				SourceInfo: "",
			}
			i.Packages[p].types[enumName] = t1
		}

	}
}

func (i *TSInfo) getType(p string, t *doc.Type, src []TSSourceFile) {
	var isTypescript = strings.HasPrefix(t.Doc, "Typescript:")
	if isTypescript {

		command := strings.TrimPrefix(t.Doc, "Typescript:")
		command = strings.TrimSpace(command)
		command = strings.Trim(command, "\n")

		fmt.Println("Typescript = ", command)
	}
	for _, spec := range t.Decl.Specs {
		if len(t.Consts) > 0 {
			fmt.Println(t.Consts[0].Doc)
			i.getConst(p, t.Consts[0], src)

			continue
		}
		switch spec.(type) {
		case *ast.TypeSpec:
			typeSpec := spec.(*ast.TypeSpec)

			switch typeSpec.Type.(type) {
			case *ast.StructType:
				v := TSStruct{
					Name:       typeSpec.Name.Name,
					Typescript: isTypescript,
					Fields:     []TSSField{},
					SourceInfo: getSourceInfo(int(typeSpec.Name.NamePos), src),
				}
				v.getStruct(typeSpec, src)
				i.Packages[p].structs[typeSpec.Name.Name] = v
			default:
				t := TSType{
					Name:       typeSpec.Name.Name,
					Typescript: isTypescript,
					Type:       getFieldInfo(typeSpec.Type),
					TsType:     getFieldTsInfo(typeSpec.Type),
					dependOn:   toBeImported(typeSpec.Type.(ast.Expr)),
					SourceInfo: getSourceInfo(int(typeSpec.Name.NamePos), src),
				}
				i.Packages[p].types[typeSpec.Name.Name] = t
			}
		}
	}
}

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
					exiterrorf.ExitErrorf(err)
				}

				for pkg, f := range packages {
					if _, ok := i.Packages[pkg]; !ok {
						i.Packages[pkg] = TSInfoPakage{structs: make(map[string]TSStruct), types: make(map[string]TSType), enums: make(map[string]TSEnum)}
					}

					var src = []TSSourceFile{}
					for n := range f.Files {

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
									if strings.Contains(l.Source, "// Typescript:") {

										if strings.Contains(l.Source, "TStype=") {
											s := strings.TrimPrefix(l.Source, "// Typescript: TStype=")
											a := strings.Split(s, "=")
											if len(a) == 2 {
												t := TSType{
													Name:       strings.Trim(a[0], " "),
													Typescript: true,
													Type:       "UserDefined",
													TsType:     strings.Trim(a[1], " "),
													dependOn:   false,
													SourceInfo: getSourceInfo(int(l.Pos), src),
												}
												i.Packages[pkg].types[strings.Trim(a[0], " ")] = t
											}

										}
									}
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
						i.getType(pkg, t, src)
					}

					for _, c := range p.Consts {
						i.getConst(pkg, c, src)
					}
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
