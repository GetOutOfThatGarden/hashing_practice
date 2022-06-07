package main

// I'm practicing this code thorugh this tutorial: https://www.developer.com/languages/cryptography-in-go/
// Recently I learned by there are / in the package that are imported.
// I the below packages, the crypto/aes package referers to the aes directory whitin the crypto package.
// I think this is a way of saving space, so that I don't need to import every directory within the crypto package.
import (
	"crypto/aes"
	"fmt"
)

func encryptMessage(key string, message string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	// I just learned that the below 'make()' function is used to create a dynamic array.
	// The 'len()' callback funtion within the make() function returns the length of the string passed in to it as an integer.
	// It is another  built in function.
	msgByte := make([]byte, len(message))
	// Right now, I don't know what the '[]byte' thing is.
	// Is it an array that was set up with the make() fucnction from the line above?
	// I'm a bit confused about the Encrypt function below...
	// It says it is part of the cipher.Block package in the desciption whenI hover over it...
	// But I did not import that package. Only fmt, crypto/aes and crypto/hex.
	// How do I figure this out? If I try to run this, will it import that package for me?
	c.Encrypt(msgByte, []byte(message))

}
