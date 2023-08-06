import { faBackwardFast } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const BackwardFastIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faBackwardFast} />
    </div>
  );
};
