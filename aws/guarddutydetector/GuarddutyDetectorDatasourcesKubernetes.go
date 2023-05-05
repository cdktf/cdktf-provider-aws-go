package guarddutydetector


type GuarddutyDetectorDatasourcesKubernetes struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/guardduty_detector#audit_logs GuarddutyDetector#audit_logs}
	AuditLogs *GuarddutyDetectorDatasourcesKubernetesAuditLogs `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

