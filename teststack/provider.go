package teststack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/xuyaming0800/terraform-provider-teststack/teststack/demo/demo"
)

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key": "ak",
		"secret_key": "sk",
		"region":     "cn-beijing-1",
	}
}

func ProviderConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
		Region:    d.Get("region").(string),
	}
	client, err := config.Client()
	return client, err
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TESTSTACK_ACCESS_KEY", nil),
				Description: descriptions["access_key"],
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TESTSTACK_SECRET_KEY", nil),
				Description: descriptions["secret_key"],
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TESTSTACK_REGION", nil),
				Description: descriptions["region"],
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"teststack_demos": demo.DataSourceTeststackDemos(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"teststack_demo": demo.ResourceTeststackDemo(),
		},
		ConfigureFunc: ProviderConfigure,
	}
}
