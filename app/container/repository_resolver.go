package container

var repositoryResolverObject Resolver

type repositoryResolver struct {
}

func newRepositoryResolverObj() Resolver {
	repositoryObj := new(repositoryResolver)
	return repositoryObj
}

func GetRepositoryResolverObj() Resolver {
	if repositoryResolverObject == nil {
		repositoryResolverObject = newRepositoryResolverObj()
	}
	return repositoryResolverObject
}

func (r *repositoryResolver) Resolve() Container {
	repository := Container{
		Repository: Repository{},
	}
	return repository
}
