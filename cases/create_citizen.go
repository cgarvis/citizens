package cases

import "fmt"

import "github.com/cgarvis/citizens/entities"

func CreateCitizen (uid, secret string) (Citizen, error) {
    var citizen Citizen

    if uid == "" {
        return citizen, fmt.Errorf("UID is required")
    }

    if secret == "" {
        return citizen, fmt.Errorf("Secret is required")
    }

    if _, found := Store.FetchCitizenByUID(uid); found {
        return citizen, fmt.Errorf("UID is already taken: %v", uid)
    }

    citizen.UID = uid
    citizen.encryptedSecret = entities.EncryptSecret([]byte(secret))

    if err := Store.SaveCitizen(citizen); err != nil {
        return citizen, fmt.Errorf("There was a problem saving citizen")
    }

    return citizen, nil
}
