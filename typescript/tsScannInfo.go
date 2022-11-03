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
	types   map[string]TSInfoType
}

type TSInfoType struct {
	Name       string
	Type       string
	TsType     string
	Typescript bool
	dependOn   bool
}

type TSInfo struct {
	packages map[string]TSInfoPakage
}

/// exporter
type TSModuleInfo struct {
	structs map[string]string
	types   map[string]string
}

var TsInfo = TSInfo{}

func (i *TSInfo) Populate() {
	i.packages = make(map[string]TSInfoPakage)
	err := filepath.Walk("/Users/fabio/go/src/github.com/devdiversity/go-typescript/",
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
					if _, ok := i.packages[k]; !ok {
						i.packages[k] = TSInfoPakage{structs: make(map[string]TSStruct), types: make(map[string]TSInfoType)}
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
									}
									v.getStruct(typeSpec)
									i.packages[k].structs[typeSpec.Name.Name] = v
								default:
									//dependOn:ToBeImported(field.Type.(ast.Expr)),
									/* t := TSScanType{
										Name:       typeSpec.Name.Name,
										Typescript: isTypescript,
										Type:       GetFieldInfo(typeSpec.Type),
										TsType:     GetFieldTsInfo(typeSpec.Type),
										dependOn:   ToBeImported(typeSpec.Type.(ast.Expr)),
									}
									tsInfo.packages[k].types[typeSpec.Name.Name] = t */
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

	/* for p, _ := range tsScan.packages {
		fmt.Println("\n----------------")
		fmt.Printf("Package %s\n", p)
		for k, st := range tsScan.packages[p].structs {
			if st.Typescript {
				s, _ := ToTs(tsScan.packages, p, k)
				fmt.Println(s)

			}

		}
		fmt.Println("\ntypes----------------")
		for k, st := range tsScan.packages[p].types {
			if st.Typescript {
				s, _ := ToTs2(tsScan.packages, p, k)
				fmt.Println(s)
			}

		}
		fmt.Println("\n----------------")
	} */

}

/* func ScannAll() {
	tsScan := TSScan{}
	tsScan.packages = make(map[string]TSScanPakage)
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
					if _, ok := tsScan.packages[k]; !ok {
						tsScan.packages[k] = TSScanPakage{structs: make(map[string]TSScanStruct), types: make(map[string]TSScanType)}
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
									v := TSScanStruct{
										Name:       typeSpec.Name.Name,
										Typescript: isTypescript,
										Fields:     []TSScanField{},
									}
									v.getStruct(typeSpec)
									tsScan.packages[k].structs[typeSpec.Name.Name] = v
								default:
									//dependOn:ToBeImported(field.Type.(ast.Expr)),
									t := TSScanType{
										Name:       typeSpec.Name.Name,
										Typescript: isTypescript,
										Type:       GetFieldInfo(typeSpec.Type),
										TsType:     GetFieldTsInfo(typeSpec.Type),
										dependOn:   ToBeImported(typeSpec.Type.(ast.Expr)),
									}
									tsScan.packages[k].types[typeSpec.Name.Name] = t
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

	for p, _ := range tsScan.packages {
		fmt.Println("\n----------------")
		fmt.Printf("Package %s\n", p)
		for k, st := range tsScan.packages[p].structs {
			if st.Typescript {
				s, _ := ToTs(tsScan.packages, p, k)
				fmt.Println(s)

			}

		}
		fmt.Println("\ntypes----------------")
		for k, st := range tsScan.packages[p].types {
			if st.Typescript {
				s, _ := ToTs2(tsScan.packages, p, k)
				fmt.Println(s)
			}

		}
		fmt.Println("\n----------------")
	}
} */
