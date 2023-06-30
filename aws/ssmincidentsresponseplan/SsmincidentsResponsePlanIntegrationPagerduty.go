package ssmincidentsresponseplan


type SsmincidentsResponsePlanIntegrationPagerduty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/ssmincidents_response_plan#name SsmincidentsResponsePlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/ssmincidents_response_plan#secret_id SsmincidentsResponsePlan#secret_id}.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/ssmincidents_response_plan#service_id SsmincidentsResponsePlan#service_id}.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

