import { PartialType } from "@nestjs/mapped-types";
import { Group } from "../entities/group.entity";

export class UpdateGroupRequestDto extends PartialType(Group) {}
