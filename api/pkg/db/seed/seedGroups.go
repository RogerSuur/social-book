package seed

type SeedGroup struct {
	CreatorEmail   string
	Title          string
	Description    string
	Users          []string
	ImagePath      string
	SeedEventsData []*SeedEvent
}

var SeedGroupsData = []*SeedGroup{
	{
		CreatorEmail:   "b@b.com",
		Title:          "Adventurers United",
		Description:    "Adventurers United is a group dedicated to exploring uncharted territories, unraveling mysteries, and pushing the boundaries of human exploration. Comprising fearless individuals from different backgrounds, this group seeks to discover hidden wonders, encounter diverse cultures, and document their findings for the world to marvel at. From scaling towering peaks to delving into ancient ruins, Adventurers United is fueled by the passion for discovery and the thrill of venturing into the unknown.",
		Users:          []string{"a@a.com", "b@b.com", "c@c.com", "d@d.com", "e@e.com"},
		SeedEventsData: SeedEventsDataA,
		ImagePath:      "groupA.png",
	},
	{
		CreatorEmail: "a@a.com",
		Title:        "Blissful Harmony",
		Description:  "Blissful Harmony is a group dedicated to promoting inner peace and cultivating a harmonious existence through mindfulness, meditation, and holistic practices.",
		Users:        []string{"a@a.com", "d@d.com", "e@e.com"},
		ImagePath:    "groupB.png",
	},
	{
		CreatorEmail:   "b@b.com",
		Title:          "Creative Catalysts",
		Description:    "Creative Catalysts is a gathering of visionary artists, innovators, and thinkers who believe in the transformative power of creativity. Embracing diverse forms of artistic expression, this group seeks to challenge conventions, provoke thought, and inspire change through their work. From visual arts and music to literature and performance, Creative Catalysts use their talents to ignite conversations, bridge gaps, and breathe life into new ideas. Their collective energy fuels a vibrant and dynamic creative community.",
		Users:          []string{"a@a.com", "b@b.com", "c@c.com", "d@d.com"},
		SeedEventsData: SeedEventsDataC,
		ImagePath:      "groupC.png",
	},
	{
		CreatorEmail: "c@c.com",
		Title:        "Dreamcatchers",
		Description:  "Dreamcatchers is a group dedicated to empowering individuals to pursue their dreams and aspirations. Recognizing the importance of nurturing ambitions, this group provides a supportive platform for members to share their goals, seek guidance, and celebrate achievements. Through mentorship programs, motivational workshops, and networking opportunities, Dreamcatchers aim to inspire personal growth, foster resilience, and create a community of dreamers who believe that anything is possible with dedication and perseverance.",
		Users:        []string{"c@c.com", "d@d.com", "e@e.com"},
		ImagePath:    "groupD.png",
	},
	{
		CreatorEmail: "b@b.com",
		Title:        "Ecological Guardians",
		Description:  "Ecological Guardians is a passionate group of environmental advocates committed to preserving and protecting the planet. Driven by a deep concern for the Earth's well-being, they actively engage in conservation efforts, sustainable practices, and education campaigns to raise awareness about environmental issues. From organizing clean-up drives to promoting renewable energy solutions, Ecological Guardians work tirelessly to safeguard ecosystems, combat climate change, and inspire others to adopt eco-friendly lifestyles for a greener future.",
		Users:        []string{"b@b.com"},
		ImagePath:    "groupE.png",
	},
	{
		CreatorEmail: "h@h.com",
		Title:        "Friends of Nature",
		Description:  "Friends of Nature is a community of nature enthusiasts who are passionate about exploring and preserving the natural world. Through their shared love for the environment, they engage in various outdoor activities, such as hiking, wildlife observation, and nature photography. Members of Friends of Nature also actively participate in conservation projects, organize awareness campaigns, and advocate for sustainable practices to protect ecosystems and biodiversity.",
		Users:        []string{"a@a.com", "b@b.com", "c@c.com", "d@d.com", "e@e.com", "h@h.com"},
		ImagePath:    "groupF.png",
	},
	{
		CreatorEmail: "h@h.com",
		Title:        "Global Changemakers",
		Description:  "Global Changemakers is a diverse community of individuals dedicated to making a positive impact on the world. The group focuses on addressing various global challenges, including poverty, inequality, and environmental degradation. Members of Global Changemakers collaborate on social initiatives, volunteer in local communities, and support sustainable development projects. Through their collective efforts, they strive to create lasting change and build a more equitable and sustainable future for all.",
		Users:        []string{"h@h.com"},
		ImagePath:    "",
	},
	{
		CreatorEmail: "a@a.com",
		Title:        "Health and Wellness Enthusiasts",
		Description:  "Health and Wellness Enthusiasts is a vibrant group of individuals passionate about promoting holistic well-being. The group focuses on sharing knowledge, tips, and resources related to physical fitness, mental health, nutrition, and self-care practices. Members of Health and Wellness Enthusiasts engage in discussions, organize fitness challenges, and support each other on their wellness journeys. By fostering a supportive and inclusive community, the group aims to inspire others to prioritize their health and embrace a balanced lifestyle.",
		Users:        []string{"a@a.com", "d@d.com", "e@e.com"},
		ImagePath:    "groupH.png",
	},
	{
		CreatorEmail: "c@c.com",
		Title:        "Innovation Hub",
		Description:  "Innovation Hub is a collaborative space for creative thinkers, entrepreneurs, and innovators to share ideas and explore new possibilities. The group encourages discussions on emerging technologies, startup ventures, and breakthrough innovations across various industries. Members of the Innovation Hub network, exchange expertise, provide feedback, and foster a supportive environment for turning ideas into reality. With a shared passion for innovation, the group aims to inspire and drive positive change in the world through entrepreneurial endeavors.",
		Users:        []string{"c@c.com", "e@e.com"},
		ImagePath:    "",
	},
	{
		CreatorEmail: "j@j.com",
		Title:        "Joyful Journeyers",
		Description:  "Joyful Journeyers is a community of travel enthusiasts who believe in the transformative power of exploration and cultural immersion. The group shares travel experiences, tips, and recommendations to inspire others to embark on enriching journeys. Members of Joyful Journeyers engage in conversations about different destinations, local customs, and sustainable travel practices. They also organize group trips and collaborate on travel-related projects aimed at promoting responsible tourism and cross-cultural understanding. Joyful Journeyers encourages members to embrace the joy of travel and create unforgettable experiences.",
		Users:        []string{"b@b.com", "j@j.com", "f@f.com", "h@h.com"},
		ImagePath:    "",
	},
}
