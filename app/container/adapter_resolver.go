package container

var adapterResolverObject Resolver

type adapterResolver struct {
}

func newAdapterResolverObj() Resolver {
	adapterObj := new(adapterResolver)
	return adapterObj
}

func GetAdapterResolverObj() Resolver {
	if adapterResolverObject == nil {
		adapterResolverObject = newAdapterResolverObj()
	}
	return adapterResolverObject
}

func (a *adapterResolver) Resolve() Container {
	adapter := Container{
		Adapter: Adapter{},
	}
	return adapter
}

func resolveDBAdapter() {

}
