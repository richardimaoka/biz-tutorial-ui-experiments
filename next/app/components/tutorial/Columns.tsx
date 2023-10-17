import { Column } from "./Column";
import styles from "./Columns.module.css";
import { TutorialColumnProps } from "./definitions";

interface Props {
  columns: TutorialColumnProps[];
}

export async function Columns(props: Props) {
  return (
    <div className={styles.component}>
      {props.columns.map((c) => (
        <div key={c.kind} className={styles.column}>
          <Column column={c} />
        </div>
      ))}
    </div>
  );
}
