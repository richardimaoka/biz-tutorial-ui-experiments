import { faVideo } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const VideoIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faVideo} />
    </div>
  );
};
