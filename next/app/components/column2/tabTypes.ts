export type TabName =
  | "Terminal"
  | "SourceCode"
  | "ImageDescription"
  | "BackgroundImage"
  | "Markdown"
  | "YouTube"
  | "Browser"
  | "DevTools";

export type TabProperties = {
  isSelected: boolean;
  name: TabName;
  href: string;
};
