package globalacceleratorendpointgroup


type GlobalacceleratorEndpointGroupPortOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/globalaccelerator_endpoint_group#endpoint_port GlobalacceleratorEndpointGroup#endpoint_port}.
	EndpointPort *float64 `field:"required" json:"endpointPort" yaml:"endpointPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/globalaccelerator_endpoint_group#listener_port GlobalacceleratorEndpointGroup#listener_port}.
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
}

