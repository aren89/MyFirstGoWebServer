package person

import "log"

func ManagePersonToSave(model PersonModel) (int, error) {
	log.Println(model)
	return 1, nil
}
