import { source_code_pro } from "../fonts/fonts";
import { DirectoryIcon } from "../icons/DirectoryIcon";
import styles from "./TerminalHeader.module.css";

interface Props {
  currentDirectory: string;
}

export function TerminalHeader(props: Props) {
  return (
    <div className={styles.component}>
      {props.currentDirectory && <DirectoryIcon />}
      <span className={source_code_pro.className}>
        {props.currentDirectory ? props.currentDirectory : "Terminal"}
      </span>
    </div>
  );
}
