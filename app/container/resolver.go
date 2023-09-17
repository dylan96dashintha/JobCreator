package container

// Resolve resolves the entire container.
//
// The order of resolution is very important.
//Low level dependencies need to be resolved before high level dependencies.
// It generally happens in this order.
// 		- Adapters
// 		- Repositories
// 		- Services
func Resolve() *Container {

	adapterResolverObj := GetAdapterResolverObj()
	repositoryResolverObj := GetRepositoryResolverObj()
	supportServiceResolverObj := GetSupportServiceResolverObj()

	return &Container{
		Adapter:        adapterResolverObj.Resolve().Adapter,
		Repository:     repositoryResolverObj.Resolve().Repository,
		SupportService: supportServiceResolverObj.Resolve().SupportService,
	}
}
