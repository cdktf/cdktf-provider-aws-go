package vpc


type VpcEndpointDnsOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint#dns_record_ip_type VpcEndpoint#dns_record_ip_type}.
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
}

