import {
	IGetUserResponse,
	IGetVotesResponse,
} from "@eurovision-game-monorepo/core";
import { Injectable, Logger, NotFoundException } from "@nestjs/common";

import { Collection, MongoClient, ObjectId, ServerApiVersion } from "mongodb";
import { Admin } from "../modules/admin/entities/admin.entity";
import { Score } from "../modules/score/entities/score.entity";
import { initialVotes } from "./RepoClient.config";

enum CollectionTypes {
	VOTES = "votes",
	USERS = "users",
	ADMIN = "admin",
	SCORES = "scores",
}

enum DatabaseTypes {
	EUROVISION_GAME = "eurovision-game",
}

// TODO: move IGetVotes & IEditVotes to types file
interface IGetVotes {
	username: string;
}

interface IEditVotes extends IGetVotes {
	votes: IGetVotesResponse;
}

// TODO: refactor repoClient to separate modular repositories
@Injectable()
export class RepoClient {
	private readonly client: MongoClient;
	private votesCollection: Collection<IGetVotesResponse>;
	private usersCollection: Collection<IGetUserResponse>;
	private adminCollection: Collection<Admin>;
	private scoresCollection: Collection<Score>;

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
				.collection<IGetVotesResponse>(CollectionTypes.VOTES);
			this.usersCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<IGetUserResponse>(CollectionTypes.USERS);
			this.adminCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Admin>(CollectionTypes.ADMIN);
			this.scoresCollection = this.client
				.db(DatabaseTypes.EUROVISION_GAME)
				.collection<Score>(CollectionTypes.SCORES);
		});

		this.client.on("error", (err) =>
			Logger.log("[RepoClient] ERROR: " + err)
		);
	}

	// VOTES

	async getVotesByUserName({ username }: IGetVotes) {
		return await this.votesCollection.findOne({ username });
	}

	async editVotesByUserName({ username, votes }: IEditVotes) {
		// TODO: add error handling
		const response = await this.votesCollection.updateOne(
			{ username },
			{
				$set: {
					...votes,
				},
			},
			{ upsert: true }
		);

		return response;
	}
	async createNewUserVotes(username: string) {
		// TODO: add error handling
		const response = await this.votesCollection.insertOne({
			username,
			...initialVotes,
		});

		return response;
	}

	// USERS

	async getUserByUsername(username: string) {
		return await this.usersCollection.findOne({ username });
	}

	async getUserById(_id: ObjectId) {
		return await this.usersCollection.findOne({ _id });
	}

	async createUser(username: string, password: string) {
		return await this.usersCollection.insertOne({ username, password });
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
		const { year, type } = scoreInstance;
		const updatedScore = await this.scoresCollection.updateOne(
			{
				year,
				type,
			},
			scoreInstance
		);

		if (!updatedScore) throw new NotFoundException();

		return updatedScore;
	}
}
