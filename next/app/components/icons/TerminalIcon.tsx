import { faTerminal } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const TerminalIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faTerminal} />
    </div>
  );
};
