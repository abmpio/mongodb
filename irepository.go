package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 一个抽象的用来处理任意类型的mongodb的仓储基类
type IRepository interface {
	IEntityFind
	IEntityCreate
	IEntityUpdate
	IEntityDelete
	IEntityIndex

	// aggregate
	Aggregate(pipeline interface{}, dataList interface{}, opts ...AggregateOption) (err error)

	// replace*
	ReplaceById(id primitive.ObjectID, doc interface{}, opts ...*options.ReplaceOptions) (err error)
	Replace(filter interface{}, doc interface{}, opts ...*options.ReplaceOptions) (err error)

	GetName() (name string)
	GetCollection() (c *mongo.Collection)
}

type IEntityFind interface {
	CountByFilter(filter interface{}) (count int64, err error)

	// find
	FindAll(opts ...FindOption) IFindResult
	FindByObjectId(id primitive.ObjectID) IFindResult
	FindOne(filter interface{}, opts ...FindOneOption) IFindResult
	FindByFilter(filter interface{}, opts ...FindOption) IFindResult
}

type IEntityCreate interface {
	// create
	Create(data interface{}, opts ...*options.InsertOneOptions) (id primitive.ObjectID, err error)
	CreateMany(itemList []interface{}, opts ...*options.InsertManyOptions) (ids []primitive.ObjectID, err error)
}

// update
type IEntityUpdate interface {
	FindOneAndUpdate(entity IEntity, opts ...*options.FindOneAndUpdateOptions) error
	FindOneAndUpdateWithId(objectId primitive.ObjectID, update interface{}, opts ...*options.FindOneAndUpdateOptions) error
	UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) error
	UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) error
}

type IEntityDelete interface {
	// delete
	DeleteOne(id primitive.ObjectID, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteOneByFilter(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type IEntityIndex interface {
	// index
	CreateIndex(indexDefine EntityIndexDefine, indexOptions *options.IndexOptions) (string, error)
	CreateIndexes(indexDefineList []EntityIndexDefine, indexOptions *options.IndexOptions) ([]string, error)
	MustCreateIndex(indexDefine EntityIndexDefine, indexOptions *options.IndexOptions)
	MustCreateIndexes(indexDefineList []EntityIndexDefine, indexOptions *options.IndexOptions)
	DeleteIndex(name string) (err error)
	DeleteAllIndexes() (err error)
	ListIndexes() (indexes []map[string]interface{}, err error)
}
