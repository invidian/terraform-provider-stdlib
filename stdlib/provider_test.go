package stdlib_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/invidian/terraform-provider-stdlib/stdlib"
)

func TestProvider(t *testing.T) {
	if err := stdlib.Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("internal validation: %v", err)
	}
}
