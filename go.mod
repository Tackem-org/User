module github.com/Tackem-org/User

go 1.17

require (
	github.com/Tackem-org/Global v0.0.0-00010101000000-000000000000
	github.com/Tackem-org/Proto v0.0.0-00010101000000-000000000000
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce
	google.golang.org/grpc v1.43.0
	gorm.io/driver/sqlite v1.2.6
	gorm.io/gorm v1.22.5
)

require github.com/xhit/go-str2duration/v2 v2.0.0 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace (
	github.com/Tackem-org/Global => ../Global
	github.com/Tackem-org/Proto => ../Proto
)
