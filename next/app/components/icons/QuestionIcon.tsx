import { faQuestion } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./style.module.css";

export const QuestionIcon = (): JSX.Element => {
  return (
    <div className={styles.icon}>
      <FontAwesomeIcon fixedWidth icon={faQuestion} />
    </div>
  );
};
