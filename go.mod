module github.com/company/go-enterprise-template

go 1.22

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/spf13/viper v1.18.2
	github.com/sirupsen/logrus v1.9.3
	gorm.io/gorm v1.25.5
	gorm.io/driver/mysql v1.5.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/files v1.0.1
	github.com/swaggo/swag v1.16.2
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/go-playground/validator/v10 v10.16.0
	github.com/prometheus/client_golang v1.17.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	github.com/elastic/go-elasticsearch/v8 v8.11.1
	github.com/golang-migrate/migrate/v4 v4.17.0
	golang.org/x/crypto v0.17.0
	github.com/google/uuid v1.5.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)