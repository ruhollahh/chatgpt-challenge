package promptparam

type InsertIfNotExistRequest struct {
	ID      string
	Content string
}

type InsertIfNotExistResponse struct {
	Inserted bool
}
