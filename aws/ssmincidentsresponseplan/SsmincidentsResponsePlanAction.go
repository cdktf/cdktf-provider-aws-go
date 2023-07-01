package ssmincidentsresponseplan


type SsmincidentsResponsePlanAction struct {
	// ssm_automation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/ssmincidents_response_plan#ssm_automation SsmincidentsResponsePlan#ssm_automation}
	SsmAutomation interface{} `field:"optional" json:"ssmAutomation" yaml:"ssmAutomation"`
}

