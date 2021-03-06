// Crypto operations designed to be slow (internally utilizing PBKDF2)
package slowcrypto

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

func WithPassword(password string) passwordCrypto {
	return passwordCrypto(password)
}

type passwordCrypto string

// envelope = <24 bytes of nonce> <ciphertext>
func (p passwordCrypto) Encrypt(plaintext []byte) ([]byte, error) {
	return p.encryptWithRandom(plaintext, rand.Reader)
}

func (p passwordCrypto) encryptWithRandom(plaintext []byte, random io.Reader) ([]byte, error) {
	// You must use a different nonce for each message you encrypt with the
	// same key. Since the nonce here is 192 bits long, a random value
	// provides a sufficiently small probability of repeats.
	var nonce [24]byte
	if _, err := io.ReadFull(random, nonce[:]); err != nil {
		return nil, err
	}

	// using Seal() nonce as PBKDF2 salt
	encryptionKey := passwordTo256BitEncryptionKey100k(string(p), nonce[:])

	nonceAndCiphertextEnvelope := secretbox.Seal(nonce[:], plaintext, &nonce, &encryptionKey)

	return nonceAndCiphertextEnvelope, nil
}

func (p passwordCrypto) Decrypt(nonceAndCiphertextEnvelope []byte) ([]byte, error) {
	// When you decrypt, you must use the same nonce and key you used to
	// encrypt the message. One way to achieve this is to store the nonce
	// alongside the encrypted message. Above, we stored the nonce in the first
	// 24 bytes of the encrypted text.
	// 24 bytes of nonce seems fine https://security.stackexchange.com/a/112592
	var decryptNonce [24]byte
	copy(decryptNonce[:], nonceAndCiphertextEnvelope[:24])

	// using Seal() nonce as PBKDF2 salt
	encryptionKey := passwordTo256BitEncryptionKey100k(string(p), decryptNonce[:])

	plaintextBytes, ok := secretbox.Open(nil, nonceAndCiphertextEnvelope[24:], &decryptNonce, &encryptionKey)
	if !ok {
		return nil, errors.New("decryption error. wrong password?")
	}

	return plaintextBytes, nil
}

func passwordTo256BitEncryptionKey100k(password string, salt []byte) [32]byte {
	var derivedKey [32]byte
	copy(derivedKey[:], Pbkdf2Sha256100kDerive([]byte(password), salt))

	return derivedKey
}

func Pbkdf2Sha256100kDerive(key []byte, salt []byte) []byte {
	return pbkdf2.Key(
		key,
		salt,
		100*1000,
		sha256.Size,
		sha256.New)
}
