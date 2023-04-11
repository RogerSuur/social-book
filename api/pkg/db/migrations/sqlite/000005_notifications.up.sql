CREATE TABLE IF NOT EXISTS notification_types(
	id INTEGER PRIMARY KEY,
    entity TEXT NOT NULL,
    description TEXT NOT NULL,
    template TEXT NOT NULL
);

INSERT INTO notification_types (id, entity, description, template)
VALUES 
(0, "users", "user follow request", "User {username} has sent you a follow request."),
(1, "groups", "group invitation", "User {username} invites you to join the group {groupname}"),
(2, "groups", "group join request", "User {username} has sent you a request to join the group {groupname}"),
(3, "group_events", "group event", "User {username} has created an event {eventname} for the group {groupname}");



CREATE TABLE IF NOT EXISTS notification_details(
	id INTEGER PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    notification_type_id INTEGER NOT NULL,
    entity_id INTEGER NOT NULL,
	created_at CURRENT_TIMESTAMP DATETIME NOT NULL,
	FOREIGN KEY (sender_id) 
		REFERENCES users (id)
	FOREIGN KEY (notification_type_id) 
		REFERENCES notification_types (id)
);


CREATE TABLE IF NOT EXISTS notifications(
	id INTEGER PRIMARY KEY,
    receiver_id INTEGER NOT NULL,
    notification_details_id INTEGER NOT NULL,
    seen_at DATETIME,
    reaction BOOL,
	FOREIGN KEY (receiver_id) 
		REFERENCES users (id)
	FOREIGN KEY (notification_details_id) 
		REFERENCES notification_details (id)
);