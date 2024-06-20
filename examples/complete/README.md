# complete

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5.0, <= 1.5.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 3.57.0 |
| <a name="requirement_random"></a> [random](#requirement\_random) | >= 3.4.3 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_appmesh"></a> [appmesh](#module\_appmesh) | git::https://github.com/launchbynttdata/tf-aws-module_primitive-appmesh | 1.0.1 |
| <a name="module_virtual_router"></a> [virtual\_router](#module\_virtual\_router) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | Name to use for the virtual router. Must be between 1 and 255 characters in length. | `string` | n/a | yes |
| <a name="input_app_mesh_name"></a> [app\_mesh\_name](#input\_app\_mesh\_name) | Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length. | `string` | n/a | yes |
| <a name="input_listeners"></a> [listeners](#input\_listeners) | Listeners that the virtual router is expected to receive inbound traffic from. Currently only one listener is supported per virtual router. | `any` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of custom tags to be attached to this resource | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | ID of the Virtual router. |
| <a name="output_arn"></a> [arn](#output\_arn) | ARN of the Virtual router |
| <a name="output_name"></a> [name](#output\_name) | Name of the Virtual router |
| <a name="output_mesh_name"></a> [mesh\_name](#output\_mesh\_name) | Name of the App Mesh |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
