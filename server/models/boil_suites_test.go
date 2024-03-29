// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGS)
	t.Run("Favorites", testFavorites)
	t.Run("HubUsers", testHubUsers)
	t.Run("Hubs", testHubs)
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("UserFavorites", testUserFavorites)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSDelete)
	t.Run("Favorites", testFavoritesDelete)
	t.Run("HubUsers", testHubUsersDelete)
	t.Run("Hubs", testHubsDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("UserFavorites", testUserFavoritesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSQueryDeleteAll)
	t.Run("Favorites", testFavoritesQueryDeleteAll)
	t.Run("HubUsers", testHubUsersQueryDeleteAll)
	t.Run("Hubs", testHubsQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("UserFavorites", testUserFavoritesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSSliceDeleteAll)
	t.Run("Favorites", testFavoritesSliceDeleteAll)
	t.Run("HubUsers", testHubUsersSliceDeleteAll)
	t.Run("Hubs", testHubsSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("UserFavorites", testUserFavoritesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSExists)
	t.Run("Favorites", testFavoritesExists)
	t.Run("HubUsers", testHubUsersExists)
	t.Run("Hubs", testHubsExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("UserFavorites", testUserFavoritesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSFind)
	t.Run("Favorites", testFavoritesFind)
	t.Run("HubUsers", testHubUsersFind)
	t.Run("Hubs", testHubsFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("UserFavorites", testUserFavoritesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSBind)
	t.Run("Favorites", testFavoritesBind)
	t.Run("HubUsers", testHubUsersBind)
	t.Run("Hubs", testHubsBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("UserFavorites", testUserFavoritesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSOne)
	t.Run("Favorites", testFavoritesOne)
	t.Run("HubUsers", testHubUsersOne)
	t.Run("Hubs", testHubsOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("UserFavorites", testUserFavoritesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSAll)
	t.Run("Favorites", testFavoritesAll)
	t.Run("HubUsers", testHubUsersAll)
	t.Run("Hubs", testHubsAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("UserFavorites", testUserFavoritesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSCount)
	t.Run("Favorites", testFavoritesCount)
	t.Run("HubUsers", testHubUsersCount)
	t.Run("Hubs", testHubsCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("UserFavorites", testUserFavoritesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSHooks)
	t.Run("Favorites", testFavoritesHooks)
	t.Run("HubUsers", testHubUsersHooks)
	t.Run("Hubs", testHubsHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("UserFavorites", testUserFavoritesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSInsert)
	t.Run("ChatMSGS", testChatMSGSInsertWhitelist)
	t.Run("Favorites", testFavoritesInsert)
	t.Run("Favorites", testFavoritesInsertWhitelist)
	t.Run("HubUsers", testHubUsersInsert)
	t.Run("HubUsers", testHubUsersInsertWhitelist)
	t.Run("Hubs", testHubsInsert)
	t.Run("Hubs", testHubsInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("UserFavorites", testUserFavoritesInsert)
	t.Run("UserFavorites", testUserFavoritesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ChatMSGToHubUsingHub", testChatMSGToOneHubUsingHub)
	t.Run("ChatMSGToUserUsingUser", testChatMSGToOneUserUsingUser)
	t.Run("HubUserToHubUsingHub", testHubUserToOneHubUsingHub)
	t.Run("HubUserToUserUsingUser", testHubUserToOneUserUsingUser)
	t.Run("HubToUserUsingUser", testHubToOneUserUsingUser)
	t.Run("UserFavoriteToFavoriteUsingFavorite", testUserFavoriteToOneFavoriteUsingFavorite)
	t.Run("UserFavoriteToUserUsingUser", testUserFavoriteToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("FavoriteToUserFavorites", testFavoriteToManyUserFavorites)
	t.Run("HubToChatMSGS", testHubToManyChatMSGS)
	t.Run("HubToHubUsers", testHubToManyHubUsers)
	t.Run("UserToChatMSGS", testUserToManyChatMSGS)
	t.Run("UserToHubUsers", testUserToManyHubUsers)
	t.Run("UserToHubs", testUserToManyHubs)
	t.Run("UserToUserFavorites", testUserToManyUserFavorites)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ChatMSGToHubUsingChatMSGS", testChatMSGToOneSetOpHubUsingHub)
	t.Run("ChatMSGToUserUsingChatMSGS", testChatMSGToOneSetOpUserUsingUser)
	t.Run("HubUserToHubUsingHubUsers", testHubUserToOneSetOpHubUsingHub)
	t.Run("HubUserToUserUsingHubUsers", testHubUserToOneSetOpUserUsingUser)
	t.Run("HubToUserUsingHubs", testHubToOneSetOpUserUsingUser)
	t.Run("UserFavoriteToFavoriteUsingUserFavorites", testUserFavoriteToOneSetOpFavoriteUsingFavorite)
	t.Run("UserFavoriteToUserUsingUserFavorites", testUserFavoriteToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("FavoriteToUserFavorites", testFavoriteToManyAddOpUserFavorites)
	t.Run("HubToChatMSGS", testHubToManyAddOpChatMSGS)
	t.Run("HubToHubUsers", testHubToManyAddOpHubUsers)
	t.Run("UserToChatMSGS", testUserToManyAddOpChatMSGS)
	t.Run("UserToHubUsers", testUserToManyAddOpHubUsers)
	t.Run("UserToHubs", testUserToManyAddOpHubs)
	t.Run("UserToUserFavorites", testUserToManyAddOpUserFavorites)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSReload)
	t.Run("Favorites", testFavoritesReload)
	t.Run("HubUsers", testHubUsersReload)
	t.Run("Hubs", testHubsReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("UserFavorites", testUserFavoritesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSReloadAll)
	t.Run("Favorites", testFavoritesReloadAll)
	t.Run("HubUsers", testHubUsersReloadAll)
	t.Run("Hubs", testHubsReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("UserFavorites", testUserFavoritesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSSelect)
	t.Run("Favorites", testFavoritesSelect)
	t.Run("HubUsers", testHubUsersSelect)
	t.Run("Hubs", testHubsSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("UserFavorites", testUserFavoritesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSUpdate)
	t.Run("Favorites", testFavoritesUpdate)
	t.Run("HubUsers", testHubUsersUpdate)
	t.Run("Hubs", testHubsUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("UserFavorites", testUserFavoritesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("ChatMSGS", testChatMSGSSliceUpdateAll)
	t.Run("Favorites", testFavoritesSliceUpdateAll)
	t.Run("HubUsers", testHubUsersSliceUpdateAll)
	t.Run("Hubs", testHubsSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("UserFavorites", testUserFavoritesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
