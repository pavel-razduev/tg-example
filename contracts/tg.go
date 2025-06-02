// @tg version=v1.0.0
// @tg title=`some API`
// @tg packageJSON=github.com/goccy/go-json
// @tg description=`A service which provide some API`
// @tg servers=`https://some-service.stage;stage|https://some-service.prod;prod`
//
//go:generate tg transport --services . --ifaces Service1  --out ../internal/transport/service1
//go:generate tg transport --services . --ifaces Service2 --out ../internal/transport/service2

//go:generate tg client -go --services . --ifaces Service1 --outPath ../pkg/clients/service1
//go:generate tg client -go --services . --ifaces Service2 --outPath ../pkg/clients/service2

//go:generate tg swagger --services . --ifaces Service1 --outFile ../api/service1-swagger.yaml
//go:generate tg swagger --services . --ifaces Service2 --outFile ../api/service2-swagger.yaml

//go:generate goimports -l -w ../pkg/clients
//go:generate goimports -l -w ../internal/transport/service1
//go:generate goimports -l -w ../internal/transport/service2
package contracts
