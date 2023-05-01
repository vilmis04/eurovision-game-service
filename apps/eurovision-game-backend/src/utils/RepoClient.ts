import {
	GameTypes,
	IGetVotesRequest,
	RoleTypes,
} from "@eurovision-game-monorepo/core";
import { Injectable, Logger } from "@nestjs/common";

import {
	Collection,
	Condition,
	MongoClient,
	ObjectId,
	ServerApiVersion,
} from "mongodb";
import { Admin } from "../modules/admin/entities/admin.entity";
import { Country } from "../modules/country/entities/country.entity";
import { Votes } from "../modules/votes/entities/Votes";
import { Group } from "../modules/group/entities/group.entity";
import { UpdateGroupRequestDto } from "../modules/group/dto/update-group.request.dto";
import { User } from "../modules/users/entities/user.entity";

enum CollectionTypes {
	VOTES = "votes",
	USERS = "users",
	ADMIN = "admin",
	SCORES = "scores",
	COUNTRY = "country",
	GROUP = "group",
}

enum DatabaseTypes {
	EUROVISION_GAME = "eurovision-game",
}

// TODO: refactor repoClient to separate modular repositories
@Injectable()
export class RepoClient {
	private readonly client: MongoClient;
	private votesCollection: Collection<Votes>;
	private usersCollection: Collection<User>;
	private adminCollection: Collection<Admin>;
	private countryCollection: Collection<Country>;
	private groupCollection: Collection<Group>;

	constructor() {
		this.client = new MongoClient(`${process.env.MONGO_URI}`, {
			serverApi: ServerApiVersion.v1,
		});

		// TODO: add error handling
		this.client.connect();

		this.client.on("connectionReady", () => {
			Logger.log("[RepoClient] Connection OK");
			this.votesCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Votes>(CollectionTypes.VOTES);
			this.usersCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<User>(CollectionTypes.USERS);
			this.adminCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Admin>(CollectionTypes.ADMIN);
			this.countryCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Country>(CollectionTypes.COUNTRY);
			this.groupCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Group>(CollectionTypes.GROUP);
		});

		this.client.on("error", (err) =>
			Logger.log("[RepoClient] ERROR: " + err)
		);
	}

	// VOTES

	async getVotes({ username, type, year }: IGetVotesRequest) {
		return await this.votesCollection.findOne({ username, type, year });
	}

	async updateVotes({ username, type, year, votes }: Votes) {
		const response = await this.votesCollection.updateOne(
			{ username, type, year },
			{
				$set: {
					votes,
				},
			},
			{ upsert: true }
		);

		return response;
	}
	async createVotes(voteData: IGetVotesRequest) {
		const votes = { ...voteData, votes: {} };
		await this.votesCollection.insertOne(votes);

		return votes;
	}

	// USERS

	async getUserByUsername(username: string) {
		return await this.usersCollection.findOne({ username });
	}

	async getUserById(_id: ObjectId) {
		return await this.usersCollection.findOne({ _id });
	}

	async createUser(username: string, password: string) {
		return await this.usersCollection.insertOne({
			username,
			password,
			roles: [RoleTypes.USER],
			groups: [],
		});
	}

	async updateUser(username: string, user: Partial<User>) {
		return await this.usersCollection.updateOne(
			{
				username,
			},
			{ $set: { ...user } }
		);
	}

	// ADMIN

	async getAdminConfig() {
		const adminWithId = await this.adminCollection.find({}).toArray();

		const { _id, ...admin } = adminWithId[0]; // 0 because there is only one admin entry

		return admin;
	}

	async updateAdminConfig(adminConfig: Admin) {
		const adminArray = await this.adminCollection.find({}).toArray();

		const updatedAdmin = {
			...adminArray[0],
			...adminConfig,
		};

		const adminResponse = await this.adminCollection.updateOne(
			{ _id: updatedAdmin._id },
			{ $set: adminConfig }
		);

		return adminResponse;
	}

	// COUNTRY

	async createCountry(country: Country) {
		return await this.countryCollection.insertOne(country);
	}

	async createManyCountries(countries: Country[]) {
		return await this.countryCollection.insertMany(countries);
	}

	async findCountries(year: string, type: GameTypes) {
		return this.countryCollection.find({ year, type }).toArray();
	}

	async getOneCountry(year: string, name: string) {
		return await this.countryCollection.findOne({ year, name });
	}

	async removeCountry(year: string, name: string) {
		return await this.countryCollection.deleteOne({ year, name });
	}

	async updateCountry(country: Partial<Country>) {
		const { name, year } = country;
		return await this.countryCollection.updateOne(
			{ name, year },
			{ $set: { ...country } }
		);
	}

	// GROUP

	async createGroup(group: Group) {
		return await this.groupCollection.insertOne(group);
	}

	async removeGroup(id: Condition<ObjectId>) {
		return await this.groupCollection.deleteOne({
			_id: new ObjectId(`${id}`),
		});
	}

	async findGroup(id: Condition<ObjectId>) {
		return await this.groupCollection.findOne({ _id: id });
	}

	async findAllUserOwnedGroups(yearCreated: string, owner: string) {
		return await this.groupCollection
			.find({ owner, yearCreated })
			.toArray();
	}

	async updateGroup(
		id: Condition<ObjectId>,
		updateGroupDto: UpdateGroupRequestDto
	) {
		return await this.groupCollection.findOneAndUpdate(
			{ _id: id },
			{ $set: { ...updateGroupDto } }
		);
	}
}
