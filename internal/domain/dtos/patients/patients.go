package patients

type SearchRequest struct {
	Name  string `uri:"name"`
	Email string `uri:"email"`
}

type SearchReply struct {
	Patients []*Patient
}

type GetRequest struct {
	ID int `uri:"id"`
}

type CreateRequest struct {
	Name  string
	Email string
}

type UpdateRequest struct {
	ID    int `uri:"id"`
	Name  string
	Email string
}

type Patient struct {
	ID    int
	Name  string
	Email string
}

type DeleteRequest struct {
	ID int `uri:"id"`
}

type Empty struct{}
