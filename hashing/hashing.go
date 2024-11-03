package hashing

import (
	"Bittorrent/parse"
	"crypto/sha1"
	"fmt"

	"github.com/jackpal/bencode-go"
)

func InfoHash(torrentData parse.BencodeTorrent) ([]byte, error) {

	hash := sha1.New()
	//encode the torrentData.Info and write to the hash object.
	hashErr := bencode.Marshal(hash, torrentData.Info)
	sum := hash.Sum(nil)
	//In Go, hash.Sum is used to finalize and retrieve the result of a hash computation,
	if hashErr != nil {
		fmt.Println("failed to hash data: %w", hashErr)
		return nil, fmt.Errorf("failed to parse bencoded data: %w", hashErr)
	}
	return sum, nil
}
