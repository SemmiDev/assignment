package db

// DBName is the name of the database.
type DBName = string

// Data is the data representation of a database.
type Data = []byte

// Loader is the interface for loading a database.
type Loader interface {
	// Load loads the database.
	Load(DBName) ([]byte, error)
}

// Resetter is the interface for resetting a database.
type Resetter interface {
	// Reset resets the database.
	Reset(DBName, Data) error
}

// Saver is the interface for saving a database.
type Saver interface {
	// Save saves the database.
	Save(DBName, Data) error
}

// DB is the interface for a database. it combines the Loader, Resetter, and Saver interfaces.
type DB interface {
	Loader
	Resetter
	Saver
}
