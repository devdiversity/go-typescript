package typescript

/*
type TSScanPakage struct {
	structs map[string]TSStruct
	types   map[string]TSScanType
}

type TSScanType struct {
	Name       string
	Type       string
	TsType     string
	Typescript bool
	dependOn   bool
}

type TSScan struct {
	packages map[string]TSScanPakage
}

/// exporter
type TSModule struct {
	structs map[string]string
	types   map[string]string
}

type TSSouces map[string]TSModule

var tsSouces = TSSouces{}

func (ts TSSouces) findStruct(p string, n string) bool {
	if _, ok := ts[p]; !ok {
		if _, ok := ts[p].structs[n]; !ok {
			return true
		}
	}
	return false
}
func (ts TSSouces) findType(p string, n string) bool {
	if _, ok := ts[p]; !ok {
		if _, ok := ts[p].types[n]; !ok {
			return true
		}
	}
	return false
}
func (ts TSSouces) find(p string, n string) bool {
	return ts.findType(p, n) || ts.findStruct(p, n)
}

/* func ToTs(tsScan map[string]TSScanPakage, p string, k string) (string, []string) {
	var result = ""
	var dependencies = []string{}
	s := tsScan[p].structs[k]
	if len(s.Fields) > 0 {
		result += fmt.Sprintf("\n%s interface {\n", k)
		for _, v := range tsScan[p].structs[k].Fields {
			var name = v.JsonName
			if v.JsonName == "-" || v.JsonName == "" {
				name = v.Name
			}
			if v.Expand {
				sp := strings.Split(v.Name, ".")
				for _, v := range tsScan[sp[0]].structs[sp[1]].Fields {
					result += fmt.Sprintf("\t%s: %s;\n", v.Name, v.TsType)
				}
			} else {
				result += fmt.Sprintf("\t%s: %s;\n", name, v.TsType)
			}
			if v.DependOn {
				dependencies = append(dependencies, v.TsType)
			}

		}
		result += fmt.Sprintf("}\n")
	}
	if len(dependencies) > 0 {
		result += "// dependencies\n"
		for _, v := range dependencies {
			pk := p
			st := string(v)
			if strings.Contains(st, ".") {
				sp := strings.Split(st, ".")
				if len(sp) == 2 {
					pk = sp[0]
					st = sp[1]
				}
			}
			s, _ := ToTs(tsScan, pk, st)
			result += s
		}
	}

	return result, dependencies
}

func ToTs2(tsScan map[string]TSScanPakage, p string, k string) (string, []string) {
	var result = ""
	var dependencies = []string{}
	s := tsScan[p].types[k]

	result = fmt.Sprintf("%s %s", s.Name, s.TsType)

	if s.dependOn {
		pk := p
		v := s.TsType
		if strings.Contains(s.TsType, ".") {
			sp := strings.Split(s.TsType, ".")
			if len(sp) == 2 {
				pk = sp[0]
				v = sp[1]
			}
		}
		s, _ := ToTs2(tsScan, pk, v)
		result += s

	}

	return result, dependencies
}

func getStructInfo(ts *ast.TypeSpec) string {
	if st, ok := ts.Type.(*ast.StructType); ok {
		fmt.Println(ts.Name.Name)
		for _, field := range st.Fields.List {
			fmt.Printf("%s %s\n", field.Names[0], GetFieldInfo(field.Type.(ast.Expr)))
		}
	}
	return fmt.Sprintf("%s", ts.Name.Name)
}

*/

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
