import { Injectable, Req, Res } from "@nestjs/common";
import { UpdateGroupRequestDto } from "./dto/update-group.request.dto";
import { CreateGroupRequestDto } from "./dto/create-group.request.dto";
import { RepoClient } from "../../utils/RepoClient";
import { Request, Response } from "express";
import { JwtUtils } from "../../utils/JwtUtils";
import { ObjectId, WithId } from "mongodb";
import { UsersService } from "../users/users.service";
import { Group } from "./entities/group.entity";
import { VotesService } from "../votes/votes.service";
import { AdminService } from "../admin/admin.service";

@Injectable()
export class GroupService {
	constructor(
		private readonly repoClient: RepoClient,
		private readonly jwtUtils: JwtUtils,
		private readonly usersService: UsersService,
		private readonly votesService: VotesService,
		private readonly adminService: AdminService
	) {}

	async create({ name }: CreateGroupRequestDto, @Req() request: Request) {
		const { username } = this.jwtUtils.getUser(request);
		const user = await this.repoClient.getUserByUsername(username);
		if (!user) throw new Error("No user found");

		const { insertedId } = await this.repoClient.createGroup({
			name,
			members: [username],
			owner: username,
			yearCreated: new Date().getFullYear().toString(),
		});

		const updatedGroupList = [...user.groups, `${insertedId}`];
		await this.usersService.updateUser(username, {
			groups: updatedGroupList,
		});
	}

	async remove(id: string, @Req() request: Request) {
		const groupToDelete = await this.findOne(id);
		// TODO: fix error handling
		if (!groupToDelete) throw new Error("Group not found");
		await Promise.all(
			groupToDelete.members.map(async (member) => {
				const user = await this.repoClient.getUserByUsername(member);
				if (!user) return;
				const updatedGroups = user.groups.filter(
					(groupId) => id !== groupId
				);

				await this.usersService.updateUser(member, {
					groups: updatedGroups,
				});
			})
		);

		const response = await this.repoClient.removeGroup(new ObjectId(id));
		if (!response.acknowledged || response.deletedCount !== 1) {
			throw new Error("Group was not deleted, internal server error");
		}

		return response;
	}

	async findOne(id: string) {
		return await this.repoClient.findGroup(new ObjectId(id));
	}

	async findAllOwned(year: string, request: Request) {
		const { username } = this.jwtUtils.getUser(request);

		return await this.repoClient.findAllUserOwnedGroups(year, username);
	}

	async findAllJoined(year: string, request: Request) {
		const { username } = this.jwtUtils.getUser(request);
		const user = await this.repoClient.getUserByUsername(username);
		if (!user) throw new Error("No user found");

		const joinedGroups = await Promise.all(
			user.groups.map(async (groupId) => {
				const group = await this.findOne(groupId);
				if (group && group.yearCreated === year) return group;
			})
		);

		return joinedGroups as WithId<Group>[];
	}

	async update(id: string, updateGroupDto: UpdateGroupRequestDto) {
		const groupToUpdate = await this.findOne(id);
		const updatedGroup = { ...groupToUpdate, ...updateGroupDto };

		return await this.repoClient.updateGroup(
			new ObjectId(id),
			updatedGroup
		);
	}

	async generateInvitationLink(id: string) {
		const token = await this.jwtUtils.encryptLink(id);
		const splitToken = token.split(".").join("/");
		// TODO: update with correct link and move to .env
		const link = `http://localhost:3000/groups/join/${splitToken}`;

		return link;
	}

	async addGroupToJoin(response: Response, token: string) {
		response.cookie("group", token, {
			maxAge: 1000 * 24 * 3600,
			httpOnly: true,
		});
	}

	async joinGroup(
		@Req() request: Request,
		@Res({ passthrough: true }) response: Response
	) {
		const groupToken = request.cookies?.group as string | undefined;
		if (!groupToken) return;

		const { username } = this.jwtUtils.getUser(request);
		const groupId = await this.jwtUtils.decryptLink(groupToken);

		const groupToJoin = await this.findOne(groupId);
		if (!groupToJoin) throw new Error("Group not found");
		const { members } = groupToJoin;

		const user = await this.repoClient.getUserByUsername(username);
		if (!user) throw new Error("User not found");
		const { groups } = user;

		if (!groups.includes(groupId) && !members.includes(username)) {
			await this.update(groupId, {
				members: [...groupToJoin.members, username],
			});
			await this.usersService.updateUser(username, {
				groups: [...groups, groupId],
			});
		}

		response.cookie("group", "", { maxAge: 1, httpOnly: true });
	}

	async leaveGroup(@Req() request: Request, id: string) {
		const { username } = this.jwtUtils.getUser(request);
		const user = await this.repoClient.getUserByUsername(username);

		if (!user) throw new Error("No user found");
		const groupToLeave = await this.findOne(id);
		if (!groupToLeave) throw new Error("No group found");

		const updatedMembers = groupToLeave.members.filter(
			(member) => member !== username
		);
		const updatedGroups = user.groups.filter((groupId) => groupId !== id);

		await this.update(id, { members: updatedMembers });
		await this.usersService.updateUser(username, { groups: updatedGroups });
	}

	async getGroupUserVotes(groups: WithId<Group>[]) {
		const members = groups.reduce(
			(arr, { members }) => [...arr, ...members],
			[] as string[]
		);
		const uniqueMembers = [...new Set(members)];
		const { year } = await this.adminService.getAdminConfig();

		return await this.votesService.getGroupVotes(uniqueMembers, year);
	}
}
