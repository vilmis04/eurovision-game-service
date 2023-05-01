import {
	Controller,
	Get,
	Post,
	Body,
	Param,
	Delete,
	Req,
	Put,
	Res,
	Patch,
} from "@nestjs/common";
import { GroupService } from "./group.service";
import { UpdateGroupRequestDto } from "./dto/update-group.request.dto";
import { CreateGroupRequestDto } from "./dto/create-group.request.dto";
import { Request, Response } from "express";
import { RootPaths } from "../../types/paths";

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

	@Post("join-group")
	joinGroup(
		@Req() request: Request,
		@Res({ passthrough: true }) response: Response
	) {
		return this.groupService.joinGroup(request, response);
	}

	@Patch("leave/:id")
	leaveGroup(@Req() request: Request, @Param("id") id: string) {
		return this.groupService.leaveGroup(request, id);
	}
}
