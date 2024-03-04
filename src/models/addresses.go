package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	CustomerId string `json:"customerId" bson:"customerId"`
	AddressId  string `json:"addressId" bson:"addressId"`
	Name       string `json:"name,omitempty" bson:"name"`
	Number     string `json:"number,omitempty" bson:"number"`
	StreetName string `json:"streetName" bson:"streetName"`
	Locality   string `json:"locality" bson:"locality"`
	PostTown   string `json:"postTown" bson:"postTown"`
	PostCode   string `json:"postCode" bson:"postCode"`
}

func (address *Address) Save() (*Address, error) {
	err := Database.Model(&address).Create(&address).Error

	if err != nil {
		return &Address{}, err
	}
	return address, nil
}

func FetchAddressById(customerId string, addressId string) (*Address, error) {
	var address Address
	err := Database.Where("customer_id =? AND address_id = ?", customerId, addressId).First(&address).Error

	if err != nil {
		return &Address{}, err
	}
	return &address, nil

}

func FetchCustomerAddresses(customerId string) (*[]Address, error) {
	var addresses []Address

	err := Database.Where("customer_id =?", customerId).Find(&addresses).Error

	if err != nil {
		return &[]Address{}, err
	}
	return &addresses, nil
}

func (address *Address) UpdateCustomerAddress(customerId string, addressId string) (*Address, error) {
	err := Database.Model(&Address{}).Where("customer_id = ? AND address_id = ?", customerId, addressId).Updates(address).Error
	if err != nil {
		return &Address{}, err
	}
	return address, nil
}

func DeleteAddress(customerId string, addressId string) error {
	err := Database.Model(&Address{}).Where("customer_id = ? AND address_id = ?", customerId, addressId).Delete(&Address{}).Error
	if err != nil {
		return err
	}
	return nil
}
