package database

import (
	"SocialNetworkRestApi/api/pkg/enums"
)

type SeedUser struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
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
	PrivacyType   enums.PrivacyType
	CommentSet    []*SeedComment
	LoremComments int
}

type SeedComment struct {
	UserEmail string
	Content   string
	// ImagePath string
}

var SeedUserData = []*SeedUser{
	{
		FirstName:       "Ann",
		LastName:        "Addams",
		Email:           "a@a.com",
		Nickname:        "AnnieA",
		About:           "Hey there, I'm Ann! I'm a fitness enthusiast and nutrition coach. I'm passionate about helping others achieve their health and wellness goals through personalized meal plans and exercise routines. In my free time, I enjoy trying out new healthy recipes, practicing yoga, and spending time with my two rescue dogs.",
		IsPublic:        true,
		PostSet:         SeedPostsDataSetA,
		FollowingEmails: []string{"c@c.com", "d@d.com"},
	},
	{
		FirstName:       "Benjamin",
		LastName:        "Button",
		Email:           "b@b.com",
		Nickname:        "Buttons",
		About:           "Hi, I'm Benjamin! I'm a freelance writer and avid traveler. I love exploring new cultures and cuisines, and I'm always on the lookout for my next adventure. When I'm not writing or traveling, you can find me hiking, reading a good book, or practicing my photography skills.",
		IsPublic:        true,
		PostSet:         SeedPostsDataSetB,
		FollowingEmails: []string{"a@a.com", "d@d.com"},
	},
	{
		FirstName:       "Carlos",
		LastName:        "Cortez",
		Email:           "c@c.com",
		Nickname:        "Carlito",
		About:           "Hi, my name is Carlos! I'm a software developer with a love for all things tech. I specialize in building mobile and web applications, and I'm always looking for new and innovative ways to solve complex problems through code. When I'm not coding, you can find me playing video games or tinkering with my latest DIY project.",
		IsPublic:        false,
		PostSet:         SeedPostsDataSetC,
		FollowingEmails: []string{"a@a.com", "d@d.com"},
	},
	{
		FirstName: "Deanna",
		LastName:  "Davis",
		Email:     "d@d.com",
		Nickname:  "DeeDee",
		About:     "Hi there, I'm Deanna Davis. I'm a freelance writer and digital marketer with a passion for creating compelling content that connects with audiences. When I'm not working, you can usually find me hiking with my dog or experimenting with new vegan recipes in the kitchen.",
		IsPublic:  true,
	},
	{
		FirstName: "Ethan",
		LastName:  "Evans",
		Email:     "e@e.com",
		Nickname:  "EvansCode",
		About:     "Hey, I'm Ethan Evans. I'm a software engineer with over 10 years of experience in the industry. I'm passionate about using technology to solve real-world problems and improve people's lives. When I'm not coding, I enjoy playing basketball and exploring new restaurants in the city.",
	},
	{
		FirstName: "Lorem",
		LastName:  "Ipsum",
		Email:     "l@l.com",
		Nickname:  "LoremGenerator",
		About:     "",
	},
}

var SeedPostsDataSetA = []*SeedPost{
	{
		Content:       "Just finished a grueling but rewarding 10K race! Who else loves the feeling of crossing the finish line?",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Excited to announce that my new cookbook, \"Healthy Eats for Busy Lives\", is now available for pre-order! It's packed with easy and nutritious recipes for people on the go.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:     "Who else loves practicing yoga? I just got certified as a yoga teacher and I can't wait to share my love for this practice with others.",
		PrivacyType: enums.Public,
		CommentSet:  SeedCommentDataSetA1,
	},
	{
		Content:       "Happy World Vegan Day! Being vegan has been one of the best decisions I've ever made for my health and the environment.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Just got back from an amazing week-long wellness retreat in Costa Rica. Yoga, healthy food, and beautiful scenery - what more could you ask for?",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "I recently tried a new plant-based burger and it was delicious! Who says vegan food has to be boring?",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Who else struggles with meal planning? I'm hosting a free webinar next week on how to create a personalized meal plan that works for you. Sign up now!",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Just finished a challenging but rewarding HIIT workout. Who else loves a good sweat session?",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Happy International Women's Day! Let's celebrate the amazing women in our lives and work towards a more equal and just world.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Who else loves taking their dog for a walk? My two rescue dogs, Luna and Max, are my favorite workout buddies.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
}

var SeedPostsDataSetB = []*SeedPost{
	{
		Content:       "Just got back from an incredible trip to Japan! The food, the people, and the culture were all amazing. Can't wait to go back someday!",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Who else loves a good outdoor adventure? Just went on a challenging hike up Mount Kilimanjaro and it was definitely worth it. What's your favorite hike?",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Excited to share that my latest travel article was published in National Geographic! It's all about the hidden gems of Barcelona. Check it out if you're planning a trip there soon!",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Recently started a photography project where I take a photo of something beautiful every day. It's been a great way to appreciate the little things in life.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Happy National Book Lovers Day! I just finished \"The Overstory\" by Richard Powers and it's definitely one of my new favorites. What are you currently reading?",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "I'm officially a certified scuba diver! I've always been fascinated by the ocean and it was incredible to see all the marine life up close.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Anyone else a fan of street food? I recently tried the best tacos al pastor I've ever had in Mexico City. Already planning my next trip back!",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Just got back from an amazing trip to Bali. The beaches, the temples, and the food were all incredible. Can't wait to go back someday!",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Just finished a 10-day silent meditation retreat and it was one of the most challenging and rewarding experiences of my life. Highly recommend it to anyone interested in mindfulness and inner peace.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Happy Earth Day! Let's all do our part to protect our planet and make it a better place for future generations.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
}

var SeedPostsDataSetC = []*SeedPost{
	{
		Content:       "Just launched my latest app, \"TaskMaster\", on the App Store! It's a productivity app that helps you stay on top of your to-do list. Check it out!",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Who else loves a good hackathon? Just won first place at the HackNY hackathon with my team. Can't wait for the next one.",
		PrivacyType:   enums.Public,
		LoremComments: 2,
	},
	{
		Content:       "Excited to announce that I've been accepted into the Google Developer Expert program for mobile development! It's an honor to be part of this community of experts.",
		PrivacyType:   enums.Public,
		LoremComments: 3,
	},
	{
		Content:       "Just wrapped up a successful project with a Fortune 500 company. It was a challenging but rewarding experience, and I'm proud of what our team accomplished.",
		PrivacyType:   enums.Public,
		LoremComments: 0,
	},
	{
		Content:       "Excited to share that I'll be speaking at the upcoming TechCrunch Disrupt conference about the future of mobile development. Can't wait to share my insights with the tech community!",
		PrivacyType:   enums.Public,
		LoremComments: 22,
	},
	{
		Content:       "Feeling a bit burnt out lately. The tech industry can be so demanding sometimes, and I feel like I'm always on call. Trying to take some time for self-care and relaxation, but it's tough when there's always another deadline looming.",
		PrivacyType:   enums.Private,
		LoremComments: 0,
	},
	{
		Content:       "Had a heart-to-heart with my mentor today about imposter syndrome. It's something that's been weighing on me lately, but it was helpful to hear that even seasoned developers experience it from time to time.",
		PrivacyType:   enums.Private,
		LoremComments: 0,
	},
	{
		Content:       "Dealing with a bit of imposter syndrome lately. I keep worrying that I'm not skilled enough or experienced enough to tackle the projects I'm working on. Trying to remind myself that I wouldn't have gotten this far if I didn't have the skills and knowledge to back it up.",
		PrivacyType:   enums.Private,
		LoremComments: 0,
	},
	{
		Content:       "Just got back from a weekend getaway with my partner. It was so nice to unplug from work and spend some quality time together.",
		PrivacyType:   enums.Private,
		LoremComments: 0,
	},
	{
		Content:       "Feeling grateful for my team today. We've been working on a really challenging project, but everyone has been pulling their weight and pushing us towards success. It's great to work with such talented and dedicated individuals.",
		PrivacyType:   enums.Private,
		LoremComments: 0,
	},
}

// 6 rows of ChatGPT generated comments for user A
var SeedCommentDataSetA1 = []*SeedComment{
	{
		UserEmail: "b@b.com",
		Content:   " Congrats on getting certified, that's awesome! Can't wait to attend one of your classes and learn from the best.",
	},
	{
		UserEmail: "a@a.com",
		Content:   "@Benjamin, Thank you!",
	},
	{
		UserEmail: "c@c.com",
		Content:   "I'm a big fan of yoga too! It's such a great way to unwind and de-stress after a long day at work. Congrats on becoming a teacher!",
	},
	{
		UserEmail: "a@a.com",
		Content:   "@Carlos, yoga has been a lifesaver for me in terms of managing my stress and anxiety, and I'm sure you'll find it helpful too.",
	},
	{
		UserEmail: "d@d.com",
		Content:   "I've been wanting to try yoga for ages, but never found the right teacher. Looking forward to attending one of your classes and finally giving it a go!",
	},
	{
		UserEmail: "a@a.com",
		Content:   "@Deanna, I was in the same boat as you before I found the right teacher - trust me, it's worth the wait! Looking forward to seeing you all at the studio soon.",
	},
	{
		UserEmail: "e@e.com",
		Content:   "Yoga is a game-changer! I've been practicing for years and it's helped me maintain both my physical and mental health. Excited to see you share your knowledge and passion with others.",
	},
}

// // 22 rows of lorem ipsum comments from the same person
// var SeedCommentDataSetC1 = []*SeedComment{
// 	{
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	},
// 	{
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	}, {
// 		UserEmail: "b@b.com",
// 		Content:   faker.Sentence(),
// 	},
// }
