package secretsmanager


type SecretsmanagerSecretRotationRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret#automatically_after_days SecretsmanagerSecret#automatically_after_days}.
	AutomaticallyAfterDays *float64 `field:"required" json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
}

