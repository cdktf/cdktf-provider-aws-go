package connectquickconnect


type ConnectQuickConnectQuickConnectConfigUserConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/connect_quick_connect#contact_flow_id ConnectQuickConnect#contact_flow_id}.
	ContactFlowId *string `field:"required" json:"contactFlowId" yaml:"contactFlowId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/connect_quick_connect#user_id ConnectQuickConnect#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

