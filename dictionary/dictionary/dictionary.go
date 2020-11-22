package dictionary

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"time"
)

// Dictionary structure de la bdd
type Dictionary struct {
	db *badger.DB
}

// Entry données
type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

// Fonction de formatage des données
func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

// New nouvelle instance de dictionary
func New(dir string) (*Dictionary, error) {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("./badger").WithLoggingLevel(badger.WARNING))
	if err != nil {
		return nil, err
	}
	dict := &Dictionary{
		db: db,
	}

	return dict, nil
}

// Close fermeture de la bdd
func (d *Dictionary) Close() {
	d.db.Close()
}
