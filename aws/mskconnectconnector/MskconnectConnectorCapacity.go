package mskconnectconnector


type MskconnectConnectorCapacity struct {
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#autoscaling MskconnectConnector#autoscaling}
	Autoscaling *MskconnectConnectorCapacityAutoscaling `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// provisioned_capacity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#provisioned_capacity MskconnectConnector#provisioned_capacity}
	ProvisionedCapacity *MskconnectConnectorCapacityProvisionedCapacity `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

