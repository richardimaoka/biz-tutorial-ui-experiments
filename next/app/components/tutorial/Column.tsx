import styles from "./Columns.module.css";
import { TutorialColumnProps } from "./definitions";

interface Props {
  column: TutorialColumnProps;
}

export async function Columns(props: Props) {
  return <div className={styles.component}></div>;
}
