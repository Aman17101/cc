package store

// type DBOperations struct{

// }
type DBOperations interface{

	SaveRecord(record interface{}) 
}