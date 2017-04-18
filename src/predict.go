package main

import (
	"errors"

	"github.com/boltdb/bolt"
)

func predictOneWord(predictBase []string) string {
	var result = ""
	err := myBoltDB.View(func(tx *bolt.Tx) error {
		currBucket := tx.Bucket([]byte(dbBucketName))

		if currBucket == nil {
			panic("Root Bucket does not exist")
		}

		for i := 0; i < gramNum-1; i++ {
			currBucket = currBucket.Bucket([]byte(predictBase[i]))
			if currBucket == nil {
				return errors.New("Text Not found")
			}
		}

		candidate := make(map[string]int)
		currBucket.ForEach(func(k, v []byte) error {
			candidate[string(k)] = byteArrayToInt(v)
			return nil
		})

		total := 0
		maxNum := 0
		for k, v := range candidate {
			total += v
			if v > maxNum {
				maxNum = v
				result = k
			}
		}

		return nil
	})

	if err != nil {
		return ""
	}

	return result
}
