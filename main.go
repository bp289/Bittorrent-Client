package main

import (
	"Bittorrent/hashing"
	"Bittorrent/parse"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/jackpal/bencode-go"
)

// Open parses a torrent file
func Open(reader io.Reader) (*parse.BencodeTorrent, error) {
	bto := parse.BencodeTorrent{}
	err := bencode.Unmarshal(reader, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	command := os.Args[1]

	switch command {
	case "decode":

		bencodedValue := os.Args[2]

		decoded_, err := bencode.Decode(bytes.NewReader([]byte(bencodedValue)))

		// decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded_)
		fmt.Println(string(jsonOutput))

	case "info":
		inputFile := os.Args[2]

		file, err := os.Open(inputFile)

		if err != nil {
			fmt.Printf("error reading file %v\n", err)
			return
		}

		torrentData, unmarshErr := parse.Parse(file)

		if unmarshErr != nil {
			fmt.Println("failed to parse bencoded data: %w", unmarshErr)
			return
		}

		hashing.HashPieces(*torrentData)

		// infoHash, hashErr := hashing.InfoHash(*torrentData)

		// //In Go, hash.Sum is used to finalize and retrieve the result of a hash computation,
		// if hashErr != nil {
		// 	fmt.Println("failed to hash data: %w", hashErr)
		// 	return
		// }

		// // to understand the %x read: https://pkg.go.dev/fmt
		// fmt.Printf("Info Hash: %x\n", infoHash)

		// fmt.Println("Piece Length:", torrentData.Info.PieceLength)
		// fmt.Println("Total Length:", torrentData.Info.Length)
		// fmt.Println("File Name:", torrentData.Info.Name)
		// fmt.Println("Announce URL:", torrentData.Announce)

	}

}
