package demo

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

/*

Import
DEMO can be imported using the id, e.g.
```
$ terraform import Teststack_demo.default demo-12345678
```

*/

func ResourceTeststackDemo() *schema.Resource {
	return &schema.Resource{
		Create: resourceTeststackDemoCreate,
		Read:   resourceTeststackDemoRead,
		Delete: resourceTeststackDemoDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Name of the Demo.",
			},
		},
	}
}

func resourceTeststackDemoCreate(d *schema.ResourceData, meta interface{}) (err error) {
	d.SetId("1000")
	return resourceTeststackDemoRead(d, meta)
}

func resourceTeststackDemoRead(d *schema.ResourceData, meta interface{}) (err error) {
	d.SetId("1000")
	return d.Set("name", "demo_name")
}

func resourceTeststackDemoDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return err
}
