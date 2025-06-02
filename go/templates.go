package main // Or your appropriate package name

// PageData is a general struct for data passed to the main index.html
type PageData struct {
	Title       string
	HeaderData  HeaderData
	ChatData    ChatWindowData
}

// HeaderData holds data for templates/header.html
type HeaderData struct {
	Title      string
	Stylesheets []string
	Scripts     []string
}

// ChatWindowData holds data for templates/chat_window.html
type ChatWindowData struct {
	FormPostURL       string // e.g., "/generate"
	InitialMessages   []Message // For pre-populating messages
	// Any other specific data the chat window might need
}

// Message is a generic struct for a chat message
// Used by ChatWindowData and potentially by image generation logic
type Message struct {
	ID                string // Optional: for unique identification
	Content           string
	IsBot             bool
	ImageCarouselData *ImageCarouselWrapperData // Pointer, as not all messages have carousels
}

// ImageCarouselWrapperData holds data for templates/image_carousel_wrapper.html
type ImageCarouselWrapperData struct {
	MessageText   string         // The original text prompt for the image
	FirstImage    ImageSlideData // Data for the initially displayed image
	// Potentially IDs or classes for targeting if needed by HTMX/JS from Go
}

// ImageSlideData holds data for templates/image_slide.html
type ImageSlideData struct {
	ImageURL      string
	AltText       string
	ImageID       string // Optional: for carousel navigation
}

// You might also have a struct for data passed to the /generate/image endpoint,
// which could include the MessageText.
// For example:
// type GenerateImageRequest struct {
//	 MessageText string `json:"message_text"`
// }

// Note: The actual message structure for HTMX responses from /generate might be simpler,
// directly returning HTML snippets. These structs are for the initial page load via Go templates.
// The server would decide which template to render (e.g. image_slide.html or image_carousel_wrapper.html)
// and populate it with the necessary data.
