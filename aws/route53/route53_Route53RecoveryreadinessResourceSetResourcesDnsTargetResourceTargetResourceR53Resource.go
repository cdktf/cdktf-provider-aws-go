package route53


type Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53recoveryreadiness_resource_set#domain_name Route53RecoveryreadinessResourceSet#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53recoveryreadiness_resource_set#record_set_id Route53RecoveryreadinessResourceSet#record_set_id}.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
}

