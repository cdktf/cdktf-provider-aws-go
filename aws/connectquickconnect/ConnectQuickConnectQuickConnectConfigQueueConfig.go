package connectquickconnect


type ConnectQuickConnectQuickConnectConfigQueueConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/connect_quick_connect#contact_flow_id ConnectQuickConnect#contact_flow_id}.
	ContactFlowId *string `field:"required" json:"contactFlowId" yaml:"contactFlowId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/connect_quick_connect#queue_id ConnectQuickConnect#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

