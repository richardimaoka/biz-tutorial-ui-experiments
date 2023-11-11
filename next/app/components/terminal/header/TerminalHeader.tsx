import { CurrentDirectory } from "./CurrentDirectory";
import styles from "./TerminalHeader.module.css";
import { TerminalTabs } from "./TerminalTabs";

interface Props {
  currentDirectory: string;
  tabs?: {
    name: string;
    href: string;
  }[];
  selectTab: string;
}

export type TerminalHeaderProps = Props;

export function TerminalHeader(props: Props) {
  return (
    <div className={styles.component}>
      <CurrentDirectory currentDirectory={props.currentDirectory} />
      {props.tabs && props.tabs.length > 1 && (
        <TerminalTabs tabs={props.tabs} selectTab={props.selectTab} />
      )}
    </div>
  );
}
