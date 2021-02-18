package pages

// Object is a theoretically usable interface
// that can be used to represent every type of
// page in the SSG
type Object interface {
	ID() string
	String() string
}
