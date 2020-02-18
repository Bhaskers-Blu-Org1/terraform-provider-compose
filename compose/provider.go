package compose

import (
	"github.com/compose/gocomposeapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider implements a terraform provider for compose.io
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Compose.io API Token",
				DefaultFunc: schema.EnvDefaultFunc("COMPOSE_API_TOKEN", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"compose_account": dataSourceAccount(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"compose_deployment": resourceDeployment(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return composeapi.NewClient(d.Get("api_token").(string))
		},
	}
}
