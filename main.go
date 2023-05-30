package main

import (
	"context"
	"log"

	// "github.com/benjaminlai/opa-sidecar/policy"
	"github.com/benjaminlai/opa-sidecar/policy"
	"github.com/open-policy-agent/opa/rego"
)

var (
	policyPath   = "policies/sample.rego"
	defaultQuery = "x = data.sample.allow"
)

type input struct {
	Issuer   string `json:"issuer"`
	Audience string `json:"audience"`
}

func main() {
	s := input{
		Issuer:   "service-b",
		Audience: "service-a",
	}

	input := map[string]interface{}{
		"issuer":   s.Issuer,
		"audience": s.Audience,
	}

	policy, err := policy.ReadPolicy(policyPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	query, err := rego.New(
		rego.Query(defaultQuery),
		rego.Module(policyPath, string(policy)),
	).PrepareForEval(ctx)
	if err != nil {
		log.Fatalf("initial rego error: %v", err)
	}

	ok, _ := result(ctx, query, input)
	log.Println(ok)
}

func result(ctx context.Context, query rego.PreparedEvalQuery, input map[string]interface{}) (bool, error) {
	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalf("evaluation error: %v", err)
	} else if len(results) == 0 {
		log.Fatal("undefined result", err)
		// Handle undefined result.
	} else if result, ok := results[0].Bindings["x"].(bool); !ok {
		log.Fatalf("unexpected result type: %v", result)
	}

	return results[0].Bindings["x"].(bool), nil
}
