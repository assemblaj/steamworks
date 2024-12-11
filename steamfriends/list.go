package steamfriends

import (
	"github.com/assemblaj/steamworks"
	"github.com/assemblaj/steamworks/internal"
)

type FriendFlags internal.EFriendFlags

type Friend struct {
	SteamID     steamworks.SteamID
	PersonaName string
}

const (
	FriendFlagNone                 FriendFlags = 0
	FriendFlagBlocked              FriendFlags = 1
	FriendFlagFriendshipRequested  FriendFlags = 2
	FriendFlagImmediate            FriendFlags = 4
	FriendFlagClanMember           FriendFlags = 8
	FriendFlagOnGameServer         FriendFlags = 16
	FriendFlagRequestingFriendship FriendFlags = 128
	FriendFlagRequestingInfo       FriendFlags = 256
	FriendFlagIgnored              FriendFlags = 512
	FriendFlagIgnoredFriend        FriendFlags = 1024
	FriendFlagChatMember           FriendFlags = 4096
	FriendFlagAll                  FriendFlags = 65535
)

func GetFriendCount(friendFlag FriendFlags) int32 {
	return internal.SteamAPI_ISteamFriends_GetFriendCount(int32(friendFlag))
}

func GetFriendsList(friendFlag FriendFlags) []Friend {
	friendList := make([]Friend, 0)
	count := GetFriendCount(friendFlag)
	for i := int32(0); i < count; i++ {
		friendSteamID := internal.SteamAPI_ISteamFriends_GetFriendByIndex(i, int32(friendFlag))
		personaName := internal.GoString(internal.SteamAPI_ISteamFriends_GetFriendPersonaName(friendSteamID))
		friendList = append(friendList, Friend{steamworks.SteamID(friendSteamID), personaName})
	}
	return friendList
}
