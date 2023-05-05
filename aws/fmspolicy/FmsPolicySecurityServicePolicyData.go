package fmspolicy


type FmsPolicySecurityServicePolicyData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/fms_policy#type FmsPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/fms_policy#managed_service_data FmsPolicy#managed_service_data}.
	ManagedServiceData *string `field:"optional" json:"managedServiceData" yaml:"managedServiceData"`
}

