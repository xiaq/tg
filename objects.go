package tgbot

// GENERATED AUTOMATICALLY BY objects.py

type Update struct {
	UpdateID int `json:"update_id"`
	Message *Message `json:"message"`
	InlineQuery *InlineQuery `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
}

type GetUpdatesRequest struct {
	Offset *int `json:"offset,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Timeout *int `json:"timeout,omitempty"`
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName *string `json:"last_name"`
	Username *string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Title *string `json:"title"`
	Username *string `json:"username"`
	FirstName *string `json:"first_name"`
	LastName *string `json:"last_name"`
}

type Message struct {
	MessageID int `json:"message_id"`
	From *User `json:"from"`
	Date int `json:"date"`
	Chat Chat `json:"chat"`
	ForwardFrom *User `json:"forward_from"`
	ForwardDate *int `json:"forward_date"`
	ReplyToMessage *Message `json:"reply_to_message"`
	Text *string `json:"text"`
	Audio *Audio `json:"audio"`
	Document *Document `json:"document"`
	Photo []PhotoSize `json:"photo"`
	Sticker *Sticker `json:"sticker"`
	Video *Video `json:"video"`
	Voice *Voice `json:"voice"`
	Caption *string `json:"caption"`
	Contact *Contact `json:"contact"`
	Location *Location `json:"location"`
	NewChatParticipant *User `json:"new_chat_participant"`
	LeftChatParticipant *User `json:"left_chat_participant"`
	NewChatTitle *string `json:"new_chat_title"`
	NewChatPhoto []PhotoSize `json:"new_chat_photo"`
	DeleteChatPhoto bool `json:"delete_chat_photo"`
	GroupChatCreated bool `json:"group_chat_created"`
	SupergroupChatCreated bool `json:"supergroup_chat_created"`
	ChannelChatCreated bool `json:"channel_chat_created"`
	MigrateToChatID *int `json:"migrate_to_chat_id"`
	MigrateFromChatID *int `json:"migrate_from_chat_id"`
}

type PhotoSize struct {
	FileID string `json:"file_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	FileSize *int `json:"file_size"`
}

type Audio struct {
	FileID string `json:"file_id"`
	Duration int `json:"duration"`
	Performer *string `json:"performer"`
	Title *string `json:"title"`
	MimeType *string `json:"mime_type"`
	FileSize *int `json:"file_size"`
}

type Document struct {
	FileID string `json:"file_id"`
	Thumb *PhotoSize `json:"thumb"`
	FileName *string `json:"file_name"`
	MimeType *string `json:"mime_type"`
	FileSize *int `json:"file_size"`
}

type Sticker struct {
	FileID string `json:"file_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Thumb *PhotoSize `json:"thumb"`
	FileSize *int `json:"file_size"`
}

type Video struct {
	FileID string `json:"file_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Duration int `json:"duration"`
	Thumb *PhotoSize `json:"thumb"`
	MimeType *string `json:"mime_type"`
	FileSize *int `json:"file_size"`
}

type Voice struct {
	FileID string `json:"file_id"`
	Duration int `json:"duration"`
	MimeType *string `json:"mime_type"`
	FileSize *int `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName string `json:"first_name"`
	LastName *string `json:"last_name"`
	UserID *int `json:"user_id"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type UserProfilePhotos struct {
	TotalCount int `json:"total_count"`
	Photos [][]PhotoSize `json:"photos"`
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]string `json:"keyboard"`
	ResizeKeyboard *bool `json:"resize_keyboard"`
	OneTimeKeyboard *bool `json:"one_time_keyboard"`
	Selective *bool `json:"selective"`
}

type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective *bool `json:"selective"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective *bool `json:"selective"`
}

type SendMessageRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	Text string `json:"text"`
	ParseMode *string `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool `json:"disable_web_page_preview,omitempty"`
	ReplyToMessageID *int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // ReplyKeyboardMarkup|ReplyKeyboardHide|ForceReply
}

type ForwardMessageRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	FromChatID interface{} `json:"from_chat_id"` // int|string
	MessageID int `json:"message_id"`
}

type SendPhotoRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	Photo interface{} `json:"photo"` // InputFile|string
	Caption *string `json:"caption,omitempty"`
	ReplyToMessageID *int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // ReplyKeyboardMarkup|ReplyKeyboardHide|ForceReply
}

type SendDocumentRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	Document interface{} `json:"document"` // InputFile|string
	ReplyToMessageID *int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // ReplyKeyboardMarkup|ReplyKeyboardHide|ForceReply
}

type SendStickerRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	Sticker interface{} `json:"sticker"` // InputFile|string
	ReplyToMessageID *int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // ReplyKeyboardMarkup|ReplyKeyboardHide|ForceReply
}

type SendVideoRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	Video interface{} `json:"video"` // InputFile|string
	Duration *int `json:"duration,omitempty"`
	Caption *string `json:"caption,omitempty"`
	ReplyToMessageID *int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // ReplyKeyboardMarkup|ReplyKeyboardHide|ForceReply
}

type SendVoiceRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	Voice interface{} `json:"voice"` // InputFile|string
	Duration *int `json:"duration,omitempty"`
	ReplyToMessageID *int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // ReplyKeyboardMarkup|ReplyKeyboardHide|ForceReply
}

type SendLocationRequest struct {
	ChatID interface{} `json:"chat_id"` // int|string
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ReplyToMessageID *int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"` // ReplyKeyboardMarkup|ReplyKeyboardHide|ForceReply
}

type GetUserProfilePhotosRequest struct {
	UserID int `json:"user_id"`
	Offset *int `json:"offset,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

type GetFileRequest struct {
	FileID string `json:"file_id"`
}

type InlineQuery struct {
	ID string `json:"id"`
	From User `json:"from"`
	Query string `json:"query"`
	Offset string `json:"offset"`
}

type AnswerInlineQueryRequest struct {
	InlineQueryID string `json:"inline_query_id"`
	Results []InlineQueryResult `json:"results"`
	CacheTime *int `json:"cache_time,omitempty"`
	IsPersonal *bool `json:"is_personal,omitempty"`
	NextOffset *string `json:"next_offset,omitempty"`
}

type InlineQueryResult interface{} // InlineQueryResultArticle|InlineQueryResultPhoto|InlineQueryResultGif|InlineQueryResultMpeg4Gif|InlineQueryResultVideo

type InlineQueryResultArticle struct {
	Type string `json:"type"`
	ID string `json:"id"`
	Title string `json:"title"`
	MessageText string `json:"message_text"`
	ParseMode *string `json:"parse_mode"`
	DisableWebPagePreview *bool `json:"disable_web_page_preview"`
	URL *string `json:"url"`
	HideURL *bool `json:"hide_url"`
	Description *string `json:"description"`
	ThumbURL *string `json:"thumb_url"`
	ThumbWidth *int `json:"thumb_width"`
	ThumbHeight *int `json:"thumb_height"`
}

type InlineQueryResultPhoto struct {
	Type string `json:"type"`
	ID string `json:"id"`
	PhotoURL string `json:"photo_url"`
	PhotoWidth *int `json:"photo_width"`
	PhotoHeight *int `json:"photo_height"`
	ThumbURL string `json:"thumb_url"`
	Title *string `json:"title"`
	Description *string `json:"description"`
	Caption *string `json:"caption"`
	MessageText *string `json:"message_text"`
	ParseMode *string `json:"parse_mode"`
	DisableWebPagePreview *bool `json:"disable_web_page_preview"`
}

type InlineQueryResultGif struct {
	Type string `json:"type"`
	ID string `json:"id"`
	GIFURL string `json:"gif_url"`
	GIFWidth *int `json:"gif_width"`
	GIFHeight *int `json:"gif_height"`
	ThumbURL string `json:"thumb_url"`
	Title *string `json:"title"`
	Caption *string `json:"caption"`
	MessageText *string `json:"message_text"`
	ParseMode *string `json:"parse_mode"`
	DisableWebPagePreview *bool `json:"disable_web_page_preview"`
}

type InlineQueryResultMpeg4Gif struct {
	Type string `json:"type"`
	ID string `json:"id"`
	MPEG4URL string `json:"mpeg4_url"`
	MPEG4Width *int `json:"mpeg4_width"`
	MPEG4Height *int `json:"mpeg4_height"`
	ThumbURL string `json:"thumb_url"`
	Title *string `json:"title"`
	Caption *string `json:"caption"`
	MessageText *string `json:"message_text"`
	ParseMode *string `json:"parse_mode"`
	DisableWebPagePreview *bool `json:"disable_web_page_preview"`
}

type InlineQueryResultVideo struct {
	Type string `json:"type"`
	ID string `json:"id"`
	VideoURL string `json:"video_url"`
	MimeType string `json:"mime_type"`
	MessageText string `json:"message_text"`
	ParseMode *string `json:"parse_mode"`
	DisableWebPagePreview *bool `json:"disable_web_page_preview"`
	VideoWidth *int `json:"video_width"`
	VideoHeight *int `json:"video_height"`
	VideoDuration *int `json:"video_duration"`
	ThumbURL string `json:"thumb_url"`
	Title string `json:"title"`
	Description *string `json:"description"`
}

type ChosenInlineResult struct {
	ResultID string `json:"result_id"`
	From User `json:"from"`
	Query string `json:"query"`
}

