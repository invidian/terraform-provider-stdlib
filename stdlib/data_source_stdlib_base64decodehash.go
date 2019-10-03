package stdlib

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceBase64DecodeHash() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBase64DecodeHashRead,
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sha1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha256": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBase64DecodeHashRead(d *schema.ResourceData, m interface{}) error {
	b, err := base64.StdEncoding.DecodeString(d.Get("data").(string))
	if err != nil {
		fmt.Errorf("failed to base64 decode 'data': %w", err)
	}
	id := fmt.Sprintf("%x", sha256.Sum256(b))
	d.Set("sha1", fmt.Sprintf("%x", sha1.Sum(b)))
	d.Set("sha256", id)
	d.SetId(id)
	return nil
}
