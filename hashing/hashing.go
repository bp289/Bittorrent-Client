package hashing

import (
	"Bittorrent/parse"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/jackpal/bencode-go"
)

func InfoHash(torrentData parse.TorrentDetails) ([]byte, error) {

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

func HashPieces(torrentData parse.TorrentDetails) ([]string, error) {

	pieces := torrentData.Info.Pieces
	var pieceHashes []string
	//each pieces hash is 20 bytes long, therefore we iterate over the pieces in steps of 20
	for i := 0; i < len(pieces); i += 20 {

		if i+20 > len(pieces) {
			break
		}
		// the pieces are of length 20 therefore we iterate 20 times
		pieceHash := pieces[i : i+20]
		pieceHashes = append(pieceHashes, hex.EncodeToString([]byte(pieceHash)))

		// fmt.Printf("Piece %d hash: %x\n", i/20, pieceHash)
	}

	fmt.Printf("Piece Length: %v bytes\n", torrentData.Info.PieceLength)
	fmt.Printf("Total Number of Piece Hashes: %v\n", len(pieceHashes))
	fmt.Printf("Length: %v\n", torrentData.Info.Length)

	return pieceHashes, nil
}
