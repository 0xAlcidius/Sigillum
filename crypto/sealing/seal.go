package sealing

type ExecuteSealInterface interface {
	ExecuteSeal(key []byte, ciphertext []byte) ([]byte, error)
}
