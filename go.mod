module github.com/projectriff/cli

go 1.12

require (
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d
	github.com/boz/go-logutil v0.1.0
	github.com/boz/kail v0.10.1
	github.com/buildpack/pack v0.2.0
	github.com/fatih/color v1.7.0
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.3.0
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/knative/pkg v0.0.0-20190624141606-d82505e6c5b4
	github.com/mitchellh/go-homedir v1.1.0
	github.com/projectriff/system v0.0.0-20190627032734-afb8a6d9901d
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8
	k8s.io/api v0.0.0-20190627205229-acea843d18eb
	k8s.io/apimachinery v0.0.0-20190627205106-bc5732d141a8
	k8s.io/client-go v0.0.0-20190627205436-11059204e07c
)

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.10.2 // indirect
	github.com/Azure/go-autorest v0.0.0-20190226174127-bca49d5b51a5 // indirect
	github.com/go-logr/logr v0.1.0 // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/knative/build v0.6.0 // indirect
	github.com/knative/serving v0.6.0 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/trisberg/service v0.0.0-20190705200517-f4cbd53354b4
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/apiextensions-apiserver v0.0.0-20190627210706-e1f2db1c2108 // indirect
	k8s.io/kube-openapi v0.0.0-20190502190224-411b2483e503 // indirect
	sigs.k8s.io/controller-runtime v0.1.12 // indirect
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
)

replace (
	// kail wants to upgrade these deps, don't let it
	k8s.io/api => k8s.io/api v0.0.0-20190226173710-145d52631d00
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190226174127-78295b709ec6
)
