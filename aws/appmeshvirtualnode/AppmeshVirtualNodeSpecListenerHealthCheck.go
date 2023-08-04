package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerHealthCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_virtual_node#healthy_threshold AppmeshVirtualNode#healthy_threshold}.
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_virtual_node#interval_millis AppmeshVirtualNode#interval_millis}.
	IntervalMillis *float64 `field:"required" json:"intervalMillis" yaml:"intervalMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_virtual_node#protocol AppmeshVirtualNode#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_virtual_node#timeout_millis AppmeshVirtualNode#timeout_millis}.
	TimeoutMillis *float64 `field:"required" json:"timeoutMillis" yaml:"timeoutMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_virtual_node#unhealthy_threshold AppmeshVirtualNode#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_virtual_node#path AppmeshVirtualNode#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/appmesh_virtual_node#port AppmeshVirtualNode#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

