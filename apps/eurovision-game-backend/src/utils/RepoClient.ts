import {
	IGetUserResponse,
	IGetVotesResponse,
} from "@eurovision-game-monorepo/core";
import { Injectable, Logger } from "@nestjs/common";

import { MongoClient, ObjectId, ServerApiVersion } from "mongodb";
import { initialVotes } from "./RepoClient.config";

enum CollectionTypes {
	VOTES = "votes",
	USERS = "users",
}

enum DatabaseTypes {
	EUROVISION_GAME = "eurovision-game",
}

interface IGetVotes {
	username: string;
}

interface IEditVotes extends IGetVotes {
	votes: IGetVotesResponse;
}

@Injectable()
export class RepoClient {
	private readonly client: MongoClient;
	constructor() {
		this.client = new MongoClient(`${process.env.MONGO_URI}`, {
			serverApi: ServerApiVersion.v1,
		});

		// TODO: add error handling
		this.client.connect();

		this.client.on("connectionReady", () =>
			Logger.log("[RepoClient] Connection OK")
		);
		this.client.on("error", (err) =>
			Logger.log("[RepoClient] ERROR: " + err)
		);
	}

	// VOTES
	async getVotesByUserName({ username }: IGetVotes) {
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection<IGetVotesResponse>(CollectionTypes.VOTES);

		const votes = await collection.findOne({ username });

		return votes;
	}

	async editVotesByUserName({ username, votes }: IEditVotes) {
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection<IGetVotesResponse>(CollectionTypes.VOTES);

		// TODO: add error handling
		const response = await collection.updateOne(
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
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection<IGetVotesResponse>(CollectionTypes.VOTES);

		// TODO: add error handling
		const response = await collection.insertOne({
			username,
			...initialVotes,
		});

		return response;
	}

	// USERS
	async getUserByUsername(username: string) {
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection<IGetUserResponse>(CollectionTypes.USERS);

		const user = await collection.findOne({ username });

		return user;
	}

	async getUserById(_id: ObjectId) {
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection<IGetUserResponse>(CollectionTypes.USERS);

		const user = await collection.findOne({ _id });

		return user;
	}

	async createUser(username: string, password: string) {
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection<IGetUserResponse>(CollectionTypes.USERS);

		const user = await collection.insertOne({ username, password });

		return user;
	}
}
