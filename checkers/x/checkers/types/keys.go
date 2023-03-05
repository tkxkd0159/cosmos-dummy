package types

import (
	"encoding/binary"
	"time"
)

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// StoredGameKey returns the store key to retrieve a StoredGame from the index fields
func StoredGameKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

var _ binary.ByteOrder

const (
	// StoredGameKeyPrefix is the prefix to retrieve all StoredGame
	StoredGameKeyPrefix = "StoredGame/value/"
	SystemInfoKey       = "SystemInfo/value/"
)

const (
	NoFifoIndex           = "-1"
	MaxTurnDuration1Day   = time.Duration(24 * 3_600 * 1000_000_000)
	MaxTurnDuration1Min   = time.Duration(1 * 60 * 1000_000_000)
	DefaultDeadlineLayout = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
	GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
	GameCreatedEventCreator   = "creator"          // Subsidiary information
	GameCreatedEventGameIndex = "game-index"       // What game is relevant
	GameCreatedEventBlack     = "black"            // Is it relevant to me?
	GameCreatedEventRed       = "red"              // Is it relevant to me?

	MovePlayedEventType      = "move-played"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventCapturedX = "captured-x"
	MovePlayedEventCapturedY = "captured-y"
	MovePlayedEventWinner    = "winner"
	MovePlayedEventBoard     = "board"

	GameRejectedEventType      = "game-rejected"
	GameRejectedEventCreator   = "creator"
	GameRejectedEventGameIndex = "game-index"
)
