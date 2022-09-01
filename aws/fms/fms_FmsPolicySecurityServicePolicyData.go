package fms


type FmsPolicySecurityServicePolicyData struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy#type FmsPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fms_policy#managed_service_data FmsPolicy#managed_service_data}.
	ManagedServiceData *string `field:"optional" json:"managedServiceData" yaml:"managedServiceData"`
}

