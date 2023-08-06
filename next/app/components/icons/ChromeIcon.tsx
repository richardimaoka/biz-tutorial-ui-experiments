import { faChrome } from "@fortawesome/free-brands-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const ChromeIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faChrome} />
    </div>
  );
};
