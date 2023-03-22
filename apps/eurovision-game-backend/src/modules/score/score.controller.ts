import { Controller, Get, Post, Body, Patch, Param } from "@nestjs/common";
import { ScoreService } from "./score.service";
import { Score } from "./entities/score.entity";
import { RootPaths } from "../../types/paths";

@Controller(RootPaths.SCORES)
export class ScoreController {
	constructor(private readonly scoreService: ScoreService) {}

	@Post()
	async create(@Body() scoreInstance: Score) {
		return await this.scoreService.create(scoreInstance);
	}

	@Get(":year/:type")
	async findOne(
		@Param("year") year: Score["year"],
		@Param("type") type: Score["type"]
	) {
		return await this.scoreService.findOne(year, type);
	}

	@Patch()
	async update(@Body() scoreInstance: Score) {
		return await this.scoreService.update(scoreInstance);
	}
}
