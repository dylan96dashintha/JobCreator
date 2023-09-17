package container

var supportServiceResolverObject Resolver

type supportServiceResolver struct {
}

func newSupportServiceResolverObj() Resolver {
	supportServiceResolverObj := new(supportServiceResolver)
	return supportServiceResolverObj
}

func GetSupportServiceResolverObj() Resolver {
	if supportServiceResolverObject == nil {
		supportServiceResolverObject = newSupportServiceResolverObj()
	}
	return supportServiceResolverObject
}

func (s *supportServiceResolver) Resolve() Container {
	supportService := Container{
		SupportService: SupportService{},
	}
	return supportService
}
