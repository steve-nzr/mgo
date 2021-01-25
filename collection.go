package mgo

import "github.com/steve-nzr/mgo/bson"

// Collection stores documents
//
// Relevant documentation:
//
//    https://docs.mongodb.com/manual/core/databases-and-collections/#collections
//
type Collection interface {
	Name() string
	FullName() string
	Database() Database

	Bulk() *Bulk
	With(s *Session) Collection
	EnsureIndexKey(key ...string) error
	EnsureIndex(index Index) error
	DropIndex(key ...string) error
	DropIndexName(name string) error
	DropAllIndexes() error
	Indexes() (indexes []Index, err error)
	Find(query interface{}) *Query
	Repair() *Iter
	FindId(id interface{}) *Query
	Pipe(pipeline interface{}) *Pipe
	NewIter(session *Session, firstBatch []bson.Raw, cursorId int64, err error) *Iter
	Insert(docs ...interface{}) error
	Update(selector interface{}, update interface{}) error
	UpdateId(id interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error)
	Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error)
	UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error)
	Remove(selector interface{}) error
	RemoveId(id interface{}) error
	RemoveAll(selector interface{}) (info *ChangeInfo, err error)
	DropCollection() error
	Create(info *CollectionInfo) error
	Count() (n int, err error)
	Watch(pipeline interface{},
		options ChangeStreamOptions) (*ChangeStream, error)
}
