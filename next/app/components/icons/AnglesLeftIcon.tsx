import { faAnglesLeft } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const AnglesLeftIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faAnglesLeft} />
    </div>
  );
};
