module demo

go 1.19

replace github.com/kaliazur/demo-clean-architecture/v2 => ./

require (
	github.com/aws/aws-sdk-go v1.44.306
	github.com/cucumber/godog v0.12.6
	github.com/google/uuid v1.3.0
	github.com/kaliazur/demo-clean-architecture/v2 v2.0.0-00010101000000-000000000000
)

require (
	github.com/cucumber/gherkin-go/v19 v19.0.3 // indirect
	github.com/cucumber/messages-go/v16 v16.0.1 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
