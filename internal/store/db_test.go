package store

import (
	"testing"
)

func TestGetTrip(t *testing.T) {
	// t.Parallel()

	// path := filepath.Join(t.TempDir(), "testgettrip.json")
	// var err error
	// db, err = jsonfile.New[Store](path)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if _, err = os.Stat(path); err == nil {
	// 	defer os.Remove(path)
	// }

	// want := seedTrip()[0]
	// sampleData := []models.Trip{want}

	// db.Write(func(s *Store) error {
	// 	s.Trips = sampleData
	// 	return nil
	// })

	// // check to see trip got added to db
	// got := GetTrip("test")
	// if !reflect.DeepEqual(*got, want) {
	// 	t.Errorf("got %+v, want %+v", *got, want)
	// }

	// got = GetTrip("nonexistent trip")
	// if got != nil {
	// 	t.Errorf("got %+v, wanted nil", *got)
	// }

	// // check to see db is only what we added to it
	// db.Read(func(data *Store) {
	// 	if !reflect.DeepEqual(data.Trips, sampleData) {
	// 		t.Errorf("new trip not added to db")
	// 	}
	// })

}
