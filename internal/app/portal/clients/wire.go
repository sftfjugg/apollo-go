package clients

import "github.com/google/wire"

// ProviderSet wire provider of clients clients
var ProviderSet = wire.NewSet(NewUicService, NewLimosService, NewHttpClient)
