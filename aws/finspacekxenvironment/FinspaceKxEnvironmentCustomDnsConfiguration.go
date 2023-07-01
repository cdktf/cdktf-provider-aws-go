package finspacekxenvironment


type FinspaceKxEnvironmentCustomDnsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/finspace_kx_environment#custom_dns_server_ip FinspaceKxEnvironment#custom_dns_server_ip}.
	CustomDnsServerIp *string `field:"required" json:"customDnsServerIp" yaml:"customDnsServerIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/finspace_kx_environment#custom_dns_server_name FinspaceKxEnvironment#custom_dns_server_name}.
	CustomDnsServerName *string `field:"required" json:"customDnsServerName" yaml:"customDnsServerName"`
}

