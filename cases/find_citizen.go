package cases

import "errors"

func FindCitizen(uid string) (Citizen, error) {
    var citizen Citizen

    if citizen, ok := Store.FetchCitizenByUID(uid); ok {
        return citizen, nil
    }

    return citizen, errors.New("Citizen not found")
}

