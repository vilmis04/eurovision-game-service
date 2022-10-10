import React from "react";
// import reportWebVitals from "./reportWebVitals";
import { BrowserRouter, Routes, Route } from "react-router-dom";
// import { ThemeProvider } from "@mui/system";
import HomePage from "./pages/HomePage";
import GlobalStyles from "./components/GlobalStyles/GlobalStyles";

function App() {
  return (
    <React.StrictMode>
      <BrowserRouter>
        <GlobalStyles />
        <Routes>
          <Route path="/" element={<HomePage />}></Route>
        </Routes>
      </BrowserRouter>
    </React.StrictMode>
  );
}

export default App;
