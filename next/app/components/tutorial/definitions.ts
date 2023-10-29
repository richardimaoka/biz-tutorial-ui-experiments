import { TerminalComponentProps } from "../terminal2/__TerminalComponent";

export type ColumnName =
  | "Terminal"
  | "SourceCode"
  | "ImageDescription"
  | "BackgroundImage"
  | "Markdown"
  | "YouTube"
  | "Browser"
  | "DevTools";

export const columnWidthPx = 768;

export type TerminalColumnProps = {
  kind: "Terminal";
} & TerminalComponentProps;

export type TutorialColumnProps = TerminalColumnProps;
