package seed

type SeedGroup struct {
	CreatorEmail string
	Title        string
	Description  string
	Users        []string
	// ImagePath string
}

// type SeedGroupUsers struct {
// 	UserEmail string
// }

var SeedGroupsData = []*SeedGroup{
	{
		CreatorEmail: "b@b.com",
		Title:        "Adventurers United",
		Description:  "Adventurers United is a group dedicated to exploring uncharted territories, unraveling mysteries, and pushing the boundaries of human exploration. Comprising fearless individuals from different backgrounds, this group seeks to discover hidden wonders, encounter diverse cultures, and document their findings for the world to marvel at. From scaling towering peaks to delving into ancient ruins, Adventurers United is fueled by the passion for discovery and the thrill of venturing into the unknown.",
		Users:        []string{"a@a.com", "b@b.com", "c@c.com", "d@d.com", "e@e.com"},
	},
	{
		CreatorEmail: "a@a.com",
		Title:        "Blissful Harmony",
		Description:  "Blissful Harmony is a group dedicated to promoting inner peace and cultivating a harmonious existence through mindfulness, meditation, and holistic practices.",
		Users:        []string{"a@a.com", "d@d.com", "e@e.com"},
	},
	{
		CreatorEmail: "b@b.com",
		Title:        "Creative Catalysts",
		Description:  "Creative Catalysts is a gathering of visionary artists, innovators, and thinkers who believe in the transformative power of creativity. Embracing diverse forms of artistic expression, this group seeks to challenge conventions, provoke thought, and inspire change through their work. From visual arts and music to literature and performance, Creative Catalysts use their talents to ignite conversations, bridge gaps, and breathe life into new ideas. Their collective energy fuels a vibrant and dynamic creative community.",
		Users:        []string{"a@a.com", "b@b.com", "c@c.com", "d@d.com"},
	},
	{
		CreatorEmail: "c@c.com",
		Title:        "Dreamcatchers",
		Description:  "Dreamcatchers is a group dedicated to empowering individuals to pursue their dreams and aspirations. Recognizing the importance of nurturing ambitions, this group provides a supportive platform for members to share their goals, seek guidance, and celebrate achievements. Through mentorship programs, motivational workshops, and networking opportunities, Dreamcatchers aim to inspire personal growth, foster resilience, and create a community of dreamers who believe that anything is possible with dedication and perseverance.",
		Users:        []string{"c@c.com", "d@d.com", "e@e.com"},
	},
	{
		CreatorEmail: "b@b.com",
		Title:        "Ecological Guardians",
		Description:  "Ecological Guardians is a passionate group of environmental advocates committed to preserving and protecting the planet. Driven by a deep concern for the Earth's well-being, they actively engage in conservation efforts, sustainable practices, and education campaigns to raise awareness about environmental issues. From organizing clean-up drives to promoting renewable energy solutions, Ecological Guardians work tirelessly to safeguard ecosystems, combat climate change, and inspire others to adopt eco-friendly lifestyles for a greener future.",
		Users:        []string{"b@b.com"},
	},
}
