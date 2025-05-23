// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"
)

type AuthorizeResult struct {
	Success bool `json:"success"`
	// A textual status message describing the result, can be used to show an error message when `success` is false
	Status string `json:"status"`
	// An access token used to authenticate new API requests as the newly authorized user. Is present when success is true
	Token *string `json:"token,omitempty"`
}

type Coordinates struct {
	// GPS latitude in degrees
	Latitude float64 `json:"latitude"`
	// GPS longitude in degrees
	Longitude float64 `json:"longitude"`
}

type MediaDownload struct {
	// A description of the role of the media file
	Title    string    `json:"title"`
	MediaURL *MediaURL `json:"mediaUrl"`
}

type Mutation struct {
}

type Notification struct {
	// A key used to identify the notification, new notification updates with the same key, should replace the old notifications
	Key  string           `json:"key"`
	Type NotificationType `json:"type"`
	// The text for the title of the notification
	Header string `json:"header"`
	// The text for the body of the notification
	Content string `json:"content"`
	// A value between 0 and 1 when the notification type is `Progress`
	Progress *float64 `json:"progress,omitempty"`
	// Whether or not the message of the notification is positive, the UI might reflect this with a green color
	Positive bool `json:"positive"`
	// Whether or not the message of the notification is negative, the UI might reflect this with a red color
	Negative bool `json:"negative"`
	// Time in milliseconds before the notification should close
	Timeout *int `json:"timeout,omitempty"`
}

// Used to specify how to sort items
type Ordering struct {
	// A column in the database to order by
	OrderBy        *string         `json:"order_by,omitempty"`
	OrderDirection *OrderDirection `json:"order_direction,omitempty"`
}

// Used to specify pagination on a list of items
type Pagination struct {
	// How many items to maximally fetch
	Limit *int `json:"limit,omitempty"`
	// How many items to skip from the beginning of the query, specified by the `Ordering`
	Offset *int `json:"offset,omitempty"`
}

type Query struct {
}

type ScannerResult struct {
	Finished bool     `json:"finished"`
	Success  bool     `json:"success"`
	Progress *float64 `json:"progress,omitempty"`
	Message  *string  `json:"message,omitempty"`
}

type SearchResult struct {
	// The string that was searched for
	Query string `json:"query"`
	// A list of albums that matched the query
	Albums []*Album `json:"albums"`
	// A list of media that matched the query
	Media []*Media `json:"media"`
}

// Credentials used to identify and authenticate a share token
type ShareTokenCredentials struct {
	Token    string  `json:"token"`
	Password *string `json:"password,omitempty"`
}

type Subscription struct {
}

// A group of media from the same album and the same day, that is grouped together in a timeline view
// NOTE: It isn't used. Just copy from the old schema.graphql.
type TimelineGroup struct {
	// The full album containing the media in this timeline group
	Album *Album `json:"album"`
	// The media contained in this timeline group
	Media []*Media `json:"media"`
	// The total amount of media in this timeline group
	MediaTotal int `json:"mediaTotal"`
	// The day shared for all media in this timeline group
	Date time.Time `json:"date"`
}

// Supported language translations of the user interface
type LanguageTranslation string

const (
	LanguageTranslationEnglish              LanguageTranslation = "English"
	LanguageTranslationFrench               LanguageTranslation = "French"
	LanguageTranslationItalian              LanguageTranslation = "Italian"
	LanguageTranslationSwedish              LanguageTranslation = "Swedish"
	LanguageTranslationDanish               LanguageTranslation = "Danish"
	LanguageTranslationSpanish              LanguageTranslation = "Spanish"
	LanguageTranslationPolish               LanguageTranslation = "Polish"
	LanguageTranslationUkrainian            LanguageTranslation = "Ukrainian"
	LanguageTranslationGerman               LanguageTranslation = "German"
	LanguageTranslationRussian              LanguageTranslation = "Russian"
	LanguageTranslationTraditionalChineseTw LanguageTranslation = "TraditionalChineseTW"
	LanguageTranslationTraditionalChineseHk LanguageTranslation = "TraditionalChineseHK"
	LanguageTranslationSimplifiedChinese    LanguageTranslation = "SimplifiedChinese"
	LanguageTranslationPortuguese           LanguageTranslation = "Portuguese"
	LanguageTranslationBasque               LanguageTranslation = "Basque"
	LanguageTranslationTurkish              LanguageTranslation = "Turkish"
)

var AllLanguageTranslation = []LanguageTranslation{
	LanguageTranslationEnglish,
	LanguageTranslationFrench,
	LanguageTranslationItalian,
	LanguageTranslationSwedish,
	LanguageTranslationDanish,
	LanguageTranslationSpanish,
	LanguageTranslationPolish,
	LanguageTranslationUkrainian,
	LanguageTranslationGerman,
	LanguageTranslationRussian,
	LanguageTranslationTraditionalChineseTw,
	LanguageTranslationTraditionalChineseHk,
	LanguageTranslationSimplifiedChinese,
	LanguageTranslationPortuguese,
	LanguageTranslationBasque,
	LanguageTranslationTurkish,
}

func (e LanguageTranslation) IsValid() bool {
	switch e {
	case LanguageTranslationEnglish, LanguageTranslationFrench, LanguageTranslationItalian, LanguageTranslationSwedish, LanguageTranslationDanish, LanguageTranslationSpanish, LanguageTranslationPolish, LanguageTranslationUkrainian, LanguageTranslationGerman, LanguageTranslationRussian, LanguageTranslationTraditionalChineseTw, LanguageTranslationTraditionalChineseHk, LanguageTranslationSimplifiedChinese, LanguageTranslationPortuguese, LanguageTranslationBasque, LanguageTranslationTurkish:
		return true
	}
	return false
}

func (e LanguageTranslation) String() string {
	return string(e)
}

func (e *LanguageTranslation) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageTranslation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageTranslation", str)
	}
	return nil
}

func (e LanguageTranslation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *LanguageTranslation) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e LanguageTranslation) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}

// Specified the type a particular notification is of
type NotificationType string

const (
	// A regular message with no special additions
	NotificationTypeMessage NotificationType = "Message"
	// A notification with an attached progress indicator
	NotificationTypeProgress NotificationType = "Progress"
	// Close a notification with a given key
	NotificationTypeClose NotificationType = "Close"
)

var AllNotificationType = []NotificationType{
	NotificationTypeMessage,
	NotificationTypeProgress,
	NotificationTypeClose,
}

func (e NotificationType) IsValid() bool {
	switch e {
	case NotificationTypeMessage, NotificationTypeProgress, NotificationTypeClose:
		return true
	}
	return false
}

func (e NotificationType) String() string {
	return string(e)
}

func (e *NotificationType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationType", str)
	}
	return nil
}

func (e NotificationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *NotificationType) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e NotificationType) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}

// Used to specify which order to sort items in
type OrderDirection string

const (
	// Sort accending A-Z
	OrderDirectionAsc OrderDirection = "ASC"
	// Sort decending Z-A
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *OrderDirection) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e OrderDirection) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}
