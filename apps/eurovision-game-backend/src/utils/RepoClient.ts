import {
	GameTypes,
	IGetUserResponse,
	IGetVotesRequest,
	RoleTypes,
} from "@eurovision-game-monorepo/core";
import { Injectable, Logger, NotFoundException } from "@nestjs/common";

import { Collection, MongoClient, ObjectId, ServerApiVersion } from "mongodb";
import { Admin } from "../modules/admin/entities/admin.entity";
import { Score } from "../modules/score/entities/score.entity";
import { Country } from "../modules/country/entities/country.entity";
import { Votes } from "../modules/votes/entities/Votes";

enum CollectionTypes {
	VOTES = "votes",
	USERS = "users",
	ADMIN = "admin",
	SCORES = "scores",
	COUNTRY = "country",
}

enum DatabaseTypes {
	EUROVISION_GAME = "eurovision-game",
}

// TODO: refactor repoClient to separate modular repositories
@Injectable()
export class RepoClient {
	private readonly client: MongoClient;
	private votesCollection: Collection<Votes>;
	private usersCollection: Collection<IGetUserResponse>;
	private adminCollection: Collection<Admin>;
	private scoresCollection: Collection<Score>;
	private countryCollection: Collection<Country>;

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
				.collection<IGetUserResponse>(CollectionTypes.USERS);
			this.adminCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Admin>(CollectionTypes.ADMIN);
			this.scoresCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Score>(CollectionTypes.SCORES);
			this.countryCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Country>(CollectionTypes.COUNTRY);
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

	// SCORES

	async createScore(scoreInstance: Score) {
		return await this.scoresCollection.insertOne(scoreInstance);
	}

	async findScore(year: Score["year"], type: Score["type"]) {
		const score = await this.scoresCollection.findOne({ year, type });

		if (!score) throw new NotFoundException({ message: "Score not found" });

		return score;
	}

	async updateScore(scoreInstance: Score) {
		const { year, type, countries } = scoreInstance;
		const updatedScore = await this.scoresCollection.updateOne(
			{
				year,
				type,
			},
			{ $set: { year, type, countries } }
		);

		if (!updatedScore) throw new NotFoundException();

		return updatedScore;
	}

	// COUNTRY

	async createCountry(country: Country) {
		return await this.countryCollection.insertOne(country);
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
}
