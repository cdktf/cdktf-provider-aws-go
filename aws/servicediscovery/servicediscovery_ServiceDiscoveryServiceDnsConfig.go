package servicediscovery


type ServiceDiscoveryServiceDnsConfig struct {
	// dns_records block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_service#dns_records ServiceDiscoveryService#dns_records}
	DnsRecords interface{} `field:"required" json:"dnsRecords" yaml:"dnsRecords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_service#namespace_id ServiceDiscoveryService#namespace_id}.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_service#routing_policy ServiceDiscoveryService#routing_policy}.
	RoutingPolicy *string `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
}

