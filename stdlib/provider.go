package stdlib

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider exports terraform-provider-stdlib, which can be used in tests
// for other providers.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"stdlib_base64decodehash": dataSourceBase64DecodeHash(),
		},
	}
}
