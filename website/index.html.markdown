---
layout: "teststack"
page_title: "Provider: teststack"
sidebar_current: "docs-teststack-index"
description: |-
The Teststack provider is used to interact with many resources supported by Teststack. The provider needs to be configured with the proper credentials before it can be used.
---

# Teststack Provider

The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** This guide requires an available Teststack account or sub-account with project to create resources.

## Example Usage
```hcl
# Configure the Teststack Provider
provider "teststack" {
  access_key = "your ak"
  secret_key = "your sk"
  region = "cn-beijing"
}

# Query Demo
data "teststack_demos" "default"{
  ids = ["demo-123456"]
}

#Create Demo
resource "teststack_demo" "foo" {
  name = "tf-test-1"
}

```

## Authentication

The Teststack provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding an `public_key` and `private_key` in-line in the
teststack provider block:

Usage:

```hcl
provider "teststack" {
   access_key = "your ak"
   secret_key = "your sk"
   region = "cn-beijing"
}
```

### Environment variables

You can provide your credentials via `TESTSTACK_ACCESS_KEY` and `TESTSTACK_SECRET_KEY`
environment variables, representing your teststack public key and private key respectively.
`TESTSTACK_REGION` is also used, if applicable:

```hcl
provider "teststack" {
  
}
```

Usage:

```hcl
$ export TESTSTACK_ACCESS_KEY="your_public_key"
$ export TESTSTACK_SECRET_KEY="your_private_key"
$ export TESTSTACK_REGION="cn-beijing"
$ terraform plan
```

