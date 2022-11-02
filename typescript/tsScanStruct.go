package typescript

import (
	"fmt"
	"go/ast"
	"strings"
)

type TSScanField struct {
	Name     string
	Type     string
	TsType   string
	JsonName string
	Expand   bool
	DependOn bool
}

type TSScanStruct struct {
	Name       string
	Typescript bool
	Fields     []TSScanField
}

func isNativeType(t string) bool {
	switch t {
	case "uint8", "uint16", "uint32", "uint64", "uint",
		"int8", "int16", "int32", "int64", "int",
		"float32", "float64":
		return true
	case "bool":
		return true
	case "string":
		return true
	}
	return false
}

func toBeImported(t ast.Expr) bool {
	switch ft := t.(type) {
	case *ast.Ident:
		return !isNativeType(ft.Name)
	case *ast.SelectorExpr:
		return true

	}
	return false
}

func typeToTypescript(k string) string {
	switch k {
	case "uint8", "uint16", "uint32", "uint64", "uint",
		"int8", "int16", "int32", "int64", "int",
		"float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	case "string":
		return "string"
	}
	return k
}

func getFieldInfo(t ast.Expr) string {
	result := ""
	switch ft := t.(type) {
	case *ast.Ident:
		result = ft.Name
	case *ast.SelectorExpr:
		result = fmt.Sprintf("%s.%s", ft.X, ft.Sel)
	case *ast.ArrayType:
		result = fmt.Sprintf("[]%s", getFieldInfo(ft.Elt))
	case *ast.StarExpr:
		result = fmt.Sprintf("*%s", getFieldInfo(ft.X))
	case *ast.MapType:
		result = fmt.Sprintf("map[%s]%s", ft.Key, getFieldInfo(ft.Value))
	case *ast.InterfaceType:
		result = "interface{}"
	default:
		fmt.Println(t)
	}
	return result
}

func getFieldTsInfo(t ast.Expr) string {
	result := ""
	switch ft := t.(type) {
	case *ast.Ident:
		result = typeToTypescript(ft.Name)
	case *ast.SelectorExpr:
		result = typeToTypescript(fmt.Sprintf("%s.%s", ft.X, ft.Sel))
	case *ast.ArrayType:
		result = fmt.Sprintf("%s[]", typeToTypescript(getFieldTsInfo(ft.Elt)))
	case *ast.StarExpr:
		result = fmt.Sprintf("Nullable<%s>", typeToTypescript(getFieldTsInfo(ft.X)))
	case *ast.MapType:
		result = fmt.Sprintf("Record<%s , %s>", typeToTypescript(fmt.Sprintf("%s", ft.Key)), typeToTypescript(getFieldTsInfo(ft.Value)))
	case *ast.InterfaceType:
		result = "unknown"
	default:
		fmt.Println(ft)
	}
	return result
}

func (s *TSScanStruct) getStruct(ts *ast.TypeSpec) {
	if st, ok := ts.Type.(*ast.StructType); ok {
		for _, field := range st.Fields.List {
			tag := ""
			if field.Tag != nil {
				tag = field.Tag.Value
			}

			jsonName := ""
			tagJson := TSTagJson{}
			if ok := tagJson.parse(tag); ok {
				jsonName = tagJson[0]
			}

			if len(field.Names) > 0 {
				var f = TSScanField{
					Name:     field.Names[0].String(),
					JsonName: jsonName,
					Type:     getFieldInfo(field.Type.(ast.Expr)),
					TsType:   getFieldTsInfo(field.Type.(ast.Expr)),
					Expand:   strings.Contains(tag, "`ts:\"expand\"`"),
					DependOn: toBeImported(field.Type.(ast.Expr)),
				}
				s.Fields = append(s.Fields, f)
			} else {
				if se, ok := field.Type.(*ast.SelectorExpr); ok {
					var f = TSScanField{
						Name:     fmt.Sprintf("%s.%s", se.X, se.Sel),
						JsonName: jsonName,
						Type:     getFieldInfo(field.Type.(ast.Expr)),
						TsType:   "",
						Expand:   strings.Contains(tag, "`ts:\"expand\"`"),
						DependOn: toBeImported(field.Type.(ast.Expr)),
					}
					s.Fields = append(s.Fields, f)
				} else {
					fmt.Printf("field type %T", field.Type)
				}
			}
		}
	}
}
