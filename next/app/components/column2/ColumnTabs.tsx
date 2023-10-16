import styles from "./ColumnTabs.module.css";
import { ColumnTab } from "./ColumnTab";
import { TabName } from "./tabTypes";

type Props = {
  tabs: {
    name: TabName;
    href: string;
  }[];
  selectTab: TabName;
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
