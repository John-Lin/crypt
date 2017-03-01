package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Subcommands
	encCommand := flag.NewFlagSet("enc", flag.ExitOnError)
	decCommand := flag.NewFlagSet("dec", flag.ExitOnError)

	// Enc subcommand flag pointers
	encFilenamePtr := encCommand.String("f", "", "Specify a filename.")

	// Dec subcommand flag pointers
	decFilenamePtr := decCommand.String("f", "", "Specify a filename.")
	keyPtr := decCommand.String("key", "", "Specify a key for decrypting.")

	if len(os.Args) < 2 {
		fmt.Println("enc or dec subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "enc":
		encCommand.Parse(os.Args[2:])
	case "dec":
		decCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if encCommand.Parsed() {
		if *encFilenamePtr == "" {
			fmt.Println("Please specify a filename to encrypt.")
			encCommand.PrintDefaults()
			os.Exit(1)
		}
		// Key generate
		newKey := NewEncryptionKey256()
		keyEncoded := b64.StdEncoding.EncodeToString(newKey)
		fmt.Println("KEY: " + string(keyEncoded))

		// Read file content
		secretBytes, err := ioutil.ReadFile(*encFilenamePtr)
		check(err)

		// Encrypting
		ciphertext, err := Encrypt(secretBytes, newKey)
		check(err)

		werr := ioutil.WriteFile(*encFilenamePtr+"Enc", ciphertext, 0644)
		check(werr)
		// fmt.Println("Path:", *filenamePtr)
		// fmt.Println("Encrypt:", *encBoolPtr)

	}
	if decCommand.Parsed() {
		if *keyPtr == "" {
			fmt.Println("Please provide a key to decrypt.")
			decCommand.PrintDefaults()
			os.Exit(1)
		}
		if *decFilenamePtr == "" {
			fmt.Println("Please specify a filename.")
			decCommand.PrintDefaults()
			os.Exit(1)
		}
		// fmt.Println("key:", *keyPtr)
		// fmt.Println("Decrypt:", *decBoolPtr)

		keyDecoded, _ := b64.StdEncoding.DecodeString(*keyPtr)

		// Read encrypted file content
		encryptedSecret, err := ioutil.ReadFile(*decFilenamePtr)
		check(err)

		// Decrypting
		plaintext, err := Decrypt(encryptedSecret, keyDecoded)
		check(err)

		werr := ioutil.WriteFile(*decFilenamePtr+"Dec", plaintext, 0644)
		check(werr)

	}

}
