package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo-value-"
)

const (
	GameCreatedEventType      = "new-game-created"
	GameCreatedEventCreator   = "creator"
	GameCreatedEventGameIndex = "game-index"
	GameCreatedEventBlack     = "black"
	GameCreatedEventRed       = "red"
)

const (
	MovePlayedEventType      = "move-played"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventCapturedX = "captured-x"
	MovePlayedEventCapturedY = "captured-y"
	MovePlayedEventWinner    = "winner"
	MovePlayedEventBoard     = "board"
)

const (
	MaxTurnDuration = time.Duration(24 * 3_600 * 1000_000_000) // 1 day
	DeadlineLayout  = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
	NoFifoIndex = "-1"
)
