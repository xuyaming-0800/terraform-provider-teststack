---
subcategory: "DEMO"
layout: "teststack"
page_title: "Teststack: teststack_demo"
sidebar_current: "docs-teststack-resource-demo"
description: |-
  Provides a resource to manage demo
---
# teststack_demo
Provides a resource to manage demo
## Example Usage
```hcl
resource "teststack_demo" "foo" {
  name = "demo_name"
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The Name of the Demo.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
DEMO can be imported using the id, e.g.
```
$ terraform import Teststack_demo.default demo-12345678
```

