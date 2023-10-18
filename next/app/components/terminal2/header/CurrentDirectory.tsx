import { source_code_pro } from "../../fonts/fonts";
import { DirectoryIcon } from "../../icons/DirectoryIcon";
import styles from "./CurrentDirectory.module.css";

interface Props {
  currentDirectory: string; //if "", then "Terminal" is displayed instead of current directory name
}

export function CurrentDirectory(props: Props) {
  return (
    <div className={styles.component}>
      {props.currentDirectory && <DirectoryIcon />}
      <span>
        {props.currentDirectory ? props.currentDirectory : "Terminal"}
      </span>
    </div>
  );
}
