package cases

import "errors"

type DataStore interface {
    FetchCitizenByUID(uid string) (Citizen, bool)
    FetchCitizens() []Citizen
    SaveCitizen(Citizen) error
}

var Store DataStore

type memoryStore struct {
    citizens    map[string]Citizen
}

func (m memoryStore) FetchCitizenByUID(uid string) (Citizen, bool) {
    if citizen, ok := m.citizens[uid]; ok {
        return citizen, true
    }

    return Citizen{}, false
}

func (m memoryStore) FetchCitizens() []Citizen {
    citizens := []Citizen{}
    for _, citizen := range m.citizens {
        citizens = append(citizens, citizen)
    }
    return citizens
}

func (m memoryStore) SaveCitizen(citizen Citizen) error {
    if citizen.UID == "" {
        return errors.New("Citizen needs to have a UID")
    }

    m.citizens[citizen.UID] = citizen
    return nil
}

func NewMemoryStore() memoryStore {
    citizens := make(map[string]Citizen)
    return memoryStore{citizens}
}

