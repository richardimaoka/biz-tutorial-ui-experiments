import { faFolder } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const DirectoryIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon className={styles.folder} icon={faFolder} />
    </div>
  );
};
