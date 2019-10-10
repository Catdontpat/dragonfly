package world

import "github.com/dragonfly-tech/dragonfly/dragonfly/world/chunk"

// Provider represents a value that may provide world data to a World value. It usually does the reading and
// writing of the world data so that the World may use it.
type Provider interface {
	// WorldName returns the name of the world that the provider provides for. When setting the provider of a
	// World, the World will replace its current name with this one.
	WorldName() string
	// LoadChunk attempts to load a chunk from the chunk position passed. If successful, a non-nil chunk is
	// returned and exists is true and err nil. If no chunk was saved at the chunk position passed, the chunk
	// returned is nil, and so is the error. If the chunk did exist, but if the data was invalid, nil is
	// returned for the chunk and true, with a non-nil error.
	// If exists ends up false, the chunk at the position is instead newly generated by the world.
	LoadChunk(position ChunkPos) (c *chunk.Chunk, exists bool, err error)
	// SaveChunk saves a chunk at a specific position in the provider. If writing was not successful, an error
	// is returned.
	SaveChunk(position ChunkPos, c *chunk.Chunk) error
}

// NoIOProvider implements a Provider while not performing any disk I/O. It generates values on the run and
// dynamically, instead of reading and writing data.
type NoIOProvider struct{}

// SaveChunk ...
func (p NoIOProvider) SaveChunk(position ChunkPos, c *chunk.Chunk) error {
	return nil
}

// LoadChunk ...
func (p NoIOProvider) LoadChunk(position ChunkPos) (*chunk.Chunk, bool, error) {
	return nil, false, nil
}

// WorldName ...
func (p NoIOProvider) WorldName() string {
	return "World"
}