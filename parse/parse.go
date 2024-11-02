package parse

import (
	"fmt"
	"io"
	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
    Pieces      string `bencode:"pieces"`
    PieceLength int    `bencode:"piece length"`
    Length      int    `bencode:"length"`
    Name        string `bencode:"name"`
}

type bencodeTorrent struct {
    Announce string      `bencode:"announce"`
    Info     bencodeInfo `bencode:"info"`
}

func Parse(reader io.Reader) (*bencodeTorrent, error) {
	
	data := bencodeTorrent{}
    err := bencode.Unmarshal(reader, &data)
    
	if err != nil {
        return nil, fmt.Errorf("failed to parse bencoded data: %w", err)
    }
	return &data, nil
}


