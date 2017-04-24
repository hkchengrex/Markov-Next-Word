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

		rootRunes := make([]rune, gramNum)
		rootRunes[0] = '\n'
		currBucket := rootBucket
		//Create root set
		for i := 1; i < gramNum-1; i++ {
			var rootCandidate []pair
			sum := 0
			currBucket.ForEach(func(k, v []byte) error {
				fmt.Println(len(string(k)))
				fmt.Println(len([]rune(string(k))))
				fmt.Println([]rune(string(k))[0])
				fmt.Println(len(v))
				num := byteArrayToInt(v)
				sum += num
				rootCandidate = append(rootCandidate, pair{[]rune(string(k))[0], sum})
				return nil
			})

			randPick := rand.Intn(sum)
			for _, v := range rootCandidate {
				if v.cdf > randPick {
					rootRunes[i] = v.key
					break
				}
			}
		}

		for ch := 0; ch < length; ch++ {
			for i := 0; i < gramNum; i++ {
				currBucket = currBucket.Bucket([]byte(string(rootRunes[i])))
				if currBucket == nil {
					return errors.New("Not found.")
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
				for _, v := range candidate {
					if v.cdf > randPick {
						rootRunes[i] = v.key
						break
					}
				}
			}

			//Output one rune
			result += string(rootRunes[gramNum-1])
			for i := 0; i < gramNum-1; i++ {
				rootRunes[i] = rootRunes[i+1]
			}
		}

		return nil
	})

	if err != nil {
		return ""
	}

	return result
}
