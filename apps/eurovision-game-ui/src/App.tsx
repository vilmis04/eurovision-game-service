// import reportWebVitals from "./reportWebVitals";
import { BrowserRouter, Routes, Route } from "react-router-dom";
// import { ThemeProvider } from "@mui/system";
import GlobalStyles from "./components/GlobalStyles/GlobalStyles";
import { StrictMode, lazy, Suspense } from "react";
import { Provider } from "react-redux";
import { store } from "./redux/store";
import { paths } from "./paths";
import PageLayout from "./components/PageLayout/PageLayout";

const VotingPage = lazy(() => import("./pages/VotingPage/VotingPage"));
const LoginPage = lazy(() => import("./pages/LoginPage/LoginPage"));
const SignupPage = lazy(() => import("./pages/SignupPage/SignupPage"));
const AdminPage = lazy(() => import("./pages/AdminPage/AdminPage"));

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
		path: paths.admin,
		element: <AdminPage />,
	},
];

const App: React.FC = () => (
	<StrictMode>
		<Provider store={store}>
			<BrowserRouter>
				<GlobalStyles />
				<Suspense>
					<Routes>
						<Route path={paths.home} element={<PageLayout />}>
							{routes.map((routeData) => (
								<Route key={routeData.path} {...routeData} />
							))}
						</Route>
					</Routes>
				</Suspense>
			</BrowserRouter>
		</Provider>
	</StrictMode>
);

export default App;
