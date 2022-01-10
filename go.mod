module github.com/Tackem-org/User

go 1.17

require (
	github.com/Tackem-org/Global v0.0.0-00010101000000-000000000000
	github.com/Tackem-org/Proto v0.0.0-00010101000000-000000000000
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.42.0
	gorm.io/driver/sqlite v1.2.6
	gorm.io/gorm v1.22.4
)

require github.com/xhit/go-str2duration/v2 v2.0.0 // indirect

require (
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace (
	github.com/Tackem-org/Global => ../Global
	github.com/Tackem-org/Proto => ../Proto
)
