package cloudfront


type CloudfrontFieldLevelEncryptionConfigContentTypeProfileConfig struct {
	// content_type_profiles block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_field_level_encryption_config#content_type_profiles CloudfrontFieldLevelEncryptionConfig#content_type_profiles}
	ContentTypeProfiles *CloudfrontFieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfiles `field:"required" json:"contentTypeProfiles" yaml:"contentTypeProfiles"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_field_level_encryption_config#forward_when_content_type_is_unknown CloudfrontFieldLevelEncryptionConfig#forward_when_content_type_is_unknown}.
	ForwardWhenContentTypeIsUnknown interface{} `field:"required" json:"forwardWhenContentTypeIsUnknown" yaml:"forwardWhenContentTypeIsUnknown"`
}

