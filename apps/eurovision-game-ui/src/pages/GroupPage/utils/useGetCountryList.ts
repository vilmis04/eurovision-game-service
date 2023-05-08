import {
	GameTypes,
	ICountry,
	IFinalsList,
	VoteTypes,
	mapSemiCountryList,
} from "@eurovision-game-monorepo/core";
import { useGetCountryListQuery } from "../../@modules/country.api";
import { useEffect, useState } from "react";

const getPosition = (index: number): VoteTypes => {
	const position = index + 1;
	if (position === 1) return VoteTypes.P1;
	if (position > 1 && position <= 5) return VoteTypes.P2_5;
	if (position > 5 && position <= 10) return VoteTypes.P6_10;
	if (position > 10 && position <= 15) return VoteTypes.P11_15;
	if (position > 15 && position <= 20) return VoteTypes.P16_20;
	return VoteTypes.P21_25;
};

const mapFinalsList = (countryList: ICountry[] | undefined) => {
	const list = countryList || [];
	return list.length > 0
		? [...list]
				.sort((a, b) => b.finalScore - a.finalScore)
				.map(({ name }, index) => ({
					name,
					position: getPosition(index),
				}))
		: [];
};

export const useGetCountryList = (
	type?: GameTypes,
	configYear?: string,
	calculateFinalsScore?: boolean
) => {
	const [semiWinnersList1, setSemiWinnersList1] = useState<string[]>([]);
	const [semiWinnersList2, setSemiWinnersList2] = useState<string[]>([]);
	const [finalsList, setFinalsList] = useState<IFinalsList[]>([]);
	const currentYear = new Date().getFullYear().toString();
	const year = configYear || currentYear;
	const { data: semiList1, isFetching: isFetchingSemi1 } =
		useGetCountryListQuery(
			{
				type: GameTypes.SEMI_1,
				year,
			},
			{
				pollingInterval: type === GameTypes.SEMI_1 ? 30000 : 0,
			}
		);
	const { data: semiList2, isFetching: isFetchingSemi2 } =
		useGetCountryListQuery(
			{
				type: GameTypes.SEMI_2,
				year,
			},
			{
				pollingInterval: type === GameTypes.SEMI_2 ? 30000 : 0,
			}
		);
	const { data: finalData, isFetching: isFetchingFinalList } =
		useGetCountryListQuery({
			type: GameTypes.FINAL,
			year,
		});

	const isFetching =
		isFetchingSemi1 || isFetchingSemi2 || isFetchingFinalList;

	useEffect(() => {
		if (!isFetching) {
			setSemiWinnersList1(mapSemiCountryList(semiList1));
			setSemiWinnersList2(mapSemiCountryList(semiList2));
			setFinalsList(calculateFinalsScore ? mapFinalsList(finalData) : []);
		}
	}, [isFetching]);

	return { semiWinnersList1, semiWinnersList2, finalsList, isFetching };
};
