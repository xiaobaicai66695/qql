package logic

import (
	"context"
	"qql/apps/user/rpc/internal/svc"
	"qql/apps/user/rpc/user"
	"reflect"
	"testing"
)

func TestNewRegisterLogic(t *testing.T) {
	type args struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	tests := []struct {
		name string
		args args
		want *RegisterLogic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegisterLogic(tt.args.ctx, tt.args.svcCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegisterLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "1", args: args{in: &user.RegisterReq{
				Phone:    "15691172798",
				Nickname: "xiaobaicai",
				Password: "210618",
				Avatar:   "png.jpg",
				Sex:      1,
			}}, wantPrint: true, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantPrint) {
				t.Errorf("Register() got = %v, want %v", got, tt.wantPrint)
			}
		})
	}
}
