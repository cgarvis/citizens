package cases

import "errors"

import "github.com/cgarvis/citizens/entities"

func UpdateCitizen(uid string, attributes Citizen) (Citizen, error) {
    var citizen Citizen

    if citizen, ok := Store.FetchCitizenByUID(uid); ok {
        citizen.secret = entities.EncryptSecret(attributes.secret)
        Store.SaveCitizen(citizen)
        return citizen, nil
    }

    return citizen, errors.New("Could not find citizen")
}
