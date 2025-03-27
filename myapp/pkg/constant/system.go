package constant

const (
	RequestID                       = "RequestID"
	FC_REGION                       = "FC_REGION"
	ALIBABA_CLOUD_ACCESS_KEY_ID     = "ALIBABA_CLOUD_ACCESS_KEY_ID"
	ALIBABA_CLOUD_ACCESS_KEY_SECRET = "ALIBABA_CLOUD_ACCESS_KEY_SECRET"
	ALIBABA_CLOUD_SECURITY_TOKEN    = "ALIBABA_CLOUD_SECURITY_TOKEN"
	APP_REDIS_URL                   = "APP_REDIS_URL"
	LOG_LEVEL                       = "LOG_LEVEL"
	SEND_TO_CS_ALERT                = "SEND_TO_CS_ALERT"
	STAT_DOMAIN_HASH_KEEP_DURANG    = 60 * 60 * 24 * 7 // 7day, unit: sec
	STAT_SESSION_DURANG_SEC         = 15 * 60          // 15mins, unit: sec
	SENT_ALERT_INTERVAL_SEC         = 15 * 60
)
