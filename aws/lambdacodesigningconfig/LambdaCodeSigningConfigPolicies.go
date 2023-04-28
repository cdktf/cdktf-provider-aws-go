package lambdacodesigningconfig


type LambdaCodeSigningConfigPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/lambda_code_signing_config#untrusted_artifact_on_deployment LambdaCodeSigningConfig#untrusted_artifact_on_deployment}.
	UntrustedArtifactOnDeployment *string `field:"required" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

