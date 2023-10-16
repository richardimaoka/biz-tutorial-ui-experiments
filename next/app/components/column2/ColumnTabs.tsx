import styles from "./ColumnTabs.module.css";
import { ColumnTab, ColumnTabProps } from "./ColumnTab";

type Props = {
  tabs: ColumnTabProps[];
};

export async function ColumnTabs(props: Props) {
  return (
    <div className={styles.component}>
      {props.tabs.map((x) => (
        <ColumnTab
          key={x.name}
          href={x.href}
          isSelected={x.isSelected}
          name={x.name}
        />
      ))}
    </div>
  );
}
