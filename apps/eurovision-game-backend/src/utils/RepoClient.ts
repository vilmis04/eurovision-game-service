import { Injectable } from "@nestjs/common";

import { MongoClient, ServerApiVersion } from "mongodb";

// const { MongoClient, ServerApiVersion } = require("mongodb");
// const uri = process.env.MONGO_URI;
// const client = new MongoClient(uri, {
// 	useNewUrlParser: true,
// 	useUnifiedTopology: true,
// 	serverApi: ServerApiVersion.v1,
// });
// client.connect((err) => {
// 	const collection = client.db("test").collection("devices");
// 	// perform actions on the collection object
// 	client.close();
// });

enum CollectionTypes {
	VOTES = "votes",
}

enum DatabaseTypes {
	EUROVISION_GAME = "eurovision-game",
}

@Injectable()
export class RepoClient {
	private readonly client: MongoClient;
	constructor() {
		this.client = new MongoClient(`${process.env.MONGO_URI}`, {
			// useNewUrlParser: true,
			// useUnifiedTopology: true,
			serverApi: ServerApiVersion.v1,
		});

		// TODO: add error handling
		this.client.connect();
	}

	public getVotesByUserName = async ({ username }: { username: string }) => {
		const collection = this.client
			.db(DatabaseTypes.EUROVISION_GAME)
			.collection(CollectionTypes.VOTES);

		// TODO: add error handling
		const votes = await collection.findOne({ username });
		return votes;
	};
}
