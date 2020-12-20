package global

const (
	//dburi       = "mongodb+srv://overstandman:247_Secured@cluster0-4newh.mongodb.net/protoblogdb?retryWrites=true&w=majority"
	dbUrl       = "mongodb://localhost:27017"
	dbname      = "protoblogdb"
	performance = 100
)

// []byte cant be constant
var (
	jwtSecret = []byte("superuwaifo")
)
