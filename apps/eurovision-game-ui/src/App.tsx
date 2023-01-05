// import reportWebVitals from "./reportWebVitals";
import { BrowserRouter, Routes, Route } from "react-router-dom";
// import { ThemeProvider } from "@mui/system";
import GlobalStyles from "./components/GlobalStyles/GlobalStyles";
import React, { lazy, Suspense } from "react";
import { Provider } from "react-redux";
import { store } from "./redux/store";

const VotingPage = lazy(() => import("./pages/VotingPage/VotingPage"));
const LoginPage = lazy(() => import("./pages/LoginPage/LoginPage"));

const App: React.FC = () => (
	<React.StrictMode>
		<Provider store={store}>
			<BrowserRouter>
				<GlobalStyles />
				<Suspense>
					<Routes>
						<Route path="/" element={<VotingPage />} />
						<Route path="/login" element={<LoginPage />} />
					</Routes>
				</Suspense>
			</BrowserRouter>
		</Provider>
	</React.StrictMode>
);

export default App;
