import {
	Controller,
	Get,
	Post,
	Body,
	Param,
	Delete,
	Req,
	Put,
	Patch,
} from "@nestjs/common";
import { GroupService } from "./group.service";
import { UpdateGroupRequestDto } from "./dto/update-group.request.dto";
import { CreateGroupRequestDto } from "./dto/create-group.request.dto";
import { Request } from "express";
import { RootPaths } from "../../types/paths";
import { Group } from "./entities/group.entity";
import { WithId } from "mongodb";

@Controller(RootPaths.GROUPS)
export class GroupController {
	constructor(private readonly groupService: GroupService) {}

	@Post()
	create(
		@Body() createGroupDto: CreateGroupRequestDto,
		@Req() request: Request
	) {
		return this.groupService.create(createGroupDto, request);
	}

	@Delete(":id")
	remove(@Param("id") id: string, @Req() request: Request) {
		return this.groupService.remove(id, request);
	}

	@Get(":id")
	findOne(@Param("id") id: string) {
		return this.groupService.findOne(id);
	}

	@Get("all/:year")
	findAll(@Param("year") year: string, @Req() request: Request) {
		return this.groupService.findAllJoined(year, request);
	}

	@Put(":id")
	update(
		@Param("id") id: string,
		@Body() updateGroupDto: UpdateGroupRequestDto
	) {
		return this.groupService.update(id, updateGroupDto);
	}

	@Post("invitation-link/:groupId")
	generateInvitationLink(@Param("groupId") id: string) {
		return this.groupService.generateInvitationLink(id);
	}

	@Patch("leave/:id")
	leaveGroup(@Req() request: Request, @Param("id") id: string) {
		return this.groupService.leaveGroup(request, id);
	}

	@Post("scores")
	getGroupUserVotes(@Body() groups: WithId<Group>[]) {
		return this.groupService.getGroupUserVotes(groups);
	}
}
