package main

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/benjaminlai/opa-sidecar/policy"
	"github.com/open-policy-agent/opa/rego"
)

var query rego.PreparedEvalQuery

func setup() {
	var err error
	p, err := policy.ReadPolicy(policyPath)
	if err != nil {
		log.Fatal(err)
	}

	query, err = rego.New(
		rego.Query(defaultQuery),
		rego.Module(policyPath, string(p)),
	).PrepareForEval(context.TODO())

	if err != nil {
		log.Fatalf("initial rego error: %v", err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func Test_result(t *testing.T) {
	ctx := context.TODO()
	type args struct {
		ctx   context.Context
		query rego.PreparedEvalQuery
		input map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test_admin",
			args: args{
				ctx:   ctx,
				query: query,
				input: map[string]interface{}{
					"issuer":   "service-b",
					"audience": "service-a",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := result(tt.args.ctx, tt.args.query, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("result() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("result() = %v, want %v", got, tt.want)
			}
		})
	}
}
