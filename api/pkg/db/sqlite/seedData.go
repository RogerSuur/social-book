package database

import "SocialNetworkRestApi/api/pkg/enums"

type SeedUser struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Nickname  string
	About     string
	//ImagePath string
	IsPublic        bool
	PostSet         []*SeedPost
	FollowingEmails []string
}

type SeedPost struct {
	Id      int
	Content string
	// ImagePath   string
	PrivacyType enums.PrivacyType
}

type SeedComment struct {
}

var SeedUserData = []*SeedUser{
	{
		FirstName:       "Ann",
		LastName:        "Addams",
		Email:           "a@a.com",
		Password:        "a123",
		Nickname:        "AnnieA",
		About:           "Hey there, I'm Ann! I'm a fitness enthusiast and nutrition coach. I'm passionate about helping others achieve their health and wellness goals through personalized meal plans and exercise routines. In my free time, I enjoy trying out new healthy recipes, practicing yoga, and spending time with my two rescue dogs.",
		IsPublic:        true,
		PostSet:         SeedPostsDataSetA,
		FollowingEmails: []string{"c@c.com"},
	},
	{
		FirstName: "Benjamin",
		LastName:  "Button",
		Email:     "b@b.com",
		Password:  "b123",
		Nickname:  "Buttons",
		About:     "Hi, I'm Benjamin! I'm a freelance writer and avid traveler. I love exploring new cultures and cuisines, and I'm always on the lookout for my next adventure. When I'm not writing or traveling, you can find me hiking, reading a good book, or practicing my photography skills.",
		IsPublic:  true,
		PostSet:   SeedPostsDataSetB,
	},
	{
		FirstName: "Carlos",
		LastName:  "Cortez",
		Email:     "c@c.com",
		Password:  "c123",
		Nickname:  "Carlito",
		About:     "Hi, my name is Carlos! I'm a software developer with a love for all things tech. I specialize in building mobile and web applications, and I'm always looking for new and innovative ways to solve complex problems through code. When I'm not coding, you can find me playing video games or tinkering with my latest DIY project.",
		IsPublic:  false,
		PostSet:   SeedPostsDataSetC,
	},
}

var SeedPostsDataSetA = []*SeedPost{
	{
		Content:     "Just finished a grueling but rewarding 10K race! Who else loves the feeling of crossing the finish line?",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Excited to announce that my new cookbook, \"Healthy Eats for Busy Lives\", is now available for pre-order! It's packed with easy and nutritious recipes for people on the go.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Who else loves practicing yoga? I just got certified as a yoga teacher and I can't wait to share my love for this practice with others.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Happy World Vegan Day! Being vegan has been one of the best decisions I've ever made for my health and the environment.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Just got back from an amazing week-long wellness retreat in Costa Rica. Yoga, healthy food, and beautiful scenery - what more could you ask for?",
		PrivacyType: enums.Public,
	},
	{
		Content:     "I recently tried a new plant-based burger and it was delicious! Who says vegan food has to be boring?",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Who else struggles with meal planning? I'm hosting a free webinar next week on how to create a personalized meal plan that works for you. Sign up now!",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Just finished a challenging but rewarding HIIT workout. Who else loves a good sweat session?",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Happy International Women's Day! Let's celebrate the amazing women in our lives and work towards a more equal and just world.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Who else loves taking their dog for a walk? My two rescue dogs, Luna and Max, are my favorite workout buddies.",
		PrivacyType: enums.Public,
	},
}

var SeedPostsDataSetB = []*SeedPost{
	{
		Content:     "Just got back from an incredible trip to Japan! The food, the people, and the culture were all amazing. Can't wait to go back someday!",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Who else loves a good outdoor adventure? Just went on a challenging hike up Mount Kilimanjaro and it was definitely worth it. What's your favorite hike?",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Excited to share that my latest travel article was published in National Geographic! It's all about the hidden gems of Barcelona. Check it out if you're planning a trip there soon!",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Recently started a photography project where I take a photo of something beautiful every day. It's been a great way to appreciate the little things in life.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Happy National Book Lovers Day! I just finished \"The Overstory\" by Richard Powers and it's definitely one of my new favorites. What are you currently reading?",
		PrivacyType: enums.Public,
	},
	{
		Content:     "I'm officially a certified scuba diver! I've always been fascinated by the ocean and it was incredible to see all the marine life up close.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Anyone else a fan of street food? I recently tried the best tacos al pastor I've ever had in Mexico City. Already planning my next trip back!",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Just got back from an amazing trip to Bali. The beaches, the temples, and the food were all incredible. Can't wait to go back someday!",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Just finished a 10-day silent meditation retreat and it was one of the most challenging and rewarding experiences of my life. Highly recommend it to anyone interested in mindfulness and inner peace.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Happy Earth Day! Let's all do our part to protect our planet and make it a better place for future generations.",
		PrivacyType: enums.Public,
	},
}

var SeedPostsDataSetC = []*SeedPost{
	{
		Content:     "Just launched my latest app, \"TaskMaster\", on the App Store! It's a productivity app that helps you stay on top of your to-do list. Check it out!",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Who else loves a good hackathon? Just won first place at the HackNY hackathon with my team. Can't wait for the next one.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Excited to announce that I've been accepted into the Google Developer Expert program for mobile development! It's an honor to be part of this community of experts.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Just wrapped up a successful project with a Fortune 500 company. It was a challenging but rewarding experience, and I'm proud of what our team accomplished.",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Excited to share that I'll be speaking at the upcoming TechCrunch Disrupt conference about the future of mobile development. Can't wait to share my insights with the tech community!",
		PrivacyType: enums.Public,
	},
	{
		Content:     "Feeling a bit burnt out lately. The tech industry can be so demanding sometimes, and I feel like I'm always on call. Trying to take some time for self-care and relaxation, but it's tough when there's always another deadline looming.",
		PrivacyType: enums.Private,
	},
	{
		Content:     "Had a heart-to-heart with my mentor today about imposter syndrome. It's something that's been weighing on me lately, but it was helpful to hear that even seasoned developers experience it from time to time.",
		PrivacyType: enums.Private,
	},
	{
		Content:     "Dealing with a bit of imposter syndrome lately. I keep worrying that I'm not skilled enough or experienced enough to tackle the projects I'm working on. Trying to remind myself that I wouldn't have gotten this far if I didn't have the skills and knowledge to back it up.",
		PrivacyType: enums.Private,
	},
	{
		Content:     "Just got back from a weekend getaway with my partner. It was so nice to unplug from work and spend some quality time together.",
		PrivacyType: enums.Private,
	},
	{
		Content:     "Feeling grateful for my team today. We've been working on a really challenging project, but everyone has been pulling their weight and pushing us towards success. It's great to work with such talented and dedicated individuals.",
		PrivacyType: enums.Private,
	},
}
