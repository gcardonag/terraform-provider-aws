---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_account_settings"
description: |-
  Terraform resource for managing an AWS QuickSight Account Settings.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_account_settings

Terraform resource for managing an AWS QuickSight Account Settings.

~> Deletion of this resource will not modify any settings, only remove the resource from state.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws. import QuicksightAccountSettings
from imports.aws.quicksight_account_subscription import QuicksightAccountSubscription
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        subscription = QuicksightAccountSubscription(self, "subscription",
            account_name="quicksight-terraform",
            authentication_method="IAM_AND_QUICKSIGHT",
            edition="ENTERPRISE",
            notification_email="notification@email.com"
        )
        QuicksightAccountSettings(self, "example",
            depends_on=[subscription],
            termination_protection_enabled=False
        )
```

## Argument Reference

This resource supports the following arguments:

* `default_namespace` - (Optional) The default namespace for this Amazon Web Services account. Currently, the default is `default`.
* `termination_protection_enabled` - (Optional) A boolean value that determines whether or not an Amazon QuickSight account can be deleted. If `true`, it does not allow the account to be deleted and results in an error message if a user tries to make a DeleteAccountSubscription request. If `false`, it will allow the account to be deleted.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `aws_account_id` - The ID for the AWS account that contains the settings.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import QuickSight Account Settings using the AWS account ID. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws. import QuicksightAccountSettings
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightAccountSettings.generate_config_for_import(self, "example", "012345678901")
```

Using `terraform import`, import QuickSight Account Settings using the AWS account ID. For example:

```console
% terraform import aws_quicksight_account_settings.example "012345678901"
```

<!-- cache-key: cdktf-0.20.8 input-8aa9ac9dd39e973eed9c876e9e1da7b45f81c9fd9e6f6f876709f215fb774274 -->