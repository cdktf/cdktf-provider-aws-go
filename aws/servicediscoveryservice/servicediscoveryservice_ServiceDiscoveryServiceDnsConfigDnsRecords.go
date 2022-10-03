package servicediscoveryservice


type ServiceDiscoveryServiceDnsConfigDnsRecords struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_service#ttl ServiceDiscoveryService#ttl}.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_service#type ServiceDiscoveryService#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

