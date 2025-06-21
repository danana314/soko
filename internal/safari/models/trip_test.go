package models

import (
	"regexp"
	"testing"
	"time"
)

func TestNewTrip(t *testing.T) {
	trip := NewTrip()
	
	if trip == nil {
		t.Fatal("NewTrip() returned nil")
	}
	
	if trip.Id == "" {
		t.Error("NewTrip() should generate an ID")
	}
	
	// Verify ID format (shortuuid)
	matched, err := regexp.MatchString(`^[A-Za-z0-9]{20,25}$`, trip.Id)
	if err != nil {
		t.Errorf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("NewTrip() ID = %q, doesn't match expected shortuuid format", trip.Id)
	}
}

func TestNewUser(t *testing.T) {
	user := NewUser()
	
	if user == nil {
		t.Fatal("NewUser() returned nil")
	}
	
	if user.Id == "" {
		t.Error("NewUser() should generate an ID")
	}
	
	// Verify ID format (shortuuid)
	matched, err := regexp.MatchString(`^[A-Za-z0-9]{20,25}$`, user.Id)
	if err != nil {
		t.Errorf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("NewUser() ID = %q, doesn't match expected shortuuid format", user.Id)
	}
}

func TestNewExpense(t *testing.T) {
	expense := NewExpense()
	
	if expense == nil {
		t.Fatal("NewExpense() returned nil")
	}
	
	if expense.Id == "" {
		t.Error("NewExpense() should generate an ID")
	}
	
	// Verify ID format (shortuuid) 
	matched, err := regexp.MatchString(`^[A-Za-z0-9]{20,25}$`, expense.Id)
	if err != nil {
		t.Errorf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("NewExpense() ID = %q, doesn't match expected shortuuid format", expense.Id)
	}
}

func TestUniqueIDs(t *testing.T) {
	t.Run("trips have unique IDs", func(t *testing.T) {
		ids := make(map[string]struct{})
		for range 10 {
			trip := NewTrip()
			if _, exists := ids[trip.Id]; exists {
				t.Errorf("NewTrip() generated duplicate ID: %s", trip.Id)
			}
			ids[trip.Id] = struct{}{}
		}
	})
	
	t.Run("users have unique IDs", func(t *testing.T) {
		ids := make(map[string]struct{})
		for range 10 {
			user := NewUser()
			if _, exists := ids[user.Id]; exists {
				t.Errorf("NewUser() generated duplicate ID: %s", user.Id)
			}
			ids[user.Id] = struct{}{}
		}
	})
	
	t.Run("expenses have unique IDs", func(t *testing.T) {
		ids := make(map[string]struct{})
		for range 10 {
			expense := NewExpense()
			if _, exists := ids[expense.Id]; exists {
				t.Errorf("NewExpense() generated duplicate ID: %s", expense.Id)
			}
			ids[expense.Id] = struct{}{}
		}
	})
}

func TestSplitUserDate(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedUser  string
		expectedDate  time.Time
		expectErr     bool
	}{
		{
			name:         "valid user date string",
			input:        "user123_2023-12-25",
			expectedUser: "user123",
			expectedDate: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "different valid format",
			input:        "abc_2024-01-01", 
			expectedUser: "abc",
			expectedDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "missing underscore",
			input:     "user123-2023-12-25",
			expectErr: true,
		},
		{
			name:      "too many segments",
			input:     "user_123_2023-12-25",
			expectErr: true,
		},
		{
			name:      "invalid date format",
			input:     "user123_2023/12/25",
			expectErr: true,
		},
		{
			name:      "empty string",
			input:     "",
			expectErr: true,
		},
		{
			name:      "only underscore",
			input:     "_",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId, date, err := splitUserDate(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if userId != tt.expectedUser {
				t.Errorf("splitUserDate(%q) userId = %q, want %q", tt.input, userId, tt.expectedUser)
			}
			if !date.Equal(tt.expectedDate) {
				t.Errorf("splitUserDate(%q) date = %v, want %v", tt.input, date, tt.expectedDate)
			}
		})
	}
}

func TestNewScheduleEntry(t *testing.T) {
	users := []*User{
		{Id: "user1", Name: "Alice"},
		{Id: "user2", Name: "Bob"},
		{Id: "user3", Name: "Charlie"},
	}

	tests := []struct {
		name           string
		users          []*User
		userDateString string
		expectedUserId string
		expectedName   string
		expectedDate   time.Time
		expectErr      bool
	}{
		{
			name:           "valid schedule entry",
			users:          users,
			userDateString: "user1_2023-12-25",
			expectedUserId: "user1",
			expectedName:   "Alice",
			expectedDate:   time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			name:           "another valid entry",
			users:          users,
			userDateString: "user2_2024-01-01",
			expectedUserId: "user2",
			expectedName:   "Bob",
			expectedDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:           "user not found",
			users:          users,
			userDateString: "nonexistent_2023-12-25",
			expectErr:      true,
		},
		{
			name:           "invalid date format",
			users:          users,
			userDateString: "user1_invalid-date",
			expectErr:      true,
		},
		{
			name:           "empty users list",
			users:          []*User{},
			userDateString: "user1_2023-12-25",
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := NewScheduleEntry(tt.users, tt.userDateString)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if entry == nil {
				t.Fatal("NewScheduleEntry() returned nil")
			}
			if entry.User.Id != tt.expectedUserId {
				t.Errorf("NewScheduleEntry() User.Id = %q, want %q", entry.User.Id, tt.expectedUserId)
			}
			if entry.User.Name != tt.expectedName {
				t.Errorf("NewScheduleEntry() User.Name = %q, want %q", entry.User.Name, tt.expectedName)
			}
			if !entry.Date.Equal(tt.expectedDate) {
				t.Errorf("NewScheduleEntry() Date = %v, want %v", entry.Date, tt.expectedDate)
			}
		})
	}
}

func TestTripData_IsBooked(t *testing.T) {
	user1 := User{Id: "user1", Name: "Alice"}
	user2 := User{Id: "user2", Name: "Bob"}
	date1 := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2023, 12, 26, 0, 0, 0, 0, time.UTC)

	tripData := &TripData{
		Schedule: []*ScheduleEntry{
			{Date: date1, User: user1},
			{Date: date2, User: user1},
			{Date: date1, User: user2},
		},
	}

	tests := []struct {
		name     string
		user     User
		date     time.Time
		expected bool
	}{
		{
			name:     "user is booked on date",
			user:     user1,
			date:     date1,
			expected: true,
		},
		{
			name:     "user is booked on different date",
			user:     user1,
			date:     date2,
			expected: true,
		},
		{
			name:     "different user is booked on same date",
			user:     user2,
			date:     date1,
			expected: true,
		},
		{
			name:     "user not booked on date",
			user:     user2,
			date:     date2,
			expected: false,
		},
		{
			name:     "nonexistent user",
			user:     User{Id: "nonexistent", Name: "Ghost"},
			date:     date1,
			expected: false,
		},
		{
			name:     "nonexistent date",
			user:     user1,
			date:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tripData.IsBooked(tt.user, tt.date)
			if result != tt.expected {
				t.Errorf("IsBooked(%v, %v) = %v, want %v", tt.user, tt.date, result, tt.expected)
			}
		})
	}
}

func TestTripData_IsBooked_EmptySchedule(t *testing.T) {
	tripData := &TripData{
		Schedule: []*ScheduleEntry{},
	}

	user := User{Id: "user1", Name: "Alice"}
	date := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)

	result := tripData.IsBooked(user, date)
	if result {
		t.Error("IsBooked() should return false for empty schedule")
	}
}

func TestTripData_IsBooked_NilSchedule(t *testing.T) {
	tripData := &TripData{
		Schedule: nil,
	}

	user := User{Id: "user1", Name: "Alice"}
	date := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)

	result := tripData.IsBooked(user, date)
	if result {
		t.Error("IsBooked() should return false for nil schedule")
	}
}