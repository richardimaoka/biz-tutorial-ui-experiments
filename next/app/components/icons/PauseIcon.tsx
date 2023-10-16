import { faPause } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export function PauseIcon() {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faPause} />
    </div>
  );
}
