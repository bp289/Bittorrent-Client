package parse

import (
	"fmt"
	"io"

	"github.com/jackpal/bencode-go"
)

type BencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type TorrentDetails struct {
	Announce    string      `bencode:"announce"`
	Info        BencodeInfo `bencode:"info"`
	PieceHashes []string
}

func Parse(reader io.Reader) (*TorrentDetails, error) {

	data := TorrentDetails{}
	err := bencode.Unmarshal(reader, &data)

	if err != nil {
		return nil, fmt.Errorf("failed to parse bencoded data: %w", err)
	}
	return &data, nil
}

type BencodeTorrentWithPH struct {
	TorrentDetails // Embedding BencodeTorrent to inherit its fields
	PieceHashes    [][]string
}
