package lib

type ChatterinoConfig struct {
	Behavior BehavoirConfig `json:"behavior"`
	Misc     MiscConfig     `json:"misc"`
	Ui       UiConfig       `json:"ui"`
	//Accounts      AccountsConfig        `json:"accounts"`
	Highlighting  HighlightingConfig    `json:"highlighting"`
	Filtering     FilteringList         `json:"filtering"`
	Links         LinksConfig           `json:"links"`
	Emotes        EmotesConfig          `json:"emotes"`
	Appearance    AppearanceConfig      `json:"appearance"`
	Notifications NotificationsSettings `json:"notifications"`
	Whispers      WhispersConfig        `json:"whispers"`
	Debug         DebugConfig           `json:"debug"`
	External      ExternalsConfig       `json:"external"`
	StreamerMode  StreamerModeConfig    `json:"streamerMode"`
	Logging       LoggingConfig         `json:"logging"`
	//Hotkeys       HotkeysConfig         `json:"hotkeys"`
	Similarity SimilarityConfig `json:"similarity"`
	Nicknames  []NicknameConfig `json:"nicknames"`
}

type BehavoirConfig struct {
	Autorun   bool `json:"autorun"`
	ShowJoins bool `json:"showJoins"`
	ShowParts bool `json:"showParts"`
}

type MiscConfig struct {
	CurrentVersion     string           `json:"currentVersion"`
	Beta               bool             `json:"beta"`
	Twitch             MiscTwitchConfig `json:"twitch"`
	ExperimentalIrc    bool             `json:"experimentalIrc"`
	RestartOnCrash     bool             `json:"restartOnCrash"`
	LockNotebookLayout bool             `json:"lockNotebookLayout"`
}

type MiscTwitchConfig struct {
	LoadMessageHistroyOnConnect bool `json:"loadMessageHistoryOnConnect"`
}

type UiConfig struct {
	LastSelectChannelTab int `json:"lastSelectChannelTab"`
	LastSelectIrcConn    int `json:"lastSelectIrcConn"`
}

type HighlightingConfig struct {
	Badges          []BadgeHighlight      `json:"badges"`
	Users           []IndividualHighlight `json:"users"`
	SelfHighlight   SelfHighlightRule     `json:"selfHighlight"`
	AlwaysPlaySound bool                  `json:"alwaysPlaySound"`
	Highlights      []IndividualHighlight `json:"highlights"`
}

type BadgeHighlight struct {
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
	ShowInMentions bool   `json:"showInMentions"`
	Alert          bool   `json:"alert"`
	Sound          bool   `json:"sound"`
	SoundUrl       string `json:"soundUrl"`
	Color          string `json:"color"`
}

type IndividualHighlight struct {
	Pattern        string `json:"pattern"`
	ShowInMentions bool   `json:"showInMentions"`
	Alert          bool   `json:"alert"`
	Sound          bool   `json:"sound"`
	Regex          bool   `json:"regex"`
	Case           bool   `json:"case"`
	SoundUrl       string `json:"soundUrl"`
	Color          string `json:"color"`
}

type SelfHighlightRule struct {
	EnableSound bool `json:"enableSound"`
}

type FilteringList struct {
	Filters []FilterConfig `json:"filters"`
}

type FilterConfig struct {
	Name   string `json:"name"`
	Filter string `json:"filter"`
	Id     string `json:"id"`
}

type LinksConfig struct {
	LinkInfoTooltip   bool `json:"linkInfoTooltip"`
	UnshortLinks      bool `json:"unshortLinks"`
	DoubleClickToOpen bool `json:"doubleClickToOpen"`
}

type EmotesConfig struct {
	EmojiSet string `json:"emojiSet"`
}

type AppearanceConfig struct {
	Splitheader SplitheaderAppearanceConfig `json:"splitheader"`
	Messages    MessagesAppearanceConfig    `json:"messages"`
	BoldScale   float32                     `json:"boldScale"`
	UiScale2    float32                     `json:"UiScale2"`
}

type SplitheaderAppearanceConfig struct {
	ShowUptime      bool `json:"showUptime"`
	ShowViewerCount bool `json:"showViewerCount"`
}

type MessagesAppearanceConfig struct {
	FindAllUsernames  bool   `json:"findAllUsernames"`
	ShowMessageLength bool   `json:"showMessageLength"`
	ColorizeNicknames bool   `json:"colorizeNicknames"`
	ShowTimestamps    bool   `json:"showTimestamps"`
	TimestampFormat   string `json:"timestampFormat"`
}

type NotificationsSettings struct {
	Twitch          []string `json:"twitch"`
	EnablePlaySound bool     `json:"enablePlaySound"`
}

type WhispersConfig struct {
	HighlightInlineWhispers bool `json:"highlightInlineWhispers"`
}

type DebugConfig struct {
	ShowUnhandledIrcMessages bool `json:"showUnhandledIrcMessages"`
}

type ExternalsConfig struct {
	Streamlink StreamlinkConfig `json:"streamlink"`
}

type StreamlinkConfig struct {
	UseCustomPath bool   `json:"useCustomPath"`
	CustomPath    string `json:"customPath"`
}

type StreamerModeConfig struct {
	MuteMentions        bool `json:"muteMentions"`
	HideUsercardAvatars bool `json:"hideUsercardAvatars"`
	HideLinkThumbnails  bool `json:"hideLinkThumbnails"`
	Enabled             int  `json:"enabled"`
}

type LoggingConfig struct {
	Path    string `json:"path"`
	Enabled bool   `json:"enabled"`
}

type SimilarityConfig struct {
	HideSimilarMaxMessagesToCheck int `json:"hideSimilarMaxMessagesToCheck"`
}

type NicknameConfig struct {
	Name            string `json:"name"`
	Replace         string `json:"replace"`
	IsRegex         bool   `json:"isRegex"`
	IsCaseSensitive bool   `json:"isCaseSensitive"`
}
