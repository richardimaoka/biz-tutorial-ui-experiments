import { faFolderTree } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const FolderTreeIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faFolderTree} />
    </div>
  );
};
