package container

type Container struct {
	Adapter        Adapter
	Repository     Repository
	SupportService SupportService
}

// Adapter : Adapters hold resolved adapter instances.
// These are wrappers around third party libraries.
//All adapters will be of a corresponding adapter interface type.
type Adapter struct {
}

// Repository :Repositories hold resolved repository instances.
// These are used to connect to databases. All repositories will be of a corresponding repository interface type.
// They act as an abstraction layer between the application and the database.
// Generally a single repository represents a single table in the database along with all the actions
// that can be performed on that table.
type Repository struct {
}

// SupportService :Services hold resolved service instances.
// These are abstractions to third party APIs. All services will be
//of a corresponding service interface type.
type SupportService struct {
}
