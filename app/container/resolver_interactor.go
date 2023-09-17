package container

type Resolver interface {
	Resolve() Container
}
