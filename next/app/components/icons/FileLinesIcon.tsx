import { faFileLines } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const FileLinesIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faFileLines} />
    </div>
  );
};
