import styles from "./ColumnTabs.module.css";
import { ColumnTab } from "./ColumnTab";
import { ColumnName } from "./definitions";

type Props = {
  tabs: {
    name: ColumnName;
    href: string;
  }[];
  selectTab: ColumnName;
};

export async function ColumnTabs(props: Props) {
  return (
    <div className={styles.component}>
      {props.tabs.map((x) => (
        <ColumnTab
          key={x.name}
          href={x.href}
          isSelected={x.name === props.selectTab}
          name={x.name}
        />
      ))}
    </div>
  );
}
