package acme

import (
	"sync"

	"github.com/go-acme/lego/v4/challenge"
)

var (
	challengeStore = make(map[string]string)
	storeMutex     sync.RWMutex
)

type HTTP01Provider struct{}

func (p *HTTP01Provider) Present(domain, token, keyAuth string) error {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	challengeStore[token] = keyAuth
	return nil
}

func (p *HTTP01Provider) CleanUp(domain, token, keyAuth string) error {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	delete(challengeStore, token)
	return nil
}

func Provider() challenge.Provider {
	return &HTTP01Provider{}
}

func GetChallengeResponse(token string) (string, bool) {
	storeMutex.RLock()
	defer storeMutex.RUnlock()
	keyAuth, ok := challengeStore[token]
	return keyAuth, ok
}
