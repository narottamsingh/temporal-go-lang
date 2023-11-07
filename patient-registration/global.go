package app

import "go.temporal.io/sdk/client"

var TemporalClient client.Client

func SetTemporalClient(temClient *client.Client) {
	TemporalClient = *temClient
}
