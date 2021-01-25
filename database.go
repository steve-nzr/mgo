package mgo

// Database holds collections of documents
//
// Relevant documentation:
//
//    https://docs.mongodb.com/manual/core/databases-and-collections/#databases
//
type Database interface {
	Name() string
	Session() *Session

	FindRef(ref *DBRef) *Query
	CreateView(view string, source string, pipeline interface{}, collation *Collation) error

	Run(cmd interface{}, result interface{}) error
	With(s *Session) Database
	GridFS(prefix string) *GridFS

	Login(user, pass string) error
	Logout()

	UpsertUser(user *User) error
	AddUser(username, password string, readOnly bool) error
	RemoveUser(user string) error
	DropDatabase() error

	C(name string) Collection
	CollectionNames() (names []string, err error)
}
