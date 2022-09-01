package lambdafunction


type LambdaCodeSigningConfigPolicies struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_code_signing_config#untrusted_artifact_on_deployment LambdaCodeSigningConfig#untrusted_artifact_on_deployment}.
	UntrustedArtifactOnDeployment *string `field:"required" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

