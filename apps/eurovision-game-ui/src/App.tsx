// import reportWebVitals from "./reportWebVitals";
import { BrowserRouter, Routes, Route } from "react-router-dom";
// import { ThemeProvider } from "@mui/system";
import GlobalStyles from "./components/GlobalStyles/GlobalStyles";
import { StrictMode, lazy, Suspense } from "react";
import { Provider } from "react-redux";
import { store } from "./redux/store";
import { paths } from "./paths";
import PageLayout from "./components/PageLayout/PageLayout";
import AuthProvider from "./components/AuthProvider/AuthProvider";

const VotingPage = lazy(() => import("./pages/VotingPage/VotingPage"));
const LoginPage = lazy(() => import("./pages/LoginPage/LoginPage"));
const SignupPage = lazy(() => import("./pages/SignupPage/SignupPage"));
const AdminPage = lazy(() => import("./pages/AdminPage/AdminPage"));
const GroupPage = lazy(() => import("./pages/GroupPage/GroupPage"));
const JoinGroup = lazy(() => import("./components/JoinGroup/JoinGroup"));
const ForbiddenScreen = lazy(
	() => import("./components/ForbiddenScreen/ForbiddenScreen")
);

const routes = [
	{
		path: paths.home,
		element: <VotingPage />,
	},
	{
		path: paths.login,
		element: <LoginPage />,
	},
	{
		path: paths.signup,
		element: <SignupPage />,
	},
	{
		path: paths.groups,
		element: <GroupPage />,
	},
	{
		path: "/groups/join/:token1/:token2/:token3",
		element: <JoinGroup />,
	},
];

const protectedRoutes = [
	{
		path: paths.admin,
		element: <AuthProvider />,
		subpaths: [
			{
				element: <AdminPage />,
				index: true,
			},
		],
	},
];

const App: React.FC = () => {
	return (
		<StrictMode>
			<Provider store={store}>
				<BrowserRouter>
					<GlobalStyles />
					<Suspense>
						<Routes>
							<Route path={paths.home} element={<PageLayout />}>
								{routes.map((routeData) => (
									<Route
										key={routeData.path}
										{...routeData}
									/>
								))}
							</Route>
							{protectedRoutes.map(
								({ element, path, subpaths }) => (
									<Route
										key={path}
										element={element}
										path={path}
									>
										{subpaths.map(
											(
												{ element, index: isIndex },
												index
											) => (
												<Route
													key={index}
													element={element}
													index={isIndex}
												/>
											)
										)}
									</Route>
								)
							)}
							<Route
								path={paths.forbidden}
								element={<ForbiddenScreen />}
							/>
						</Routes>
					</Suspense>
				</BrowserRouter>
			</Provider>
		</StrictMode>
	);
};

export default App;
