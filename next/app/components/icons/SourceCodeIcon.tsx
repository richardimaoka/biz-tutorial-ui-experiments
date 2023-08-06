import { faCode } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const SourceCodeIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faCode} />
    </div>
  );
};
