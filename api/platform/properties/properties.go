package properties

import (
	"context"
	"fmt"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
	"gorm.io/gorm"
)

type Property struct {
	// gorm.Model
	AuctionType     string  `json:"auction_type"`
	JudgementAmount float64 `json:"judgement_amount"`
	Address         string  `json:"address"`
	AssessedValue   float64 `json:"assessedvalue"`
	Description     string  `json:"description"`
	ZipCode         float64 `json:"zipcode"`
}

func CreateProperty(db *gorm.DB, property *Property) (err error) {
	err = db.Create(property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllProperties(db *gorm.DB, property *[]Property) (err error) {
	err = db.Find(property).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProperty(db *gorm.DB, property *Property, address string) (err error) {
	err = db.Where("address = ?", address).First(property).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProperty(db *gorm.DB, property *Property) (err error) {
	db.Save(property)
	return nil
}

func DeleteProperty(db *gorm.DB, property *Property, address string) (err error) {
	db.Where("address = ?", address).Delete(property)
	return nil
}

func GetDescription(address string) (description string) {
	ctx := context.Background()
	results, err := googlesearch.Search(ctx, address)
	if err != nil {
		return "No information on property found!"
	}
	if len(results) > 0 {
		description = results[0].Description
		splitResult := strings.Split(description, ". ")
		if len(splitResult) > 1 {
			description = strings.TrimSpace(splitResult[1])
			fmt.Println(description)
		} else {
			fmt.Println("Desired substring not found in the input string.")
		}
	}
	return description
}
