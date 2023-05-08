module.exports = {
	apps: [
		{
			name: "ev-game-be",
			script: "./dist/apps/eurovision-game-backend/main.js",
			instances: 1,
			exec_mode: "cluster",
			env: {
				NODE_ENV: "production",
				PORT: "",
				MONGO_URI: "",
				SECRET_KEY: "",
				BASE_URL: "http://134.122.81.2",
			},
		},
	],
};
