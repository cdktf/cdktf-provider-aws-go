package kendraindex


type KendraIndexDocumentMetadataConfigurationUpdatesRelevance struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#duration KendraIndex#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#freshness KendraIndex#freshness}.
	Freshness interface{} `field:"optional" json:"freshness" yaml:"freshness"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#importance KendraIndex#importance}.
	Importance *float64 `field:"optional" json:"importance" yaml:"importance"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#rank_order KendraIndex#rank_order}.
	RankOrder *string `field:"optional" json:"rankOrder" yaml:"rankOrder"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#values_importance_map KendraIndex#values_importance_map}.
	ValuesImportanceMap *map[string]*float64 `field:"optional" json:"valuesImportanceMap" yaml:"valuesImportanceMap"`
}

