package types

//bson:"_id" => for db table
//json:"id" => for json conversion
//json:"id,omitempty" => ignore if field is empty
//json:"-" => ignore field while JSON marshaling

type User struct {
	ID        string `bson:"_id" json:"-"`
	FirstName string `bson:"first_name" json:"firstName"`
	LastName  string `bson:"last_name" json:"lastName"`
}
