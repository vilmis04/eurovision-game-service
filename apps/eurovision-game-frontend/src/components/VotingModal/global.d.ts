import { SxProps, Theme } from "@mui/material";

declare global {
  interface ISx {
    [key: string]: SxProps<Theme>;
  }
}

export {};
