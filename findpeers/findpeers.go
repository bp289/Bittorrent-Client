package findpeers

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jackpal/bencode-go"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "*+-=_<>,.({[]})&^%$#@!~|"

type TorrentPeers struct {
	Interval int
	Peers    string
}

func FindPeers(trackerUrl string, infoHash []byte) ([]string, error) {
	port := 6881
	peer_id := generatePeerId()

	uploaded := 0
	downloaded := 0
	left := 0

	compact := 1

	query := url.Values{}
	query.Add("info_hash", string(infoHash))
	query.Add("peer_id", peer_id)
	query.Add("port", strconv.Itoa(port))
	query.Add("uploaded", strconv.Itoa(uploaded))
	query.Add("downloaded", strconv.Itoa(downloaded))
	query.Add("left", strconv.Itoa(left))
	query.Add("compact", strconv.Itoa(compact))

	resp, err := http.Get(trackerUrl)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var responseDecoded TorrentPeers

	bencodeErr := bencode.Unmarshal(resp.Body, &responseDecoded)

	if bencodeErr != nil {
		log.Fatalln(bencodeErr)
		return nil, err
	}

	var peers = make([]string, 0)
	peerData := []byte(responseDecoded.Peers)

	//Each peer is represented using 6 bytes.
	//first 4 are peers ip last two are peers port
	//increment 6 each time
	for i := 0; i < len(peerData); i += 6 {
		//as we need ip to be from 0 - 3 not 0 - 4
		ipStart := i
		ipEndPlusOne := i + 4
		ipBytes := peerData[ipStart:ipEndPlusOne]

		portStart := i + 4
		portEnd := i + 6
		portBytes := peerData[portStart:portEnd]

		ipStr := net.IP(ipBytes).String()
		//bigEdian means the most significant, (higher value) comes first.
		portStr := binary.BigEndian.Uint16(portBytes)

		ipPortStr := ipStr + ":" + fmt.Sprintf("%d", portStr)
		peers = append(peers, ipPortStr)
	}

	return peers, nil
}

func generatePeerId() string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, 20)

	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
