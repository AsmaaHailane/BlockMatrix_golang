package blockmatrix

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"sync"
)

type blockmatrix struct {
	Dimension int
	//HashAlgorithm string
	BlockData     [][]bytes.Buffer
	BlockHashes   [][]string
	RowHashes     []string
	HashOfRows    string
	ColumnHashes  []string
	HashOfColumns string
	HashOfMatrix  string
	RowLocks      []sync.RWMutex
	ColumnLocks   []sync.RWMutex
}

func initBlockMatrix(Dimension int) *blockmatrix {

	bm := new(blockmatrix)
	bm.Dimension = Dimension

	bm.BlockData = make([][]bytes.Buffer, Dimension)
	for i := 0; i < Dimension; i++ {
		bm.BlockData[i] = make([]bytes.Buffer, Dimension)
	}

	bm.BlockHashes = make([][]string, Dimension)
	for i := 0; i < Dimension; i++ {
		bm.BlockHashes[i] = make([]string, Dimension)
	}

	bm.RowHashes = make([]string, Dimension)
	bm.ColumnHashes = make([]string, Dimension)

	bm.RowLocks = make([]sync.RWMutex, Dimension)
	bm.ColumnLocks = make([]sync.RWMutex, Dimension)

	bm.fillDiagonalWithRandomData()
	return bm
}

// function that provide hash of a given string

func (bm *blockmatrix) hashOfString(DataStr string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(DataStr)))
}

func (bm *blockmatrix) hashOfBytes(Data bytes.Buffer) string {

	var HashStr string
	HashStr = fmt.Sprintf("%x", sha256.Sum256(Data.Bytes()))

	return HashStr
}

// function that calculate the hash of row
func (bm *blockmatrix) calculateRowHash(row int) {

	var Hashes string

	for column := 0; column < bm.Dimension; column++ {
		if row != column {
			// Hashes += bm.BlockData[row][column]
			bm.BlockHashes[row][column] = bm.hashOfBytes(bm.BlockData[row][column])
			Hashes += bm.BlockHashes[row][column]
		}
	}

	bm.RowHashes[row] = bm.hashOfString(Hashes)

}

// function that calculate the hash of column
func (bm *blockmatrix) calculateColumnHash(column int) {

	var Hashes string

	for row := 0; row < bm.Dimension; row++ {
		if column != row {
			bm.BlockHashes[row][column] = bm.hashOfBytes(bm.BlockData[row][column])
			Hashes += bm.BlockHashes[row][column]
		}
	}

	bm.ColumnHashes[column] = bm.hashOfString(Hashes)

}

func (bm *blockmatrix) fillDiagonalWithRandomData() error {

	RandomData := make([]byte, 64)

	for k := 0; k < bm.Dimension; k++ {

		_, err := rand.Read(RandomData)
		if err != nil {
			return err
		}

		bm.BlockData[k][k].Write(RandomData)
		bm.BlockHashes[k][k] = bm.hashOfBytes(bm.BlockData[k][k])
	}

	return nil
}
