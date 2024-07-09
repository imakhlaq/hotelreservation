package types

//bson:"_id" => what should be db colum name
//json:"id" => for json conversion
//json:"id,omitempty" => ignore if field is empty
//json:"-" => ignore field while JSON marshaling
//bson:"_id,omitempty =>Because we want mongodb to create the id for us

type User struct {
	ID        string `bson:"_id,omitempty" json:"-"`
	FirstName string `bson:"first_name" json:"firstName"`
	LastName  string `bson:"last_name" json:"lastName"`
}
