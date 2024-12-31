// Example model schema from the Drizzle docs
// https://orm.drizzle.team/docs/sql-schema-declaration

import { int, sqliteTableCreator, text } from "drizzle-orm/sqlite-core";

/**
 * This is an example of how to use the multi-project schema feature of Drizzle ORM. Use the same
 * database instance for multiple projects.
 *
 * @see https://orm.drizzle.team/docs/goodies#multi-project-schema
 */
export const createTable = sqliteTableCreator(
  (name) => `contract-rummy_${name}`,
);

export const games = createTable("games", {
  id: int("id", { mode: "number" }).primaryKey({ autoIncrement: true }),
  date: text("date", { length: 256 }),
});

export const users = createTable("users", {
  id: int("id", { mode: "number" }).primaryKey({ autoIncrement: true }),
  name: text("name", { length: 256 }),
});

export const scores = createTable("scores", {
  id: int("id", { mode: "number" }).primaryKey({ autoIncrement: true }),
  gameId: int("gameId", { mode: "number" }),
  roundNo: int("roundNo", { mode: "number" }),
  userId: int("userId", { mode: "number" }),
  score: int("score", { mode: "number" }),
});
