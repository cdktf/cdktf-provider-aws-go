package codepipeline


type CodepipelineWebhookAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codepipeline_webhook#allowed_ip_range CodepipelineWebhook#allowed_ip_range}.
	AllowedIpRange *string `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codepipeline_webhook#secret_token CodepipelineWebhook#secret_token}.
	SecretToken *string `field:"optional" json:"secretToken" yaml:"secretToken"`
}
