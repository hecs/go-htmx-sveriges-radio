package main

type PageData struct {
	Title     string
	Programs  []Program
	Copyright string `json:"copyright"`
}

type LiveAudio struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	StatKey string `json:"statkey"`
}

type Channel struct {
	Image         string     `json:"image"`
	ImageTemplate string     `json:"imagetemplate"`
	Color         string     `json:"color"`
	SiteURL       string     `json:"siteurl"`
	LiveAudio     *LiveAudio `json:"liveaudio"`
	ScheduleURL   string     `json:"scheduleurl"`
	ChannelType   string     `json:"channeltype"`
	XMLTvID       string     `json:"xmltvid"`
	ID            int        `json:"id"`
	Name          string     `json:"name"`
}

type ProgramCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Program struct {
	Description          string           `json:"description"`
	ProgramCategory      *ProgramCategory `json:"programcategory"`
	Payoff               string           `json:"payoff"`
	BroadcastInfo        string           `json:"broadcastinfo"`
	Email                string           `json:"email"`
	Phone                string           `json:"phone"`
	ProgramURL           string           `json:"programurl"`
	ProgramImage         string           `json:"programimage"`
	ProgramImageTemplate string           `json:"programimagetemplate"`
	SocialImage          string           `json:"socialimage"`
	SocialImageTemplate  string           `json:"socialimagetemplate"`
	Channel              *Channel         `json:"channel"`
	Archived             bool             `json:"archived"`
	HasOndemand          bool             `json:"hasondemand"`
	HasPod               bool             `json:"haspod"`
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
}

type Podfile struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	FileSizeInBytes int      `json:"filesizeinbytes"`
	Program         *Program `json:"program"`
	Duration        int      `json:"duration"`
	//	PublishDate     time.Time `json:"publishdateutc"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	StatKey string `json:"statkey"`
}

type Playlist struct {
	Song     *Song    `json:"song"`
	NextSong *Song    `json:"nextsong"`
	Channel  *Channel `json:"channel"`
}

// Song represents a song
type Song struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Artist      string `json:"artist"`
	Composer    string `json:"composer"`
	Conductor   string `json:"conductor"`
}

type ScheduledEpisode struct {
	EpisodeID int    `json:"episodeid"`
	Title     string `json:"title"`
	//	StartTime time.Time `json:"starttimeutc"`
	//	EndTime   time.Time `json:"endtimeutc"`
	Program *Program `json:"program"`
	Channel *Channel `json:"channel"`
}

type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SportBroadcast struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	LocalStartTime  string     `json:"localstarttime"`
	LocalStopTime   string     `json:"localstoptime"`
	Hometeam        *IDName    `json:"hometeam"`
	Awayteam        *IDName    `json:"awayteam"`
	League          *IDName    `json:"league"`
	Season          *IDName    `json:"season"`
	Arena           *IDName    `json:"arena"`
	Sport           *IDName    `json:"sport"`
	Publisher       *IDName    `json:"publisher"`
	Channel         *IDName    `json:"channel"`
	LiveAudio       *LiveAudio `json:"liveaudio"`
	MobileLiveAudio *LiveAudio `json:"mobileliveaudio"`
}

type Show struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	//	Date             time.Time  `json:"dateutc"`
	Type             string     `json:"type"`
	Program          *Program   `json:"program"`
	ImageURL         string     `json:"imageurl"`
	ImageURLTemplate string     `json:"imageurltemplate"`
	Broadcast        *Broadcast `json:"broadcast"`
}

// Asset represents a radio asset
type Asset struct {
	Duration int `json:"duration"`
	//	PublishDate time.Time `json:"publishdateutc"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	StatKey string `json:"statkey"`
}

// Broadcast represents a radio broadcast
type Broadcast struct {
	//	AvailableStop  time.Time `json:"availablestoputc"`
	Playlist       Asset   `json:"playlist"`
	Broadcastfiles []Asset `json:"broadcastfiles"`
}

type BroadcastTime struct {
	//StartTimeUTC time.Time `json:"starttimeutc"`
	//EndTimeUTC   time.Time `json:"endtimeutc"`
}

type ListenPodfile struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	FileSizeInBytes int     `json:"filesizeinbytes"`
	Program         Program `json:"program"`
	//	AvailableFromUTC time.Time `json:"availablefromutc"`
	Duration int `json:"duration"`
	//	PublishDateUTC   time.Time `json:"publishdateutc"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	StatKey string `json:"statkey"`
}

type DownloadPodfile struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	FileSizeInBytes int     `json:"filesizeinbytes"`
	Program         Program `json:"program"`
	//	AvailableFromUTC time.Time `json:"availablefromutc"`
	Duration int `json:"duration"`
	//	PublishDateUTC   time.Time `json:"publishdateutc"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	StatKey string `json:"statkey"`
}

type Episode struct {
	ID                int             `json:"id"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	URL               string          `json:"url"`
	Program           Program         `json:"program"`
	AudioPreference   string          `json:"audiopreference"`
	AudioPriority     string          `json:"audiopriority"`
	AudioPresentation string          `json:"audiopresentation"`
	PublishDateUTC    string          `json:"publishdateutc"`
	ImageURL          string          `json:"imageurl"`
	ImageURLTemplate  string          `json:"imageurltemplate"`
	Photographer      string          `json:"photographer"`
	BroadcastTime     BroadcastTime   `json:"broadcasttime"`
	ListenPodfile     ListenPodfile   `json:"listenpodfile"`
	DownloadPodfile   DownloadPodfile `json:"downloadpodfile"`
}

type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalHits  int `json:"totalhits"`
	TotalPages int `json:"totalpages"`
}

type EpisodesResponse struct {
	Copyright  string     `json:"copyright"`
	Episodes   []Episode  `json:"episodes"`
	Pagination Pagination `json:"pagination"`
}
