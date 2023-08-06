import { faFile } from "@fortawesome/free-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const FileIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faFile} />
    </div>
  );
};
