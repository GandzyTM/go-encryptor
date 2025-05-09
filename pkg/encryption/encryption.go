package encryption

import (
	"gitlab.pikabiduskibidi.ru/box/go-encryption/internal/encryption"
	"gitlab.pikabiduskibidi.ru/box/go-encryption/internal/interfaces"
	"gitlab.pikabiduskibidi.ru/box/go-encryption/internal/sensitive"
	"gitlab.pikabiduskibidi.ru/box/go-encryption/pkg/config"
)

// Encryptor предоставляет публичный API для шифрования
type Encryptor struct {
	encryptor interfaces.Encryptor
	handler   interfaces.FieldEncryptor
}

// NewEncryptor создает новый экземпляр Encryptor
func NewEncryptor(cfg *config.Config) (*Encryptor, error) {
	provider := encryption.NewEncryptorProvider()
	enc, err := provider.ProvideEncryptor(cfg)
	if err != nil {
		return nil, err
	}

	handler := sensitive.NewFieldEncryptor(enc)

	return &Encryptor{
		encryptor: enc,
		handler:   handler,
	}, nil
}

// EncryptString шифрует строку
func (e *Encryptor) EncryptString(data string) (string, error) {
	return e.encryptor.Encrypt(data)
}

// DecryptString расшифровывает строку
func (e *Encryptor) DecryptString(data string) (string, error) {
	return e.encryptor.Decrypt(data)
}

// EncryptFields шифрует поля в структуре, помеченные тегом encrypted:"true"
func (e *Encryptor) EncryptFields(data interface{}) error {
	return e.handler.HandleFields(data, true)
}

// DecryptFields расшифровывает поля в структуре, помеченные тегом encrypted:"true"
func (e *Encryptor) DecryptFields(data interface{}) error {
	return e.handler.HandleFields(data, false)
}
