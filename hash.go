// // package main

// // import (
// // 	"crypto/rand"
// // 	"crypto/sha512"
// // 	"encoding/base64"
// // )

// // // Generate 16 bytes randomly and securely using the
// // // Cryptographically secure pseudorandom number generator (CSPRNG)
// // // in the crypto.rand package
// // func generateRandomSalt(saltSize int) []byte {
// // 	var salt = make([]byte, saltSize)

// // 	_, err := rand.Read(salt[:])

// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	return salt
// // }

// // // Combine password and salt then hash them using the SHA-512
// // // hashing algorithm and then return the hashed password
// // // as a base64 encoded string
// // func hashPassword(password string, salt []byte) string {
// // 	// Convert password string to byte slice
// // 	var passwordBytes = []byte(password)

// // 	// Create sha-512 hasher
// // 	var sha512Hasher = sha512.New()

// // 	// Append salt to password
// // 	salt= []byte{4, 17 ,178, 104, 25, 207, 47, 37, 106, 40, 231, 24, 174, 36, 12, 104}
// // 	passwordBytes = append(passwordBytes, salt...)

// // 	// Write password bytes to the hasher
// // 	sha512Hasher.Write(passwordBytes)

// // 	// Get the SHA-512 hashed password
// // 	var hashedPasswordBytes = sha512Hasher.Sum(nil)

// // 	// Convert the hashed password to a base64 encoded string
// // 	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

// // 	return base64EncodedPasswordHash
// }