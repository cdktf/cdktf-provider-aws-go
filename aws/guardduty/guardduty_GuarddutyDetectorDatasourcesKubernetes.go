package guardduty


type GuarddutyDetectorDatasourcesKubernetes struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/guardduty_detector#audit_logs GuarddutyDetector#audit_logs}
	AuditLogs *GuarddutyDetectorDatasourcesKubernetesAuditLogs `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

