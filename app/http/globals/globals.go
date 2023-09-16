package globals

const SOURCE_KEY = "registrationSource"
const SOURCE_NPTA = "NPTA"
const (
	Bearer                        = "bearer"
	JWTTokenContextKey contextKey = "JWTToken"
)

var (
	JWTTokenSecret = []byte("secretkey")
)

const (
	UUID = "uuid"
)

type contextKey string
