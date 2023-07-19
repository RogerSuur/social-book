package seed

import "time"

type SeedEvent struct {
	CreatorEmail string
	CreatedAt    time.Time
	EventTime    time.Time
	TimeSpan     time.Duration
	Title        string
	Description  string
}

type SeedEventAttendance struct {
	UserId      int64
	EventId     int64
	IsAttending bool
}

var SeedEventsDataA = []*SeedEvent{
	{
		CreatorEmail: "b@b.com",
		CreatedAt:    time.Now().Add(time.Hour * -2),
		EventTime:    time.Now().Add(time.Hour * 100),
		TimeSpan:     time.Hour * 2,
		Title:        "Lost City Expedition",
		Description:  "Join us on an epic quest as we journey deep into uncharted lands to uncover the secrets of a long-lost civilization. Are you ready for the adventure of a lifetime?",
	},
	{
		CreatorEmail: "b@b.com",
		CreatedAt:    time.Now().Add(time.Hour * -10),
		EventTime:    time.Now().Add(time.Hour * 48),
		TimeSpan:     time.Minute * 45,
		Title:        "Mystic Forest Trek",
		Description:  "Step into the mystical realm of ancient forests as we embark on a challenging trek, encountering rare flora and fauna along the way. Brace yourself for enchanting encounters and breathtaking vistas.",
	},
	{
		CreatorEmail: "a@a.com",
		CreatedAt:    time.Now().Add(time.Hour * -24),
		EventTime:    time.Now().Add(time.Hour * 400),
		TimeSpan:     time.Hour * 48,
		Title:        "Summit Conquest: Mount Everest",
		Description:  "Conquer the world's highest peak with a team of daring adventurers. Experience the ultimate test of endurance and witness awe-inspiring views from the top of Mount Everest. Are you up for the challenge?",
	},
	{
		CreatorEmail: "a@a.com",
		CreatedAt:    time.Now().Add(time.Hour * -25),
		EventTime:    time.Now().Add(time.Hour * 72),
		TimeSpan:     time.Hour * 1,
		Title:        "Underwater Odyssey: Dive into the Abyss",
		Description:  "Plunge into the depths of the ocean as we explore unexplored underwater realms teeming with vibrant marine life and hidden treasures. Discover the wonders that lie beneath the surface in this thrilling underwater adventure.",
	},
	{
		CreatorEmail: "a@a.com",
		CreatedAt:    time.Now().Add(time.Hour * -10),
		EventTime:    time.Now().Add(time.Hour * 72),
		TimeSpan:     time.Hour * 24,
		Title:        "Ancient Temple Quest: Unveiling Mysteries",
		Description:  "Embark on a quest to uncover the secrets of an ancient temple shrouded in myth and legend. Navigate through intricate puzzles, decipher cryptic symbols, and delve into the past as we unlock the mysteries that lie within these sacred walls.",
	},
}

var SeedEventsDataC = []*SeedEvent{
	{
		CreatorEmail: "d@d.com",
		CreatedAt:    time.Now(),
		EventTime:    time.Now().Add(time.Hour * 48),
		TimeSpan:     time.Hour,
		Title:        "Artistic Fusion: Collaborative Exhibition",
		Description:  "Witness the convergence of creative minds as visionary artists from various disciplines come together to showcase their collaborative works. Experience the power of artistic fusion and be inspired by the vibrant energy that emerges when diverse talents unite.",
	},
}

var SeedEventAttendanceDataAccepted = []*SeedEventAttendance{
	{
		UserId:      2,
		EventId:     1,
		IsAttending: true,
	},
	{
		UserId:      2,
		EventId:     2,
		IsAttending: true,
	},
	{
		UserId:      1,
		EventId:     3,
		IsAttending: true,
	},
	{
		UserId:      1,
		EventId:     4,
		IsAttending: true,
	},
	{
		UserId:      1,
		EventId:     5,
		IsAttending: true,
	},
	{
		UserId:      2,
		EventId:     5,
		IsAttending: true,
	},
	{
		UserId:      3,
		EventId:     5,
		IsAttending: true,
	},
	{
		UserId:      1,
		EventId:     1,
		IsAttending: true,
	},
}

var SeedEventAttendanceDataPending = []*SeedEventAttendance{
	{
		UserId:      1,
		EventId:     2,
		IsAttending: false,
	},
	{
		UserId:      3,
		EventId:     2,
		IsAttending: false,
	},
	{
		UserId:      3,
		EventId:     1,
		IsAttending: false,
	},
	{
		UserId:      2,
		EventId:     3,
		IsAttending: false,
	},
}
