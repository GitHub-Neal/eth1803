package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main() {
	db,err:=bolt.Open("mlly.db",0600,nil)
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b,err:=tx.CreateBucket([]byte("MyBucket"))
		if err!=nil{
			return fmt.Errorf("create bucket:%s",err)
		}
		if b!=nil{
			err=b.Put([]byte("1"),[]byte("110"))
			if err!=nil{
				return err
			}
		}
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte("MyBucket"))
		if b!=nil{
			value:=b.Get([]byte("1"))
			fmt.Printf("value:%s\n",string(value))
		}
		return nil
	})

}
