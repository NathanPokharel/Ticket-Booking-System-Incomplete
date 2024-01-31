package data

func GetGroups() ([]Group, error) {
	var groups []Group
	if err := DB.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func GetBuses(groupID uint) ([]Bus, error) {
	var buses []Bus
	if err := DB.Where("group_id = ?", groupID).Find(&buses).Error; err != nil {
		return nil, err
	}
	return buses, nil
}

func GetSeats(busID uint) ([]Seat, error) {
	var seats []Seat
	if err := DB.Where("bus_id = ?", busID).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}
