package opensearchdomain


type OpensearchDomainAutoTuneOptionsMaintenanceScheduleDuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opensearch_domain#unit OpensearchDomain#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opensearch_domain#value OpensearchDomain#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

