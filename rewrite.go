package astTool

import (
	"errors"
	"go/ast"
	"go/token"
)

const (
	HEAD = "head"
	TAIL = "tail"
)

/*
	替换节点
*/
func (h *HappyAst) ReplaceNode(pos token.Pos, newNode ast.Node) error {
	if pos == token.NoPos {
		return errors.New("pos not exist")
	}
	ast.Inspect(h.Ast, func(node ast.Node) bool {
		switch node.(type) {
		case ast.Node:
			if !node.Pos().IsValid() {
				return true
			}
			if node.Pos() == pos {
				node = newNode
				return false
			}
		}
		return true
	})
	return nil
}

/*
	添加语句
	@deprecated
*/
func (h *HappyAst) AddStmt(bmt *ast.BlockStmt, location string, stmt ast.Stmt) {
	if location == HEAD {
		tempStmtList := make([]ast.Stmt, 0)
		tempStmtList = append(tempStmtList, stmt)
		for _, v := range bmt.List {
			tempStmtList = append(tempStmtList, v)
		}
	} else {
		tempStmtList := bmt.List
		tempStmtList = append(tempStmtList, stmt)
		bmt.List = tempStmtList
	}
}

/*
	添加赋值语句
	bmt 语句块节点
	location 在bmt.List数组中的位置
	stmt 待放置的赋值语句节点
*/
func (h *HappyAst) AddAssignStmt(bmt *ast.BlockStmt, location int, stmt ast.Stmt) {
	tempStmtList := make([]ast.Stmt, 0)

	tempStmtList = append(tempStmtList, bmt.List[:location]...)
	tempStmtList = append(tempStmtList, stmt)
	tempStmtList = append(tempStmtList, bmt.List[location:]...)
	bmt.List = tempStmtList
}

/*
	添加函数声明
	location 在root.Decls数组中的位置
	stmt 待放置的赋值语句节点
*/
func (h *HappyAst) AppendFundDecl(location int, decl ast.Decl) {
	tempDeclList := make([]ast.Decl, 0)

	tempDeclList = append(tempDeclList, h.Ast.Decls[:location]...)
	tempDeclList = append(tempDeclList, decl)
	tempDeclList = append(tempDeclList, h.Ast.Decls[location:]...)
	h.Ast.Decls = tempDeclList
}

/*
	添加声明
	@deprecated
*/
func (h *HappyAst) AddDecl(location string, appends ...ast.Decl) {
	if location == HEAD {
		tempDecls := make([]ast.Decl, 0)
		for _, v := range appends {
			tempDecls = append(tempDecls, v)
		}
		for _, val := range h.Ast.Decls {
			tempDecls = append(tempDecls, val)
		}
		h.Ast.Decls = tempDecls
		return
	}
	if location == TAIL {
		tempDecls := make([]ast.Decl, 0)
		tempDecls = h.Ast.Decls
		for _, v := range appends {
			tempDecls = append(tempDecls, v)
		}
		h.Ast.Decls = tempDecls
		return
	}
}

/*
type svc interface {
  UserGet()
}

=>

type svc interface {
  UserGet()
  RoleGet()
}
*/
func (h *HappyAst) AddFieldOfFuncType(bmt *ast.FieldList, location int, field *ast.Field) {
	tempFieldList := make([]*ast.Field, 0)

	tempFieldList = append(tempFieldList, bmt.List[:location]...)
	tempFieldList = append(tempFieldList, field)
	tempFieldList = append(tempFieldList, bmt.List[location:]...)
	bmt.List = tempFieldList
}

/*
参数:
bmt,被插入的fieldList
location,在bmt中的顺序. -1尾部添加  0头部添加
field,插入的field

示例:
type PartnerSvcEndpoints struct {
	modelFetchEndpoint  kitendpoint.Endpoint
}

=>

type PartnerSvcEndpoints struct {
	modelFetchEndpoint  kitendpoint.Endpoint
	gameFetchEndpoint  kitendpoint.Endpoint
}
*/
func (h *HappyAst) AddField(bmt *ast.FieldList, location int, field *ast.Field) {
	tempFieldList := make([]*ast.Field, 0)

	if location == -1 {
		tempFieldList = append(tempFieldList, bmt.List[:]...)
		tempFieldList = append(tempFieldList, field)
	} else {
		tempFieldList = append(tempFieldList, bmt.List[:location]...)
		tempFieldList = append(tempFieldList, field)
		tempFieldList = append(tempFieldList, bmt.List[location:]...)
	}

	bmt.List = tempFieldList
}
