CREATE TABLE `contract-rummy_games` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`date` text(256)
);
--> statement-breakpoint
CREATE TABLE `contract-rummy_scores` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`gameId` integer,
	`roundNo` integer,
	`userId` integer,
	`score` integer
);
--> statement-breakpoint
CREATE TABLE `contract-rummy_users` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`name` text(256)
);
