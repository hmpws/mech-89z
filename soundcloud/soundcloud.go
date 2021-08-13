package soundcloud

import (
   "fmt"
   "net/http"
   "strings"
)

// API is a wrapper for the SoundCloud private API used internally for soundcloud.com
type API struct {
	client              *client
	StripMobilePrefix   bool
	ConvertFirebaseURLs bool
}

// APIOptions are the options for creating an API struct
type APIOptions struct {
	ClientID            string       // optional and a new one will be fetched if not provided
	HTTPClient          *http.Client // the HTTP client to make requests with
	StripMobilePrefix   bool         // whether or not to convert mobile URLs to regular URLs
	ConvertFirebaseURLs bool         // whether or not to convert SoundCloud firebase URLs to regular URLs
}

// New returns a pointer to a new SoundCloud API struct.
func New(options APIOptions) (*API, error) {
	if options.ClientID == "" {
		var err error
		options.ClientID, err = FetchClientID()
		if err != nil {
                        return nil, err
		}
	}
	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}
	return &API{
		client:              newClient(options.ClientID, options.HTTPClient),
		StripMobilePrefix:   options.StripMobilePrefix,
		ConvertFirebaseURLs: options.ConvertFirebaseURLs,
	}, nil
}

// GetDownloadURL retuns the URL to download a track. This is useful if you
// want to implement your own downloading algorithm. If the track has a
// publicly available download link, that link will be preferred and the
// streamType parameter will be ignored. streamType can be either "hls" or
// "progressive", defaults to "progressive"
func (sc *API) GetDownloadURL(url string, streamType string) (string, error) {
   url, err := sc.prepareURL(url)
   if err != nil {
      return "", err
   }
   streamType = strings.ToLower(streamType)
   if streamType == "" {
      streamType = "progressive"
   }
   if IsURL(url, false, false) {
      info, err := sc.client.getTrackInfo(GetTrackInfoOptions{
      URL: url,
      })
      if err != nil {
         return "", err
      }
      if len(info) == 0 {
         return "", fmt.Errorf("%v fail", url)
      }
      if info[0].Downloadable && info[0].HasDownloadsLeft {
      downloadURL, err := sc.client.getDownloadURL(info[0].ID)
      if err != nil {
         return "", err
      }
      return downloadURL, nil
      }
      for _, transcoding := range info[0].Media.Transcodings {
      if strings.ToLower(transcoding.Format.Protocol) == streamType {
      mediaURL, err := sc.client.getMediaURL(transcoding.URL)
      if err != nil {
      return "", err
      }
      return mediaURL, nil
      }
      }
      mediaURL, err := sc.client.getMediaURL(info[0].Media.Transcodings[0].URL)
      if err != nil {
      return "", err
      }
      return mediaURL, nil
   }
   return "", fmt.Errorf("%v is not a track URL", url)
}

func (sc *API) prepareURL(url string) (string, error) {
   if sc.StripMobilePrefix {
      if IsMobileURL(url) {
         url = StripMobilePrefix(url)
      }
   }
   return url, nil
}

// Track represents the JSON response of a track's info
type Track struct {
   ArtworkURL        string `json:"artwork_url"`
   CommentCount      int64  `json:"comment_count"`
   Commentable       bool   `json:"commentable"`
   CreatedAt         string `json:"created_at"`
   Description       string `json:"description"`
   DisplayDate       string `json:"display_date"`
   DownloadCount     int64  `json:"download_count"`
   Downloadable      bool   `json:"downloadable"`
   DurationMS        int64  `json:"duration"`
   FullDurationMS    int64  `json:"full_duration"`
   Genre             string `json:"genre"`
   HasDownloadsLeft  bool   `json:"has_downloads_left"`
   ID                int64  `json:"id"`
   Kind              string `json:"kind"`
   LabelName         string `json:"label_name"`
   LastModified      string `json:"last_modified"`
   LikesCount        int64  `json:"likes_count"`
   Media             Media  `json:"media"`
   MonetizationModel string `json:"monetization_model"`
   Permalink         string `json:"permalink"`
   PermalinkURL      string `json:"permalink_url"`
   PlaybackCount     int64  `json:"playback_count"`
   Policy            string `json:"polic"`
   Public            bool   `json:"public"`
   RepostsCount      int64  `json:"reposts_count"`
   SecretToken       string `json:"secret_token"`
   Streamable        bool   `json:"streamable"`
   TagList           string `json:"tag_list"`
   Title             string `json:"title"`
   URI               string `json:"uri"`
   User              User   `json:"user"`
   UserID            int64  `json:"user_id"`
   WaveformURL       string `json:"waveform_url"`
}

// Media contains an array of transcoding for a track
type Media struct {
	Transcodings []Transcoding `json:"transcodings"`
}

// Transcoding contains information about the transcoding of a track
type Transcoding struct {
	URL     string            `json:"url"`
	Preset  string            `json:"preset"`
	Snipped bool              `json:"snipped"`
	Format  TranscodingFormat `json:"format"`
}

// TranscodingFormat contains the protocol by which the track is delivered ("progressive" or "HLS"), and
// the mime type of the track
type TranscodingFormat struct {
	Protocol string `json:"protocol"`
	MimeType string `json:"mime_type"`
}

// Playlist represents the JSON response of a playlist
type Playlist struct {
	ArtworkURL     string  `json:"artwork_url"`
	CreatedAt      string  `json:"created_at"`
	Description    string  `json:"description"`
	DurationMS     int64   `json:"duration"`
	EmbeddeableBy  string  `json:"embeddable_by"`
	Genre          string  `json:"genre"`
	ID             int64   `json:"id"`
	Kind           string  `json:"kind"`
	LabelName      string  `json:"label_name"`
	LastModified   string  `json:"last_modified"`
	License        string  `json:"license"`
	LikesCount     int     `json:"likes_count"`
	ManagedByFeeds bool    `json:"managed_by_feeds"`
	Permalink      string  `json:"permalink"`
	PermalinkURL   string  `json:"permalink_url"`
	Public         bool    `json:"public"`
	SecretToken    string  `json:"secret_token"`
	Sharing        string  `json:"private"`
	TagList        string  `json:"tag_list"`
	Title          string  `json:"title"`
	URI            string  `json:"uri"`
	UserID         int64   `json:"user_id"`
	SetType        string  `json:"set_type"`
	IsAlbum        bool    `json:"is_album"`
	PublishedAt    string  `json:"published_at"`
	DisplayDate    string  `json:"display_date"`
	User           User    `json:"user"`
	Tracks         []Track `json:"tracks"`
	TrackCount     int     `json:"track_count"`
}

// User represents the JSON payload for user data
type User struct {
	ID              int64  `json:"id"`
	AvatarURL       string `json:"avatar_url"`
	City            string `json:"city"`
	CommentsCount   int64  `json:"comments_count"`
	CountryCode     string `json:"country_code"`
	CreatedAt       string `json:"created_at"`
	Description     string `json:"description"`
	FollowersCount  int64  `json:"followers_count"`
	FollowingsCount int64  `json:"followings_count"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PermalinkURL    string `json:"permalink_url"`
	URI             string `json:"uri"`
	Username        string `json:"username"`
	Kind            string `json:"kind"`
	Likes           int    `json:"likes_count"`
	PlaylistLikes   int    `json:"playlist_likes_count"`
	Verified        bool   `json:"verified"`
}

// MediaURLResponse is the JSON response of retrieving media information of a track
type MediaURLResponse struct {
	URL string `json:"url"`
}

// DownloadURLResponse is the JSON respose of retrieving media information of a publicly downloadable track
type DownloadURLResponse struct {
	URL string `json:"redirectUri"`
}

// PaginatedQuery is the JSON response for a paginated query
type PaginatedQuery struct {
	Collection   []map[string]interface{} `json:"collection"`
	TotalResults int                      `json:"total_results"`
	NextHref     string                   `json:"next_href"`
	QueryURN     string                   `json:"query_urn"`
}

// Like is the JSON response for a like
type Like struct {
	CreatedAt string   `json:"created_at"`
	Kind      string   `json:"kind"`
	Track     Track    `json:"track"`
	Playlist  Playlist `json:"playlist"`
}
