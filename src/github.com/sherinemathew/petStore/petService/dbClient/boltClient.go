package dbclient

import (
	"github.com/sherinemathew/petStore/petService/service/model"
	"github.com/boltdb/bolt"
    "log"
    "strconv"
    "encoding/json"
    "fmt"
)

type IBoltClient interface {
        OpenBoltDb()
        QueryPet(petId string) (model.Pet, error)
        Seed()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("pets.db", 0600, nil)
	if err != nil {
			log.Fatal(err)
	}
}

func (bc *BoltClient) QueryPet(petId string) (model.Pet, error) {

	pet := model.Pet{}

	// Read an object from the pet table using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("PetBucket"))

			// Read the pet record value
			petBytes := b.Get([]byte(petId))
			if petBytes == nil {
					return fmt.Errorf("No pet found for " + petId)
			}
			json.Unmarshal(petBytes, &pet)
			return nil
	})
	// If there were an error, return the error
	if err != nil {
			return model.Pet{}, err
	}
	// Return the pet struct and nil as error.
	return pet, nil
}

// Create dummy pet records
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedPets()
}

// Creates a pet table in BoltDB.
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucket([]byte("PetBucket"))
			if err != nil {
					return fmt.Errorf("create bucket failed: %s", err)
			}
			return nil
	})
}

// Create dummy pet records
func (bc *BoltClient) seedPets() {
	fmt.Printf("seeding")

	total := 50
	for i := 0; i < total; i++ {

			// Generate a key
			id := 1 + i;
			key := strconv.Itoa(id);

			// Create an instance of Pet struct
			pet := model.Pet{
				Id: id,
				Name: "Pet_" + strconv.Itoa(i),
				PhotoURLs: []string{}, Status: "available",
				Tags: nil,
			}

			// Serialize the struct to JSON
			jsonBytes, _ := json.Marshal(pet)

			// Write the data to the Pet table
			bc.boltDB.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("PetBucket"))
					err := b.Put([]byte(key), jsonBytes)
					return err
			})
	}
	fmt.Printf("Inserted %v fake pets...\n", total)
}