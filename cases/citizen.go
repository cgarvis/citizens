package cases

type Citizen struct {
    UID             string
    Secret          string
    encryptedSecret []byte
}
