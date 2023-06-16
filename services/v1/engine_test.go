package v1

import (
	"ir-web/conf"
	"ir-web/internal/wrapper"
	v1_req "ir-web/models/protoreq/v1"
	"testing"
)

func init() {
	config, err := conf.Read("../../conf/config.json")
	if err != nil {
		panic("read " + err.Error())
	}
	conf.InitGlobalConfig(config)
}

func TestQuery(t *testing.T) {
	type args struct {
		ctx     *wrapper.Context
		reqBody interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				ctx: &wrapper.Context{},
				reqBody: &v1_req.EngineRequest{
					Token: "hello",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Query(tt.args.ctx, tt.args.reqBody); (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
