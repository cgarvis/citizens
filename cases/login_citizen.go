 package cases

 import "errors"

 import "github.com/cgarvis/citizens/entities"

 var Key []byte

 func LoginCitizen(uid, plain_secret string) (string, error) {
    var (
        citizen Citizen
        ok bool
    )
    if citizen, ok = Store.FetchCitizenByUID(uid); !ok {
        return "", errors.New("Citizen not found")
    }

    if !entities.CompareSecret(citizen.encryptedSecret, []byte(plain_secret)) {
        return "", errors.New("Secret does not match")
    }

    token := entities.GenerateToken(Key, citizen.UID)

    return token, nil
 }
