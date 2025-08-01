// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubernetesNetworkConfig struct {
	// elastic_load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/eks_cluster#elastic_load_balancing EksCluster#elastic_load_balancing}
	ElasticLoadBalancing *EksClusterKubernetesNetworkConfigElasticLoadBalancing `field:"optional" json:"elasticLoadBalancing" yaml:"elasticLoadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/eks_cluster#ip_family EksCluster#ip_family}.
	IpFamily *string `field:"optional" json:"ipFamily" yaml:"ipFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/eks_cluster#service_ipv4_cidr EksCluster#service_ipv4_cidr}.
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
}

