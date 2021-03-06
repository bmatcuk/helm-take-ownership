module github.com/bmatcuk/helm-take-ownership

go 1.13

require (
	cloud.google.com/go v0.1.1-0.20160913182117-3b1ae45394a2
	github.com/Azure/go-autorest v8.0.0+incompatible
	github.com/BurntSushi/toml v0.3.0
	github.com/Masterminds/semver v1.3.1
	github.com/Masterminds/sprig v2.14.1+incompatible
	github.com/PuerkitoBio/purell v1.0.0
	github.com/PuerkitoBio/urlesc v0.0.0-20160726150825-5bd2802263f2
	github.com/aokoli/goutils v1.0.1
	github.com/beorn7/perks v0.0.0-20160229213445-3ac7bf7a47d1
	github.com/davecgh/go-spew v1.1.1-0.20170626231645-782f4967f2dc
	github.com/dgrijalva/jwt-go v3.0.1-0.20160705203006-01aeca54ebda+incompatible
	github.com/docker/distribution v2.6.0-rc.1.0.20170726174610-edc3ab29cdff+incompatible
	github.com/docker/docker v1.4.2-0.20170731201938-4f3616fb1c11
	github.com/docker/go-connections v0.3.0
	github.com/docker/go-units v0.3.2-0.20170127094116-9e638d38cf69
	github.com/docker/spdystream v0.0.0-20160310174837-449fdfce4d96
	github.com/emicklei/go-restful v1.1.4-0.20170410110728-ff4f55a20633
	github.com/emicklei/go-restful-swagger12 v0.0.0-20170208215640-dcef7f557305
	github.com/evanphx/json-patch v0.0.0-20170719203123-944e07253867
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d
	github.com/fatih/camelcase v0.0.0-20160318181535-f6a740d52f96
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/jsonpointer v0.0.0-20160704185906-46af16f9f7b1
	github.com/go-openapi/jsonreference v0.0.0-20160704190145-13c6e3589ad9
	github.com/go-openapi/spec v0.0.0-20160808142527-6aced65f8501
	github.com/go-openapi/swag v0.0.0-20160704191624-1d0bd113de87
	github.com/gobwas/glob v0.2.2
	github.com/gogo/protobuf v0.0.0-20170330071051-c0656edd0d9e
	github.com/golang/glog v0.0.0-20141105023935-44145f04b68c
	github.com/golang/groupcache v0.0.0-20160516000752-02826c3e7903
	github.com/golang/protobuf v0.0.0-20161109072736-4bd1920723d7
	github.com/google/btree v0.0.0-20160524151835-7d79101e329e
	github.com/google/gofuzz v0.0.0-20161122191042-44d81051d367
	github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d
	github.com/gophercloud/gophercloud v0.0.0-20170831144856-2bf16b94fdd9
	github.com/gregjones/httpcache v0.0.0-20170728041850-787624de3eb7
	github.com/hashicorp/golang-lru v0.0.0-20160207214719-a0d98a5f2880
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/huandu/xstrings v0.0.0-20171208101919-37469d0c81a7
	github.com/imdario/mergo v0.0.0-20141206190957-6633656539c1
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/json-iterator/go v0.0.0-20170829155851-36b14963da70
	github.com/juju/ratelimit v0.0.0-20170523012141-5b9ff8664717
	github.com/mailru/easyjson v0.0.0-20160728113105-d5b7844b561a
	github.com/matttproud/golang_protobuf_extensions v0.0.0-20150406173934-fc2b8d3a73c4
	github.com/opencontainers/go-digest v0.0.0-20170106003457-a6d0ee40d420
	github.com/opencontainers/image-spec v1.0.0-rc6.0.20170604055404-372ad780f634
	github.com/pborman/uuid v0.0.0-20150603214016-ca53cad383ca
	github.com/peterbourgon/diskv v2.0.1+incompatible
	github.com/prometheus/client_golang v0.8.0
	github.com/prometheus/client_model v0.0.0-20150212101744-fa8ad6fec335
	github.com/prometheus/common v0.0.0-20170427095455-13ba4ddd0caa
	github.com/prometheus/procfs v0.0.0-20170519190837-65c1f6f8f0fc
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v0.0.1
	github.com/spf13/pflag v1.0.1-0.20171106142849-4c012f6dcd95
	golang.org/x/crypto v0.0.0-20170825220121-81e90905daef
	golang.org/x/net v0.0.0-20170809000501-1c05540f6879
	golang.org/x/oauth2 v0.0.0-20170412232759-a6bd8cefa181
	golang.org/x/sys v0.0.0-20170901181214-7ddbeae9ae08
	golang.org/x/text v0.0.0-20170810154203-b19bf474d317
	google.golang.org/appengine v1.0.1-0.20171212223047-5bee14b453b4
	gopkg.in/inf.v0 v0.9.0
	gopkg.in/yaml.v2 v2.0.0
	k8s.io/api v0.0.0-20170922112058-fe29995db376
	k8s.io/apiextensions-apiserver v0.0.0-20180118134117-0b14e80bb6cc
	k8s.io/apimachinery v0.0.0-20170925234155-019ae5ada31d
	k8s.io/apiserver v0.0.0-20180116185515-74bf92070805
	k8s.io/client-go v5.0.0+incompatible
	k8s.io/helm v2.7.2+incompatible
	k8s.io/kube-openapi v0.0.0-20170830100654-868f2f29720b
	k8s.io/kubernetes v1.8.0
	k8s.io/metrics v0.0.0-20180118134523-78dff7e0cde0
	k8s.io/utils v0.0.0-20170719031128-9fdc871a36f3
	vbom.ml/util v0.0.0-20170409195630-256737ac55c4
)
