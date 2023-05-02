package constants

const (
	EnvGrpcPort                   = "GRPC_PORT"
	EnvValidationApiUrl           = "VALIDATION_API"
	EnvValidationApiKey           = "VALIDATION_API_KEY"
	EnvEventStoreConnectionString = "EVENT_STORE_CONNECTION_STRING"
	EnvJaegerHostPort             = "JAEGER_HOST_PORT"

	ConfigPath = "CONFIG_PATH"

	Yaml = "yaml"
	Tcp  = "tcp"

	GRPC     = "GRPC"
	SIZE     = "SIZE"
	URI      = "URI"
	STATUS   = "STATUS"
	HTTP     = "HTTP"
	ERROR    = "ERROR"
	METHOD   = "METHOD"
	METADATA = "METADATA"
	REQUEST  = "REQUEST"
	REPLY    = "REPLY"
	TIME     = "TIME"

	Topic        = "topic"
	Partition    = "partition"
	Message      = "message"
	WorkerID     = "workerID"
	Offset       = "offset"
	Time         = "time"
	GroupName    = "GroupName"
	StreamID     = "StreamID"
	EventID      = "EventID"
	EventType    = "EventType"
	EventNumber  = "EventNumber"
	CreatedDate  = "CreatedDate"
	UserMetadata = "UserMetadata"

	GraphProjection        = "(Graph Projection)"
	DataEnricherProjection = "(Data Enricher Projection)"

	Validate        = "validate"
	FieldValidation = "field validation"
	RequiredHeaders = "required header"
	Base64          = "base64"
	Unmarshal       = "unmarshal"
	Uuid            = "uuid"
	Cookie          = "cookie"
	Token           = "token"
	Bcrypt          = "bcrypt"
	Redis           = "redis"

	EsAll = "$all"
)
