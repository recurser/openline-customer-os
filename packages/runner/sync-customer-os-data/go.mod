module github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data

go 1.20

//replace github.com/openline-ai/openline-customer-os/packages/server/events-processing-common => ../../server/events-processing-common
//replace github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module => ../../server/customer-os-common-module

require (
	github.com/caarlos0/env/v6 v6.10.1
	github.com/google/uuid v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/neo4j/neo4j-go-driver/v5 v5.14.0
	github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module v0.0.0-20231016163826-6e7df2192aae
	github.com/openline-ai/openline-customer-os/packages/server/events-processing-common v0.0.0-20231016163826-6e7df2192aae
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	golang.org/x/oauth2 v0.13.0
	golang.org/x/text v0.13.0
	google.golang.org/grpc v1.59.0
	gorm.io/driver/postgres v1.5.4
	gorm.io/gorm v1.25.5
	zgo.at/zcache v1.2.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
