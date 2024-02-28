module example.com/p2p_parser_jq

go 1.20

require (
	example.com/p2p_parser v0.0.0-00010101000000-000000000000
	github.com/jnovack/flag v1.16.0
	github.com/prometheus/client_golang v1.16.0
	github.com/sirupsen/logrus v1.9.3

)

require (
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/AthenZ/athenz v1.10.39 // indirect
	github.com/DataDog/zstd v1.5.0 // indirect
	github.com/apache/pulsar-client-go v0.10.0 // indirect
	github.com/ardielle/ardielle-go v1.5.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.4.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/itchyny/timefmt-go v0.1.5 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/linkedin/goavro/v2 v2.9.8 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

require (
	github.com/itchyny/gojq v0.12.13
	golang.org/x/sys v0.8.0 // indirect
	example.com/gojq_extention v0.0.0-00010101000000-000000000000
)

replace example.com/p2p_parser => ../p2p_parser
replace example.com/gojq_extention => ../gojq_extention