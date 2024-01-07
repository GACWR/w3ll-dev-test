package types

const (
	// ModuleName defines the module name
	ModuleName = "will"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_will"
)

var (
	ParamsKey = []byte("p_will")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
