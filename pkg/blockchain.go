package blockchain

import (
	"log"
)

type Blockchain struct {
	initialized bool
}

func (bc *Blockchain) InitializeNetwork() {
	log.Println("Blockchain network initialization started...")
	// Simulate initialization (e.g., setting up genesis block, peers, etc.)
	// Here, we'll just wait a bit to simulate some setup time
	bc.initialized = true
	log.Println("Blockchain network initialization completed.")
}

func (bc *Blockchain) IsInitialized() bool {
	return bc.initialized
}
