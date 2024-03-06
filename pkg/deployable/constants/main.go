package constants

type CTX_ATTRIBUTES string

const (
	IS_DEPLOYABLE_CTX CTX_ATTRIBUTES = "is_deployable"
	DEPLOYABLE_REF    CTX_ATTRIBUTES = "deployable_ref"
	LOGGER_REF        CTX_ATTRIBUTES = "logger_ref"
	CONTROLLABLE_TYPE CTX_ATTRIBUTES = "controllable_type"
	CONTROLLABLE_NAME CTX_ATTRIBUTES = "controller_name"
	CTX_ID            CTX_ATTRIBUTES = "ctx_id"
)

type CONTROLLABLE_TYPES string

const (
	HTTP_SERVER CONTROLLABLE_TYPES = "HTTP_SERVER"
)
