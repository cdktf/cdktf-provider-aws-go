package kendraexperience


type KendraExperienceConfigurationContentSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#data_source_ids KendraExperience#data_source_ids}.
	DataSourceIds *[]*string `field:"optional" json:"dataSourceIds" yaml:"dataSourceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#direct_put_content KendraExperience#direct_put_content}.
	DirectPutContent interface{} `field:"optional" json:"directPutContent" yaml:"directPutContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#faq_ids KendraExperience#faq_ids}.
	FaqIds *[]*string `field:"optional" json:"faqIds" yaml:"faqIds"`
}

