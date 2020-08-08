package plex

import (
	"strings"
)

const (
	// LibraryOnDeck A new item is added that appears in the user’s On Deck.
	// A poster is also attached to this event.
	LibraryOnDeck          Name = "library.on.deck"

	// LibraryNew A new item is added to a library to which the user has access.
	// A poster is also attached to this event.
	LibraryNew             Name = "library.new"

	// MediaPause Media playback pauses.
	MediaPause             Name = "media.pause"

	// MediaPlay Media starts playing. An appropriate poster is attached.
	MediaPlay              Name = "media.play"

	// MediaRate Media is rated. A poster is also attached to this event.
	MediaRate              Name = "media.rate"

	// MediaResume Media playback resumes.
	MediaResume            Name = "media.resume"

	// MediaScrobble Media is viewed (played past the 90% mark).
	MediaScrobble          Name = "media.scrobble"

	// MediaStop Media playback stops.
	MediaStop              Name = "media.stop"

	// AdminDatabaseBackup A database backup is completed successfully via Scheduled Tasks.
	AdminDatabaseBackup    Name = "admin.database.backup"

	// AdminDatabaseCorrupted Corruption is detected in the server database.
	AdminDatabaseCorrupted Name = "admin.database.corrupted"

	// DeviceNew A device accesses the owner’s server for any reason, which may come from
	// background connection testing and doesn’t necessarily indicate active browsing or playback.
	DeviceNew              Name = "device.new"

	// PlaybackStarted Playback is started by a shared user for the server.
	// A poster is also attached to this event.
	PlaybackStarted        Name = "playback.started"

	// EpisodeMetadataType Event is related to a TV Show
	EpisodeMetadataType = "episode"

	// MovieMetadataType Event is related to a Movie
	MovieMetadataType   = "movie"
)

// Event This contains event, user, and owner attributes.
// The event attribute holds the name of the event, as specified above.
// The user and owner flags indicate whether the event is sent because the user has a
// webhook configured (user flag is set) or because the server owner has a webhook set (the owner flag is set).
// This can be very useful for certain scenarios.
// Note that if a server owner triggers an event, both user and owner flags will be set.
type Event struct {
	Name     Name     `json:"event"`
	User     bool     `json:"user"`
	Owner    bool     `json:"owner"`
	Account  Account  `json:"Account"`
	Server   Server   `json:"Server"`
	Player   Player   `json:"Player"`
	Metadata Metadata `json:"Metadata"`
}

// Account The account object contains information about the user who triggered the event, including ID, title,
// and in some cases, a URL for the user’s avatar image. Note that the owner ID will always be 1.
type Account struct {
	ID    int    `json:"id"`
	Thumb string `json:"thumb"`
	Title string `json:"title"`
}

// Player The player object contains information about the player which generated the event, if applicable.
// It contains the title of the player, its UUID, the public IP address of the player,
// and a local flag indicating whether or not it was on the same network as the server.
type Player struct {
	Local         bool   `json:"local"`
	PublicAddress string `json:"publicAddress"`
	Title         string `json:"title"`
	UUID          string `json:"uuid"`
}

// Server The server object contains information about the server from which the event was generated,
// including the title and UUID.
type Server struct {
	Title string `json:"title"`
	UUID  string `json:"uuid"`
}

// Metadata This last object contains detailed information about the media.
type Metadata struct {
	LibrarySectionType   string `json:"librarySectionType"`
	RatingKey            string `json:"ratingKey"`
	Key                  string `json:"key"`
	ParentRatingKey      string `json:"parentRatingKey"`
	GrandparentRatingKey string `json:"grandparentRatingKey"`
	GUID                 string `json:"guid"`
	LibrarySectionID     int    `json:"librarySectionID"`
	Type                 string `json:"type"`
	Title                string `json:"title"`
	GrandparentKey       string `json:"grandparentKey"`
	ParentKey            string `json:"parentKey"`
	GrandparentTitle     string `json:"grandparentTitle"`
	ParentTitle          string `json:"parentTitle"`
	Summary              string `json:"summary"`
	Index                int    `json:"index"`
	ParentIndex          int    `json:"parentIndex"`
	RatingCount          int    `json:"ratingCount"`
	Thumb                string `json:"thumb"`
	Art                  string `json:"art"`
	ParentThumb          string `json:"parentThumb"`
	GrandparentThumb     string `json:"grandparentThumb"`
	GrandparentArt       string `json:"grandparentArt"`
	AddedAt              int    `json:"addedAt"`
	UpdatedAt            int    `json:"updatedAt"`
}

// Name Event Name
type Name string

func (name Name) String() string {
	return strings.ToLower(string(name))
}
