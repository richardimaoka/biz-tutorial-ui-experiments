import { faChevronCircleUp } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export function ChevronDownIcon() {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faChevronCircleUp} />
    </div>
  );
}
