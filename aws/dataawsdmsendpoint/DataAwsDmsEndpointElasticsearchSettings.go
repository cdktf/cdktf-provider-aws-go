package dataawsdmsendpoint


type DataAwsDmsEndpointElasticsearchSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/dms_endpoint#endpoint_uri DataAwsDmsEndpoint#endpoint_uri}.
	EndpointUri *string `field:"required" json:"endpointUri" yaml:"endpointUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/data-sources/dms_endpoint#service_access_role_arn DataAwsDmsEndpoint#service_access_role_arn}.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

