package vault

type KMSManager interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

// Passthrough encryption (no encryption, that is)

func NewPassthroughKeyManager() *PassthroughKeyManager {
	return &PassthroughKeyManager{}
}

type PassthroughKeyManager struct{}

func (k PassthroughKeyManager) Encrypt(in []byte) ([]byte, error) { return in, nil }
func (k PassthroughKeyManager) Decrypt(in []byte) ([]byte, error) { return in, nil }
