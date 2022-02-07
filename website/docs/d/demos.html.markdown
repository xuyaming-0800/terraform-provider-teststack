---
subcategory: "DEMO"
layout: "teststack"
page_title: "Teststack: teststack_demos"
sidebar_current: "docs-teststack-datasource-demos"
description: |-
  Use this data source to query detailed information of demos
---
# teststack_demos
Use this data source to query detailed information of demos
## Example Usage
```hcl
data "teststack_demos" "foo" {
  ids = ["demo-123456"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Demo Ids.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `demos` - The collection of Demo query.
* `total_count` - The total count of Demo query.


