// @tg version=v1.0.0
// @tg title=`some API`
// @tg packageJSON=github.com/goccy/go-json
// @tg description=`A service which provide some API`
// @tg servers=`https://some-service.stage;stage|https://some-service.prod;prod`
//
//go:generate tg transport --services . --out ../internal/transport
//go:generate tg client -go --services . --outPath ../pkg/clients/example
//go:generate tg swagger --services . --outFile ../api/service-swagger.yaml
//go:generate goimports -l -w ../internal/transport
//go:generate goimports -l -w ../pkg/clients/example
package contracts
