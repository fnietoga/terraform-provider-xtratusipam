package xtratusipam

import (
	"context"

	cli "terraform-provider-xtratusipam/ipamclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The root url of the APIM REST API solution to be used, without the /api url suffix",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("XTRATUSIPAM_TOKEN", nil),
				Description: "The bearer token to be used when authenticating to the API. Must be also assigned at XTRATUSIPAM_TOKEN environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"xtratusipam_reservation": resourceReservation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"xtratusipam_reservations": dataSourceReservations(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	hostUrl := d.Get("api_url").(string)
	token := d.Get("token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if hostUrl != "" && token != "" {
		c, err := cli.NewClient(&hostUrl, &token)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}

	c, err := cli.NewClient(nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
