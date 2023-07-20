package lib

import (
	"github.com/dave/jennifer/jen"
	"github.com/heartBrokenGod/beer/model"
)

type CreateFunc struct {
	Package  model.Package
	Function model.Function
}

func CreateFunction(createFunc CreateFunc) *jen.Statement {

	f := jen.Func().Id(*createFunc.Function.Name)

	inpParams := []jen.Code{}
	for _, inputParam := range createFunc.Function.InputParams {
		p := jen.Id(*inputParam.Name)
		switch *inputParam.Type {
		case "string":
			p.String()
		case "int":
			p.Int()
		default:
			p.Any()
		}
		inpParams = append(inpParams, p)
	}

	f.Params(inpParams...)

	returnParams := []jen.Code{}
	for _, returnParam := range createFunc.Function.ReturnParams {
		p := jen.Id(*returnParam.Name)
		switch *returnParam.Type {
		case "string":
			p.String()
		case "int":
			p.Int()
		default:
			p.Any()
		}
		returnParams = append(returnParams, p)
	}

	f.Params(returnParams...)

	f.Block(
		jen.Comment("add your code here"),
		jen.Return(jen.Empty()),
	)

	return f
}
