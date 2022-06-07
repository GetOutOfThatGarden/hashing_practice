// This is the dev branch to the tutorial from the master branch.
// I'd like to implement some sort of user input into the code. This will be a string to be encoded and decoded.
// Any of my original comments to the code can be found there. Any comments you see here are related to the current branch.

package main

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"os"
)

func encryptMessage(key string, message string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(message))
	c.Encrypt(msgByte, []byte(message))
	return hex.EncodeToString(msgByte)
}

func decryptMessage(key string, message string) string {
	txt, _ := hex.DecodeString(message)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	msgByte := make([]byte, len(txt))

	c.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])
	return msg
}

func inputText() string {

	var secretMessage string

	fmt.Println("Please enter text to be encoded:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	secretMessage = scanner.Text()

	return secretMessage

}

func main() {
	plainText := inputText()
	key := "this_must_be_of_32_byte_length!!"

	emsg := encryptMessage(key, plainText)
	dmesg := decryptMessage(key, emsg)

	fmt.Println("Encrypted Message is: ", emsg)
	fmt.Println("Decrypted Message is: ", dmesg)

}
