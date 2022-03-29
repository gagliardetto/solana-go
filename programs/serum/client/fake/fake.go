package fake

import "github.com/gagliardetto/solana-go"

type MarketState struct {
	Participants map[string]*Participant
}

func Create() (*MarketState, error) {
	var err error
	ans := new(MarketState)
	ans.Participants = make(map[string]*Participant)
	nameList := []string{
		"Dee", "Mac", "Charlie", "Dennis", "Frank", "Waitress", "Ford",
	}
	for i := 0; i < len(nameList); i++ {
		x := new(Participant)
		x.Name = nameList[i]
		ans.Participants[x.Name] = x

		x.Key, err = solana.NewRandomPrivateKey()
		if err != nil {
			return nil, err
		}

	}

	return ans, nil
}

func (ms *MarketState) IterateParticipants(callback func(*Participant) error) error {
	var err error
	for _, p := range ms.Participants {
		err = callback(p)
		if err != nil {
			return err
		}
	}
	return nil
}

type Participant struct {
	Name string
	Key  solana.PrivateKey
}

func (p *Participant) PublicKey() solana.PublicKey {
	return p.Key.PublicKey()
}
