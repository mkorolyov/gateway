// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gateway

type Post struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Profile struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Posts     []*Post `json:"posts"`
}
