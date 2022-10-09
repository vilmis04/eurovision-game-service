import { GlobalStyles as MUIGlobalStyles } from "@mui/material";

const GlobalStyles = () => (
  <MUIGlobalStyles
    styles={{
      html: {
        "-webkit-font-smoothing": "antialiased",
        "-moz-osx-font-smoothing": "grayscale",
        padding: 0,
        margin: 0,
        boxSizing: "border-box",
      },

      body: {
        padding: 0,
        margin: 0,
      },
    }}
  />
);

export default GlobalStyles;
