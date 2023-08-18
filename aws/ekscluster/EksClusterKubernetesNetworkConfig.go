package ekscluster


type EksClusterKubernetesNetworkConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/eks_cluster#ip_family EksCluster#ip_family}.
	IpFamily *string `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/eks_cluster#service_ipv4_cidr EksCluster#service_ipv4_cidr}.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
}

