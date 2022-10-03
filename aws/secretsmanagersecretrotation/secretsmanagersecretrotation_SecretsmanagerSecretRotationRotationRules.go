package secretsmanagersecretrotation


type SecretsmanagerSecretRotationRotationRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret_rotation#automatically_after_days SecretsmanagerSecretRotation#automatically_after_days}.
	AutomaticallyAfterDays *float64 `field:"required" json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
}

