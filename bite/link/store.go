package link

// Store persists and retrieves [Link] values in an in-memory map.
type Store struct {
	links map[LinkKey]Link
}
