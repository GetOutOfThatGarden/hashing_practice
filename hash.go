package main

// I'm practicing this code thorugh this tutorial: https://www.developer.com/languages/cryptography-in-go/
// Recently I learned by there are / in the package that are imported.
// I the below packages, the crypto/aes package referers to the aes directory whitin the crypto package.
// I think this is a way of saving space, so that I don't need to import every directory within the crypto package.
import (
	"crypto/aes"
	"encoding/hex"
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
	// I think the below line is converting the byte array into a string.(GCP). It would appear it is a string being returned.
	return hex.EncodeToString(msgByte)
}

func decryptMessage(key string, message string) string {
	// I'm not sure what the underscore _ is on the line below is for. Is it an err? The hex.DecodeString() function apparently "returns the byte repesented by the hexadecimal string 's'."
	txt, _ := hex.DecodeString(message)
	// I'm still getting my head around the declaration of a second variable in declarations like below. I know err is related to errors, but is it a variable? Is it built in variable? I guess the function returns two variables, so the second variable it returns is the error.
	c, err := aes.NewCipher([]byte(key))
	// I presume the the below line is checking to see if an error exists? If the error is not an empty variable, then it will print the error.
	if err != nil {
		fmt.Println(err)
	}
	// Below I see that the make() function is used again, like in the previous function. The first argument is the length of the array. (GCP) And the second is the len() function is returning an int that is the length of the string passed in.
	msgByte := make([]byte, len(txt))
	// Below the Decrypt function is being called on the c variable declared above.
	c.Decrypt(msgByte, []byte(txt))

	// Ok, in the below line, it looks like he is declaring a variable named 'msg'. I see he is using := (walrus operator) to declare the variable. He is specifying the type of the variable, which i didn't think was necessary when using a walrus operator. Maybe its only because of the complex code after it? What does (msgByte[:]) mean? Is it the value of the array at that specific index? What exactly is : in the square brackets for?
	msg := string(msgByte[:])
	// This function returns the variable called msg, which is a string.
	return msg
}

func main() {
	plainText := "This is a secret"
	key := "this_must_be_of_32_byte_length!!"

	// Ok, below we are calling on the functions we made earlier in the code. The arguments to the first function is declared in the lines above.
	emsg := encryptMessage(key, plainText)
	// And for the below function, the second argument is the the emsg variable in the function above.
	dmesg := decryptMessage(key, emsg)

	// Below are the encrypted and decrypted messages that are being printed to terminal
	fmt.Println("Encrypted Message is: ", emsg)
	fmt.Println("Decrypted Message is: ", dmesg)
	.
}
