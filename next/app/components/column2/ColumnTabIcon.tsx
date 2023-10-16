import { TabName } from "./tabTypes";
import { ChromeIcon } from "../icons/ChromeIcon";
import { FileLinesIcon } from "../icons/FileLinesIcon";
import { SourceCodeIcon } from "../icons/SourceCodeIcon";
import { TerminalIcon } from "../icons/TerminalIcon";
import { VideoIcon } from "../icons/VideoIcon";

interface Props {
  name: TabName;
}

export function ColumnTabIcon({ name }: Props): JSX.Element {
  switch (name) {
    case "BackgroundImage":
      return <ChromeIcon />;
    case "Browser":
      return <ChromeIcon />;
    case "DevTools":
      return <ChromeIcon />;
    case "ImageDescription":
      return <FileLinesIcon />;
    case "Markdown":
      return <FileLinesIcon />;
    case "SourceCode":
      return <SourceCodeIcon />;
    case "Terminal":
      return <TerminalIcon />;
    case "YouTube":
      return <VideoIcon />;
  }
}
