// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResources struct {
	// dns_target_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/route53recoveryreadiness_resource_set#dns_target_resource Route53RecoveryreadinessResourceSet#dns_target_resource}
	DnsTargetResource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResource `field:"optional" json:"dnsTargetResource" yaml:"dnsTargetResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/route53recoveryreadiness_resource_set#readiness_scopes Route53RecoveryreadinessResourceSet#readiness_scopes}.
	ReadinessScopes *[]*string `field:"optional" json:"readinessScopes" yaml:"readinessScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/route53recoveryreadiness_resource_set#resource_arn Route53RecoveryreadinessResourceSet#resource_arn}.
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

