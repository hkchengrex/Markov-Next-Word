package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/boltdb/bolt"
)

func writeForLength(length int) string {
	var result = ""
	err := myBoltDB.View(func(tx *bolt.Tx) error {
		rootBucket := tx.Bucket([]byte(dbBucketName))
		if rootBucket == nil {
			fmt.Println("Root Bucket does not exist")
		}

		type pair struct {
			key rune
			cdf int
		}

		rootString := obtainStartOfText()

		for ch := 0; ch < length; ch++ {
			currBucket := rootBucket.Bucket([]byte(rootString))
			if currBucket == nil {
				return errors.New(rootString + " Not found.")
			}

			var candidate []pair
			sum := 0
			currBucket.ForEach(func(k, v []byte) error {
				num := byteArrayToInt(v)
				sum += num
				candidate = append(candidate, pair{[]rune(string(k))[0], sum})
				return nil
			})

			randPick := rand.Intn(sum)
			var picked rune
			for _, v := range candidate {
				if v.cdf > randPick {
					picked = v.key
					break
				}
			}

			if picked == rune(3) {
				//If end of text is reached
				break
			}

			//Output one rune
			result += string(picked)
			r := []rune(rootString)
			for i := 0; i < gramNum-2; i++ {
				r[i] = r[i+1]
			}
			r[gramNum-2] = picked
			rootString = string(r)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return result
}
