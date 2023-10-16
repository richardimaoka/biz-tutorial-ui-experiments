import { ColumnContents } from "./ColumnContents";
import { ColumnHeader } from "./ColumnHeader";
import { ColumnName } from "./tabTypes";
import styles from "./Column.module.css";

interface Props {
  selectColumn: ColumnName;
  children: React.ReactNode;
  tabs: {
    name: ColumnName;
    href: string;
  }[];
  // selectColumn?: string;
  // skipAnimation?: boolean;
}

export async function Column(props: Props) {
  return (
    <div className={styles.component}>
      <div className={styles.header}>
        <ColumnHeader tabs={props.tabs} selectTab={props.selectColumn} />
      </div>
      <div className={styles.contents}>
        <ColumnContents>{props.children}</ColumnContents>
      </div>
    </div>
  );
}
