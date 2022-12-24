import { IGetVotesResponse } from "@eurovision-game-monorepo/core";
import { Injectable } from "@nestjs/common";

import { MongoClient, ServerApiVersion } from "mongodb";

enum CollectionTypes {
	VOTES = "votes",
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
	}

	public getVotesByUserName = async ({ username }: IGetVotes) => {
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection<IGetVotesResponse>(CollectionTypes.VOTES);

		// TODO: add error handling
		const votes = await collection.findOne({ username });
		return votes;
	};

	public editVotesByUserName = async ({ username, votes }: IEditVotes) => {
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
	};
}
