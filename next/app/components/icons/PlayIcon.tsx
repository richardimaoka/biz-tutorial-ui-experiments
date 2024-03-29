import { faPlay } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export function PlayIcon() {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faPlay} />
    </div>
  );
}
