package main

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/qcodelabsllc/qreeket/messaging/network"
	"github.com/qcodelabsllc/qreeket/messaging/utils"
	"log"
	"strings"
)

func main() {
	// This line loads the environment variables from the ".env" file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load environment variables: %+v\n", err)
	}

	// encrypt & decrypt message
	key := []byte(strings.ReplaceAll(uuid.NewString(), "-", ""))
	encryptedMessage, err := utils.EncryptMessage(key, "Hello World!")
	if err != nil {
		log.Fatalf("unable to encrypt message: %+v\n", err)
	}
	log.Printf("encrypted message: %s\n", string(encryptedMessage))
	decryptedMessage, err := utils.DecryptMessage(key, encryptedMessage)
	if err != nil {
		log.Fatalf("unable to decrypt message: %+v\n", err)
	}
	log.Printf("decrypted message: %s\n", decryptedMessage)

	// This line initializes the grpc server
	network.InitServer()
}
