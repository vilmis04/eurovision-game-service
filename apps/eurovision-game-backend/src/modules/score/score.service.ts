import { Injectable } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { Score } from "./entities/score.entity";

@Injectable()
export class ScoreService {
	constructor(private readonly repoClient: RepoClient) {}

	async create(scoreInstance: Score) {
		return await this.repoClient.createScore(scoreInstance);
	}

	async findOne(year: Score["year"], type: Score["type"]) {
		return await this.repoClient.findScore(year, type);
	}

	async update(scoreInstance: Score) {
		return await this.repoClient.updateScore(scoreInstance);
	}
}
