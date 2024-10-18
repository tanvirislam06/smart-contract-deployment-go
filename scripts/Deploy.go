package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"tanvirislam06/tanvirislam-qa-assessment/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Output struct {
	ContractAddress string `json:"contract_address"`
	DeployedAddress string `json:"deployed_address"`
	ValueSet        string `json:"value_set"`
}

func main() {
	valueToSet := flag.String("value", "", "The value to be set in the contract")
	rpcURL := flag.String("rpc", "", "The RPC API key")
	privateKeyHex := flag.String("private", "", "The wallet private key")
	flag.Parse()

	if *valueToSet == "" || *rpcURL == "" || *privateKeyHex == "" {
		log.Fatal("All parameters (value, rpc, private) are required")
	}

	client, err := ethclient.Dial(*rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(*privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := bindings.DeployBindings(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contract deployed to: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	// Set the value in the contract
	// var valueToSet32 [32]byte
	// copy(valueToSet32[:], (*valueToSet)[:32])
	// tx, err = instance.SetBytes32(auth, valueToSet32)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("SetBytes32 transaction hash: %s\n", tx.Hash().Hex())

	valueBytes32, err := instance.GetBytes32(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GetBytes32 value: %s\n", string(valueBytes32[:]))

	output := Output{
		ContractAddress: address.Hex(),
		DeployedAddress: fromAddress.Hex(),
		ValueSet:        string(valueBytes32[:]),
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Output written to output.json:")
	fmt.Println(string(jsonData))
}
