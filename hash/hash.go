package hash

var (
	DefaultHasher = NewBcrypt()
)

type Hasher interface {
	Make(value string) (string, error)
	Check(value, hash string) bool
	NeedsRehash(hash string) bool
}
