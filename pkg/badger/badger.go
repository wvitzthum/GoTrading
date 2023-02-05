package badger

import (
	badger "github.com/dgraph-io/badger/v3"
	"golang.org/x/exp/slog"
)

type BadgerClient struct {
	logger *slog.Logger
	client *badger.DB
}

func InitBadger(logger *slog.Logger, path string) (*BadgerClient, error) {
  	db, err := badger.Open(badger.DefaultOptions(path))
	return &BadgerClient{logger: logger, client: db}, err
}

func (b *BadgerClient) InsertValue(key, val string) error {
	err := b.client.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(val))
		return err
	})
	return err
}
func (b *BadgerClient) Close() error {
	return b.client.Close()
}

func (b *BadgerClient) RetrieveValuesForKey(symbol string) error {
	b.client.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(symbol)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		  item := it.Item()
		  k := item.Key()
		  err := item.Value(func(v []byte) error {
			b.logger.Info("Found item", string(k), string(v))
			return nil
		  })
		  if err != nil {
			return err
		  }
		}
		return nil
	  })
	return nil
}