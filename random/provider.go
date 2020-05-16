package random

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"random_id":       resourceId(),
			"random_shuffle":  resourceShuffle(),
			"random_pet":      resourcePet(),
			"random_string":   resourceString(),
			"random_password": resourcePassword(),
			"random_integer":  resourceInteger(),
			"random_uuid":     resourceUuid(),
		},
	}
}
