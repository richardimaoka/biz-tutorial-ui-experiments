import { faCirclePlay } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export function CirclePlayIcon() {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faCirclePlay} />
    </div>
  );
}
