package stdlib

import (
	"crypto/sha1" // #nosec:G505
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	fieldData   = "data"
	fieldSHA1   = "sha1"
	fieldSHA256 = "sha256"
)

func dataSourceBase64DecodeHash() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBase64DecodeHashRead,
		Schema: map[string]*schema.Schema{
			fieldData: {
				Type:     schema.TypeString,
				Required: true,
			},
			fieldSHA1: {
				Type:     schema.TypeString,
				Computed: true,
			},
			fieldSHA256: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBase64DecodeHashRead(d *schema.ResourceData, m interface{}) error {
	b, err := base64.StdEncoding.DecodeString(d.Get(fieldData).(string))
	if err != nil {
		return fmt.Errorf("decoding %q field: %w", fieldData, err)
	}

	id := fmt.Sprintf("%x", sha256.Sum256(b))

	if err := d.Set(fieldSHA1, fmt.Sprintf("%x", sha1.Sum(b))); err != nil { // nolint:gosec
		return fmt.Errorf("setting %q field: %w", fieldSHA1, err)
	}

	if err := d.Set(fieldSHA256, id); err != nil {
		return fmt.Errorf("setting %q field: %w", fieldSHA256, err)
	}

	d.SetId(id)

	return nil
}
