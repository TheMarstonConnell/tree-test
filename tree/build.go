package tree

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/sha3"
	"io"
	"os"
)

func buildTree(buf io.Reader, chunkSize int64) []byte {
	size := 0

	data := make([][]byte, 0)
	chunks := make([][]byte, 0)

	index := 0

	for {
		b := make([]byte, chunkSize)
		read, _ := buf.Read(b)

		if read == 0 {
			break
		}

		b = b[:read]

		size += read

		chunks = append(chunks, b)

		hexedData := hex.EncodeToString(b)

		hash := sha256.New()
		hash.Write([]byte(fmt.Sprintf("%d%s", index, hexedData))) // appending the index and the data
		hashName := hash.Sum(nil)

		data = append(data, hashName)

		index++
	}

	tree, err := merkletree.NewTree(
		merkletree.WithData(data),
		merkletree.WithHashType(sha3.New512()),
		merkletree.WithSalt(false),
	)
	if err != nil {
		panic(err)
	}

	r := tree.Root()

	return r
}

func loadFile(fileName string) []byte {

	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}

func BuildTree(fileName string, chunkSize int64) {
	data := loadFile(fileName)
	buf := bytes.NewBuffer(data)

	root := buildTree(buf, chunkSize)

	fmt.Printf("Merkle Root: %x\n", root)

}
