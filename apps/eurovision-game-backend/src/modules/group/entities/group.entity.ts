import { IGroup } from "@eurovision-game-monorepo/core";

export class Group {
	name: IGroup["name"];
	members: IGroup["members"];
	owner: IGroup["owner"];
	yearCreated: IGroup["yearCreated"];
}
