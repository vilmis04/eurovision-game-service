import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose";
import { HydratedDocument } from "mongoose";

export type VoteDocument = HydratedDocument<Vote>;

@Schema()
export class Vote {
	@Prop({ required: true })
	username: { type: String; required: true };
	@Prop()
	armenia: String;
	@Prop()
	australia: String;
	@Prop()
	azerbaijan: String;
	@Prop()
	belgium: String;
	@Prop()
	czech_republic: String;
	@Prop()
	estonia: String;
	@Prop()
	finland: String;
	@Prop()
	france: String;
	@Prop()
	germany: String;
	@Prop()
	greece: String;
	@Prop()
	iceland: String;
	@Prop()
	italy: String;
	@Prop()
	lithuania: String;
	@Prop()
	moldova: String;
	@Prop()
	netherlands: String;
	@Prop()
	norway: String;
	@Prop()
	poland: String;
	@Prop()
	portugal: String;
	@Prop()
	romania: String;
	@Prop()
	serbia: String;
	@Prop()
	spain: String;
	@Prop()
	sweden: String;
	@Prop()
	switzerland: String;
	@Prop()
	ukraine: String;
	@Prop()
	united_kingdom: String;
}

export const VoteSchema = SchemaFactory.createForClass(Vote);
