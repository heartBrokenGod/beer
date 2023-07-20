package lib

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/heartBrokenGod/beer/model"
)

func ToPtr[T any](val T) *T {
	return &val
}

func TestCreateFunction(t *testing.T) {
	type args struct {
		createFunc CreateFunc
	}
	tests := []struct {
		name string
		args args
		want *jen.Statement
	}{
		{
			name: "",
			args: args{
				createFunc: CreateFunc{
					Package: "some package",
					Function: model.Function{
						Name: ToPtr("sqare"),
						InputParams: []model.InputParam{
							{
								Name: ToPtr("num"),
								Type: ToPtr("int"),
							},
						},
						ReturnParams: []model.ReturnParam{
							{
								Name: ToPtr("square"),
								Type: ToPtr("int"),
							},
						},
					},
				},
			},
			want: jen.Empty(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateFunction(tt.args.createFunc); true {
				fmt.Printf("%#v", got)
			}
		})
	}
}
