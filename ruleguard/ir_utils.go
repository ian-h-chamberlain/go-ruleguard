package ruleguard

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/types"

	"github.com/quasilyte/go-ruleguard/ruleguard/ir"
	"github.com/quasilyte/go-ruleguard/ruleguard/irconv"
)

type ConvertedAST struct {
	File     *ir.File
	Package  *types.Package
	TypeInfo *types.Info
}

func convertAST(ctx *LoadContext, imp *goImporter, filename string, src []byte) (ConvertedAST, error) {
	parserFlags := parser.ParseComments
	f, err := parser.ParseFile(ctx.Fset, filename, src, parserFlags)
	if err != nil {
		return ConvertedAST{}, fmt.Errorf("parse file error: %w", err)
	}

	typechecker := types.Config{Importer: imp}
	typesInfo := &types.Info{
		Types:     map[ast.Expr]types.TypeAndValue{},
		Uses:      map[*ast.Ident]types.Object{},
		Defs:      map[*ast.Ident]types.Object{},
		Instances: map[*ast.Ident]types.Instance{},
	}
	pkg, err := typechecker.Check("gorules", ctx.Fset, []*ast.File{f}, typesInfo)
	if err != nil {
		return ConvertedAST{}, fmt.Errorf("typechecker error: %w", err)
	}
	irconvCtx := &irconv.Context{
		Pkg:   pkg,
		Types: typesInfo,
		Fset:  ctx.Fset,
		Src:   src,
	}
	irfile, err := irconv.ConvertFile(irconvCtx, f)
	if err != nil {
		return ConvertedAST{}, fmt.Errorf("irconv error: %w", err)
	}
	return ConvertedAST{irfile, pkg, typesInfo}, nil
}
